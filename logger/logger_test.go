package logger

import (
	"bytes"
	"flag"
	"fmt"
	"k8s.io/klog/v2"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	err := flag.Set("logtostderr", "false")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	tests := []struct {
		format         string
		args           interface{}
		expectedOutput string
	}{
		{
			format:         "my password is %s",
			args:           password,
			expectedOutput: fmt.Sprintf("my password is %s", secretValue),
		},
		{
			format:         "my secret is %s",
			args:           secret,
			expectedOutput: fmt.Sprintf("my secret is %s", secretValue),
		},
	}

	buf := new(bytes.Buffer)
	klog.SetOutput(buf)

	for _, test := range tests {
		buf.Reset()

		klog.Infof(test.format, test.args)

		actualOutput := buf.String()
		if !strings.Contains(actualOutput, test.expectedOutput) {
			t.Errorf("klog.Infof(%q, %v) = %q, expectedFormat %q", test.format, test.args, actualOutput, test.expectedOutput)
		}
	}
}
