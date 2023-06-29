package test

import (
	"github.com/digital-ai/release-integration-sdk-go/runner"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/test"
	"os"
	"testing"
)

var testMap = map[string]runner.Runner{
	"simple":                     independentOutputRunner,
	"fail_empty_input":           independentOutputRunner,
	"fail":                       failedTestRunner,
	"process_string":             stringPropertyRunner,
	"process_bool":               simpleRunnerWith(task.NewResult().Bool("the-bool", true)),
	"process_date":               datePropertyRunner,
	"process_date_string":        simpleRunnerWith(task.NewResult().DateString("the-date", "Jan _2 2006 15:04:05.000000000", "Jan 11 2023 10:44:45.000000142")),
	"process_json":               simpleRunnerWith(task.NewResult().Json("the-json", jsonPayload)),
	"process_string_json":        simpleRunnerWith(task.NewResult().JsonString("the-json-string", ".metadata.name", jsonPayload)),
	"process_bool_json":          simpleRunnerWith(task.NewResult().JsonBool("the-json-bool", ".just-random-bool", jsonPayload)),
	"process_bool_error":         simpleRunnerWith(task.NewResult().JsonBool("the-json-bool", ".metadata.creationTimestamp", jsonPayload)),
	"process_date_json":          simpleRunnerWith(task.NewResult().JsonDate("the-json-date", ".metadata.creationTimestamp", "2006-01-02T15:04:05Z", jsonPayload)),
	"process_date_format_error":  simpleRunnerWith(task.NewResult().JsonDate("the-json-date", ".metadata.creationTimestamp", "2006-01-02T15:04:05", jsonPayload)),
	"process_node_json":          simpleRunnerWith(task.NewResult().JsonNode("the-json", ".metadata.labels", jsonPayload)),
	"process_node_error_missing": simpleRunnerWith(task.NewResult().JsonNode("the-json", ".metadata.label", jsonPayload)),
	"process_composed":           composedPropertyRunner,
	"command":                    commandRunner,
	"command_secure":             commandRunner,
	"command_fail":               commandRunner,
	"no_such_command":            commandRunner,
	// TODO add abort test
}

var fixtures = test.CreateExecutorTestSet("testdata", testMap)

func TestSpec(t *testing.T) {
	os.Setenv("TZ", "UTC")
	test.ConveyTest(t, fixtures)
}
