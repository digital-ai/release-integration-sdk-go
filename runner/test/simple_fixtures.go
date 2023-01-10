package test

import (
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/runner"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

var independentOutputRunner = runner.NewSimpleRunner(
	func(_ task.InputContext) *runner.Result {
		return runner.NewResult().String("test", "the-result")
	})

var failedTestRunner = runner.NewSimpleRunner(
	func(_ task.InputContext) *runner.Result {
		return runner.NewErrorResult(fmt.Errorf("This simulates an error"))
	})

var successProcessingRunner = runner.NewSimpleRunner(
	func(input task.InputContext) *runner.Result {
		return runner.NewResult().String("the-type", input.Task.Type)
	})
