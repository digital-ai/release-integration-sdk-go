package runner

import (
	"context"
	"errors"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/logger"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"
)

// AbortContextFieldKey is a context key used to share data, such as a task ID,
// between an ongoing task and a potential abort task.
const AbortContextFieldKey = "abortContextMap"

// RunFunction represents the function signature for running a task.
type RunFunction func(task.InputContext) *task.Result

// CommandFactoryBuilder represents the function signature for building a command factory.
type CommandFactoryBuilder func(input task.InputContext) (command.CommandFactory, error)

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

// NewCommandRunner creates a new instance of CommandRunner
func NewCommandRunner(factoryBuilder CommandFactoryBuilder) *CommandRunner {
	var runner CommandRunner
	runner.commandFactoryBuilder = factoryBuilder
	return &runner
}

// Run executes the task using the command-based approach.
// The command is then executed asynchronously and waits for
// either an abort signal or for the execution to complete.
//
// If an abort signal is received, it deserializes an abort command and executes it.
// The task result reflects the output of the abort command.
//
// If the command finishes executing before an abort signal is received, the task result
// reflects the output of the executed command.
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
	rootCtx := context.WithValue(context.Background(), AbortContextFieldKey, make(map[string]interface{}))
	cancelableCtx, cancel := context.WithCancel(rootCtx)

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
		if abortExec == nil {
			cancel()
			klog.Infoln("Aborted without abort command being explicitly defined")
			abortResult := task.NewResult()
			return returnResult.Aborted(abortResult)
		}
		execResult, err := abortExec.FetchResult(rootCtx)
		defer cancel()
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
		task.HandleError(err, nil, nil)
		return
	}

	logger.AddSecrets(inputContext)
	executionResult := runner.Run(inputContext)

	resultMap, err := executionResult.Get()
	records := executionResult.GetRecords()
	if err != nil {
		switch err.(type) {
		case *task.AbortError:
			klog.Infof("Executing aborted - returning abort result")
			task.HandleAbort(resultMap)
		default:
			klog.Errorf("Failed executing runner function %v", err)
			task.HandleError(err, resultMap, records)
		}
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

func PutValueToContextMap(ctx context.Context, contextValueKey string, key string, value interface{}) error {
	dataContext, ok := ctx.Value(contextValueKey).(map[string]interface{})
	if !ok {
		return errors.New("expected context value of type 'map[string]interface{}'")
	}
	dataContext[key] = value
	return nil
}

func GetValueFromContext(ctx context.Context, contextValueKey string, valueKey string) (interface{}, error) {
	ctxMap, ok := ctx.Value(contextValueKey).(map[string]interface{})
	if !ok {
		return nil, errors.New("expected context value of type 'map[string]interface{}'")
	}

	extractedValue, ok := ctxMap[valueKey]
	if !ok {
		return nil, errors.New("expected value in context map")
	}

	return extractedValue, nil
}
