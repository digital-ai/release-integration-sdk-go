package release_api

import (
	"encoding/json"
	"net/url"
	"path"
)

func (req ReleasesReleaseIdVariablesGetRequest) ReleasesReleaseIdVariablesGet() ([]Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(req.releaseId), "/variables")

	result, err := req.client.Get(path)
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

func (req ReleasesReleaseIdVariablesPostRequest) ReleasesReleaseIdVariablesPost() (*Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(req.releaseId), "/variables")

	body, marshalErr := json.Marshal(req.variable)
	if marshalErr != nil {
		return nil, marshalErr
	}

	result, err := req.client.Post(path, body)
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

func (req ReleasesReleaseIdVariablesPutRequest) ReleasesReleaseIdVariablesPut() (*Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(req.releaseId), "/variables")

	body, marshalErr := json.Marshal(req.variable)
	if marshalErr != nil {
		return nil, marshalErr
	}

	result, err := req.client.Put(path, body)
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

func (req ReleasesVariableIdGetRequest) ReleasesVariableIdGet() (*Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(req.variableId))

	result, err := req.client.Get(path)
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

func (req ReleasesVariableIdPutRequest) ReleasesVariableIdPut() (*Variable, error) {
	path := path.Join("api/v1/releases/", url.PathEscape(req.variableId))

	body, marshalErr := json.Marshal(req.variable)
	if marshalErr != nil {
		return nil, marshalErr
	}

	result, err := req.client.Put(path, body)
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

func (req ReleasesVariableIdDeleteRequest) ReleasesVariableIdDelete() error {
	path := path.Join("api/v1/releases/", url.PathEscape(req.variableId))
	_, err := req.client.Delete(path)
	return err
}
