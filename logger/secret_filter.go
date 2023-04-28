package logger

import (
	"github.com/digital-ai/release-integration-sdk-go/task"
	"strings"
)

const secretValue = "************"

var secrets []string

type secretLogFilter struct{}

func (f *secretLogFilter) Filter(args []interface{}) []interface{} {
	return args
}

func (f *secretLogFilter) FilterF(format string, args []interface{}) (string, []interface{}) {
	return format, filterSecrets(args)
}

func (f *secretLogFilter) FilterS(msg string, keysAndValues []interface{}) (string, []interface{}) {
	return msg, filterSecrets(keysAndValues)
}

func filterSecrets(args []interface{}) []interface{} {
	for i, arg := range args {
		v, ok := arg.(string)
		if ok {
			args[i] = maskSecretsInString(v)
		}
	}
	return args
}

func maskSecretsInString(input string) string {
	for _, secret := range secrets {
		if strings.Contains(input, secret) {
			input = strings.Replace(input, secret, secretValue, -1)
		}
	}
	return input
}

func AddSecrets(inputContext task.InputContext) {
	for _, property := range inputContext.Task.Properties {
		if property.IsPassword {
			secrets = append(secrets, string(property.Value))
		}
	}
}
