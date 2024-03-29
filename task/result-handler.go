package task

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/k8s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

const SizeOf1Mb = 1024 * 1024

// callbackUrl is the environment variable for the callback URL to push results.
var callbackUrl = os.Getenv(CallbackURL)

// resultSecretKey is the environment variable for the result secret name.
var resultSecretKey = os.Getenv(ResultSecretName)

// HandleSuccess handles successful task execution by processing the result and reporting records.
func HandleSuccess(result map[string]interface{}, records interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         0,
		OutputProperties: result,
		ReportingRecords: records,
	}
	handleResult(outputContext)
}

// HandleError handles task execution error by processing the result, error message and reporting records.
func HandleError(err error, result map[string]interface{}, records []interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         -1,
		OutputProperties: result,
		JobErrorMessage:  err.Error(),
		ReportingRecords: records,
	}
	handleResult(outputContext)
}

func HandleAbort(result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         134,
		OutputProperties: result,
		JobErrorMessage:  (&AbortError{}).Error(),
	}
	handleResult(outputContext)
}

// SkipResultHandler represents a result handler that should be skipped.
type SkipResultHandler struct{}

// Error returns the error message for the SkipResultHandler.
func (m *SkipResultHandler) Error() string {
	return "skipping result handler"
}

// handleResult handles the task output context by encrypting and processing it with different result handlers.
func handleResult(outputContext TaskOutputContext) {
	data, _ := json.Marshal(outputContext)
	encryptedData, encryptErr := Encrypt(data)
	if encryptErr != nil {
		klog.Fatalf("error encrypting data [%v]", encryptErr)
	}

	done := make(chan string, 3)
	success := make(chan bool)

	go func() {
		err := pushResult(encryptedData, shouldRetryCallbackRequest(encryptedData))
		handleResultHandlerError("HTTP Push", done, success, err)
	}()
	go func() {
		err := writeToSecret(encryptedData)
		handleResultHandlerError("Secret", done, success, err)
	}()
	go func() {
		err := writeOutput(encryptedData)
		handleResultHandlerError("Output File", done, success, err)
	}()
	for i := 0; i < 3; i++ {
		klog.Infof("%s", <-done)
	}
	select {
	case <-success:
		klog.Infof("Successfully processed, result with one of the result handlers")
	default:
		klog.Fatalf("Result couldn't be processed, exiting execution with error.")
	}
}

// handleResultHandlerError handles the error from a result handler and updates the done and success channels accordingly.
func handleResultHandlerError(handler string, done chan string, success chan bool, err error) {
	if err == nil {
		done <- fmt.Sprintf("Successfully finished result return using %s.", handler)
		success <- true
	} else if errors.Is(err, &SkipResultHandler{}) {
		done <- fmt.Sprintf("Skipped %s result handler because of missing environment variable, other processors can pick the result up.", handler)
	} else {
		done <- fmt.Sprintf("Error occurred while trying to use %s result handler: %s", handler, err)
	}
}

// splitSecretResourceData splits the secret entry string into namespace, name, and key.
func splitSecretResourceData(secretEntry string) (string, string, string, error) {
	split := strings.Split(secretEntry, ":")
	if len(split) != 3 {
		return "", "", "", nil
	}
	return split[0], split[1], split[2], nil
}

// writeToSecret writes the encrypted data to the result secret if the resultSecretKey is set.
func writeToSecret(encryptedData []byte) error {
	if len(resultSecretKey) > 0 {
		if len(encryptedData) >= SizeOf1Mb {
			return errors.New("result size exceeds 1Mb and is too big to store in secret")
		}
		namespace, name, key, err := splitSecretResourceData(resultSecretKey)
		if err != nil {
			klog.Warningf("Cannot resolve value of Result Secret Name and Key %s, skipping - output written to output file", err)
			return err
		}
		clientset, err := k8s.GetClientset()
		if err != nil {
			klog.Warningf("Cannot create clientset for handling Result Secret: %s", err)
			return err
		}

		secret, err := clientset.CoreV1().Secrets(namespace).Get(context.Background(), name, v1.GetOptions{})
		if err != nil {
			klog.Warningf("Cannot fetch Result Secret: %s", err)
			return err
		}
		if secret.Data == nil {
			secret.Data = make(map[string][]byte)
		}
		secret.Data[key] = encryptedData
		_, err = clientset.CoreV1().Secrets(namespace).Update(context.TODO(), secret, v1.UpdateOptions{})
		if err != nil {
			klog.Warningf("Cannot update Result Secret: %s", err)
			return err
		}
		return nil
	} else {
		return new(SkipResultHandler)
	}
}

