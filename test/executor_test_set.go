package test

import (
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/runner"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path"
	"strings"
	"testing"
)

const (
	pluginVersion = "test"
	buildDate     = "now"
)

// CreateExecutorTestSet creates a set of executor tests based on the specified base path and runners.
func CreateExecutorTestSet(basePath string, tests map[string]runner.Runner) []ExecutorTest {
	var testSet []ExecutorTest
	for test, runner := range tests {
		executorTest, err := CreateExecutorTest(path.Join(basePath, test)).WithRunnerFunction(runner).Build()
		if err != nil {
			panic(err)
		}
		testSet = append(testSet, *executorTest)
	}
	return testSet
}

// ConveyTest executes the executor tests using Convey from goconvey.
func ConveyTest(t *testing.T, tests []ExecutorTest) {
	Convey("Given input file", t, func() {
		for _, fixture := range tests {
			t.Setenv("INPUT_LOCATION", fixture.GetInputFilepath())
			t.Setenv("OUTPUT_LOCATION", fixture.GetOutputFilepath())

			Convey(fmt.Sprintf("%s", fixture.GetInputFilepath()), func() {
				runner.Execute(pluginVersion, buildDate, fixture.Runner)
				Convey("output file should contain expected value", func() {
					output, outErr := fixture.LoadOutputResult()
					if outErr != nil {
						fmt.Println("error loading output result", outErr)
					}
					for _, assert := range fixture.asserts {
						So(strings.TrimSpace(string(output)), assert, strings.TrimSpace(string(fixture.Expected)))
					}
				})
				Reset(func() {
					err := os.Remove(fixture.GetOutputFilepath())
					if err != nil {
						fmt.Printf("Test clean-up failed for %s, %v\n", fixture.GetOutputFilepath(), err)
					}
				})
			})

		}
	})
}
