package task

import (
	"encoding/json"
	"strconv"
)

const (
	success      = "success"
	resultJson   = "resultJson"
	errorMessage = "errorMessage"
)

type CommandExecutionResult struct {
	Success      bool
	ResultJson   json.RawMessage
	ErrorMessage string
}

func MakeSuccessCommandExecutionResultMap(result json.RawMessage) map[string]string {
	resultMap := make(map[string]string)
	resultMap[success] = strconv.FormatBool(true)
	resultMap[resultJson] = string(result)
	return resultMap
}

func MakeFailCommandExecutionResultMap(err error) map[string]string {
	resultMap := make(map[string]string)
	resultMap[success] = strconv.FormatBool(false)
	resultMap[errorMessage] = err.Error()
	return resultMap
}
