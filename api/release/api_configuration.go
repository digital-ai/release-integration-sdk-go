package release

import (
	"encoding/json"
	"github.com/xebialabs/go-remote-runner-wrapper/http"
)

const (
	configurationGlobalVariablesPath      = "api/v1/config/Configuration/variables/global"
	configurationGlobalVariableValuesPath = "/api/v1/config/Configuration/variableValues/global"
)

func AddGlobalVariable(client *http.HttpClient, variable *Variable1) (*Variable, error) {
	body, marshalErr := json.Marshal(variable)
	if marshalErr != nil {
		return nil, marshalErr
	}

	result, err := client.Post(configurationGlobalVariablesPath, body)
	if err != nil {
		return nil, err
	}
	var createdVar *Variable
	err = json.Unmarshal(result, &createdVar)
	if err != nil {
		return nil, err
	}
	return createdVar, nil
}

func GetGlobalVariables(client *http.HttpClient) ([]Variable, error) {
	result, err := client.Get(configurationGlobalVariablesPath)
	if err != nil {
		return nil, err
	}
	var variables []Variable
	err = json.Unmarshal(result, &variables)
	if err != nil {
		return nil, err
	}
	return variables, nil
}

func GetGlobalVariableValues(client *http.HttpClient) (map[string]string, error) {
	result, err := client.Get(configurationGlobalVariableValuesPath)
	if err != nil {
		return nil, err
	}
	var variableValues map[string]string
	err = json.Unmarshal(result, &variableValues)
	if err != nil {
		return nil, err
	}
	return variableValues, nil
}
