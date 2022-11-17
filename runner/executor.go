package runner

import (
	"encoding/json"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"k8s.io/klog"
	"os"
)

type runFn func(map[string]json.RawMessage) (map[string]interface{}, error)

var InputLocation = os.Getenv("INPUT_LOCATION")
var OutputLocation = os.Getenv("OUTPUT_LOCATION")

func Execute(pluginVersion string, buildDate string, run runFn) {
	klog.Infof("PluginVersion:\t%s", pluginVersion)
	klog.Infof("BuildDate:\t%s", buildDate)

	propertiesMap, err := task.Deserialize(InputLocation)
	if err != nil {
		klog.Errorf("Failed executing runner function", err)
		task.SerializeError(OutputLocation, err)
	}

	executionResult, err := run(propertiesMap)
	if err != nil {
		klog.Errorf("Failed executing runner function", err)
		task.SerializeError(OutputLocation, err)
	}
	klog.Infof("Finished executing runner function")
	task.Serialize(OutputLocation, executionResult)
}
