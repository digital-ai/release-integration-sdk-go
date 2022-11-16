package runner

import (
	"encoding/json"
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"os"
)

type runFn func(map[string]json.RawMessage) (map[string]interface{}, error)

var InputLocation = os.Getenv("INPUT_LOCATION")
var OutputLocation = os.Getenv("OUTPUT_LOCATION")

func Execute(pluginVersion string, buildDate string, run runFn) {
	fmt.Println("PluginVersion:\t", pluginVersion)
	fmt.Println("BuildDate:\t", buildDate)

	propertiesMap := task.Deserialize(InputLocation)

	executionResult, err := run(propertiesMap)
	if err != nil {
		fmt.Println("Error executing the runner function", err)
		task.SerializeError(OutputLocation, err)
	}

	task.Serialize(OutputLocation, executionResult)
}
