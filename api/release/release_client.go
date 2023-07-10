package release

import (
	"encoding/base64"
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"net/url"
)

// NewReleaseApiClient creates a new instance of the openapi.APIClient for interacting with the Release API.
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

// basicAuth returns the Basic Authentication string for the given username and password.
// It takes the username and password as input and encodes them in the format required for Basic Authentication.
// The function returns the encoded string that can be used in HTTP headers for authentication.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
