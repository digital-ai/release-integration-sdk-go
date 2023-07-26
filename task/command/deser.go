package command

import (
	"github.com/digital-ai/release-integration-sdk-go/task"
)

// CommandType represents the type of a command.
type CommandType string

// AbortCommandType represents the type of command that is to be used for aborting a running task.
// It differentiates between standard commands and abort commands, allowing different handling.
type AbortCommandType CommandType

// CommandWrapper represents a wrapper for a command.
type CommandWrapper struct {
	CommandType CommandType `json:"type"`
	Properties  interface{}
}

func deserializeCommand(factory CommandFactory, properties []task.PropertyDefinition, commandType CommandType) (CommandExecutor, error) {
	var inputs []task.PropertyDefinition
	for _, property := range properties {
		if property.Category == task.InputCategory || property.Category == task.OutputCategory {
			inputs = append(inputs, property)
		}
	}

	command, err := factory.InitCommand(commandType)
	if err != nil {
		return nil, err
	}
	unmarshalErr := task.UnmarshalProperties(inputs, command)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return command, nil
}

// DeserializeCommand deserializes a command from the task context using the provided command factory.
func DeserializeCommand(factory CommandFactory, taskContext task.TaskContext) (CommandExecutor, error) {
	return deserializeCommand(factory, taskContext.Properties, CommandType(taskContext.Type))
}

func DeserializeAbortCommand(factory CommandFactory, taskContext task.TaskContext) (CommandExecutor, error) {
	return deserializeCommand(factory, taskContext.Properties, AbortCommand(CommandType(taskContext.Type)))
}
