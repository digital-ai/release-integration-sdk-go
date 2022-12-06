package runner

import (
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"k8s.io/klog"
	"os"
)

type runFn func(task.InputContext) (map[string]interface{}, error)

var InputLocation = os.Getenv("INPUT_LOCATION")
var OutputLocation = os.Getenv("OUTPUT_LOCATION")

func Execute(pluginVersion string, buildDate string, run runFn) {
	klog.Infof("PluginVersion:\t%s", pluginVersion)
	klog.Infof("BuildDate:\t%s", buildDate)

	var taskContext task.InputContext
	if err := task.Deserialize(InputLocation, &taskContext); err != nil {
		klog.Errorf("Failed to deserialize input %v", err)
		task.SerializeError(OutputLocation, fmt.Errorf("failed to deserialize input: %w", err))
		return
	}

	executionResult, err := run(taskContext)
	if err != nil {
		klog.Errorf("Failed executing runner function %v", err)
		task.SerializeError(OutputLocation, fmt.Errorf("failed to execute run function: %w", err))
		return
	}
	klog.Infof("Finished executing runner function")
	task.Serialize(OutputLocation, executionResult)
}
