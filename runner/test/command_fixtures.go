package test

import (
	"encoding/json"
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/runner"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"github.com/xebialabs/go-remote-runner-wrapper/task/command"
)

type MockFactory struct {
	ResultField       string
	SecureResultField string
}
type MockSuccessCommand struct {
	FieldLabel string
	Parameter  string `json:"parameter"`
}
type MockSuccessSecureCommand struct {
	FieldLabel string
	Parameter  string `json:"parameter"`
}
type MockFailedCommand struct {
	Parameter string `json:"parameter"`
}

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

func (command MockSuccessCommand) FetchResult() (*task.Result, error) {
	return mockResponse(command.FieldLabel, command.Parameter)
}

func (command MockSuccessSecureCommand) FetchResult() (*task.Result, error) {
	return mockResponse(command.FieldLabel, command.Parameter)
}

func (command MockFailedCommand) FetchResult() (*task.Result, error) {
	return nil, fmt.Errorf("simulate error")
}

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
	func(input task.TaskContext) (command.CommandFactory, error) {
		return MockFactory{
			"result",
			"secureResult",
		}, nil
	},
)