// pushResult pushes the encrypted data to the callback URL if the callbackUrl is set.
func pushResult(encryptedData []byte, retryOnFailure bool) error {
	if len(callbackUrl) > 0 {
		callBackUrl, err := base64.StdEncoding.DecodeString(callbackUrl)
		if err != nil {
			klog.Warningf("Cannot decode Callback URL %s", err)
			return err
		}
		url := string(callBackUrl)
		response, httpError := http.Post(url, "application/json", bytes.NewReader(encryptedData))
		if httpError != nil {
			klog.Warningf("Cannot finish Callback request: %s", httpError)
			if retryOnFailure {
				klog.Infof("Retry flag was set on Callback request, retrying until successful...")
				err = retryPushResultInfinitely(encryptedData)
				return err
			} else {
				return httpError
			}
		}
		if response.StatusCode != 204 {
			klog.Warningf("Got NOK HTTP Status Code %s, skipping", response.Status)
			return fmt.Errorf("got NOK response [%s] for HTTP result handling", response.Status)
		}
		return nil
	} else {
		return new(SkipResultHandler)
	}
}

// retryPushResultInfinitely keeps retrying to push encrypted data to the callback URL.
func retryPushResultInfinitely(encryptedData []byte) error {
	retryDelay := 1 * time.Second
	maxBackoff := 3 * time.Minute
	backoffFactor := 2.0

	for {
		clientset, err := k8s.GetClientset()
		if err != nil {
			klog.Warningf("Cannot create clientset for handling Result Secret: %s", err)
			return err
		}
		// we need to re-fetch callback URL since it will change when remote-runner restarts
		secret, err := clientset.CoreV1().Secrets(os.Getenv(RunnerNamespace)).Get(context.Background(), os.Getenv(InputContextSecret), v1.GetOptions{})
		if err != nil {
			klog.Warningf("Cannot fetch Result Secret: %s", err)
			return err
		}
		callbackUrl = string(secret.Data[InputContextSecretDataUrlKey])
		decodedCallbackUrl, err := base64.StdEncoding.DecodeString(callbackUrl)
		if err != nil {
			klog.Warningf("Cannot decode Callback URL %s", err)
			return err
		}
		url := string(decodedCallbackUrl)

		_, httpError := http.Post(url, "application/json", bytes.NewReader(encryptedData))
		if httpError == nil {
			return nil
		}

		klog.Warningf("Cannot finish retried Callback request: %s. Retrying in %v...", httpError, retryDelay)
		time.Sleep(retryDelay)

		retryDelay = time.Duration(math.Min(float64(retryDelay)*backoffFactor, float64(maxBackoff)))
	}
}

// writeOutput writes the encrypted data to the output file if the outputLocation is set.
func writeOutput(encryptedData []byte) error {
	outputLocation := os.Getenv(OutputLocation)
	if len(outputLocation) > 0 {
		err := os.WriteFile(outputLocation, encryptedData, 0644)
		if err != nil {
			klog.Warningf("Cannot write output to: %s [%v]", outputLocation, err)
			return err
		}
		return nil
	} else {
		return new(SkipResultHandler)
	}
}

// shouldRetryCallbackRequest checks if callback request should be retried on failure.
// it should be retried when result is too big for Secret and Output File handler is not used
func shouldRetryCallbackRequest(encryptedData []byte) bool {
	return len(encryptedData) >= SizeOf1Mb && len(os.Getenv(OutputLocation)) == 0
}
