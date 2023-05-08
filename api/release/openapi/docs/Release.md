# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**ScheduledStartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**DueDate** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PlannedDuration** | Pointer to **int32** |  | [optional] 
**FlagStatus** | Pointer to [**FlagStatus**](FlagStatus.md) |  | [optional] 
**FlagComment** | Pointer to **string** |  | [optional] 
**OverdueNotified** | Pointer to **bool** |  | [optional] 
**Flagged** | Pointer to **bool** |  | [optional] 
**StartOrScheduledDate** | Pointer to **time.Time** |  | [optional] 
**EndOrDueDate** | Pointer to **time.Time** |  | [optional] 
**Overdue** | Pointer to **bool** |  | [optional] 
**OrCalculateDueDate** | Pointer to **NullableTime** |  | [optional] 
**ComputedPlannedDuration** | Pointer to **map[string]interface{}** |  | [optional] 
**ActualDuration** | Pointer to **map[string]interface{}** |  | [optional] 
**RootReleaseId** | Pointer to **string** |  | [optional] 
**MaxConcurrentReleases** | Pointer to **int32** |  | [optional] 
**ReleaseTriggers** | Pointer to [**[]ReleaseTrigger**](ReleaseTrigger.md) |  | [optional] 
**Teams** | Pointer to [**[]Team**](Team.md) |  | [optional] 
**MemberViewers** | Pointer to **[]string** |  | [optional] 
**RoleViewers** | Pointer to **[]string** |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**Phases** | Pointer to [**[]Phase**](Phase.md) |  | [optional] 
**QueryableStartDate** | Pointer to **time.Time** |  | [optional] 
**QueryableEndDate** | Pointer to **time.Time** |  | [optional] 
**RealFlagStatus** | Pointer to [**FlagStatus**](FlagStatus.md) |  | [optional] 
**Status** | Pointer to [**ReleaseStatus**](ReleaseStatus.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**CalendarLinkToken** | Pointer to **string** |  | [optional] 
**CalendarPublished** | Pointer to **bool** |  | [optional] 
**Tutorial** | Pointer to **bool** |  | [optional] 
**AbortOnFailure** | Pointer to **bool** |  | [optional] 
**ArchiveRelease** | Pointer to **bool** |  | [optional] 
**AllowPasswordsInAllFields** | Pointer to **bool** |  | [optional] 
**DisableNotifications** | Pointer to **bool** |  | [optional] 
**AllowConcurrentReleasesFromTrigger** | Pointer to **bool** |  | [optional] 
**OriginTemplateId** | Pointer to **string** |  | [optional] 
**CreatedFromTrigger** | Pointer to **bool** |  | [optional] 
**ScriptUsername** | Pointer to **string** |  | [optional] 
**ScriptUserPassword** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to [**[]ReleaseExtension**](ReleaseExtension.md) |  | [optional] 
**StartedFromTaskId** | Pointer to **string** |  | [optional] 
**AutoStart** | Pointer to **bool** |  | [optional] 
**AutomatedResumeCount** | Pointer to **int32** |  | [optional] 
**MaxAutomatedResumes** | Pointer to **int32** |  | [optional] 
**AbortComment** | Pointer to **string** |  | [optional] 
**VariableMapping** | Pointer to **map[string]string** |  | [optional] 
**RiskProfile** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**CiUid** | Pointer to **int32** |  | [optional] 
**VariableValues** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**PasswordVariableValues** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CiPropertyVariables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**AllStringVariableValues** | Pointer to **map[string]string** |  | [optional] 
**AllReleaseGlobalAndFolderVariables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**AllVariableValuesAsStringsWithInterpolationInfo** | Pointer to [**map[string]ValueWithInterpolation**](ValueWithInterpolation.md) |  | [optional] 
**VariablesKeysInNonInterpolatableVariableValues** | Pointer to **[]string** |  | [optional] 
**VariablesByKeys** | Pointer to [**map[string]Variable**](Variable.md) |  | [optional] 
**AllVariables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**GlobalVariables** | Pointer to [**GlobalVariables**](GlobalVariables.md) |  | [optional] 
**FolderVariables** | Pointer to [**FolderVariables**](FolderVariables.md) |  | [optional] 
**AdminTeam** | Pointer to [**Team**](Team.md) |  | [optional] 
**ReleaseAttachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**CurrentPhase** | Pointer to [**Phase**](Phase.md) |  | [optional] 
**CurrentTask** | Pointer to [**Task**](Task.md) |  | [optional] 
**AllTasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**AllGates** | Pointer to [**[]GateTask**](GateTask.md) |  | [optional] 
**AllUserInputTasks** | Pointer to [**[]UserInputTask**](UserInputTask.md) |  | [optional] 
**Done** | Pointer to **bool** |  | [optional] 
**PlannedOrActive** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Defunct** | Pointer to **bool** |  | [optional] 
**Updatable** | Pointer to **bool** |  | [optional] 
**Aborted** | Pointer to **bool** |  | [optional] 
**Failing** | Pointer to **bool** |  | [optional] 
**Failed** | Pointer to **bool** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**Template** | Pointer to **bool** |  | [optional] 
**Planned** | Pointer to **bool** |  | [optional] 
**InProgress** | Pointer to **bool** |  | [optional] 
**Release** | Pointer to [**Release**](Release.md) |  | [optional] 
**ReleaseUid** | Pointer to **int32** |  | [optional] 
**DisplayPath** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**[]PlanItem**](PlanItem.md) |  | [optional] 
**AllPlanItems** | Pointer to [**[]PlanItem**](PlanItem.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ActiveTasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 
**Pending** | Pointer to **bool** |  | [optional] 

## Methods

### NewRelease

`func NewRelease() *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Release) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Release) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Release) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Release) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Release) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Release) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Release) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Release) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStartDate

`func (o *Release) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Release) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Release) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Release) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetScheduledStartDate

`func (o *Release) GetScheduledStartDate() time.Time`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *Release) GetScheduledStartDateOk() (*time.Time, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *Release) SetScheduledStartDate(v time.Time)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *Release) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Release) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Release) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Release) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Release) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDueDate

`func (o *Release) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Release) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Release) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Release) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetTitle

`func (o *Release) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Release) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Release) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Release) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Release) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Release) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Release) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Release) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *Release) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Release) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Release) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Release) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPlannedDuration

`func (o *Release) GetPlannedDuration() int32`

