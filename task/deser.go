package task

import (
	"encoding/json"
	"fmt"
	"io"
	"k8s.io/klog"
	"os"
	"strconv"
)

const (
	success      = "success"
	errorMessage = "errorMessage"
)

func Deserialize(inputLocation string) (map[string]json.RawMessage, error) {
	inputContent, err := os.Open(inputLocation)
	// if we os.Open returns an error then handle it
	if err != nil {
		klog.Errorf("Cannot open: %s [%v]", inputLocation, err)
		return nil, err
	}
	// defer the closing of our inputContent so that we can parse it later on
	defer inputContent.Close()

	content, _ := io.ReadAll(inputContent)
	byteValue, _ := Decrypt(content)

	var inputContext InputContext
	unMarshalErr := json.Unmarshal(byteValue, &inputContext)
	if unMarshalErr != nil {
		klog.Errorf("Cannot umarshal input: %v", unMarshalErr)
		return nil, unMarshalErr
	}

	propertiesMap := make(map[string]json.RawMessage)
	for _, data := range inputContext.Task.Properties {
		propertiesMap[data.Name] = data.Value
	}
	return propertiesMap, nil
}

func Serialize(outputLocation string, result map[string]interface{}) {
	outputContext := TaskOutputContext{
		ExitCode:         0,
		OutputProperties: result,
	}
	writeOutput(outputContext, outputLocation)
}

func ErrorResultMap(err error) map[string]interface{} {
	resultMap := make(map[string]interface{})
	resultMap[success] = strconv.FormatBool(false)
	resultMap[errorMessage] = err.Error()
	return resultMap
}

func SerializeError(outputLocation string, err error) {
	outputContext := TaskOutputContext{
		ExitCode:         -1,
		OutputProperties: ErrorResultMap(err),
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

func DeserializeProperties(task json.RawMessage) ([]PropertyDefinition, error) {
	if !json.Valid(task) {
		return nil, fmt.Errorf("cannot deserialize unvalid json or null value")
	}
	var taskContext TaskContext
	unMarshalErr := json.Unmarshal(task, &taskContext)
	if unMarshalErr != nil {
		klog.Errorf("Cannot umarshal properties: %v", unMarshalErr)
		return nil, fmt.Errorf("failed to unmarshal properties: %w", unMarshalErr)
	}
	return taskContext.Properties, nil
}
