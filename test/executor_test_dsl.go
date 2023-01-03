package test

import (
	"fmt"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/goconvey/convey"
	"github.com/xebialabs/go-remote-runner-wrapper/runner"
	"io"
	"os"
	"path"
)

type ExecutorTestConfig struct {
	path             string
	inputFilename    string
	outputFilename   string
	expectedFilename string
	Runner           runner.Runner
}

type ExecutorTest struct {
	ExecutorTestConfig
	Input    []byte
	Expected []byte
	asserts  []convey.Assertion
}

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

func (fixture *ExecutorTest) loadInputResult() error {
	input, err := loadFromLocation(fixture.GetInputFilepath())
	if err != nil {
		return err
	}
	fixture.Input = input
	return nil
}

func (fixture *ExecutorTest) loadExpectedResult() error {
	expected, err := loadFromLocation(path.Join(fixture.path, fixture.expectedFilename))
	if err != nil {
		return err
	}
	fixture.Expected = expected
	return nil
}

func (fixture ExecutorTest) LoadOutputResult() ([]byte, error) {
	return loadFromLocation(fixture.GetOutputFilepath())
}

func (fixture ExecutorTest) GetInputFilepath() string {
	return path.Join(fixture.path, fixture.inputFilename)
}

func (fixture ExecutorTest) GetOutputFilepath() string {
	return path.Join(fixture.path, fixture.outputFilename)
}

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

func (fixture *ExecutorTest) WithInputFile(filename string) *ExecutorTest {
	fixture.ExecutorTestConfig.inputFilename = filename
	return fixture
}

func (fixture *ExecutorTest) WithOutputFile(filename string) *ExecutorTest {
	fixture.ExecutorTestConfig.outputFilename = filename
	return fixture
}

func (fixture *ExecutorTest) WithExpectedOutputFile(filename string) *ExecutorTest {
	fixture.ExecutorTestConfig.expectedFilename = filename
	return fixture
}

func (fixture *ExecutorTest) WithRunnerFunction(function runner.Runner) *ExecutorTest {
	fixture.Runner = function
	return fixture
}

func (fixture *ExecutorTest) WithAssert(assert convey.Assertion) *ExecutorTest {
	fixture.asserts = append(fixture.asserts, assert)
	return fixture
}

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