GetPlannedDuration returns the PlannedDuration field if non-nil, zero value otherwise.

### GetPlannedDurationOk

`func (o *Release) GetPlannedDurationOk() (*int32, bool)`

GetPlannedDurationOk returns a tuple with the PlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDuration

`func (o *Release) SetPlannedDuration(v int32)`

SetPlannedDuration sets PlannedDuration field to given value.

### HasPlannedDuration

`func (o *Release) HasPlannedDuration() bool`

HasPlannedDuration returns a boolean if a field has been set.

### GetFlagStatus

`func (o *Release) GetFlagStatus() FlagStatus`

GetFlagStatus returns the FlagStatus field if non-nil, zero value otherwise.

### GetFlagStatusOk

`func (o *Release) GetFlagStatusOk() (*FlagStatus, bool)`

GetFlagStatusOk returns a tuple with the FlagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagStatus

`func (o *Release) SetFlagStatus(v FlagStatus)`

SetFlagStatus sets FlagStatus field to given value.

### HasFlagStatus

`func (o *Release) HasFlagStatus() bool`

HasFlagStatus returns a boolean if a field has been set.

### GetFlagComment

`func (o *Release) GetFlagComment() string`

GetFlagComment returns the FlagComment field if non-nil, zero value otherwise.

### GetFlagCommentOk

`func (o *Release) GetFlagCommentOk() (*string, bool)`

GetFlagCommentOk returns a tuple with the FlagComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagComment

`func (o *Release) SetFlagComment(v string)`

SetFlagComment sets FlagComment field to given value.

### HasFlagComment

`func (o *Release) HasFlagComment() bool`

HasFlagComment returns a boolean if a field has been set.

### GetOverdueNotified

`func (o *Release) GetOverdueNotified() bool`

GetOverdueNotified returns the OverdueNotified field if non-nil, zero value otherwise.

### GetOverdueNotifiedOk

`func (o *Release) GetOverdueNotifiedOk() (*bool, bool)`

GetOverdueNotifiedOk returns a tuple with the OverdueNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueNotified

`func (o *Release) SetOverdueNotified(v bool)`

SetOverdueNotified sets OverdueNotified field to given value.

### HasOverdueNotified

`func (o *Release) HasOverdueNotified() bool`

HasOverdueNotified returns a boolean if a field has been set.

### GetFlagged

`func (o *Release) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *Release) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *Release) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *Release) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetStartOrScheduledDate

`func (o *Release) GetStartOrScheduledDate() time.Time`

GetStartOrScheduledDate returns the StartOrScheduledDate field if non-nil, zero value otherwise.

### GetStartOrScheduledDateOk

`func (o *Release) GetStartOrScheduledDateOk() (*time.Time, bool)`

GetStartOrScheduledDateOk returns a tuple with the StartOrScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOrScheduledDate

`func (o *Release) SetStartOrScheduledDate(v time.Time)`

SetStartOrScheduledDate sets StartOrScheduledDate field to given value.

### HasStartOrScheduledDate

`func (o *Release) HasStartOrScheduledDate() bool`

HasStartOrScheduledDate returns a boolean if a field has been set.

### GetEndOrDueDate

`func (o *Release) GetEndOrDueDate() time.Time`

GetEndOrDueDate returns the EndOrDueDate field if non-nil, zero value otherwise.

### GetEndOrDueDateOk

`func (o *Release) GetEndOrDueDateOk() (*time.Time, bool)`

GetEndOrDueDateOk returns a tuple with the EndOrDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOrDueDate

`func (o *Release) SetEndOrDueDate(v time.Time)`

SetEndOrDueDate sets EndOrDueDate field to given value.

### HasEndOrDueDate

`func (o *Release) HasEndOrDueDate() bool`

HasEndOrDueDate returns a boolean if a field has been set.

### GetOverdue

`func (o *Release) GetOverdue() bool`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *Release) GetOverdueOk() (*bool, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *Release) SetOverdue(v bool)`

SetOverdue sets Overdue field to given value.

### HasOverdue

`func (o *Release) HasOverdue() bool`

HasOverdue returns a boolean if a field has been set.

### GetOrCalculateDueDate

`func (o *Release) GetOrCalculateDueDate() time.Time`

GetOrCalculateDueDate returns the OrCalculateDueDate field if non-nil, zero value otherwise.

### GetOrCalculateDueDateOk

`func (o *Release) GetOrCalculateDueDateOk() (*time.Time, bool)`

GetOrCalculateDueDateOk returns a tuple with the OrCalculateDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrCalculateDueDate

`func (o *Release) SetOrCalculateDueDate(v time.Time)`

SetOrCalculateDueDate sets OrCalculateDueDate field to given value.

### HasOrCalculateDueDate

`func (o *Release) HasOrCalculateDueDate() bool`

HasOrCalculateDueDate returns a boolean if a field has been set.

### SetOrCalculateDueDateNil

`func (o *Release) SetOrCalculateDueDateNil(b bool)`

 SetOrCalculateDueDateNil sets the value for OrCalculateDueDate to be an explicit nil

### UnsetOrCalculateDueDate
`func (o *Release) UnsetOrCalculateDueDate()`

UnsetOrCalculateDueDate ensures that no value is present for OrCalculateDueDate, not even an explicit nil
### GetComputedPlannedDuration

`func (o *Release) GetComputedPlannedDuration() map[string]interface{}`

GetComputedPlannedDuration returns the ComputedPlannedDuration field if non-nil, zero value otherwise.

### GetComputedPlannedDurationOk

`func (o *Release) GetComputedPlannedDurationOk() (*map[string]interface{}, bool)`

GetComputedPlannedDurationOk returns a tuple with the ComputedPlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPlannedDuration

`func (o *Release) SetComputedPlannedDuration(v map[string]interface{})`

SetComputedPlannedDuration sets ComputedPlannedDuration field to given value.

### HasComputedPlannedDuration

`func (o *Release) HasComputedPlannedDuration() bool`

HasComputedPlannedDuration returns a boolean if a field has been set.

### GetActualDuration

`func (o *Release) GetActualDuration() map[string]interface{}`

GetActualDuration returns the ActualDuration field if non-nil, zero value otherwise.

### GetActualDurationOk

`func (o *Release) GetActualDurationOk() (*map[string]interface{}, bool)`

GetActualDurationOk returns a tuple with the ActualDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDuration

`func (o *Release) SetActualDuration(v map[string]interface{})`

