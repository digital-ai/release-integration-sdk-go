package runner

import (
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/logger"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
	"k8s.io/klog/v2"
)

var ExecutionChan = make(chan bool, 1)

type RunFunction func(task.InputContext) *task.Result
type FactoryBuilder func(input task.InputContext) (command.CommandFactory, error)

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
	factory, err := runner.factoryBuilder(ctx)
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
		if result != nil {
			return result.Error(err)
		}
		return returnResult.Error(err)
	}
	klog.Infoln("Finished executing command")
	return result
}

func Execute(pluginVersion string, buildDate string, runner Runner) {
	klog.Infof("PluginVersion:\t%s", pluginVersion)
	klog.Infof("BuildDate:\t%s", buildDate)

	var inputContext task.InputContext
	if err := task.Deserialize(&inputContext); err != nil {
		klog.Errorf("Failed to deserialize input %v", err)
		task.HandleError(fmt.Errorf("failed to deserialize input: %v", err), nil)
		return
	}

	logger.AddSecrets(inputContext)
	executionResult := runner.Run(inputContext)

	resultMap, err := executionResult.Get()
	if err != nil {
		klog.Errorf("Failed executing runner function %v", err)
		task.HandleError(fmt.Errorf("failed to execute run function: %v", err), resultMap)
		return
	}
	klog.Infof("Finished executing runner function")
	task.HandleSuccess(resultMap)
}
