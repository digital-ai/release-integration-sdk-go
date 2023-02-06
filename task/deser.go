package task

import (
	"encoding/json"
	"io"
	"k8s.io/klog/v2"
	"os"
)

const (
	InputCategory = "input"
)

func Deserialize(inputLocation string, context *InputContext) error {
	inputContent, err := os.Open(inputLocation)
	// if we os.Open returns an error then handle it
	if err != nil {
		klog.Errorf("Cannot open: '%s' [%v]", inputLocation, err)
		return err
	}
	// defer the closing of our inputContent so that we can parse it later on
	defer inputContent.Close()

	content, _ := io.ReadAll(inputContent)
	byteValue, _ := Decrypt(content)

	unMarshalErr := json.Unmarshal(byteValue, context)
	if unMarshalErr != nil {
		klog.Errorf("Cannot umarshal input: %v", unMarshalErr)
		return unMarshalErr
	}

	return nil
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

func Serialize(outputLocation string, result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         0,
		OutputProperties: result,
	}
	writeOutput(outputContext, outputLocation)
}

func SerializeError(outputLocation string, result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         -1,
		OutputProperties: result,
	}
	writeOutput(outputContext, outputLocation)
}

func writeOutput(outputContext TaskOutputContext, outputLocation string) {
	data, _ := json.Marshal(outputContext)
	encryptedData, encryptErr := Encrypt(data)
	if encryptErr != nil {
		klog.Fatalf("Cannot write output to: %s error encrypting data [%v]", outputLocation, encryptErr)
	}
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
