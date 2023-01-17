package openapi

import (
	"encoding/base64"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

func NewReleaseApiClient(ctx task.ReleaseContext) *APIClient {
	conf := NewConfiguration()
	conf.DefaultHeader = map[string]string{
		"Authorization": "Basic " + basicAuth(ctx.AutomatedTaskAsUser.Username, ctx.AutomatedTaskAsUser.Password),
		"Content-Type":  "application/json",
	}
	conf.Host = ctx.Url
	conf.Scheme = "http"
	//conf.AddDefaultHeader("Authorization", "Basic "+basicAuth(ctx.AutomatedTaskAsUser.Username, ctx.AutomatedTaskAsUser.Password))
	//conf.AddDefaultHeader("Content-Type", "application/json")

	releaseApi := NewAPIClient(conf)
	return releaseApi
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
