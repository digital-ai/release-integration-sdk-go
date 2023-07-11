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
	ReportingRecords interface{}            `json:"reportingRecords,omitempty"`
}

type ReportingRecord struct {
	TargetId           string    `json:"TargetId,omitempty"`
	ServerUrl          string    `json:"serverUrl,omitempty"`
	ServerUser         string    `json:"serverUser,omitempty"`
	CreationDate       time.Time `json:"creationDate,omitempty"`
	RetryAttemptNumber int32     `json:"retryAttemptNumber,omitempty"`
	CreatedViaApi      bool      `json:"createdViaApi,omitempty"`
}

type ItsmReportingRecord struct {
	ReportingRecord
	Record    string `json:"Record,omitempty"`
	RecordUrl string `json:"RecordUrl,omitempty"`
	Title     string `json:"Title,omitempty"`
	Status    string `json:"Status,omitempty"`
	Priority  string `json:"Priority,omitempty"`
	CreatedBy string `json:"CreatedBy,omitempty"`
}

type PlanRecord struct {
	ReportingRecord
	Ticket      string    `json:"Ticket,omitempty"`
	Title       string    `json:"Title,omitempty"`
	TicketType  string    `json:"TicketType,omitempty"`
	Status      string    `json:"Status,omitempty"`
	UpdatedDate time.Time `json:"UpdatedDate,omitempty"`
	UpdatedBy   string    `json:"UpdatedBy,omitempty"`
}

type BuildRecord struct {
	ReportingRecord
	Project   string    `json:"Project,omitempty"`
	Build     string    `json:"Build,omitempty"`
	Outcome   string    `json:"Outcome,omitempty"`
	StartDate time.Time `json:"StartDate,omitempty"`
	EndDate   time.Time `json:"EndDate,omitempty"`
	Duration  string    `json:"Duration,omitempty"`
	BuildUrl  string    `json:"BuildUrl,omitempty"`
}

type DeploymentRecord struct {
	ReportingRecord
	DeploymentTask    string           `json:"DeploymentTask,omitempty"`
	DeploymentTaskUrl string           `json:"DeploymentTaskUrl,omitempty"`
	ApplicationName   string           `json:"ApplicationName,omitempty"`
	Version           string           `json:"Version,omitempty"`
	EnvironmentName   string           `json:"EnvironmentName,omitempty"`
	Status            DeploymentStatus `json:"Status,omitempty"`
}

type CodeComplianceRecord struct {
	ReportingRecord
	Project        string    `json:"Project,omitempty"`
	ProjectUrl     string    `json:"ProjectUrl,omitempty"`
	AnalysisDate   time.Time `json:"AnalysisDate,omitempty"`
	Outcome        string    `json:"Outcome,omitempty"`
	ComplianceData string    `json:"ComplianceData,omitempty"`
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
