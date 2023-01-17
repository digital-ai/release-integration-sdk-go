# go-remote-runner-wrapper

Go remote runner wrapper for projects that use the container tasks made in go.

See [Project WIKI](https://github.com/xebialabs/go-remote-runner-wrapper/wiki)

## Release Api 
Since the go-remote-runner-wrapper has support for calling Release public API, plugins can consume them.
Example of calling Release Api that should be placed in the plugin:
```go
package main

import (
	"github.com/xebialabs/go-remote-runner-wrapper/api/release"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
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

    releaseApi := release.NewReleaseClient(ctx)
    execute, err := releaseApi.DeleteVariable("variable1")
    ...
}
```