SetActualDuration sets ActualDuration field to given value.

### HasActualDuration

`func (o *Release) HasActualDuration() bool`

HasActualDuration returns a boolean if a field has been set.

### GetRootReleaseId

`func (o *Release) GetRootReleaseId() string`

GetRootReleaseId returns the RootReleaseId field if non-nil, zero value otherwise.

### GetRootReleaseIdOk

`func (o *Release) GetRootReleaseIdOk() (*string, bool)`

GetRootReleaseIdOk returns a tuple with the RootReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootReleaseId

`func (o *Release) SetRootReleaseId(v string)`

SetRootReleaseId sets RootReleaseId field to given value.

### HasRootReleaseId

`func (o *Release) HasRootReleaseId() bool`

HasRootReleaseId returns a boolean if a field has been set.

### GetMaxConcurrentReleases

`func (o *Release) GetMaxConcurrentReleases() int32`

GetMaxConcurrentReleases returns the MaxConcurrentReleases field if non-nil, zero value otherwise.

### GetMaxConcurrentReleasesOk

`func (o *Release) GetMaxConcurrentReleasesOk() (*int32, bool)`

GetMaxConcurrentReleasesOk returns a tuple with the MaxConcurrentReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentReleases

`func (o *Release) SetMaxConcurrentReleases(v int32)`

SetMaxConcurrentReleases sets MaxConcurrentReleases field to given value.

### HasMaxConcurrentReleases

`func (o *Release) HasMaxConcurrentReleases() bool`

HasMaxConcurrentReleases returns a boolean if a field has been set.

### GetReleaseTriggers

`func (o *Release) GetReleaseTriggers() []ReleaseTrigger`

GetReleaseTriggers returns the ReleaseTriggers field if non-nil, zero value otherwise.

### GetReleaseTriggersOk

`func (o *Release) GetReleaseTriggersOk() (*[]ReleaseTrigger, bool)`

GetReleaseTriggersOk returns a tuple with the ReleaseTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTriggers

`func (o *Release) SetReleaseTriggers(v []ReleaseTrigger)`

SetReleaseTriggers sets ReleaseTriggers field to given value.

### HasReleaseTriggers

`func (o *Release) HasReleaseTriggers() bool`

HasReleaseTriggers returns a boolean if a field has been set.

### GetTeams

`func (o *Release) GetTeams() []Team`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Release) GetTeamsOk() (*[]Team, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Release) SetTeams(v []Team)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *Release) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetMemberViewers

`func (o *Release) GetMemberViewers() []string`

GetMemberViewers returns the MemberViewers field if non-nil, zero value otherwise.

### GetMemberViewersOk

`func (o *Release) GetMemberViewersOk() (*[]string, bool)`

GetMemberViewersOk returns a tuple with the MemberViewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberViewers

`func (o *Release) SetMemberViewers(v []string)`

SetMemberViewers sets MemberViewers field to given value.

### HasMemberViewers

`func (o *Release) HasMemberViewers() bool`

HasMemberViewers returns a boolean if a field has been set.

### GetRoleViewers

`func (o *Release) GetRoleViewers() []string`

GetRoleViewers returns the RoleViewers field if non-nil, zero value otherwise.

### GetRoleViewersOk

`func (o *Release) GetRoleViewersOk() (*[]string, bool)`

GetRoleViewersOk returns a tuple with the RoleViewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleViewers

`func (o *Release) SetRoleViewers(v []string)`

SetRoleViewers sets RoleViewers field to given value.

### HasRoleViewers

`func (o *Release) HasRoleViewers() bool`

HasRoleViewers returns a boolean if a field has been set.

### GetAttachments

`func (o *Release) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Release) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Release) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Release) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPhases

`func (o *Release) GetPhases() []Phase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *Release) GetPhasesOk() (*[]Phase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *Release) SetPhases(v []Phase)`

SetPhases sets Phases field to given value.

### HasPhases

`func (o *Release) HasPhases() bool`

HasPhases returns a boolean if a field has been set.

### GetQueryableStartDate

`func (o *Release) GetQueryableStartDate() time.Time`

GetQueryableStartDate returns the QueryableStartDate field if non-nil, zero value otherwise.

### GetQueryableStartDateOk

`func (o *Release) GetQueryableStartDateOk() (*time.Time, bool)`

GetQueryableStartDateOk returns a tuple with the QueryableStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryableStartDate

`func (o *Release) SetQueryableStartDate(v time.Time)`

SetQueryableStartDate sets QueryableStartDate field to given value.

### HasQueryableStartDate

`func (o *Release) HasQueryableStartDate() bool`

HasQueryableStartDate returns a boolean if a field has been set.

### GetQueryableEndDate

`func (o *Release) GetQueryableEndDate() time.Time`

GetQueryableEndDate returns the QueryableEndDate field if non-nil, zero value otherwise.

### GetQueryableEndDateOk

`func (o *Release) GetQueryableEndDateOk() (*time.Time, bool)`

GetQueryableEndDateOk returns a tuple with the QueryableEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryableEndDate

`func (o *Release) SetQueryableEndDate(v time.Time)`

SetQueryableEndDate sets QueryableEndDate field to given value.

### HasQueryableEndDate

`func (o *Release) HasQueryableEndDate() bool`

HasQueryableEndDate returns a boolean if a field has been set.

### GetRealFlagStatus

`func (o *Release) GetRealFlagStatus() FlagStatus`

GetRealFlagStatus returns the RealFlagStatus field if non-nil, zero value otherwise.

### GetRealFlagStatusOk

`func (o *Release) GetRealFlagStatusOk() (*FlagStatus, bool)`

GetRealFlagStatusOk returns a tuple with the RealFlagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealFlagStatus

`func (o *Release) SetRealFlagStatus(v FlagStatus)`

SetRealFlagStatus sets RealFlagStatus field to given value.

### HasRealFlagStatus

`func (o *Release) HasRealFlagStatus() bool`

HasRealFlagStatus returns a boolean if a field has been set.

### GetStatus

`func (o *Release) GetStatus() ReleaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Release) GetStatusOk() (*ReleaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Release) SetStatus(v ReleaseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Release) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Release) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Release) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Release) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Release) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVariables

