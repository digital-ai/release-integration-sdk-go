package task

import (
	ctx "context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/k8s"
	"github.com/pkg/errors"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"net/http"
	"os"
)

// InputCategory represents the category of input properties.
const InputCategory = "input"

// OutputCategory represents the category of output properties.
const OutputCategory = "output"

// InputLocation represents the environment variable for input location.
const InputLocation = "INPUT_LOCATION"

// OutputLocation represents the environment variable for output location.
const OutputLocation = "OUTPUT_LOCATION"

// ResultSecretName represents the name of the result secret.
const ResultSecretName = "RESULT_SECRET_NAME"

// CallbackURL represents the environment variable for the callback URL.
const CallbackURL = "CALLBACK_URL"

// ReleaseURL represents the environment variable for the release URL.
const ReleaseURL = "RELEASE_URL"

// InputContextSecret represents the environment variable for the input context secret.
const InputContextSecret = "INPUT_CONTEXT_SECRET"

// RunnerNamespace represents the environment variable for the runner namespace.
const RunnerNamespace = "RUNNER_NAMESPACE"

// SessionKey represents the environment variable for the session key.
const SessionKey = "SESSION_KEY"

// InputContextSecretDataInput represents the key for input data in the input context secret.
const InputContextSecretDataInput = "input"

// InputContextSecretDataSessionKey represents the key for session key in the input context secret.
const InputContextSecretDataSessionKey = "session-key"

// InputContextSecretDataUrlKey represents the key for URL in the input context secret.
const InputContextSecretDataUrlKey = "url"

// InputContextSecretDataFetchUrlKey represents the key for Fetching task content URL in the input context secret. This
const InputContextSecretDataFetchUrlKey = "fetchUrl"

// InputContextSecretExecutionIdKey represents the key for execution ID in the input context secret.
const InputContextSecretExecutionIdKey = "execution-id"

// Deserialize deserializes the input context from the specified source into the provided InputContext object.
func Deserialize(context *InputContext) error {
	context.Release.Url = os.Getenv(ReleaseURL)
	inputLocation := os.Getenv(InputLocation)

	var content []byte
	if len(inputLocation) > 0 {
		inputContent, err := os.Open(inputLocation)
		// if we os.Open returns an error then handle it
		if err != nil {
			klog.Errorf("Cannot open: '%s' [%v]", inputLocation, err)
			return err
		}
		// defer the closing of our inputContent so that we can parse it later on
		defer func(inputContent *os.File) {
			if deferredErr := inputContent.Close(); deferredErr != nil {
				err = deferredErr
			}
		}(inputContent)

		content, err = io.ReadAll(inputContent)
		if err != nil {
			return err
		}
	} else {
		// reading input context from secret
		clientset, err := k8s.GetClientset()
		if err != nil {
			klog.Warningf("Cannot create clientset for handling Result Secret: %s", err)
			return err
		}
		secret, err := clientset.CoreV1().Secrets(os.Getenv(RunnerNamespace)).Get(ctx.Background(), os.Getenv(InputContextSecret), v1.GetOptions{})
		if err != nil {
			klog.Warningf("Cannot fetch Result Secret: %s", err)
			return err
		}

		var ok bool
		content, ok = secret.Data[InputContextSecretDataInput]
		if !ok || len(content) == 0 {
			fetchUrlBase64, fetchUrlOk := secret.Data[InputContextSecretDataFetchUrlKey]
			if !fetchUrlOk || len(fetchUrlBase64) == 0 {
				return errors.Errorf("cannot find fetch url for task")
			}
			fetchUrlBytes, err := base64.StdEncoding.DecodeString(string(fetchUrlBase64))
			if err != nil {
				klog.Warningf("Cannot finish data fetch request: %s, skipping", err)
				return errors.Errorf("cannot decode fetch url %s", err)
			}
			response, httpError := http.Get(string(fetchUrlBytes))
			if httpError != nil {
				klog.Warningf("Cannot finish fetch request: '%s', returning failure of task execution", httpError)
				return httpError
			}
			defer response.Body.Close()

			if response.StatusCode != http.StatusOK {
				return fmt.Errorf("failed to fetch data, server returned status: %s", response.Status)
			}

			// Read the response body into a byte slice
			content, err = io.ReadAll(response.Body)
			if err != nil {
				return err
			}

		}
		sessionKey = string(secret.Data[InputContextSecretDataSessionKey])
		callbackUrl = string(secret.Data[InputContextSecretDataUrlKey])
	}

	decrypted, err := Decrypt(content)
	if err != nil {
		return err
	}

	unMarshalErr := json.Unmarshal(decrypted, context)
	if unMarshalErr != nil {
		klog.Errorf("Cannot unmarshal input: %v", unMarshalErr)
		return unMarshalErr
	}

	return err
}

// DeserializeTask deserializes the input properties into the provided task instance.
func DeserializeTask(properties []PropertyDefinition, taskInstance any) error {
	var inputs []PropertyDefinition
	for _, property := range properties {
		if property.Category == InputCategory {
			inputs = append(inputs, property)
		}
	}

	return UnmarshalProperties(inputs, taskInstance)
}

// Serialize serializes the provided result into the output properties.
func Serialize(result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         0,
		OutputProperties: result,
	}
	handleResult(outputContext)
}

// SerializeError serializes the provided error and result into the output properties.
func SerializeError(err error, result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         -1,
		OutputProperties: result,
		JobErrorMessage:  err.Error(),
	}
	handleResult(outputContext)
}

// UnmarshalProperties unmarshalls the input properties into the provided prototype object.
func UnmarshalProperties(properties []PropertyDefinition, prototype interface{}) error {
	propsMap := make(map[string]json.RawMessage)
	for _, property := range properties {
		propsMap[property.Name] = property.Value
	}
	jsonMap, err := json.Marshal(propsMap)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonMap, prototype); err != nil {
		return err
	}
	return nil
}
