package command

import (
	"encoding/json"
	"strconv"
)

const (
	success      = "success"
	errorMessage = "errorMessage"
)

func MakeSuccessCommandExecutionResultMap(resultField string, result json.RawMessage) map[string]interface{} {
	resultMap := make(map[string]interface{})
	resultMap[success] = strconv.FormatBool(true)
	resultMap[resultField] = string(result)
	return resultMap
}

func MakeFailCommandExecutionResultMap(err error) map[string]interface{} {
	resultMap := make(map[string]interface{})
	resultMap[success] = strconv.FormatBool(false)
	resultMap[errorMessage] = err.Error()
	return resultMap
}