`func (o *Release) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Release) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Release) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Release) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetCalendarLinkToken

`func (o *Release) GetCalendarLinkToken() string`

GetCalendarLinkToken returns the CalendarLinkToken field if non-nil, zero value otherwise.

### GetCalendarLinkTokenOk

`func (o *Release) GetCalendarLinkTokenOk() (*string, bool)`

GetCalendarLinkTokenOk returns a tuple with the CalendarLinkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarLinkToken

`func (o *Release) SetCalendarLinkToken(v string)`

SetCalendarLinkToken sets CalendarLinkToken field to given value.

### HasCalendarLinkToken

`func (o *Release) HasCalendarLinkToken() bool`

HasCalendarLinkToken returns a boolean if a field has been set.

### GetCalendarPublished

`func (o *Release) GetCalendarPublished() bool`

GetCalendarPublished returns the CalendarPublished field if non-nil, zero value otherwise.

### GetCalendarPublishedOk

`func (o *Release) GetCalendarPublishedOk() (*bool, bool)`

GetCalendarPublishedOk returns a tuple with the CalendarPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarPublished

`func (o *Release) SetCalendarPublished(v bool)`

SetCalendarPublished sets CalendarPublished field to given value.

### HasCalendarPublished

`func (o *Release) HasCalendarPublished() bool`

HasCalendarPublished returns a boolean if a field has been set.

### GetTutorial

`func (o *Release) GetTutorial() bool`

GetTutorial returns the Tutorial field if non-nil, zero value otherwise.

### GetTutorialOk

`func (o *Release) GetTutorialOk() (*bool, bool)`

GetTutorialOk returns a tuple with the Tutorial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTutorial

`func (o *Release) SetTutorial(v bool)`

SetTutorial sets Tutorial field to given value.

### HasTutorial

`func (o *Release) HasTutorial() bool`

HasTutorial returns a boolean if a field has been set.

### GetAbortOnFailure

`func (o *Release) GetAbortOnFailure() bool`

GetAbortOnFailure returns the AbortOnFailure field if non-nil, zero value otherwise.

### GetAbortOnFailureOk

`func (o *Release) GetAbortOnFailureOk() (*bool, bool)`

GetAbortOnFailureOk returns a tuple with the AbortOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortOnFailure

`func (o *Release) SetAbortOnFailure(v bool)`

SetAbortOnFailure sets AbortOnFailure field to given value.

### HasAbortOnFailure

`func (o *Release) HasAbortOnFailure() bool`

HasAbortOnFailure returns a boolean if a field has been set.

### GetArchiveRelease

`func (o *Release) GetArchiveRelease() bool`

GetArchiveRelease returns the ArchiveRelease field if non-nil, zero value otherwise.

### GetArchiveReleaseOk

`func (o *Release) GetArchiveReleaseOk() (*bool, bool)`

GetArchiveReleaseOk returns a tuple with the ArchiveRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRelease

`func (o *Release) SetArchiveRelease(v bool)`

SetArchiveRelease sets ArchiveRelease field to given value.

### HasArchiveRelease

`func (o *Release) HasArchiveRelease() bool`

HasArchiveRelease returns a boolean if a field has been set.

### GetAllowPasswordsInAllFields

`func (o *Release) GetAllowPasswordsInAllFields() bool`

GetAllowPasswordsInAllFields returns the AllowPasswordsInAllFields field if non-nil, zero value otherwise.

### GetAllowPasswordsInAllFieldsOk

`func (o *Release) GetAllowPasswordsInAllFieldsOk() (*bool, bool)`

GetAllowPasswordsInAllFieldsOk returns a tuple with the AllowPasswordsInAllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPasswordsInAllFields

`func (o *Release) SetAllowPasswordsInAllFields(v bool)`

SetAllowPasswordsInAllFields sets AllowPasswordsInAllFields field to given value.

### HasAllowPasswordsInAllFields

`func (o *Release) HasAllowPasswordsInAllFields() bool`

HasAllowPasswordsInAllFields returns a boolean if a field has been set.

### GetDisableNotifications

`func (o *Release) GetDisableNotifications() bool`

GetDisableNotifications returns the DisableNotifications field if non-nil, zero value otherwise.

### GetDisableNotificationsOk

`func (o *Release) GetDisableNotificationsOk() (*bool, bool)`

GetDisableNotificationsOk returns a tuple with the DisableNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotifications

`func (o *Release) SetDisableNotifications(v bool)`

SetDisableNotifications sets DisableNotifications field to given value.

### HasDisableNotifications

`func (o *Release) HasDisableNotifications() bool`

HasDisableNotifications returns a boolean if a field has been set.

### GetAllowConcurrentReleasesFromTrigger

`func (o *Release) GetAllowConcurrentReleasesFromTrigger() bool`

GetAllowConcurrentReleasesFromTrigger returns the AllowConcurrentReleasesFromTrigger field if non-nil, zero value otherwise.

### GetAllowConcurrentReleasesFromTriggerOk

`func (o *Release) GetAllowConcurrentReleasesFromTriggerOk() (*bool, bool)`

GetAllowConcurrentReleasesFromTriggerOk returns a tuple with the AllowConcurrentReleasesFromTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowConcurrentReleasesFromTrigger

`func (o *Release) SetAllowConcurrentReleasesFromTrigger(v bool)`

SetAllowConcurrentReleasesFromTrigger sets AllowConcurrentReleasesFromTrigger field to given value.

### HasAllowConcurrentReleasesFromTrigger

`func (o *Release) HasAllowConcurrentReleasesFromTrigger() bool`

HasAllowConcurrentReleasesFromTrigger returns a boolean if a field has been set.

### GetOriginTemplateId

`func (o *Release) GetOriginTemplateId() string`

GetOriginTemplateId returns the OriginTemplateId field if non-nil, zero value otherwise.

### GetOriginTemplateIdOk

`func (o *Release) GetOriginTemplateIdOk() (*string, bool)`

GetOriginTemplateIdOk returns a tuple with the OriginTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTemplateId

`func (o *Release) SetOriginTemplateId(v string)`

SetOriginTemplateId sets OriginTemplateId field to given value.

### HasOriginTemplateId

`func (o *Release) HasOriginTemplateId() bool`

HasOriginTemplateId returns a boolean if a field has been set.

### GetCreatedFromTrigger

`func (o *Release) GetCreatedFromTrigger() bool`

GetCreatedFromTrigger returns the CreatedFromTrigger field if non-nil, zero value otherwise.

### GetCreatedFromTriggerOk

`func (o *Release) GetCreatedFromTriggerOk() (*bool, bool)`

GetCreatedFromTriggerOk returns a tuple with the CreatedFromTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFromTrigger

`func (o *Release) SetCreatedFromTrigger(v bool)`

SetCreatedFromTrigger sets CreatedFromTrigger field to given value.

### HasCreatedFromTrigger

`func (o *Release) HasCreatedFromTrigger() bool`

HasCreatedFromTrigger returns a boolean if a field has been set.

### GetScriptUsername

`func (o *Release) GetScriptUsername() string`

GetScriptUsername returns the ScriptUsername field if non-nil, zero value otherwise.

### GetScriptUsernameOk

`func (o *Release) GetScriptUsernameOk() (*string, bool)`

GetScriptUsernameOk returns a tuple with the ScriptUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptUsername

`func (o *Release) SetScriptUsername(v string)`

SetScriptUsername sets ScriptUsername field to given value.

### HasScriptUsername

`func (o *Release) HasScriptUsername() bool`

HasScriptUsername returns a boolean if a field has been set.

### GetScriptUserPassword

`func (o *Release) GetScriptUserPassword() string`

GetScriptUserPassword returns the ScriptUserPassword field if non-nil, zero value otherwise.

### GetScriptUserPasswordOk

`func (o *Release) GetScriptUserPasswordOk() (*string, bool)`

GetScriptUserPasswordOk returns a tuple with the ScriptUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptUserPassword

`func (o *Release) SetScriptUserPassword(v string)`

SetScriptUserPassword sets ScriptUserPassword field to given value.

### HasScriptUserPassword

`func (o *Release) HasScriptUserPassword() bool`

HasScriptUserPassword returns a boolean if a field has been set.

### GetExtensions

`func (o *Release) GetExtensions() []ReleaseExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Release) GetExtensionsOk() (*[]ReleaseExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Release) SetExtensions(v []ReleaseExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Release) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetStartedFromTaskId

