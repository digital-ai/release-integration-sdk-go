package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Deserialize(inputLocation string) map[string]json.RawMessage {
	inputContent, err := os.Open(inputLocation)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("Cannot open: ", inputLocation)
		fmt.Println(err)
	}
	// defer the closing of our inputContent so that we can parse it later on
	defer inputContent.Close()

	content, _ := ioutil.ReadAll(inputContent)
	byteValue, _ := Decrypt(content)

	var inputContext InputContext
	unMarshalErr := json.Unmarshal(byteValue, &inputContext)
	if unMarshalErr != nil {
		fmt.Println(unMarshalErr)
	}

	propertiesMap := make(map[string]json.RawMessage)
	for _, data := range inputContext.Task.Properties {
		propertiesMap[data.Name] = data.Value
	}
	return propertiesMap
}

func Serialize(outputLocation string, result TaskResult) {
	outputContext := TaskOutputContext{
		ExitCode:         0,
		OutputProperties: result,
	}

	data, _ := json.Marshal(outputContext)
	encryptedData := Encrypt(data)
	err := ioutil.WriteFile(outputLocation, encryptedData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func DeserializeProperties(task json.RawMessage) ([]PropertyDefinition, error) {
	var taskContext TaskContext
	unMarshalErr := json.Unmarshal(task, &taskContext)
	if unMarshalErr != nil {
		fmt.Println(unMarshalErr)
		return nil, unMarshalErr
	}
	return taskContext.Properties, nil
}
