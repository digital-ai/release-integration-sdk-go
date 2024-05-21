package task

import (
	"strings"
	"time"
)

// DeploymentStatus represents the status of the deployment.
type DeploymentStatus string

// ReportingRecordType a string value of synthetic type for ReportingRecord types
type ReportingRecordType string

const (
	Planned    DeploymentStatus = "planned"
	InProgress DeploymentStatus = "in_progress"
	Completed  DeploymentStatus = "completed"
	Skipped    DeploymentStatus = "skipped"
	Failed     DeploymentStatus = "failed"
	Aborted    DeploymentStatus = "aborted"
)

const (
	DeploymentRecordType     ReportingRecordType = "udm.DeploymentRecord"
	BuildRecordType          ReportingRecordType = "udm.BuildRecord"
	ItsmRecordType           ReportingRecordType = "udm.ItsmRecord"
	PlanRecordType           ReportingRecordType = "udm.PlanRecord"
	CodeComplianceRecordType ReportingRecordType = "udm.CodeComplianceRecord"
)

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
	Id                 string              `json:"id"`
	Type               ReportingRecordType `json:"type"`
	TargetId           string              `json:"targetId,omitempty"`
	ServerUrl          string              `json:"serverUrl,omitempty"`
	ServerUser         string              `json:"serverUser,omitempty"`
	CreationDate       time.Time           `json:"creationDate,omitempty"`
	RetryAttemptNumber int32               `json:"retryAttemptNumber,omitempty"`
	CreatedViaApi      bool                `json:"createdViaApi,omitempty"`
}

// ReportingRecordBuilder builder helper type for ReportingRecord.
type ReportingRecordBuilder struct {
	record ReportingRecord
}

// NewReportingRecord create builder with empty ReportingRecord.
func NewReportingRecord() *ReportingRecordBuilder {
	return &ReportingRecordBuilder{record: ReportingRecord{}}
}

// Id set Id to ReportingRecord.
func (b *ReportingRecordBuilder) Id(id string) *ReportingRecordBuilder {
	b.record.Id = id
	return b
}

// TargetId set TargetId to ReportingRecord.
func (b *ReportingRecordBuilder) TargetId(targetId string) *ReportingRecordBuilder {
	b.record.TargetId = targetId
	return b
}

// ServerUrl set ServerUrl to ReportingRecord.
func (b *ReportingRecordBuilder) ServerUrl(serverUrl string) *ReportingRecordBuilder {
	b.record.ServerUrl = serverUrl
	return b
}

// ServerUser set ServerUser to ReportingRecord.
func (b *ReportingRecordBuilder) ServerUser(serverUser string) *ReportingRecordBuilder {
	b.record.ServerUser = serverUser
	return b
}

// CreationDate set CreationDate to ReportingRecord.
func (b *ReportingRecordBuilder) CreationDate(date time.Time) *ReportingRecordBuilder {
	b.record.CreationDate = date
	return b
}

// RetryAttemptNumber set number of retry attempts to ReportingRecord.
func (b *ReportingRecordBuilder) RetryAttemptNumber(retry int32) *ReportingRecordBuilder {
	b.record.RetryAttemptNumber = retry
	return b
}

