# release-integration-sdk-go

Release integration SDK for projects that use the container tasks made in go.

See [Project WIKI](https://github.com/digital-ai/release-integration-sdk-go/wiki)

## Release Api 
Since the release-integration-sdk-go has support for calling Release public API, plugins can consume them.

Client files are generated with command:

`openapi-generator generate -i release-api.yaml -g go -o api/release/openapi/ -p enumClassPrefix=true`

Example of calling Release Api that should be placed in the plugin:
```go
package main

import (
	"github.com/digital-ai/release-integration-sdk-go"
	"github.com/digital-ai/release-integration-sdk-go/task"
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

	releaseClient := release.NewReleaseApiClient(ctx)
	result, err := releaseClient.ReleaseApi.DeleteReleaseVariable(context.TODO(), "variable1").Execute()
}
```
