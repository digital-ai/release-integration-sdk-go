package command

import (
	"github.com/digital-ai/release-integration-sdk-go/task"
)

// CommandType represents the type of a command.
type CommandType string

// CommandWrapper represents a wrapper for a command.
type CommandWrapper struct {
	CommandType CommandType `json:"type"`
	Properties  interface{}
}

// DeserializeCommand deserializes a command from the task context using the provided command factory.
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
