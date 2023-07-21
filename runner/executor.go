package runner

import (
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/logger"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
	"k8s.io/klog/v2"
	"os"
)

// RunFunction represents the function signature for running a task.
type RunFunction func(task.InputContext) *task.Result

// FactoryBuilder represents the function signature for building a command factory.
type FactoryBuilder func(input task.InputContext) (command.CommandFactory, error)

// ExecuteFunction represents the function signature for executing the runner.
type ExecuteFunction func(pluginVersion string, buildDate string, runner Runner)

// Runner is an interface for executing a task.
type Runner interface {
	Run(task.InputContext) *task.Result
}

// SimpleRunner represents a simple implementation of the Runner interface.
type SimpleRunner struct {
	run RunFunction
}

// CommandRunner represents a command-based implementation of the Runner interface.
type CommandRunner struct {
	factoryBuilder FactoryBuilder
}

// NewSimpleRunner creates a new instance of SimpleRunner.
func NewSimpleRunner(run RunFunction) *SimpleRunner {
	var runner SimpleRunner
	runner.run = run
	return &runner
}

// Run executes the task using the provided RunFunction.
func (runner SimpleRunner) Run(ctx task.InputContext) *task.Result {
	return runner.run(ctx)
}

// NewCommandRunner creates a new instance of CommandRunner.
func NewCommandRunner(factoryBuilder FactoryBuilder) *CommandRunner {
	var runner CommandRunner
	runner.factoryBuilder = factoryBuilder
	return &runner
}

// Run executes the task using the command-based approach.
func (runner CommandRunner) Run(ctx task.InputContext) *task.Result {
	returnResult := task.NewResult()
	factory, err := runner.factoryBuilder(ctx)
	if err != nil {
		return returnResult.Error(fmt.Errorf("cannot create factory from task: %v", err))
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

// execute is an internal function that executes the runner.
func execute(pluginVersion string, buildDate string, runner Runner) {
	klog.Infof("PluginVersion:\t%s", pluginVersion)
	klog.Infof("BuildDate:\t%s", buildDate)

	var inputContext task.InputContext
	if err := task.Deserialize(&inputContext); err != nil {
		klog.Errorf("Failed to deserialize input %v", err)
		task.HandleError(fmt.Errorf("failed to deserialize input: %v", err), nil, nil)
		return
	}

	logger.AddSecrets(inputContext)
	executionResult := runner.Run(inputContext)

	resultMap, err := executionResult.Get()
	records := executionResult.GetRecords()
	if err != nil {
		klog.Errorf("Failed executing runner function %v", err)
		task.HandleError(fmt.Errorf("failed to execute run function: %v", err), resultMap, records)
		return
	}
	klog.Infof("Finished executing runner function")
	task.HandleSuccess(resultMap, records)
}

// Execute executes the runner with the provided plugin version, build date, and runner implementation.
func Execute(pluginVersion string, buildDate string, runner Runner) {
	execute(pluginVersion, buildDate, runner)
	if os.Getenv(ExecutionMode) == ExecutionModeDaemon {
		StartInputContextWatcher(execute, pluginVersion, buildDate, runner)
	}
}
