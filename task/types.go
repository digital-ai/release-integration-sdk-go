package task

import (
	"encoding/json"
	"time"
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
	ReportingRecords []ReportingRecord
}

type ReportingRecord struct {
	Id               string               `json:"id,omitempty"`
	Type             string               `json:"type,omitempty"`
	Properties       []PropertyDefinition `json:"properties,omitempty"`
	Scope            string               `json:"scope,omitempty"`
	TargetId         string               `json:"targetId,omitempty"`
	ConfigurationUri string               `json:"configurationUri,omitempty"`
	VariableMapping  map[string]string    `json:"variableMapping,omitempty"`
	//VariableUsages UsagePoint `json:"variableUsages,omitempty"`
	PropertiesWithVariables []interface{} `json:"propertiesWithVariables,omitempty"`
	ServerUrl               string        `json:"serverUrl,omitempty"`
	ServerUser              string        `json:"serverUser,omitempty"`
	CreationDate            time.Time     `json:"creationDate,omitempty"`
	RetryAttemptNumber      int32         `json:"retryAttemptNumber,omitempty"`
	CreatedViaApi           bool          `json:"createdViaApi,omitempty"`
}

type ItsmReportingRecord struct {
	ReportingRecord []ReportingRecord
	Record          string
	RecordUrl       string
	Title           string
	Status          string
	Priority        string
	CreatedBy       string
}

type PlanRecord struct {
	ReportingRecord []ReportingRecord
	Ticket          string
	Title           string
	TicketType      string
	Status          string
	UpdatedDate     time.Time
	UpdatedBy       string
}

type BuildRecord struct {
	ReportingRecord []ReportingRecord
	Project         string
	Build           string
	Outcome         string
	StartDate       time.Time
	EndDate         time.Time
	Duration        string
	BuildUrl        string
}

type DeploymentRecord struct {
	ReportingRecord   []ReportingRecord
	DeploymentTask    string
	DeploymentTaskUrl string
	ApplicationName   string
	Version           string
	EnvironmentName   string
	Status            DeploymentStatus
}

type CodeComplianceRecord struct {
	ReportingRecord []ReportingRecord
	Project         string
	ProjectUrl      string
	AnalysisDate    time.Time
	Outcome         string
	ComplianceData  string
}

type DeploymentStatus string

const (
	Planned    DeploymentStatus = "planned"
	InProgress DeploymentStatus = "in_progress"
	Completed  DeploymentStatus = "completed"
	Skipped    DeploymentStatus = "skipped"
	Failed     DeploymentStatus = "failed"
	Aborted    DeploymentStatus = "aborted"
)

// LookupResultElement represents an element of a lookup result.
type LookupResultElement struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
