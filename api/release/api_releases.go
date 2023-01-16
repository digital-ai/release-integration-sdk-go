package release

import (
	"encoding/json"
	"github.com/xebialabs/go-remote-runner-wrapper/http"
	"net/url"
	"path"
)

func GetVariablesForRelease(client *http.HttpClient, releaseId string) ([]Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(releaseId), "/variables")

	result, err := client.Get(path)
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

func CreateVariableForRelease(client *http.HttpClient, releaseId string, variable *Variable1) (*Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(releaseId), "/variables")

	body, marshalErr := json.Marshal(variable)
	if marshalErr != nil {
		return nil, marshalErr
	}

	result, err := client.Post(path, body)
	if err != nil {
		return nil, err
	}
	var created *Variable
	err = json.Unmarshal(result, &created)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func UpdateVariablesForRelease(client *http.HttpClient, releaseId string, variables *[]Variable) ([]Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(releaseId), "/variables")

	body, marshalErr := json.Marshal(variables)
	if marshalErr != nil {
		return nil, marshalErr
	}

	result, err := client.Put(path, body)
	if err != nil {
		return nil, err
	}
	var created []Variable
	err = json.Unmarshal(result, &created)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func GetVariable(client *http.HttpClient, variableId string) (*Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(variableId))

	result, err := client.Get(path)
	if err != nil {
		return nil, err
	}
	var variable *Variable
	err = json.Unmarshal(result, &variable)
	if err != nil {
		return nil, err
	}
	return variable, nil
}

func UpdateVariable(client *http.HttpClient, variableId string, variable *Variable) (*Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(variableId))

	body, marshalErr := json.Marshal(variable)
	if marshalErr != nil {
		return nil, marshalErr
	}

	result, err := client.Put(path, body)
	if err != nil {
		return nil, err
	}
	var updatedVar *Variable
	err = json.Unmarshal(result, &updatedVar)
	if err != nil {
		return nil, err
	}
	return variable, nil
}

func DeleteVariable(client *http.HttpClient, variableId string) error {
	path := path.Join("api/v1/releases/", url.PathEscape(variableId))
	_, err := client.Delete(path)
	return err
}
