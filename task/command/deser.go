package command

import (
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

const (
	InputCategory = "input"
)

type CommandType string

type CommandWrapper struct {
	CommandType CommandType `json:"type"`
	Properties  interface{}
}

func DeserializeCommand(factory CommandFactory, taskContext task.TaskContext) (CommandExecutor, error) {
	var inputs []task.PropertyDefinition
	for _, property := range taskContext.Properties {
		if property.Category == InputCategory {
			inputs = append(inputs, property)
		}
	}

	spawnedCommand, err := factory.InitCommand(CommandType(taskContext.Type))
	if err != nil {
		return nil, err
	}

	commandExec, err := PopulateCommand(spawnedCommand, inputs)
	if err != nil {
		return nil, err
	}

	return commandExec, nil
}
