package openapi

import (
	"context"
	"encoding/base64"
	_ "github.com/xebialabs/go-remote-runner-wrapper/http"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"net/http"
)

type ReleaseClient struct {
	client *APIClient
	ReleaseApi
}

type ReleaseApi interface {
	GetVariablesForRelease(string) ([]Variable, *http.Response, error)
	CreateVariablesForRelease(string, Variable1) (*Variable, *http.Response, error)
	UpdateVariablesForRelease(string, []Variable) ([]Variable, *http.Response, error)
	GetVariable(string) (*Variable, *http.Response, error)
	UpdateVariable(string, Variable) (*Variable, *http.Response, error)
	DeleteVariable(string) (*http.Response, error)
	GetVariableValuesForRelease(string) (map[string]string, *http.Response, error)
	GetVariablePossibleValues(string) ([]map[string]interface{}, *http.Response, error)
	IsVariableUsed(string) (bool, *http.Response, error)
	ReplaceVariable(string, VariableOrValue) (*http.Response, error)
}

func (r *ReleaseClient) GetVariablesForRelease(releaseId string) ([]Variable, *http.Response, error) {
	return r.client.ReleasesApi.ReleasesReleaseIdVariablesGet(context.Background(), releaseId).Execute()
}

func (r *ReleaseClient) CreateVariablesForRelease(releaseId string, variable Variable1) (*Variable, *http.Response, error) {
	return r.client.ReleasesApi.ReleasesReleaseIdVariablesPost(context.Background(), releaseId).
		Variable1(variable).
		Execute()
}

func (r *ReleaseClient) UpdateVariablesForRelease(releaseId string, variables []Variable) ([]Variable, *http.Response, error) {
	return r.client.ReleasesApi.ReleasesReleaseIdVariablesPut(context.Background(), releaseId).
		Variable(variables).
		Execute()
}

func (r *ReleaseClient) GetVariable(variableId string) (*Variable, *http.Response, error) {
	return r.client.ReleasesApi.ReleasesVariableIdGet(context.Background(), variableId).Execute()
}

func (r *ReleaseClient) UpdateVariable(variableId string, variable Variable) (*Variable, *http.Response, error) {
	return r.client.ReleasesApi.ReleasesVariableIdPut(context.Background(), variableId).
		Variable(variable).
		Execute()
}

func (r *ReleaseClient) DeleteVariable(variableId string) (*http.Response, error) {
	return r.client.ReleasesApi.ReleasesVariableIdDelete(context.Background(), variableId).Execute()
}

func (r *ReleaseClient) GetVariableValuesForRelease(releaseId string) (map[string]string, *http.Response, error) {
	return r.client.ReleasesApi.ReleasesReleaseIdVariableValuesGet(context.Background(), releaseId).Execute()
}

func (r *ReleaseClient) GetVariablePossibleValues(variableId string) ([]map[string]interface{}, *http.Response, error) {
	return r.client.ReleaseApiApi.GetVariablePossibleValues(context.Background(), variableId).Execute()
}

func (r *ReleaseClient) IsVariableUsed(variableId string) (bool, *http.Response, error) {
	return r.client.ReleaseApiApi.IsVariableUsed(context.Background(), variableId).Execute()
}

func (r *ReleaseClient) ReplaceVariable(variableId string, variableOrValue VariableOrValue) (*http.Response, error) {
	return r.client.ReleaseApiApi.ReplaceVariable(context.Background(), variableId).
		VariableOrValue(variableOrValue).
		Execute()
}

func NewReleaseClient(ctx task.ReleaseContext) ReleaseApi {
	client := NewReleaseApiClient(ctx)
	return &ReleaseClient{client: client}
}

func NewReleaseApiClient(ctx task.ReleaseContext) *APIClient {
	conf := NewConfiguration()
	conf.DefaultHeader = map[string]string{
		"Authorization": "Basic " + basicAuth(ctx.AutomatedTaskAsUser.Username, ctx.AutomatedTaskAsUser.Password),
		"Content-Type":  "application/json",
	}
	conf.Host = ctx.Url
	conf.Scheme = "http"
	releaseApi := NewAPIClient(conf)
	return releaseApi
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
