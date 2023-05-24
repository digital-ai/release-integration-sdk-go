package task

import (
	ctx "context"
	"encoding/json"
	"github.com/digital-ai/release-integration-sdk-go/k8s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"os"
)

const (
	InputCategory                    = "input"
	OutputCategory                   = "output"
	InputLocation                    = "INPUT_LOCATION"
	OutputLocation                   = "OUTPUT_LOCATION"
	CallbackURL                      = "CALLBACK_URL"
	ResultSecretName                 = "RESULT_SECRET_NAME"
	ReleaseURL                       = "RELEASE_URL"
	InputContextSecret               = "INPUT_CONTEXT_SECRET"
	RunnerNamespace                  = "RUNNER_NAMESPACE"
	InputContextSecretDataInput      = "input"
	InputContextSecretDataSessionKey = "session-key"
	InputContextSecretDataUrlKey     = "url"
)

func Deserialize(context *InputContext) error {
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
	content := secret.Data[InputContextSecretDataInput]
	SessionKey = string(secret.Data[InputContextSecretDataSessionKey])
	CallbackUrl = string(secret.Data[InputContextSecretDataUrlKey])

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
