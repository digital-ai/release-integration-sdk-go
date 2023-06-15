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
	"net/http"
	"os"
	"strings"
)

var callbackUrl = os.Getenv(CallbackURL)
var resultSecretKey = os.Getenv(ResultSecretName)

func HandleSuccess(result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         0,
		OutputProperties: result,
	}
	handleResult(outputContext)
}

func HandleError(err error, result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         -1,
		OutputProperties: result,
		JobErrorMessage:  err.Error(),
	}
	handleResult(outputContext)
}

type SkipResultHandler struct{}

func (m *SkipResultHandler) Error() string {
	return "skipping result handler"
}

func handleResult(outputContext TaskOutputContext) {
	data, _ := json.Marshal(outputContext)
	encryptedData, encryptErr := Encrypt(data)
	if encryptErr != nil {
		klog.Fatalf("error encrypting data [%v]", encryptErr)
	}

	done := make(chan string, 3)
	success := make(chan bool)
	go func() {
		err := pushResult(encryptedData)
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

func splitSecretResourceData(secretEntry string) (string, string, string, error) {
	split := strings.Split(secretEntry, ":")
	if len(split) != 3 {
		return "", "", "", nil
	}
	return split[0], split[1], split[2], nil
}

func writeToSecret(encryptedData []byte) error {
	if len(resultSecretKey) > 0 {
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

func pushResult(encryptedData []byte) error {
	if len(callbackUrl) > 0 {
		callBackUrl, err := base64.StdEncoding.DecodeString(callbackUrl)
		if err != nil {
			klog.Warningf("Cannot decode Callback URL %s", err)
			return err
		}
		url := string(callBackUrl)
		// TODO retry schema maybe?
		response, httpError := http.Post(url, "application/json", bytes.NewReader(encryptedData))
		if httpError != nil {
			klog.Warningf("Cannot finish Callback request: %s, skipping", httpError)
			return httpError
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
