package test

import (
	"fmt"
	"github.com/xebialabs/go-remote-runner-wrapper/runner"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
)

var independentOutputRunner = runner.NewSimpleRunner(
	func(_ task.InputContext) (map[string]interface{}, error) {
		var result = make(map[string]interface{})
		result["test"] = "the-result"
		return result, nil
	})

var failedTestRunner = runner.NewSimpleRunner(
	func(_ task.InputContext) (map[string]interface{}, error) {
		return nil, fmt.Errorf("This simulates an error")
	})

var successProcessingRunner = runner.NewSimpleRunner(
	func(input task.InputContext) (map[string]interface{}, error) {
		var result = make(map[string]interface{})
		result["the-type"] = input.Task.Type
		return result, nil
	})
