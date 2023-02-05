package runner

import (
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/logger"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"github.com/xebialabs/go-remote-runner-wrapper/task/command"
	"k8s.io/klog"
	"os"
)

type RunFunction func(task.InputContext) *task.Result
type FactoryBuilder func(task task.TaskContext) (command.CommandFactory, error)

type Runner interface {
	Run(task.InputContext) *task.Result
}

type SimpleRunner struct {
	run RunFunction
}

type CommandRunner struct {
	factoryBuilder FactoryBuilder
}

func NewSimpleRunner(run RunFunction) *SimpleRunner {
	var runner SimpleRunner
	runner.run = run
	return &runner
}

func (runner SimpleRunner) Run(ctx task.InputContext) *task.Result {
	return runner.run(ctx)
}

func NewCommandRunner(factoryBuilder FactoryBuilder) *CommandRunner {
	var runner CommandRunner
	runner.factoryBuilder = factoryBuilder
	return &runner
}

func (runner CommandRunner) Run(ctx task.InputContext) *task.Result {
	returnResult := task.NewResult()
	factory, err := runner.factoryBuilder(ctx.Task)
	if err != nil {
		return returnResult.Error(fmt.Errorf("cannot crete factory from task: %v", err))
	}
	exec, err := command.DeserializeCommand(factory, ctx.Task)
	if err != nil {
		return returnResult.Error(fmt.Errorf("cannot deserialize input: %v", err))
	}

	result, err := exec.FetchResult()
	if err != nil {
		klog.Infof("Finished executing command with error %v", err)
		return returnResult.Error(err)
	}
	klog.Infoln("Finished executing command")
	return result
}

func Execute(pluginVersion string, buildDate string, runner Runner) {

	klog.Infof("PluginVersion:\t%s", pluginVersion)
	klog.Infof("BuildDate:\t%s", buildDate)

	var InputLocation = os.Getenv("INPUT_LOCATION")
	var OutputLocation = os.Getenv("OUTPUT_LOCATION")

	var taskContext task.InputContext
	if err := task.Deserialize(InputLocation, &taskContext); err != nil {
		klog.Errorf("Failed to deserialize input %v", err)
		errorResult, _ := task.NewErrorResult(fmt.Errorf("failed to deserialize input: %v", err)).Get()
		task.SerializeError(OutputLocation, errorResult)
		return
	}

	logger.AddSecrets(taskContext)

	executionResult := runner.Run(taskContext)
	resultMap, err := executionResult.Get()
	if err != nil {
		klog.Errorf("Failed executing runner function %v", err)
		errorResult, _ := task.NewErrorResult(fmt.Errorf("failed to execute run function: %v", err)).Get()
		task.SerializeError(OutputLocation, errorResult)
		return
	}
	klog.Infof("Finished executing runner function")
	task.Serialize(OutputLocation, resultMap)
}