`func (o *Release) GetStartedFromTaskId() string`

GetStartedFromTaskId returns the StartedFromTaskId field if non-nil, zero value otherwise.

### GetStartedFromTaskIdOk

`func (o *Release) GetStartedFromTaskIdOk() (*string, bool)`

GetStartedFromTaskIdOk returns a tuple with the StartedFromTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedFromTaskId

`func (o *Release) SetStartedFromTaskId(v string)`

SetStartedFromTaskId sets StartedFromTaskId field to given value.

### HasStartedFromTaskId

`func (o *Release) HasStartedFromTaskId() bool`

HasStartedFromTaskId returns a boolean if a field has been set.

### GetAutoStart

`func (o *Release) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *Release) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *Release) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.

### HasAutoStart

`func (o *Release) HasAutoStart() bool`

HasAutoStart returns a boolean if a field has been set.

### GetAutomatedResumeCount

`func (o *Release) GetAutomatedResumeCount() int32`

GetAutomatedResumeCount returns the AutomatedResumeCount field if non-nil, zero value otherwise.

### GetAutomatedResumeCountOk

`func (o *Release) GetAutomatedResumeCountOk() (*int32, bool)`

GetAutomatedResumeCountOk returns a tuple with the AutomatedResumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedResumeCount

`func (o *Release) SetAutomatedResumeCount(v int32)`

SetAutomatedResumeCount sets AutomatedResumeCount field to given value.

### HasAutomatedResumeCount

`func (o *Release) HasAutomatedResumeCount() bool`

HasAutomatedResumeCount returns a boolean if a field has been set.

### GetMaxAutomatedResumes

`func (o *Release) GetMaxAutomatedResumes() int32`

GetMaxAutomatedResumes returns the MaxAutomatedResumes field if non-nil, zero value otherwise.

### GetMaxAutomatedResumesOk

`func (o *Release) GetMaxAutomatedResumesOk() (*int32, bool)`

GetMaxAutomatedResumesOk returns a tuple with the MaxAutomatedResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAutomatedResumes

`func (o *Release) SetMaxAutomatedResumes(v int32)`

SetMaxAutomatedResumes sets MaxAutomatedResumes field to given value.

### HasMaxAutomatedResumes

`func (o *Release) HasMaxAutomatedResumes() bool`

HasMaxAutomatedResumes returns a boolean if a field has been set.

### GetAbortComment

`func (o *Release) GetAbortComment() string`

GetAbortComment returns the AbortComment field if non-nil, zero value otherwise.

### GetAbortCommentOk

`func (o *Release) GetAbortCommentOk() (*string, bool)`

GetAbortCommentOk returns a tuple with the AbortComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortComment

`func (o *Release) SetAbortComment(v string)`

SetAbortComment sets AbortComment field to given value.

### HasAbortComment

`func (o *Release) HasAbortComment() bool`

HasAbortComment returns a boolean if a field has been set.

### GetVariableMapping

`func (o *Release) GetVariableMapping() map[string]string`

GetVariableMapping returns the VariableMapping field if non-nil, zero value otherwise.

### GetVariableMappingOk

`func (o *Release) GetVariableMappingOk() (*map[string]string, bool)`

GetVariableMappingOk returns a tuple with the VariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMapping

`func (o *Release) SetVariableMapping(v map[string]string)`

SetVariableMapping sets VariableMapping field to given value.

### HasVariableMapping

`func (o *Release) HasVariableMapping() bool`

HasVariableMapping returns a boolean if a field has been set.

### GetRiskProfile

`func (o *Release) GetRiskProfile() interface{}`

GetRiskProfile returns the RiskProfile field if non-nil, zero value otherwise.

### GetRiskProfileOk

`func (o *Release) GetRiskProfileOk() (*interface{}, bool)`

GetRiskProfileOk returns a tuple with the RiskProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskProfile

`func (o *Release) SetRiskProfile(v interface{})`

SetRiskProfile sets RiskProfile field to given value.

### HasRiskProfile

`func (o *Release) HasRiskProfile() bool`

HasRiskProfile returns a boolean if a field has been set.

### SetRiskProfileNil

`func (o *Release) SetRiskProfileNil(b bool)`

 SetRiskProfileNil sets the value for RiskProfile to be an explicit nil

### UnsetRiskProfile
`func (o *Release) UnsetRiskProfile()`

UnsetRiskProfile ensures that no value is present for RiskProfile, not even an explicit nil
### GetMetadata

`func (o *Release) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Release) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Release) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Release) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetArchived

