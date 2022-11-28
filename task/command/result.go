package command

import (
	"encoding/json"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"strconv"
)

const (
	success = "success"
)

func MakeSuccessCommandExecutionResultMap(resultField string, result json.RawMessage) map[string]interface{} {
	resultMap := make(map[string]interface{})
	resultMap[success] = strconv.FormatBool(true)
	var commandResultMap map[string]json.RawMessage
	err := json.Unmarshal(result, &commandResultMap)
	if err != nil {
		return MakeFailCommandExecutionResultMap(err)
	}
	resultMap[resultField] = commandResultMap
	return resultMap
}

func MakeSuccessSecureCommandExecutionResultMap(resultField string, result string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	resultMap[success] = strconv.FormatBool(true)
	resultMap[resultField] = result
	return resultMap
}

func MakeFailCommandExecutionResultMap(err error) map[string]interface{} {
	return task.ErrorResultMap(err)
}
