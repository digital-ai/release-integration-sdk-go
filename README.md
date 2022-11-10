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
