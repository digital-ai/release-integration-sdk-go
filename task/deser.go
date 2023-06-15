package task

import (
	ctx "context"
	"encoding/json"
	"github.com/digital-ai/release-integration-sdk-go/k8s"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"os"
)

const (
	InputCategory                    = "input"
	OutputCategory                   = "output"
	InputLocation                    = "INPUT_LOCATION"
	OutputLocation                   = "OUTPUT_LOCATION"
	ResultSecretName                 = "RESULT_SECRET_NAME"
	CallbackURL                      = "CALLBACK_URL"
	ReleaseURL                       = "RELEASE_URL"
	InputContextSecret               = "INPUT_CONTEXT_SECRET"
	RunnerNamespace                  = "RUNNER_NAMESPACE"
	SessionKey                       = "SESSION_KEY"
	InputContextSecretDataInput      = "input"
	InputContextSecretDataSessionKey = "session-key"
	InputContextSecretDataUrlKey     = "url"
	InputContextSecretExecutionIdKey = "execution-id"
)

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
		content = secret.Data[InputContextSecretDataInput]
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

func DeserializeTask(properties []PropertyDefinition, taskInstance any) error {
	var inputs []PropertyDefinition
	for _, property := range properties {
		if property.Category == InputCategory {
			inputs = append(inputs, property)
		}
	}

	return UnmarshalProperties(inputs, taskInstance)
}

func Serialize(result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         0,
		OutputProperties: result,
	}
	handleResult(outputContext)
}

func SerializeError(err error, result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         -1,
		OutputProperties: result,
		JobErrorMessage:  err.Error(),
	}
	handleResult(outputContext)
}

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
