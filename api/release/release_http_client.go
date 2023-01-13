package release

import (
	"github.com/xebialabs/go-remote-runner-wrapper/http"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

func GetReleaseServerHttpClient(ctx task.ReleaseContext) (*http.HttpClient, error) {
	auth := ctx.AutomatedTaskAsUser
	builder := http.NewHttpClientBuilder().
		ForHost(ctx.Url).
		WithBasicAuth(auth.Username, auth.Password)
	return builder.Build()
}
