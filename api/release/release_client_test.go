package release

import (
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xebialabs/go-remote-runner-wrapper/api/release/openapi"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"reflect"
	"testing"
)

func TestReleaseClient(t *testing.T) {
	releaseApi := NewReleaseClient(ctx)
	ConveyTest(t, releaseApi)
}

func ConveyTest(t *testing.T, client *ReleaseClient) {
	Convey("Test ReleaseApi calls", t, func() {
		httpmock.Activate()
		var1 := "var1"
		variable := openapi.Variable{Label: &var1, Title: &var1}

		Convey("running GetVariablesForRelease", func() {
			vars := []openapi.Variable{variable}

			httpmock.RegisterResponder("GET", "http://localhost:5516/api/v1/releases/release-1/variables",
				httpmock.NewJsonResponderOrPanic(200, vars))

			resp, _, _ := client.GetVariablesForRelease("release-1")
			So(resp, AssertFunc, vars)
		})

		Convey("Test ReleasesApi CreateVariablesForRelease", func() {
			httpmock.RegisterResponder("POST", "http://localhost:5516/api/v1/releases/release-1/variables",
				httpmock.NewJsonResponderOrPanic(200, variable))

			resp, _, _ := client.CreateVariablesForRelease("release-1", openapi.Variable1{Label: &var1})
			So(resp, AssertFunc, &variable)
		})

		Convey("Test ReleasesApi UpdateVariablesForRelease", func() {
			httpmock.RegisterResponder("PUT", "http://localhost:5516/api/v1/releases/release-1/variables",
				httpmock.NewJsonResponderOrPanic(200, []openapi.Variable{variable}))

			resp, _, _ := client.UpdateVariablesForRelease("release-1", []openapi.Variable{variable})
			fmt.Println(resp)
			So(resp, AssertFunc, []openapi.Variable{variable})
		})

		Convey("Test ReleasesApi GetVariable", func() {
			httpmock.RegisterResponder("GET", "http://localhost:5516/api/v1/releases/variable-1",
				httpmock.NewJsonResponderOrPanic(200, variable))

			resp, _, _ := client.GetVariable("variable-1")
			So(resp, AssertFunc, &variable)
		})

		Convey("Test ReleasesApi UpdateVariable", func() {
			updated := openapi.Variable{Label: &var1}
			httpmock.RegisterResponder("PUT", "http://localhost:5516/api/v1/releases/variable-1",
				httpmock.NewJsonResponderOrPanic(200, updated))

			resp, _, _ := client.UpdateVariable("variable-1", variable)
			So(resp, AssertFunc, &updated)
		})

		Convey("Test ReleasesApi DeleteVariable", func() {
			httpmock.RegisterResponder("DELETE", "http://localhost:5516/api/v1/releases/variable-1",
				httpmock.NewJsonResponderOrPanic(404, errors.New("not found")))

			resp, err := client.DeleteVariable("variable-1")
			So(resp.StatusCode, ShouldEqual, 404)
			So(err, ShouldEqual, err)
		})

		Convey("Test ReleasesApi GetVariableValuesForRelease", func() {
			varValues := map[string]string{"var1": "var1"}
			httpmock.RegisterResponder("GET", "http://localhost:5516/api/v1/releases/release-1/variableValues",
				httpmock.NewJsonResponderOrPanic(200, varValues))

			resp, _, _ := client.GetVariableValuesForRelease("release-1")
			So(resp, AssertFunc, varValues)
		})

		Convey("Test ReleasesApi GetVariablePossibleValues", func() {
			possibleValues := []map[string]interface{}{{"var1": "var1"}}
			httpmock.RegisterResponder("GET", "http://localhost:5516/api/v1/releases/variable-1/possibleValues",
				httpmock.NewJsonResponderOrPanic(200, possibleValues))

			resp, _, _ := client.GetVariablePossibleValues("variable-1")
			So(resp, AssertFunc, possibleValues)
		})

		Convey("Test ReleasesApi IsVariableUsed", func() {
			httpmock.RegisterResponder("GET", "http://localhost:5516/api/v1/releases/variable-1/used",
				httpmock.NewJsonResponderOrPanic(200, true))

			used, _, _ := client.IsVariableUsed("variable-1")
			So(used, ShouldBeTrue)
		})

		Convey("Test ReleasesApi ReplaceVariable", func() {
			varOrValue := openapi.VariableOrValue{Value: map[string]interface{}{"lab1": "var1"}}
			httpmock.RegisterResponder("POST", "http://localhost:5516/api/v1/releases/variable-1/replace",
				httpmock.NewJsonResponderOrPanic(200, varOrValue))

			_, err := client.ReplaceVariable("variable-1", varOrValue)
			So(err, ShouldBeNil)
		})

		Reset(func() {
			httpmock.DeactivateAndReset()
		})
	})
}

var ctx = task.ReleaseContext{
	Id: "release-1",
	AutomatedTaskAsUser: task.AutomatedTaskAsUserContext{
		Username: "admin",
		Password: "admin",
	},
	Url: "localhost:5516",
}

var AssertFunc = func(actual interface{}, expected ...interface{}) string {
	if reflect.DeepEqual(actual, expected[0]) {
		return ""
	}
	return fmt.Sprintf("actual: %s, expected: %s", string(actual.([]byte)), string(expected[0].([]byte)))
}
