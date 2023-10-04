package test

import (
	"encoding/json"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

// TestConnectionResult represents the result of a connection test.
//
// It includes a boolean flag 'Success' indicating whether the test was successful,
// and a string 'Output' containing additional information or error message.
type TestConnectionResult struct {
	Success bool   `json:"success"`
	Output  string `json:"output"`
}

// ConnectionTester is an interface for testing connections.
//
// Implementations should provide the TestConnection method for conducting the actual connection test.
type ConnectionTester interface {
	TestConnection() error
}

// TestConnection performs a connection test using the provided ConnectionTester.
//
// It returns a task.Result containing the JSON representation of the TestConnectionResult.
// If the connection test encounters an error, the 'Success' field is set to false,
// and the 'Output' field contains the error message.
func TestConnection(tester ConnectionTester) (*task.Result, error) {
	var result TestConnectionResult
	result.Success = true

	err := tester.TestConnection()
	if err != nil {
		result.Success = false
		result.Output = err.Error()
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return task.NewResult().Json(task.DefaultResponseResultField, resultJson), nil
}
