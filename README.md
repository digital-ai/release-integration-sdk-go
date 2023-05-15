# release-integration-sdk-go

Release integration SDK for projects that use the container tasks made in go.

See [Project WIKI](https://github.com/digital-ai/release-integration-sdk-go/wiki)

## Release Api 
Since the release-integration-sdk-go has support for calling Release public API, plugins can consume them.

Client files are generated with command:

```shell
openapi-generator generate -i https://raw.githubusercontent.com/digital-ai/release-api/main/rest-api/release-api.yaml -g go -o api/release/openapi/ -p enumClassPrefix=true
```

Example of calling Release Api that should be placed in the plugin:

```go
package main

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/api/release"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"k8s.io/klog/v2"
)

func main() {
	ctx := task.ReleaseContext{
		Id: "release-id",
		AutomatedTaskAsUser: task.AutomatedTaskAsUserContext{
			Username: "admin",
			Password: "admin",
		},
		Url: "localhost:5516",
	}

	releaseClient, err := release.NewReleaseApiClient(ctx)
	if err != nil {
		klog.Errorln("Error while getting release client: ", err)
	}

	_, err = releaseClient.ReleaseApi.DeleteReleaseVariable(context.TODO(), "variable1").Execute()
	if err != nil {
		klog.Errorln("Error while deleting release variable: ", err)
	}
}
```
