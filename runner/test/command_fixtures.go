package test

import (
	"encoding/json"
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/runner"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"github.com/xebialabs/go-remote-runner-wrapper/task/command"
)

type MockFactory struct{}
type MockSuccessCommand struct {
	Parameter string `json:"parameter"`
}
type MockSuccessSecureCommand struct {
	Parameter string `json:"parameter"`
}
type MockFailedCommand struct {
	Parameter string `json:"parameter"`
}

func mockResponse(param string) (json.RawMessage, bool, error) {
	result := make(map[string]string)
	result["field"] = "command-success"
	result["parameter"] = param
	slice, err := json.Marshal(result)
	if err != nil {
		return nil, false, err
	}
	return slice, false, nil
}

func (command MockSuccessCommand) FetchResult() (json.RawMessage, bool, error) {
	return mockResponse(command.Parameter)
}

func (command MockSuccessSecureCommand) FetchResult() (json.RawMessage, bool, error) {
	return mockResponse(command.Parameter)
}

func (command MockFailedCommand) FetchResult() (json.RawMessage, bool, error) {
	return nil, false, fmt.Errorf("simulate error")
}

func (mock MockFactory) InitCommand(commandType command.CommandType) (command.CommandExecutor, error) {
	switch commandType {
	case "success.Command":
		return &MockSuccessCommand{}, nil
	case "success.SecureCommand":
		return &MockSuccessSecureCommand{}, nil
	case "fail.Command":
		return &MockFailedCommand{}, nil
	}
	return nil, fmt.Errorf("cannot find command type")
}

var commandRunner = runner.NewCommandRunner(
	func(input task.TaskContext) (command.CommandFactory, error) {
		return MockFactory{}, nil
	},
	"result",
	"secureResult",
)
