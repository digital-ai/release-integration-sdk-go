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

// DeploymentRecordTaskInfo represents properties of deployment needed for audit record.
type DeploymentRecordTaskInfo struct {
	TaskId                string
	DeploymentPackage     string
	DeploymentEnvironment string
	DeploymentVersion     string
	DeploymentApplication string
	DeployedApplication   string
	Username              string
	ReleaseTaskId         string
}

// ReportingRecord represents the record produced by task.
type ReportingRecord struct {
	Id                 string    `json:"id"`
	TargetId           string    `json:"targetId,omitempty"`
	ServerUrl          string    `json:"serverUrl,omitempty"`
	ServerUser         string    `json:"serverUser,omitempty"`
	CreationDate       time.Time `json:"creationDate,omitempty"`
	RetryAttemptNumber int32     `json:"retryAttemptNumber,omitempty"`
	CreatedViaApi      bool      `json:"createdViaApi,omitempty"`
}

// ItsmReportingRecord represents the record for ITSM.
type ItsmReportingRecord struct {
	ReportingRecord
	Record    string `json:"record,omitempty"`
	RecordUrl string `json:"record_url,omitempty"`
	Title     string `json:"title,omitempty"`
	Status    string `json:"status,omitempty"`
	Priority  string `json:"priority,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
}

// PlanRecord represents the record for plan.
type PlanRecord struct {
	ReportingRecord
	Ticket      string    `json:"ticket,omitempty"`
	TicketUrl   string    `json:"ticket_url,omitempty"`
	Title       string    `json:"title,omitempty"`
	TicketType  string    `json:"ticketType,omitempty"`
	Status      string    `json:"status,omitempty"`
	UpdatedDate time.Time `json:"updatedDate,omitempty"`
	UpdatedBy   string    `json:"updatedBy,omitempty"`
}

// BuildRecord represents the record for build.
type BuildRecord struct {
	ReportingRecord
	Project   string    `json:"project,omitempty"`
	Build     string    `json:"build,omitempty"`
	Outcome   string    `json:"outcome,omitempty"`
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty"`
	Duration  string    `json:"duration,omitempty"`
	BuildUrl  string    `json:"build_url,omitempty"`
}

// DeploymentRecord represents the record for deployment.
type DeploymentRecord struct {
	ReportingRecord
	Type              string           `json:"type"`
	DeploymentTask    string           `json:"deploymentTask,omitempty"`
	DeploymentTaskUrl string           `json:"deploymentTask_url,omitempty"`
	ApplicationName   string           `json:"applicationName,omitempty"`
	Version           string           `json:"version,omitempty"`
	EnvironmentName   string           `json:"environmentName,omitempty"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

// CodeComplianceRecord represents the record for code compliance.
type CodeComplianceRecord struct {
	ReportingRecord
	Project        string    `json:"project,omitempty"`
	ProjectUrl     string    `json:"project_url,omitempty"`
	AnalysisDate   time.Time `json:"analysisDate,omitempty"`
	Outcome        string    `json:"outcome,omitempty"`
	ComplianceData string    `json:"complianceData,omitempty"`
}

// DeploymentStatus represents the status of the deployment.
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
