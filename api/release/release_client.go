package release

import (
	"encoding/base64"
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"net/url"
)

func NewReleaseApiClient(ctx task.ReleaseContext) (*openapi.APIClient, error) {
	conf := openapi.NewConfiguration()
	conf.DefaultHeader = map[string]string{
		"Authorization": "Basic " + basicAuth(ctx.AutomatedTaskAsUser.Username, ctx.AutomatedTaskAsUser.Password),
		"Content-Type":  "application/json",
	}

	baseUrl, err := url.Parse(ctx.Url)
	if err != nil {
		return nil, err
	}

	conf.Host = baseUrl.Host
	return openapi.NewAPIClient(conf), nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
