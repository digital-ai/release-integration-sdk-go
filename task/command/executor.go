package command

import (
	"github.com/digital-ai/release-integration-sdk-go/task"
)

type CommandExecutor interface {
	FetchResult() (*task.Result, error)
}

func AbortCommand(command CommandType) CommandType {
	return command + "#abort"
}
