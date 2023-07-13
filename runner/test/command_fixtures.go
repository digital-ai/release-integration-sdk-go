package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/runner"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
)

// MockFactory represents a mock implementation of the CommandFactory interface.
type MockFactory struct {
	ResultField       string
	SecureResultField string
}

// MockSuccessCommand represents a mock implementation of a successful command.
type MockSuccessCommand struct {
	FieldLabel string
	Parameter  string `json:"parameter"`
}

// MockSuccessSecureCommand represents a mock implementation of a successful secure command.
type MockSuccessSecureCommand struct {
	FieldLabel string
	Parameter  string `json:"parameter"`
}

// MockFailedCommand represents a mock implementation of a failed command.
type MockFailedCommand struct {
	Parameter string `json:"parameter"`
}

// mockResponse generates a mock task result with the specified field label and parameter.
func mockResponse(fieldLabel string, param string) (*task.Result, error) {
	result := make(map[string]string)
	result["field"] = "command-success"
	result["parameter"] = param
	slice, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return task.NewResult().Json(fieldLabel, slice), nil
}

func (command MockSuccessCommand) FetchResult(_ context.Context) (*task.Result, error) {
	return mockResponse(command.FieldLabel, command.Parameter)
}

func (command MockSuccessSecureCommand) FetchResult(_ context.Context) (*task.Result, error) {
	return mockResponse(command.FieldLabel, command.Parameter)
}

func (command MockFailedCommand) FetchResult(_ context.Context) (*task.Result, error) {
	return nil, fmt.Errorf("simulate error")
}

// InitCommand initializes the command based on the provided command type.
func (mock MockFactory) InitCommand(commandType command.CommandType) (command.CommandExecutor, error) {
	switch commandType {
	case "success.Command":
		return &MockSuccessCommand{FieldLabel: mock.ResultField}, nil
	case "success.SecureCommand":
		return &MockSuccessSecureCommand{FieldLabel: mock.SecureResultField}, nil
	case "fail.Command":
		return &MockFailedCommand{}, nil
	}
	return nil, fmt.Errorf("cannot find command type")
}

var commandRunner = runner.NewCommandRunner(
	func(input task.InputContext) (command.CommandFactory, error) {
		return MockFactory{
			"result",
			"secureResult",
		}, nil
	},
)