`func (o *Release) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Release) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Release) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Release) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCiUid

`func (o *Release) GetCiUid() int32`

GetCiUid returns the CiUid field if non-nil, zero value otherwise.

### GetCiUidOk

`func (o *Release) GetCiUidOk() (*int32, bool)`

GetCiUidOk returns a tuple with the CiUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiUid

`func (o *Release) SetCiUid(v int32)`

SetCiUid sets CiUid field to given value.

### HasCiUid

`func (o *Release) HasCiUid() bool`

HasCiUid returns a boolean if a field has been set.

### GetVariableValues

`func (o *Release) GetVariableValues() map[string]map[string]interface{}`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *Release) GetVariableValuesOk() (*map[string]map[string]interface{}, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *Release) SetVariableValues(v map[string]map[string]interface{})`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *Release) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.

### GetPasswordVariableValues

`func (o *Release) GetPasswordVariableValues() map[string]map[string]interface{}`

GetPasswordVariableValues returns the PasswordVariableValues field if non-nil, zero value otherwise.

### GetPasswordVariableValuesOk

`func (o *Release) GetPasswordVariableValuesOk() (*map[string]map[string]interface{}, bool)`

GetPasswordVariableValuesOk returns a tuple with the PasswordVariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordVariableValues

`func (o *Release) SetPasswordVariableValues(v map[string]map[string]interface{})`

SetPasswordVariableValues sets PasswordVariableValues field to given value.

### HasPasswordVariableValues

`func (o *Release) HasPasswordVariableValues() bool`

HasPasswordVariableValues returns a boolean if a field has been set.

### GetCiPropertyVariables

`func (o *Release) GetCiPropertyVariables() []Variable`

GetCiPropertyVariables returns the CiPropertyVariables field if non-nil, zero value otherwise.

### GetCiPropertyVariablesOk

`func (o *Release) GetCiPropertyVariablesOk() (*[]Variable, bool)`

GetCiPropertyVariablesOk returns a tuple with the CiPropertyVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiPropertyVariables

`func (o *Release) SetCiPropertyVariables(v []Variable)`

SetCiPropertyVariables sets CiPropertyVariables field to given value.

### HasCiPropertyVariables

`func (o *Release) HasCiPropertyVariables() bool`

HasCiPropertyVariables returns a boolean if a field has been set.

### GetAllStringVariableValues

`func (o *Release) GetAllStringVariableValues() map[string]string`

GetAllStringVariableValues returns the AllStringVariableValues field if non-nil, zero value otherwise.

### GetAllStringVariableValuesOk

`func (o *Release) GetAllStringVariableValuesOk() (*map[string]string, bool)`

GetAllStringVariableValuesOk returns a tuple with the AllStringVariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStringVariableValues

`func (o *Release) SetAllStringVariableValues(v map[string]string)`

SetAllStringVariableValues sets AllStringVariableValues field to given value.

### HasAllStringVariableValues

`func (o *Release) HasAllStringVariableValues() bool`

HasAllStringVariableValues returns a boolean if a field has been set.

### GetAllReleaseGlobalAndFolderVariables

`func (o *Release) GetAllReleaseGlobalAndFolderVariables() []Variable`

GetAllReleaseGlobalAndFolderVariables returns the AllReleaseGlobalAndFolderVariables field if non-nil, zero value otherwise.

### GetAllReleaseGlobalAndFolderVariablesOk

`func (o *Release) GetAllReleaseGlobalAndFolderVariablesOk() (*[]Variable, bool)`

GetAllReleaseGlobalAndFolderVariablesOk returns a tuple with the AllReleaseGlobalAndFolderVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReleaseGlobalAndFolderVariables

`func (o *Release) SetAllReleaseGlobalAndFolderVariables(v []Variable)`

SetAllReleaseGlobalAndFolderVariables sets AllReleaseGlobalAndFolderVariables field to given value.

### HasAllReleaseGlobalAndFolderVariables

`func (o *Release) HasAllReleaseGlobalAndFolderVariables() bool`

HasAllReleaseGlobalAndFolderVariables returns a boolean if a field has been set.

### GetAllVariableValuesAsStringsWithInterpolationInfo

`func (o *Release) GetAllVariableValuesAsStringsWithInterpolationInfo() map[string]ValueWithInterpolation`

GetAllVariableValuesAsStringsWithInterpolationInfo returns the AllVariableValuesAsStringsWithInterpolationInfo field if non-nil, zero value otherwise.

### GetAllVariableValuesAsStringsWithInterpolationInfoOk

`func (o *Release) GetAllVariableValuesAsStringsWithInterpolationInfoOk() (*map[string]ValueWithInterpolation, bool)`

GetAllVariableValuesAsStringsWithInterpolationInfoOk returns a tuple with the AllVariableValuesAsStringsWithInterpolationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllVariableValuesAsStringsWithInterpolationInfo

`func (o *Release) SetAllVariableValuesAsStringsWithInterpolationInfo(v map[string]ValueWithInterpolation)`

SetAllVariableValuesAsStringsWithInterpolationInfo sets AllVariableValuesAsStringsWithInterpolationInfo field to given value.

### HasAllVariableValuesAsStringsWithInterpolationInfo

`func (o *Release) HasAllVariableValuesAsStringsWithInterpolationInfo() bool`

HasAllVariableValuesAsStringsWithInterpolationInfo returns a boolean if a field has been set.

### GetVariablesKeysInNonInterpolatableVariableValues

`func (o *Release) GetVariablesKeysInNonInterpolatableVariableValues() []string`

GetVariablesKeysInNonInterpolatableVariableValues returns the VariablesKeysInNonInterpolatableVariableValues field if non-nil, zero value otherwise.

### GetVariablesKeysInNonInterpolatableVariableValuesOk

`func (o *Release) GetVariablesKeysInNonInterpolatableVariableValuesOk() (*[]string, bool)`

GetVariablesKeysInNonInterpolatableVariableValuesOk returns a tuple with the VariablesKeysInNonInterpolatableVariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesKeysInNonInterpolatableVariableValues

`func (o *Release) SetVariablesKeysInNonInterpolatableVariableValues(v []string)`

SetVariablesKeysInNonInterpolatableVariableValues sets VariablesKeysInNonInterpolatableVariableValues field to given value.

### HasVariablesKeysInNonInterpolatableVariableValues

`func (o *Release) HasVariablesKeysInNonInterpolatableVariableValues() bool`

