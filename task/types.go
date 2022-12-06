package task

import "encoding/json"

type PropertyDefinition struct {
	Name       string          `json:"name"`
	Value      json.RawMessage `json:"value"`
	Kind       string          `json:"kind"`
	Category   string          `json:"category"`
	IsPassword bool            `json:"isPassword"`
}

type TaskContext struct {
	Id         string               `json:"id"`
	Type       string               `json:"type"`
	Properties []PropertyDefinition `json:"properties"`
}

type ComposedProperty TaskContext

type AutomatedTaskAsUserContext struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReleaseContext struct {
	Id                  string                     `json:"id"`
	AutomatedTaskAsUser AutomatedTaskAsUserContext `json:"automatedTaskAsUser"`
	ServerUrl           string                     `json:"serverUrl"`
}

type InputContext struct {
	Release ReleaseContext `json:"release"`
	Task    TaskContext    `json:"task"`
}

type TaskOutputContext struct {
	ExitCode         int64                  `json:"exitCode"`
	OutputProperties map[string]interface{} `json:"outputProperties"`
}
