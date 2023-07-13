package command

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

// CommandExecutor represents an executor for a command.
type CommandExecutor interface {
	FetchResult(ctx context.Context) (*task.Result, error)
}

func AbortCommand(command CommandType) CommandType {
	return command + "#abort"
}
