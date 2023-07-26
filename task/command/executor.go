package command

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

// CommandExecutor represents an executor for a command.
type CommandExecutor interface {
	// FetchResult executes the command and returns the result.
	FetchResult(ctx context.Context) (*task.Result, error)
}

// AbortCommand is a wrapper for executing abort logic for a specified command
func AbortCommand(command CommandType) CommandType {
	return command + "#abort"
}
