package release

import (
	"github.com/xebialabs/go-remote-runner-wrapper/http"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

type ReleaseClient struct {
	client *http.HttpClient
	ReleaseApi
}

type ReleaseApi interface {
	GetVariablesForRelease(string) ([]Variable, error)
	CreateVariablesForRelease(string, *Variable1) (*Variable, error)
	UpdateVariablesForRelease(string, *[]Variable) ([]Variable, error)
	GetVariable(string) (*Variable, error)
	UpdateVariable(string, *Variable) (*Variable, error)
	DeleteVariable(string) error
}

func NewReleaseClient(ctx task.ReleaseContext) ReleaseApi {
	client, _ := HttpClient(ctx)
	return &ReleaseClient{client: client}
}

func HttpClient(ctx task.ReleaseContext) (*http.HttpClient, error) {
	auth := ctx.AutomatedTaskAsUser
	builder := http.NewHttpClientBuilder().
		ForHost(ctx.Url).
		WithBasicAuth(auth.Username, auth.Password)
	return builder.Build()
}

func (r *ReleaseClient) GetVariablesForRelease(releaseId string) ([]Variable, error) {
	return GetVariablesForRelease(r.client, releaseId)
}

func (r *ReleaseClient) CreateVariablesForRelease(releaseId string, variable *Variable1) (*Variable, error) {
	return CreateVariableForRelease(r.client, releaseId, variable)
}

func (r *ReleaseClient) UpdateVariablesForRelease(releaseId string, variables *[]Variable) ([]Variable, error) {
	return UpdateVariablesForRelease(r.client, releaseId, variables)
}

func (r *ReleaseClient) GetVariable(variableId string) (*Variable, error) {
	return GetVariable(r.client, variableId)
}

func (r *ReleaseClient) UpdateVariable(variableId string, variable *Variable) (*Variable, error) {
	return UpdateVariable(r.client, variableId, variable)
}

func (r *ReleaseClient) DeleteVariable(variableId string) error {
	return DeleteVariable(r.client, variableId)
}
