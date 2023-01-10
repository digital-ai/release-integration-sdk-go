package runner

import (
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"github.com/xebialabs/go-remote-runner-wrapper/task/command"
	"k8s.io/klog"
	"os"
	"strconv"
)

type RunFunction func(task.InputContext) *Result
type FactoryBuilder func(task task.TaskContext) (command.CommandFactory, error)

type Runner interface {
	Run(task.InputContext) *Result
}

type SimpleRunner struct {
	run RunFunction
}

type CommandRunner struct {
	factoryBuilder    FactoryBuilder
	ResultField       string
	SecureResultField string
}

func NewSimpleRunner(run RunFunction) *SimpleRunner {
	var runner SimpleRunner
	runner.run = run
	return &runner
}

func (runner SimpleRunner) Run(ctx task.InputContext) *Result {
	return runner.run(ctx)
}

func NewCommandRunner(factoryBuilder FactoryBuilder, resultField string, securedResultField string) *CommandRunner {
	var runner CommandRunner
	runner.factoryBuilder = factoryBuilder
	runner.ResultField = resultField
	runner.SecureResultField = securedResultField
	return &runner
}

func (runner CommandRunner) Run(ctx task.InputContext) *Result {
	returnResult := NewResult()
	factory, err := runner.factoryBuilder(ctx.Task)
	if err != nil {
		return returnResult.Error(fmt.Errorf("cannot crete factory from task: %v", err))
	}
	exec, err := command.DeserializeCommand(factory, ctx.Task)
	if err != nil {
		return returnResult.Error(fmt.Errorf("cannot deserialize input: %v", err))
	}

	result, masked, err := exec.FetchResult()
	if err != nil {
		klog.Infof("Finished executing command with error %v", err)
		return returnResult.Error(err)
	}
	klog.Infoln("Finished executing command")
	if masked {
		secureResponse, convErr := strconv.Unquote(string(result))
		if convErr != nil {
			klog.Infof("Failed to mask secure command with error %v", convErr)
			return returnResult.Error(convErr)
		}
		return returnResult.String(runner.SecureResultField, secureResponse)
	}
	return returnResult.Json(runner.ResultField, result)
}

func Execute(pluginVersion string, buildDate string, runner Runner) {

	klog.Infof("PluginVersion:\t%s", pluginVersion)
	klog.Infof("BuildDate:\t%s", buildDate)

	var InputLocation = os.Getenv("INPUT_LOCATION")
	var OutputLocation = os.Getenv("OUTPUT_LOCATION")

	var taskContext task.InputContext
	if err := task.Deserialize(InputLocation, &taskContext); err != nil {
		klog.Errorf("Failed to deserialize input %v", err)
		errorResult, _ := NewErrorResult(fmt.Errorf("failed to deserialize input: %v", err)).Get()
		task.SerializeError(OutputLocation, errorResult)
		return
	}

	executionResult := runner.Run(taskContext)
	resultMap, err := executionResult.Get()
	if err != nil {
		klog.Errorf("Failed executing runner function %v", err)
		errorResult, _ := NewErrorResult(fmt.Errorf("failed to execute run function: %v", err)).Get()
		task.SerializeError(OutputLocation, errorResult)
		return
	}
	klog.Infof("Finished executing runner function")
	task.Serialize(OutputLocation, resultMap)
}
