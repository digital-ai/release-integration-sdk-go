package command

import (
	"github.com/digital-ai/release-integration-sdk-go/task"
)

type CommandType string

type CommandWrapper struct {
	CommandType CommandType `json:"type"`
	Properties  interface{}
}

func DeserializeCommand(factory CommandFactory, taskContext task.TaskContext) (CommandExecutor, error) {
	var inputs []task.PropertyDefinition
	for _, property := range taskContext.Properties {
		if property.Category == task.InputCategory || property.Category == task.OutputCategory {
			inputs = append(inputs, property)
		}
	}

	command, err := factory.InitCommand(CommandType(taskContext.Type))
	if err != nil {
		return nil, err
	}
	unmarshalErr := task.UnmarshalProperties(inputs, command)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return command, nil
}
