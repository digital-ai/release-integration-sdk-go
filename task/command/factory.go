package command

import (
	"encoding/json"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

type CommandFactory interface {
	InitCommand(commandType CommandType) (CommandExecutor, error)
}

func PopulateCommand(command CommandExecutor, properties []task.PropertyDefinition) (CommandExecutor, error) {
	err := unmarshalCommand(properties, command)
	if err != nil {
		return nil, err
	}
	return command, nil
}

func unmarshalCommand(properties []task.PropertyDefinition, command interface{}) error {
	propsMap := make(map[string]json.RawMessage)
	for _, property := range properties {
		propsMap[property.Name] = property.Value
	}
	jsonMap, err := json.Marshal(propsMap)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonMap, command); err != nil {
		return err
	}
	return nil
}
