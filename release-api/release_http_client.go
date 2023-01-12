package release_api

import (
	"github.com/xebialabs/go-remote-runner-wrapper/http"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

func GetReleaseServerHttpClient(ctx task.ReleaseContext) (*http.HttpClient, error) {
	auth := ctx.AutomatedTaskAsUser

	httpClientConfig := &http.HttpClientConfig{
		Host: ctx.Url,
		BasicAuthentication: &http.BasicAuthentication{
			Username: auth.Username,
			Password: auth.Password,
		},
	}
	builder := http.NewHttpClientBuilder().WithHttpClientConfig(httpClientConfig)
	return builder.Build()
}
