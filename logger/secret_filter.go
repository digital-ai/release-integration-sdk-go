package logger

import (
	"github.com/digital-ai/release-integration-sdk-go/task"
	"strings"
)

// secretValue represents the masked value for secrets.
const secretValue = "************"

var secrets []string

// secretLogFilter provides methods for filtering and masking secrets in log messages.
type secretLogFilter struct{}

// Filter filters log message arguments.
func (f *secretLogFilter) Filter(args []interface{}) []interface{} {
	return args
}

// FilterF filters log message format and arguments.
func (f *secretLogFilter) FilterF(format string, args []interface{}) (string, []interface{}) {
	return format, filterSecrets(args)
}

// FilterS filters log message and key-value pairs.
func (f *secretLogFilter) FilterS(msg string, keysAndValues []interface{}) (string, []interface{}) {
	return msg, filterSecrets(keysAndValues)
}

// filterSecrets filters secrets in the given slice of arguments.
func filterSecrets(args []interface{}) []interface{} {
	for i, arg := range args {
		v, ok := arg.(string)
		if ok {
			args[i] = maskSecretsInString(v)
		}
	}
	return args
}

// maskSecretsInString masks secrets in the input string with secretValue.
func maskSecretsInString(input string) string {
	for _, secret := range secrets {
		if strings.Contains(input, secret) {
			input = strings.Replace(input, secret, secretValue, -1)
		}
	}
	return input
}

// AddSecrets adds secrets from the input task context to the secrets slice.
func AddSecrets(inputContext task.InputContext) {
	for _, property := range inputContext.Task.Properties {
		if property.IsPassword {
			secrets = append(secrets, string(property.Value))
		}
	}
}