HasVariablesKeysInNonInterpolatableVariableValues returns a boolean if a field has been set.

### GetVariablesByKeys

`func (o *Release) GetVariablesByKeys() map[string]Variable`

GetVariablesByKeys returns the VariablesByKeys field if non-nil, zero value otherwise.

### GetVariablesByKeysOk

`func (o *Release) GetVariablesByKeysOk() (*map[string]Variable, bool)`

GetVariablesByKeysOk returns a tuple with the VariablesByKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesByKeys

`func (o *Release) SetVariablesByKeys(v map[string]Variable)`

SetVariablesByKeys sets VariablesByKeys field to given value.

### HasVariablesByKeys

`func (o *Release) HasVariablesByKeys() bool`

HasVariablesByKeys returns a boolean if a field has been set.

### GetAllVariables

`func (o *Release) GetAllVariables() []Variable`

GetAllVariables returns the AllVariables field if non-nil, zero value otherwise.

### GetAllVariablesOk

`func (o *Release) GetAllVariablesOk() (*[]Variable, bool)`

GetAllVariablesOk returns a tuple with the AllVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllVariables

`func (o *Release) SetAllVariables(v []Variable)`

SetAllVariables sets AllVariables field to given value.

### HasAllVariables

`func (o *Release) HasAllVariables() bool`

HasAllVariables returns a boolean if a field has been set.

### GetGlobalVariables

`func (o *Release) GetGlobalVariables() GlobalVariables`

GetGlobalVariables returns the GlobalVariables field if non-nil, zero value otherwise.

### GetGlobalVariablesOk

`func (o *Release) GetGlobalVariablesOk() (*GlobalVariables, bool)`

GetGlobalVariablesOk returns a tuple with the GlobalVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalVariables

`func (o *Release) SetGlobalVariables(v GlobalVariables)`

SetGlobalVariables sets GlobalVariables field to given value.

### HasGlobalVariables

`func (o *Release) HasGlobalVariables() bool`

HasGlobalVariables returns a boolean if a field has been set.

### GetFolderVariables

`func (o *Release) GetFolderVariables() FolderVariables`

GetFolderVariables returns the FolderVariables field if non-nil, zero value otherwise.

### GetFolderVariablesOk

`func (o *Release) GetFolderVariablesOk() (*FolderVariables, bool)`

GetFolderVariablesOk returns a tuple with the FolderVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderVariables

`func (o *Release) SetFolderVariables(v FolderVariables)`

SetFolderVariables sets FolderVariables field to given value.

### HasFolderVariables

`func (o *Release) HasFolderVariables() bool`

HasFolderVariables returns a boolean if a field has been set.

### GetAdminTeam

`func (o *Release) GetAdminTeam() Team`

GetAdminTeam returns the AdminTeam field if non-nil, zero value otherwise.

### GetAdminTeamOk

`func (o *Release) GetAdminTeamOk() (*Team, bool)`

GetAdminTeamOk returns a tuple with the AdminTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminTeam

`func (o *Release) SetAdminTeam(v Team)`

SetAdminTeam sets AdminTeam field to given value.

### HasAdminTeam

`func (o *Release) HasAdminTeam() bool`

HasAdminTeam returns a boolean if a field has been set.

### GetReleaseAttachments

`func (o *Release) GetReleaseAttachments() []Attachment`

GetReleaseAttachments returns the ReleaseAttachments field if non-nil, zero value otherwise.

### GetReleaseAttachmentsOk

`func (o *Release) GetReleaseAttachmentsOk() (*[]Attachment, bool)`

GetReleaseAttachmentsOk returns a tuple with the ReleaseAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseAttachments

`func (o *Release) SetReleaseAttachments(v []Attachment)`

SetReleaseAttachments sets ReleaseAttachments field to given value.

### HasReleaseAttachments

`func (o *Release) HasReleaseAttachments() bool`

HasReleaseAttachments returns a boolean if a field has been set.

### GetCurrentPhase

`func (o *Release) GetCurrentPhase() Phase`

GetCurrentPhase returns the CurrentPhase field if non-nil, zero value otherwise.

### GetCurrentPhaseOk

`func (o *Release) GetCurrentPhaseOk() (*Phase, bool)`

GetCurrentPhaseOk returns a tuple with the CurrentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPhase

`func (o *Release) SetCurrentPhase(v Phase)`

SetCurrentPhase sets CurrentPhase field to given value.

### HasCurrentPhase

`func (o *Release) HasCurrentPhase() bool`

HasCurrentPhase returns a boolean if a field has been set.

### GetCurrentTask

`func (o *Release) GetCurrentTask() Task`

GetCurrentTask returns the CurrentTask field if non-nil, zero value otherwise.

### GetCurrentTaskOk

`func (o *Release) GetCurrentTaskOk() (*Task, bool)`

GetCurrentTaskOk returns a tuple with the CurrentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTask

`func (o *Release) SetCurrentTask(v Task)`

SetCurrentTask sets CurrentTask field to given value.

### HasCurrentTask

`func (o *Release) HasCurrentTask() bool`

HasCurrentTask returns a boolean if a field has been set.

### GetAllTasks

`func (o *Release) GetAllTasks() []Task`

GetAllTasks returns the AllTasks field if non-nil, zero value otherwise.

### GetAllTasksOk

`func (o *Release) GetAllTasksOk() (*[]Task, bool)`

GetAllTasksOk returns a tuple with the AllTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTasks

`func (o *Release) SetAllTasks(v []Task)`

SetAllTasks sets AllTasks field to given value.

### HasAllTasks

`func (o *Release) HasAllTasks() bool`

HasAllTasks returns a boolean if a field has been set.

### GetAllGates

`func (o *Release) GetAllGates() []GateTask`

GetAllGates returns the AllGates field if non-nil, zero value otherwise.

### GetAllGatesOk

`func (o *Release) GetAllGatesOk() (*[]GateTask, bool)`

GetAllGatesOk returns a tuple with the AllGates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGates

`func (o *Release) SetAllGates(v []GateTask)`

SetAllGates sets AllGates field to given value.

### HasAllGates

`func (o *Release) HasAllGates() bool`

HasAllGates returns a boolean if a field has been set.

### GetAllUserInputTasks

`func (o *Release) GetAllUserInputTasks() []UserInputTask`

GetAllUserInputTasks returns the AllUserInputTasks field if non-nil, zero value otherwise.

### GetAllUserInputTasksOk

