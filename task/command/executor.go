package command

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

type CommandExecutor interface {
	FetchResult(ctx context.Context) (*task.Result, error)
}

func AbortCommand(command CommandType) CommandType {
	return command + "#abort"
}
