package task

import (
	"encoding/json"
)

// PropertyDefinition represents the definition of a task property.
type PropertyDefinition struct {
	Name       string          `json:"name"`
	Value      json.RawMessage `json:"value"`
	Kind       string          `json:"kind"`
	Category   string          `json:"category"`
	IsPassword bool            `json:"password"`
}

// TaskContext represents the context of a task.
type TaskContext struct {
	Id         string               `json:"id"`
	Type       string               `json:"type"`
	Properties []PropertyDefinition `json:"properties"`
}

// ComposedProperty represents a composed property of a task.
type ComposedProperty TaskContext

// AutomatedTaskAsUserContext represents the context of an automated task as a user with credentials.
type AutomatedTaskAsUserContext struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ReleaseContext represents the context of a release.
type ReleaseContext struct {
	Id                  string                     `json:"id"`
	AutomatedTaskAsUser AutomatedTaskAsUserContext `json:"automatedTaskAsUser"`
	Url                 string                     `json:"url"`
}

// InputContext represents the input context of a task.
type InputContext struct {
	Release ReleaseContext `json:"release"`
	Task    TaskContext    `json:"task"`
}

// TaskOutputContext represents the output context of a task.
type TaskOutputContext struct {
	ExitCode         int64                  `json:"exitCode"`
	OutputProperties map[string]interface{} `json:"outputProperties,omitempty"`
	JobErrorMessage  string                 `json:"jobErrorMessage,omitempty"`
	ReportingRecords interface{}            `json:"reportingRecords,omitempty"`
}

// LookupResultElement represents an element of a lookup result.
type LookupResultElement struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
