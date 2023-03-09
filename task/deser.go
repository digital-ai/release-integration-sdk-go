package task

import (
	"encoding/json"
	"io"
	"k8s.io/klog/v2"
	"os"
)

const (
	InputCategory  = "input"
	OutputCategory = "output"
	InputLocation  = "INPUT_LOCATION"
	OutputLocation = "OUTPUT_LOCATION"
)

func Deserialize(context *InputContext) error {
	var inputLocation = os.Getenv(InputLocation)
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

	content, err := io.ReadAll(inputContent)
	if err != nil {
		return err
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

func handleResult(outputContext TaskOutputContext) {
	data, _ := json.Marshal(outputContext)
	encryptedData, encryptErr := Encrypt(data)
	if encryptErr != nil {
		klog.Fatalf("error encrypting data [%v]", encryptErr)
	}
	writeOutput(encryptedData)
}

func writeOutput(encryptedData []byte) {
	outputLocation := os.Getenv(OutputLocation)
	err := os.WriteFile(outputLocation, encryptedData, 0644)
	if err != nil {
		klog.Fatalf("Cannot write output to: %s [%v]", outputLocation, err)
	}
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
