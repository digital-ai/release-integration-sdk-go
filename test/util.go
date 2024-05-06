package test

import (
	"encoding/json"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"reflect"
	"testing"
)

// AssertRequestResult compares the actual and expected results of a task in a testing context.
func AssertRequestResult(t *testing.T, actual *task.Result, actualErr error, expected *task.Result, expectedErr error) {
	if actualErr != nil || expectedErr != nil {
		if !reflect.DeepEqual(actualErr, expectedErr) {
			t.Fatalf("Actual: [%v]; Expected: [%v]", actualErr, expectedErr)
		} else {
			t.Logf("Success!")
		}
	} else {
		mapResult, mapErr := actual.Get()
		expectedMap, expectedMapErr := expected.Get()
		if !reflect.DeepEqual(mapErr, expectedMapErr) {
			t.Fatalf("Actual: [%v]; Expected: [%v]", mapErr, expectedMapErr)
		}
		response, err := json.Marshal(mapResult)
		if err != nil {
			t.Fatalf("Error while trying to marshal actual: [%v]", err)
		}

		expectedJson, err := json.Marshal(expectedMap)
		if err != nil {
			t.Fatalf("Error while trying to marshal expected: [%v]", err)
		}

		if !reflect.DeepEqual(string(response), string(expectedJson)) {
			t.Fatalf("Actual: [%v]; Expected: [%v]", string(response), string(expectedJson))
		} else {
			t.Logf("Success!")
		}
	}
}
