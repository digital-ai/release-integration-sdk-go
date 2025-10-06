package test

import (
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/runner"
	"github.com/smarty/assertions"
	"github.com/smartystreets/goconvey/convey"
	"io"
	"os"
	"path"
)

// ExecutorTestConfig holds the configuration for an executor test.
type ExecutorTestConfig struct {
	path             string
	inputFilename    string
	outputFilename   string
	expectedFilename string
	Runner           runner.Runner
}

// ExecutorTest represents an executor test case.
type ExecutorTest struct {
	ExecutorTestConfig
	Input    []byte
	Expected []byte
	asserts  []convey.Assertion
}

// loadFromLocation loads file content from the specified location.
func loadFromLocation(location string) ([]byte, error) {
	inputContent, err := os.Open(location)
	if err != nil {
		fmt.Printf("Error loading file %s", location)
		return nil, err
	}
	defer inputContent.Close()
	content, _ := io.ReadAll(inputContent)
	return content, nil
}

// loadInputResult loads the input data for the test case.
func (fixture *ExecutorTest) loadInputResult() error {
	input, err := loadFromLocation(fixture.GetInputFilepath())
	if err != nil {
		return err
	}
	fixture.Input = input
	return nil
}

// loadExpectedResult loads the expected output for the test case.
func (fixture *ExecutorTest) loadExpectedResult() error {
	expected, err := loadFromLocation(path.Join(fixture.path, fixture.expectedFilename))
	if err != nil {
		return err
	}
	fixture.Expected = expected
	return nil
}

// LoadOutputResult loads the output result from the test case.
func (fixture ExecutorTest) LoadOutputResult() ([]byte, error) {
	return loadFromLocation(fixture.GetOutputFilepath())
}

// GetInputFilepath returns the file path of the input file.
func (fixture ExecutorTest) GetInputFilepath() string {
	return path.Join(fixture.path, fixture.inputFilename)
}

// GetOutputFilepath returns the file path of the output file.
func (fixture ExecutorTest) GetOutputFilepath() string {
	return path.Join(fixture.path, fixture.outputFilename)
}

// CreateExecutorTest creates a new executor test case.
func CreateExecutorTest(path string) *ExecutorTest {
	return &ExecutorTest{
		ExecutorTestConfig: ExecutorTestConfig{
			path:             path,
			inputFilename:    "input.json",
			expectedFilename: "expected.json",
			outputFilename:   "output.json",
		},
	}
}

// WithInputFile sets the input file name for the executor test.
func (fixture *ExecutorTest) WithInputFile(filename string) *ExecutorTest {
	fixture.ExecutorTestConfig.inputFilename = filename
	return fixture
}

// WithOutputFile sets the output file name for the executor test.
func (fixture *ExecutorTest) WithOutputFile(filename string) *ExecutorTest {
	fixture.ExecutorTestConfig.outputFilename = filename
	return fixture
}

// WithExpectedOutputFile sets the expected output file name for the executor test.
func (fixture *ExecutorTest) WithExpectedOutputFile(filename string) *ExecutorTest {
	fixture.ExecutorTestConfig.expectedFilename = filename
	return fixture
}

// WithRunnerFunction sets the runner function for the executor test.
func (fixture *ExecutorTest) WithRunnerFunction(function runner.Runner) *ExecutorTest {
	fixture.Runner = function
	return fixture
}

// WithAssert adds an assertion for the executor test.
func (fixture *ExecutorTest) WithAssert(assert convey.Assertion) *ExecutorTest {
	fixture.asserts = append(fixture.asserts, assert)
	return fixture
}

// Build creates the executor test case.
func (fixture *ExecutorTest) Build() (*ExecutorTest, error) {
	err := fixture.loadInputResult()
	if err != nil {
		return nil, err
	}
	err = fixture.loadExpectedResult()
	if err != nil {
		return nil, err
	}
	if len(fixture.asserts) == 0 {
		fixture.asserts = append(fixture.asserts, assertions.ShouldEqual)
	}
	return fixture, nil
}