`func (o *Release) GetAllUserInputTasksOk() (*[]UserInputTask, bool)`

GetAllUserInputTasksOk returns a tuple with the AllUserInputTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUserInputTasks

`func (o *Release) SetAllUserInputTasks(v []UserInputTask)`

SetAllUserInputTasks sets AllUserInputTasks field to given value.

### HasAllUserInputTasks

`func (o *Release) HasAllUserInputTasks() bool`

HasAllUserInputTasks returns a boolean if a field has been set.

### GetDone

`func (o *Release) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *Release) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *Release) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *Release) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetPlannedOrActive

`func (o *Release) GetPlannedOrActive() bool`

GetPlannedOrActive returns the PlannedOrActive field if non-nil, zero value otherwise.

### GetPlannedOrActiveOk

`func (o *Release) GetPlannedOrActiveOk() (*bool, bool)`

GetPlannedOrActiveOk returns a tuple with the PlannedOrActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedOrActive

`func (o *Release) SetPlannedOrActive(v bool)`

SetPlannedOrActive sets PlannedOrActive field to given value.

### HasPlannedOrActive

`func (o *Release) HasPlannedOrActive() bool`

HasPlannedOrActive returns a boolean if a field has been set.

### GetActive

`func (o *Release) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Release) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Release) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Release) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefunct

`func (o *Release) GetDefunct() bool`

GetDefunct returns the Defunct field if non-nil, zero value otherwise.

### GetDefunctOk

`func (o *Release) GetDefunctOk() (*bool, bool)`

GetDefunctOk returns a tuple with the Defunct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunct

`func (o *Release) SetDefunct(v bool)`

SetDefunct sets Defunct field to given value.

### HasDefunct

`func (o *Release) HasDefunct() bool`

HasDefunct returns a boolean if a field has been set.

### GetUpdatable

`func (o *Release) GetUpdatable() bool`

GetUpdatable returns the Updatable field if non-nil, zero value otherwise.

### GetUpdatableOk

`func (o *Release) GetUpdatableOk() (*bool, bool)`

GetUpdatableOk returns a tuple with the Updatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatable

`func (o *Release) SetUpdatable(v bool)`

SetUpdatable sets Updatable field to given value.

### HasUpdatable

`func (o *Release) HasUpdatable() bool`

HasUpdatable returns a boolean if a field has been set.

### GetAborted

`func (o *Release) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *Release) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *Release) SetAborted(v bool)`

SetAborted sets Aborted field to given value.

### HasAborted

`func (o *Release) HasAborted() bool`

HasAborted returns a boolean if a field has been set.

### GetFailing

`func (o *Release) GetFailing() bool`

GetFailing returns the Failing field if non-nil, zero value otherwise.

### GetFailingOk

`func (o *Release) GetFailingOk() (*bool, bool)`

GetFailingOk returns a tuple with the Failing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailing

`func (o *Release) SetFailing(v bool)`

SetFailing sets Failing field to given value.

### HasFailing

`func (o *Release) HasFailing() bool`

HasFailing returns a boolean if a field has been set.

### GetFailed

`func (o *Release) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *Release) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *Release) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *Release) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetPaused

`func (o *Release) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *Release) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *Release) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *Release) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetTemplate

`func (o *Release) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Release) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Release) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Release) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetPlanned

`func (o *Release) GetPlanned() bool`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *Release) GetPlannedOk() (*bool, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *Release) SetPlanned(v bool)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *Release) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.

### GetInProgress

`func (o *Release) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *Release) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *Release) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *Release) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetRelease

`func (o *Release) GetRelease() Release`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *Release) GetReleaseOk() (*Release, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *Release) SetRelease(v Release)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *Release) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetReleaseUid

`func (o *Release) GetReleaseUid() int32`

GetReleaseUid returns the ReleaseUid field if non-nil, zero value otherwise.

### GetReleaseUidOk

`func (o *Release) GetReleaseUidOk() (*int32, bool)`

GetReleaseUidOk returns a tuple with the ReleaseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseUid

`func (o *Release) SetReleaseUid(v int32)`

SetReleaseUid sets ReleaseUid field to given value.

### HasReleaseUid

`func (o *Release) HasReleaseUid() bool`

HasReleaseUid returns a boolean if a field has been set.

### GetDisplayPath

`func (o *Release) GetDisplayPath() string`

GetDisplayPath returns the DisplayPath field if non-nil, zero value otherwise.

### GetDisplayPathOk

`func (o *Release) GetDisplayPathOk() (*string, bool)`

GetDisplayPathOk returns a tuple with the DisplayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPath

`func (o *Release) SetDisplayPath(v string)`

SetDisplayPath sets DisplayPath field to given value.

### HasDisplayPath

`func (o *Release) HasDisplayPath() bool`

HasDisplayPath returns a boolean if a field has been set.

### GetChildren

`func (o *Release) GetChildren() []PlanItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Release) GetChildrenOk() (*[]PlanItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Release) SetChildren(v []PlanItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Release) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetAllPlanItems

`func (o *Release) GetAllPlanItems() []PlanItem`

GetAllPlanItems returns the AllPlanItems field if non-nil, zero value otherwise.

### GetAllPlanItemsOk

`func (o *Release) GetAllPlanItemsOk() (*[]PlanItem, bool)`

GetAllPlanItemsOk returns a tuple with the AllPlanItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlanItems

`func (o *Release) SetAllPlanItems(v []PlanItem)`

SetAllPlanItems sets AllPlanItems field to given value.

### HasAllPlanItems

`func (o *Release) HasAllPlanItems() bool`

HasAllPlanItems returns a boolean if a field has been set.

### GetUrl

`func (o *Release) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Release) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Release) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Release) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetActiveTasks

`func (o *Release) GetActiveTasks() []Task`

GetActiveTasks returns the ActiveTasks field if non-nil, zero value otherwise.

### GetActiveTasksOk

`func (o *Release) GetActiveTasksOk() (*[]Task, bool)`

GetActiveTasksOk returns a tuple with the ActiveTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTasks

`func (o *Release) SetActiveTasks(v []Task)`

SetActiveTasks sets ActiveTasks field to given value.

### HasActiveTasks

`func (o *Release) HasActiveTasks() bool`

HasActiveTasks returns a boolean if a field has been set.

### GetVariableUsages

`func (o *Release) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *Release) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *Release) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *Release) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.

### GetPending

`func (o *Release) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *Release) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *Release) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *Release) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


