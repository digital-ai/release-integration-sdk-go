package command

import (
	"github.com/digital-ai/release-integration-sdk-go/task"
)

// CommandExecutor represents an executor for a command.
type CommandExecutor interface {
	// FetchResult executes the command and returns the result.
	FetchResult() (*task.Result, error)
}
