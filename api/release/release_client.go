package release

import (
	"encoding/base64"
	"github.com/xebialabs/go-remote-runner-wrapper/api/release/openapi"
	_ "github.com/xebialabs/go-remote-runner-wrapper/http"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

type ReleaseClient struct {
	Client *openapi.APIClient
}

func NewReleaseClient(ctx task.ReleaseContext) *ReleaseClient {
	client := NewReleaseApiClient(ctx)
	return &ReleaseClient{Client: client}
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
