package runner

import (
	"context"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/logger"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"
)

// RunFunction represents the function signature for running a task.
type RunFunction func(task.InputContext) *task.Result

type CommandFactoryBuilder func(input task.InputContext) (command.CommandFactory, error)

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
	commandFactoryBuilder CommandFactoryBuilder
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

func NewCommandRunner(factoryBuilder CommandFactoryBuilder) *CommandRunner {
	var runner CommandRunner
	runner.commandFactoryBuilder = factoryBuilder
	return &runner
}

// Run executes the task using the command-based approach.
func (runner CommandRunner) Run(ctx task.InputContext) *task.Result {
	returnResult := task.NewResult()
	factory, err := runner.commandFactoryBuilder(ctx)
	if err != nil {
		return returnResult.Error(fmt.Errorf("cannot create factory from task: %v", err))
	}

	exec, err := command.DeserializeCommand(factory, ctx.Task)
	if err != nil {
		return returnResult.Error(fmt.Errorf("cannot deserialize input: %v", err))
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGABRT)

	resultChannel := make(chan *task.Result, 1)
	cancelableCtx, cancel := context.WithCancel(context.Background())

	go func() {
		result, err := exec.FetchResult(cancelableCtx)
		if err != nil {
			klog.Infof("Finished executing command with error %v", err)
			if result != nil {
				resultChannel <- result.Error(err)
				return
			}
			resultChannel <- returnResult.Error(err)
			return
		}
		klog.Infoln("Finished executing command")
		resultChannel <- result
	}()

	select {
	case <-signalChannel:
		abortExec, err := command.DeserializeAbortCommand(factory, ctx.Task)
		execResult, err := abortExec.FetchResult(context.Background())
		cancel()
		if err != nil {
			klog.Infof("Finished executing abort command with error %v", err)
			return returnResult.Error(err)
		}
		abortResult := task.NewResult()
		if err != nil {
			return returnResult.Error(fmt.Errorf("cannot deserialize abort command: %v", err))
		}

		abortResult.Aborted(execResult)
		return abortResult
	case result := <-resultChannel:
		cancel()
		return result
	}
}

// execute is an internal function that executes the runner.
func execute(pluginVersion string, buildDate string, runner Runner) {
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
		switch err.(type) {
		case *task.AbortError:
			klog.Infof("Executing aborted - returning abort result")
			task.HandleAbort(resultMap)
		default:
			klog.Errorf("Failed executing runner function %v", err)
			task.HandleError(fmt.Errorf("failed to execute run function: %v", err), resultMap)
		}
		return
	}
	klog.Infof("Finished executing runner function")
	task.HandleSuccess(resultMap)
}

// Execute executes the runner with the provided plugin version, build date, and runner implementation.
func Execute(pluginVersion string, buildDate string, runner Runner) {
	execute(pluginVersion, buildDate, runner)
	if os.Getenv(ExecutionMode) == ExecutionModeDaemon {
		StartInputContextWatcher(execute, pluginVersion, buildDate, runner)
	}
}
