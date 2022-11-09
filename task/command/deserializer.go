package command

import (
	"encoding/json"
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"strconv"
)

type CommandType string

type CommandWrapper struct {
	CommandType CommandType `json:"type"`
	Properties  interface{}
}

type CommandFactory interface {
	Create(commandType CommandType, properties []task.PropertyDefinition) (CommandExecutor, error)
}

func DeserializeCommand(factory CommandFactory, propertiesMap map[string]json.RawMessage) (CommandExecutor, error) {
	commandRawJson := propertiesMap["command"]
	var raw json.RawMessage
	wrapper := CommandWrapper{
		Properties: &raw,
	}

	if err := json.Unmarshal(commandRawJson, &wrapper); err != nil {
		fmt.Println("Wrapper unmarshal error ", err.Error())
		return nil, err
	}

	var properties []task.PropertyDefinition
	if err := json.Unmarshal(raw, &properties); err != nil {
		fmt.Println("Command unmarshal error ", err.Error())
		return nil, err
	}

	commandExec, err := factory.Create(wrapper.CommandType, properties)
	if err != nil {
		return nil, err
	}

	return commandExec, nil
}

func ExtractStringValue(property task.PropertyDefinition, field string) (*string, error) {
	if property.Name == field && property.Value != nil {
		unquoted, err := strconv.Unquote(string(property.Value))
		if err != nil {
			fmt.Printf("Error unmarshaling string value for %s with value [%s], skipping value.\n", field, property.Value)
			return nil, nil
		}
		return &unquoted, nil
	} else {
		return nil, nil
	}
}

func ExtractBoolValue(property task.PropertyDefinition, field string) (*bool, error) {
	if property.Name == field && property.Value != nil {
		boolVal, err := strconv.ParseBool(string(property.Value))
		if err != nil {
			fmt.Printf("Error unmarshaling bool value for %s with value [%s], skipping value.\n", field, property.Value)
			return nil, nil
		}
		return &boolVal, nil
	} else {
		return nil, nil
	}
}

func ExtractPropertiesStringValue(properties []task.PropertyDefinition, field string) (*string, error) {
	for _, property := range properties {
		if val, err := ExtractStringValue(property, field); val != nil {
			return val, nil
		} else if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func ExtractPropertiesStringArray(properties []task.PropertyDefinition, field string) (*[]string, error) {
	for _, property := range properties {
		if property.Name != field {
			continue
		}
		var array []string
		if err := json.Unmarshal(property.Value, &array); err != nil {
			fmt.Printf("Error unmarshaling string array value for %s with value [%s], skipping value.\n", field, property.Value)
			return nil, err
		}
		return &array, nil
	}
	return nil, nil
}

func ExtractPropertiesMapStringString(properties []task.PropertyDefinition, field string) (*map[string]string, error) {
	for _, property := range properties {
		if property.Name != field {
			continue
		}
		var keyValues map[string]string
		if err := json.Unmarshal(property.Value, &keyValues); err != nil {
			fmt.Printf("Error unmarshaling map[string]string value for %s with value [%s], skipping value.\n", field, property.Value)
			return nil, err
		}
		return &keyValues, nil
	}
	return nil, nil
}

func ExtractIntValue(property task.PropertyDefinition, field string) (*int64, error) {
	if property.Name == field && property.Value != nil {
		intVal, err := strconv.ParseInt(string(property.Value), 10, 0)
		if err != nil {
			fmt.Printf("Error unmarshaling int value for %s with value [%s], skipping value.\n", field, property.Value)
			return nil, nil
		}
		return &intVal, nil
	} else {
		return nil, nil
	}
}
