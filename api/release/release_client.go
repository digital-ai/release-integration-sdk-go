package release

import (
	"github.com/xebialabs/go-remote-runner-wrapper/http"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

type ReleaseClient struct {
	ctx    task.ReleaseContext
	client *http.HttpClient
	ReleaseApi
}

type ReleaseApi interface {
	ReleasesReleaseIdVariablesGet(string) ([]Variable, error)
	ReleasesReleaseIdVariablesPost(string, *Variable1) (*Variable, error)
	ReleasesReleaseIdVariablesPut(string, *Variable1) (*Variable, error)
	ReleasesVariableIdGet(string) (*Variable, error)
	ReleasesVariableIdPut(string, *Variable) (*Variable, error)
	ReleasesVariableIdDelete(string) error
}

func NewReleaseClient(ctx task.ReleaseContext) ReleaseApi {
	client, _ := HttpClient(ctx)
	return &ReleaseClient{ctx: ctx, client: client}
}

func HttpClient(ctx task.ReleaseContext) (*http.HttpClient, error) {
	auth := ctx.AutomatedTaskAsUser
	builder := http.NewHttpClientBuilder().
		ForHost(ctx.Url).
		WithBasicAuth(auth.Username, auth.Password)
	return builder.Build()
}

func (r *ReleaseClient) ReleasesReleaseIdVariablesGet(releaseId string) ([]Variable, error) {
	return ReleasesReleaseIdVariablesGet(r.client, releaseId)
}

func (r *ReleaseClient) ReleasesReleaseIdVariablesPost(releaseId string, variable *Variable1) (*Variable, error) {
	return ReleasesReleaseIdVariablesPost(r.client, releaseId, variable)
}

func (r *ReleaseClient) ReleasesReleaseIdVariablesPut(releaseId string, variable *Variable1) (*Variable, error) {
	return ReleasesReleaseIdVariablesPut(r.client, releaseId, variable)
}

func (r *ReleaseClient) ReleasesVariableIdGet(variableId string) (*Variable, error) {
	return ReleasesVariableIdGet(r.client, variableId)
}

func (r *ReleaseClient) ReleasesVariableIdPut(variableId string, variable *Variable) (*Variable, error) {
	return ReleasesVariableIdPut(r.client, variableId, variable)
}

func (r *ReleaseClient) ReleasesVariableIdDelete(variableId string) error {
	return ReleasesVariableIdDelete(r.client, variableId)
}
