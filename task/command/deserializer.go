package command

import (
	"encoding/json"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"k8s.io/klog"
)

type CommandType string

type CommandWrapper struct {
	CommandType CommandType `json:"type"`
	Properties  interface{}
}

func DeserializeCommand(factory CommandFactory, propertiesMap map[string]json.RawMessage) (CommandExecutor, error) {
	commandRawJson := propertiesMap["command"]
	var raw json.RawMessage
	wrapper := CommandWrapper{
		Properties: &raw,
	}

	if err := json.Unmarshal(commandRawJson, &wrapper); err != nil {
		klog.Errorf("Failed unmarshalling of command: %v", err)
		return nil, err
	}

	var properties []task.PropertyDefinition
	if err := json.Unmarshal(raw, &properties); err != nil {
		klog.Errorf("Failed unmarshalling of command: %v", err)
		return nil, err
	}

	spawnedCommand, err := factory.InitCommand(wrapper.CommandType)
	if err != nil {
		return nil, err
	}

	commandExec, err := PopulateCommand(spawnedCommand, properties)
	if err != nil {
		return nil, err
	}

	return commandExec, nil
}