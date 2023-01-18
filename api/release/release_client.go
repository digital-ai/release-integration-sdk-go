package release

import (
	"context"
	"encoding/base64"
	"github.com/xebialabs/go-remote-runner-wrapper/api/release/openapi"
	_ "github.com/xebialabs/go-remote-runner-wrapper/http"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"net/http"
)

type ReleaseClient struct {
	client *openapi.APIClient
	ReleaseApi
}

// TODO: remove if not needed
type ReleaseApi interface {
	GetVariablesForRelease(string) ([]openapi.Variable, *http.Response, error)
	CreateVariablesForRelease(string, openapi.Variable1) (*openapi.Variable, *http.Response, error)
	UpdateVariablesForRelease(string, []openapi.Variable) ([]openapi.Variable, *http.Response, error)
	GetVariable(string) (*openapi.Variable, *http.Response, error)
	UpdateVariable(string, openapi.Variable) (*openapi.Variable, *http.Response, error)
	DeleteVariable(string) (*http.Response, error)
	GetVariableValuesForRelease(string) (map[string]string, *http.Response, error)
	GetVariablePossibleValues(string) ([]map[string]interface{}, *http.Response, error)
	IsVariableUsed(string) (bool, *http.Response, error)
	ReplaceVariable(string, openapi.VariableOrValue) (*http.Response, error)
}

func (r *ReleaseClient) GetVariablesForRelease(releaseId string) ([]openapi.Variable, *http.Response, error) {
	return r.client.ReleaseApiApi.GetVariablesForRelease(context.Background(), releaseId).Execute()
}

func (r *ReleaseClient) CreateVariablesForRelease(releaseId string, variable openapi.Variable1) (*openapi.Variable, *http.Response, error) {
	return r.client.ReleaseApiApi.CreateVariablesForRelease(context.Background(), releaseId).
		Variable1(variable).
		Execute()
}

func (r *ReleaseClient) UpdateVariablesForRelease(releaseId string, variables []openapi.Variable) ([]openapi.Variable, *http.Response, error) {
	return r.client.ReleaseApiApi.UpdateVariablesForRelease(context.Background(), releaseId).
		Variable(variables).
		Execute()
}

func (r *ReleaseClient) GetVariable(variableId string) (*openapi.Variable, *http.Response, error) {
	return r.client.ReleaseApiApi.GetVariable(context.Background(), variableId).Execute()
}

func (r *ReleaseClient) UpdateVariable(variableId string, variable openapi.Variable) (*openapi.Variable, *http.Response, error) {
	return r.client.ReleaseApiApi.UpdateVariable(context.Background(), variableId).
		Variable(variable).
		Execute()
}

func (r *ReleaseClient) DeleteVariable(variableId string) (*http.Response, error) {
	return r.client.ReleaseApiApi.DeleteVariable(context.Background(), variableId).Execute()
}

func (r *ReleaseClient) GetVariableValuesForRelease(releaseId string) (map[string]string, *http.Response, error) {
	return r.client.ReleaseApiApi.GetVariableValuesForRelease(context.Background(), releaseId).Execute()
}

func (r *ReleaseClient) GetVariablePossibleValues(variableId string) ([]map[string]interface{}, *http.Response, error) {
	return r.client.ReleaseApiApi.GetVariablePossibleValues(context.Background(), variableId).Execute()
}

func (r *ReleaseClient) IsVariableUsed(variableId string) (bool, *http.Response, error) {
	return r.client.ReleaseApiApi.IsVariableUsed(context.Background(), variableId).Execute()
}

func (r *ReleaseClient) ReplaceVariable(variableId string, variableOrValue openapi.VariableOrValue) (*http.Response, error) {
	return r.client.ReleaseApiApi.ReplaceVariable(context.Background(), variableId).
		VariableOrValue(variableOrValue).
		Execute()
}

func NewReleaseClient(ctx task.ReleaseContext) *ReleaseClient {
	client := NewReleaseApiClient(ctx)
	return &ReleaseClient{client: client}
}

func NewReleaseApiClient(ctx task.ReleaseContext) *openapi.APIClient {
	conf := openapi.NewConfiguration()
	conf.DefaultHeader = map[string]string{
		"Authorization": "Basic " + basicAuth(ctx.AutomatedTaskAsUser.Username, ctx.AutomatedTaskAsUser.Password),
		"Content-Type":  "application/json",
	}
	conf.Host = ctx.Url
	conf.Scheme = "http"
	return openapi.NewAPIClient(conf)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
