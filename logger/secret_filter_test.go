package logger

import (
	"encoding/json"
	"github.com/xebialabs/go-remote-runner-wrapper/task"
	"os"
	"reflect"
	"testing"
)

var (
	filter   = &secretLogFilter{}
	password = "P@55w0Rd"
	secret   = "5€cr3t"
)

func TestMain(m *testing.M) {
	AddSecrets(task.InputContext{Task: task.TaskContext{
		Properties: []task.PropertyDefinition{
			{
				Name:       "password",
				Value:      json.RawMessage(password),
				IsPassword: true,
			},
			{
				Name:       "secret",
				Value:      json.RawMessage(secret),
				IsPassword: true,
			}},
	}})

	code := m.Run()
	os.Exit(code)
}
func TestAddSecrets(t *testing.T) {
	expectedSecrets := []string{password, secret}
	if !reflect.DeepEqual(secrets, expectedSecrets) {
		t.Errorf("secrets = %q, expectedSecrets %q", secrets, expectedSecrets)
	}
}

func TestSecretLogFilter(t *testing.T) {
	tests := []struct {
		format         string
		args           []interface{}
		expectedFormat string
		expectedArgs   []interface{}
	}{
		{
			format:         "my password is %s",
			args:           []interface{}{password},
			expectedFormat: "my password is %s",
			expectedArgs:   []interface{}{secretValue},
		},
		{
			format:         "my secret is %s",
			args:           []interface{}{secret},
			expectedFormat: "my secret is %s",
			expectedArgs:   []interface{}{secretValue},
		},
	}

	for _, test := range tests {
		actualFormat, actualArgs := filter.FilterF(test.format, test.args)
		if actualFormat != test.expectedFormat {
			t.Errorf("secretLogFilter.FilterF(%q, %v) = %q, expectedFormat %q", test.format, test.args, actualFormat, test.expectedFormat)
		} else if !reflect.DeepEqual(actualArgs, test.expectedArgs) {
			t.Errorf("secretLogFilter.FilterF(%q, %v) = _, %v, expectedFormat _, %v", test.format, test.args, actualArgs, test.expectedArgs)
		}
	}
}
