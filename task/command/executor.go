package command

import (
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

type CommandExecutor interface {
	FetchResult() (*task.Result, error)
}
