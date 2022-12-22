package test

import (
	"github.com/xebialabs/go-remote-runner-wrapper/runner"
	"github.com/xebialabs/go-remote-runner-wrapper/test"
	"testing"
)

var testMap = map[string]runner.Runner{
	"simple":           independentOutputRunner,
	"fail_empty_input": independentOutputRunner,
	"fail":             failedTestRunner,
	"process_input":    successProcessingRunner,
	"command":          commandRunner,
	"command_secure":   commandRunner,
	"command_fail":     commandRunner,
	"no_such_command":  commandRunner,
}

var fixtures = test.CreateExecutorTestSet("testdata", testMap)

func TestSpec(t *testing.T) {
	test.ConveyTest(t, fixtures)
}
