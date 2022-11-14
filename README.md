# Go Remote Runner Wrapper
TODO - set of common types and helpers for go integration repositories

TODO check build
```
go build ./...
```

How to use in repository with SSH key as the authentication method:
1. Add to .gitconfig
```properties
[url "ssh://git@github.com/xebialabs"]
	insteadOf = https://github.com/xebialabs
```

2. Add package to go.mod
```
require (
    github.com/xebialabs/go-remote-runner-wrapper v0.1.1
)
```
### How to test local changes
- remove `github.com/xebialabs/go-remote-runner-wrapper v0.1.1` from go.mod file
- use the go mod edit command to edit the `github.com/xebialabs/go-remote-runner-wrapper`
  module to redirect Go tools from its module path to the local directory
```
 go mod edit -replace github.com/xebialabs/go-remote-runner-wrapper=${replace-with-local-path-to-go-remote-runner-wrapper}
```
- after you run the command, the go.mod file should include a replace directive:
```
module github.com/xebialabs/xlr-argocd-remote-integration

go 1.18

replace github.com/xebialabs/go-remote-runner-wrapper => /Users/jelenakurilic/dev/go-remote-runner-wrapper
```
- run the `go mod tidy` command to synchronize the `github.com/xebialabs/go-remote-runner-wrapper` module's dependencies. The go.mod file:
```
…
require (
	…
	github.com/xebialabs/go-remote-runner-wrapper v0.0.0-00010101000000-000000000000
	…
)
```