// CreatedViaApi mark flag if the record was created via API to ReportingRecord.
func (b *ReportingRecordBuilder) CreatedViaApi(apiFlag bool) *ReportingRecordBuilder {
	b.record.CreatedViaApi = apiFlag
	return b
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

type ItsmReportingRecordBuilder struct {
	record *ItsmReportingRecord
}

// ItsmReportingRecord return empty ItsmReportingRecord for ReportingRecord.
func (b *ReportingRecordBuilder) ItsmReportingRecord() *ItsmReportingRecordBuilder {
	b.record.Type = ItsmRecordType
	return &ItsmReportingRecordBuilder{
		&ItsmReportingRecord{
			ReportingRecord: b.record,
		},
	}
}

// Record set Record field to ItsmReportingRecord.
func (r *ItsmReportingRecordBuilder) Record(record string) *ItsmReportingRecordBuilder {
	r.record.Record = record
	return r
}

// RecordUrl set RecordUrl field to ItsmReportingRecord.
func (r *ItsmReportingRecordBuilder) RecordUrl(recordUrl string) *ItsmReportingRecordBuilder {
	r.record.RecordUrl = recordUrl
	return r
}

// Title set Title field to ItsmReportingRecord.
func (r *ItsmReportingRecordBuilder) Title(title string) *ItsmReportingRecordBuilder {
	r.record.Title = title
	return r
}

// Status set Status field to ItsmReportingRecord.
func (r *ItsmReportingRecordBuilder) Status(status string) *ItsmReportingRecordBuilder {
	r.record.Status = status
	return r
}

// Priority set Priority field to ItsmReportingRecord.
func (r *ItsmReportingRecordBuilder) Priority(priority string) *ItsmReportingRecordBuilder {
	r.record.Priority = priority
	return r
}

// CreatedBy set CreatedBy field to ItsmReportingRecord.
func (r *ItsmReportingRecordBuilder) CreatedBy(createdBy string) *ItsmReportingRecordBuilder {
	r.record.CreatedBy = createdBy
	return r
}

// Build builds ItsmReportingRecord.
func (r *ItsmReportingRecordBuilder) Build() *ItsmReportingRecord {
	return r.record
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

type PlanRecordBuilder struct {
	record *PlanRecord
}

// PlanRecord create empty PlanRecord.
func (b *ReportingRecordBuilder) PlanRecord() *PlanRecordBuilder {
	b.record.Type = PlanRecordType
	return &PlanRecordBuilder{
		&PlanRecord{
			ReportingRecord: b.record,
		},
	}
}

// TicketUrl set TicketUrl field to PlanRecord.
func (r *PlanRecordBuilder) TicketUrl(ticketUrl string) *PlanRecordBuilder {
	r.record.TicketUrl = ticketUrl
	return r
}

// TicketType set TicketType field to PlanRecord.
func (r *PlanRecordBuilder) TicketType(ticketType string) *PlanRecordBuilder {
	r.record.TicketType = ticketType
	return r
}

// Ticket set Ticket field to PlanRecord.
func (r *PlanRecordBuilder) Ticket(ticket string) *PlanRecordBuilder {
	r.record.Ticket = ticket
	return r
}

// Title set Title field to PlanRecord.
func (r *PlanRecordBuilder) Title(ticket string) *PlanRecordBuilder {
	r.record.Title = ticket
	return r
}

// Status set Status field to PlanRecord.
func (r *PlanRecordBuilder) Status(status string) *PlanRecordBuilder {
	r.record.Status = status
	return r
}

// UpdatedBy set UpdatedBy field to PlanRecord.
func (r *PlanRecordBuilder) UpdatedBy(updatedBy string) *PlanRecordBuilder {
	r.record.UpdatedBy = updatedBy
	return r
}

// UpdatedDate set UpdatedDate field to PlanRecord.
func (r *PlanRecordBuilder) UpdatedDate(updatedDate time.Time) *PlanRecordBuilder {
	r.record.UpdatedDate = updatedDate
	return r
}

// Build builds PlanRecord.
func (r *PlanRecordBuilder) Build() *PlanRecord {
	return r.record
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

type BuildRecordBuilder struct {
	record *BuildRecord
}

// BuildRecord create empty BuildRecord.
func (b *ReportingRecordBuilder) BuildRecord() *BuildRecordBuilder {
	b.record.Type = BuildRecordType
	return &BuildRecordBuilder{
		&BuildRecord{
			ReportingRecord: b.record,
		},
	}
}

// Project set Project field to BuildRecord.
func (r *BuildRecordBuilder) Project(project string) *BuildRecordBuilder {
	r.record.Project = project
	return r
}

// BuildLabel set Build field to BuildRecord.
func (r *BuildRecordBuilder) BuildLabel(build string) *BuildRecordBuilder {
	r.record.Build = build
	return r
}

// Outcome set Outcome field to BuildRecord.
func (r *BuildRecordBuilder) Outcome(outcome string) *BuildRecordBuilder {
	r.record.Outcome = outcome
	return r
}

// StartDate set StartDate field to BuildRecord.
func (r *BuildRecordBuilder) StartDate(startDate time.Time) *BuildRecordBuilder {
	r.record.StartDate = startDate
	return r
}

// EndDate set EndDate field to BuildRecord.
func (r *BuildRecordBuilder) EndDate(endDate time.Time) *BuildRecordBuilder {
	r.record.EndDate = endDate
	return r
}

// Duration set Duration field to BuildRecord.
func (r *BuildRecordBuilder) Duration(duration string) *BuildRecordBuilder {
	r.record.Duration = duration
	return r
}

// BuildUrl set Duration field to BuildRecord.
func (r *BuildRecordBuilder) BuildUrl(buildUrl string) *BuildRecordBuilder {
	r.record.BuildUrl = buildUrl
	return r
}

// Build builds PlanRecord.
func (r *BuildRecordBuilder) Build() *BuildRecord {
	return r.record
}

// DeploymentRecord represents the record for deployment.
type DeploymentRecord struct {
	ReportingRecord
	DeploymentTask    string           `json:"deploymentTask,omitempty"`
	DeploymentTaskUrl string           `json:"deploymentTask_url,omitempty"`
	ApplicationName   string           `json:"applicationName,omitempty"`
	Version           string           `json:"version,omitempty"`
	EnvironmentName   string           `json:"environmentName,omitempty"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

type DeploymentRecordBuilder struct {
	record                *DeploymentRecord
	taskId                string
	serverUrl             string
	deploymentPackage     string
	deployedApplication   string
	deploymentVersion     string
	deploymentEnvironment string
}

// DeploymentRecord create empty DeploymentRecord.
func (b *ReportingRecordBuilder) DeploymentRecord() *DeploymentRecordBuilder {
	b.record.Type = DeploymentRecordType
	return &DeploymentRecordBuilder{
		record: &DeploymentRecord{
			ReportingRecord: b.record,
		},
	}
}

// TaskId set DeploymentTask field to DeploymentRecord, if taskId parameter is empty string "None" is used as default value.
func (r *DeploymentRecordBuilder) TaskId(taskId string) *DeploymentRecordBuilder {
	r.taskId = taskId
	return r
}

// DeploymentStatus set Status field to DeploymentRecord.
func (r *DeploymentRecordBuilder) DeploymentStatus(status DeploymentStatus) *DeploymentRecordBuilder {
	r.record.Status = status
	return r
}

// DeployedApplication set DeployedApplication, DeployedApplication is used as default fallback value.
func (r *DeploymentRecordBuilder) DeployedApplication(deployedApplication string) *DeploymentRecordBuilder {
	r.deployedApplication = deployedApplication
	return r
}

// DeploymentPackage set DeploymentPackage, if DeploymentPackage is not empty it will use extracted application name, otherwise deployedApplication is used as default fallback value.
func (r *DeploymentRecordBuilder) DeploymentPackage(deploymentPackage string) *DeploymentRecordBuilder {
	r.deploymentPackage = deploymentPackage
	return r
}

// DeploymentEnvironment set DeploymentEnvironment, if deployment package is not empty it will use DeploymentEnvironment value, otherwise deployedApplication will be used to extract environment name from parent path.
func (r *DeploymentRecordBuilder) DeploymentEnvironment(deploymentEnvironment string) *DeploymentRecordBuilder {
	r.deploymentEnvironment = deploymentEnvironment
	return r
}

// ApplicationVersion set Version field to DeploymentRecord.
func (r *DeploymentRecordBuilder) ApplicationVersion(deploymentVersion string) *DeploymentRecordBuilder {
	r.record.Version = deploymentVersion
	return r
}

// Build builds CodeComplianceRecord.
func (r *DeploymentRecordBuilder) Build() *DeploymentRecord {
	if len(r.taskId) > 0 {
		if len(r.record.ServerUrl) > 0 {
			r.record.DeploymentTaskUrl = strings.TrimRight(r.serverUrl, "/") + "/#/task/" + r.taskId
		} else {
			r.record.DeploymentTaskUrl = ""
		}
		r.record.DeploymentTask = r.taskId
	} else {
		r.record.DeploymentTask = "None"
	}

	if len(r.deploymentPackage) > 0 {
		if strings.HasSuffix(r.deploymentPackage, r.deploymentVersion) {
			r.record.ApplicationName = r.deploymentPackage[:len(r.deploymentPackage)-len(r.deploymentVersion)-1]
		} else {
			r.record.ApplicationName = r.deploymentPackage
		}
		r.record.EnvironmentName = r.deploymentEnvironment
	} else {
		r.record.ApplicationName = r.deployedApplication
		r.record.EnvironmentName = getCiParentPath(r.deployedApplication)
	}

	return r.record
}

func getCiParentPath(filePath string) string {
	lastSlashIndex := strings.LastIndex(filePath, "/")
	if lastSlashIndex == -1 {
		return ""
	}
	return filePath[:lastSlashIndex]
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

type CodeComplianceRecordBuilder struct {
	record *CodeComplianceRecord
}

func (b *ReportingRecordBuilder) CodeComplianceRecord() *CodeComplianceRecordBuilder {
	b.record.Type = CodeComplianceRecordType
	return &CodeComplianceRecordBuilder{
		&CodeComplianceRecord{
			ReportingRecord: b.record,
		},
	}
}

// ProjectUrl set ProjectUrl field to CodeComplianceRecord.
func (r *CodeComplianceRecordBuilder) ProjectUrl(projectUrl string) *CodeComplianceRecordBuilder {
	r.record.ProjectUrl = projectUrl
	return r
}

// Project set Project field to CodeComplianceRecord.
func (r *CodeComplianceRecordBuilder) Project(project string) *CodeComplianceRecordBuilder {
	r.record.Project = project
	return r
}

// Outcome set Outcome field to CodeComplianceRecord.
func (r *CodeComplianceRecordBuilder) Outcome(outcome string) *CodeComplianceRecordBuilder {
	r.record.Outcome = outcome
	return r
}

// ComplianceData set ComplianceData field to CodeComplianceRecord.
func (r *CodeComplianceRecordBuilder) ComplianceData(data string) *CodeComplianceRecordBuilder {
	r.record.ComplianceData = data
	return r
}

// AnalysisDate set AnalysisDate field to CodeComplianceRecord.
func (r *CodeComplianceRecordBuilder) AnalysisDate(date time.Time) *CodeComplianceRecordBuilder {
	r.record.AnalysisDate = date
	return r
}

// Build builds CodeComplianceRecord.
func (r *CodeComplianceRecordBuilder) Build() *CodeComplianceRecord {
	return r.record
}
