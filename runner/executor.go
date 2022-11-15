package runner

import (
	"encoding/json"
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"os"
)

type runFn func(map[string]json.RawMessage) (json.RawMessage, error)

var InputLocation = os.Getenv("INPUT_LOCATION")
var OutputLocation = os.Getenv("OUTPUT_LOCATION")

func Execute(pluginVersion string, buildDate string, run runFn) {
	fmt.Println("PluginVersion:\t", pluginVersion)
	fmt.Println("BuildDate:\t", buildDate)

	propertiesMap := task.Deserialize(InputLocation)

	executionResult, err := run(propertiesMap)

	var taskResult task.TaskResult
	if err != nil {
		taskResult = task.TaskResult{
			ResponseMapStringString: task.MakeFailCommandExecutionResultMap(err),
		}
	} else {
		taskResult = task.TaskResult{
			ResponseMapStringString: task.MakeSuccessCommandExecutionResultMap(executionResult),
		}
	}
	task.Serialize(OutputLocation, taskResult)
}
