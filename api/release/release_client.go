package release

import (
	"encoding/base64"
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	_ "github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

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
