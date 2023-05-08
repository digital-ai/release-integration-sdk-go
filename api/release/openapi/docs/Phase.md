# Phase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**ScheduledStartDate** | Pointer to **time.Time** |  | [optional] 
**DueDate** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
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
**ReleaseUid** | Pointer to **int32** |  | [optional] 
**Tasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**Release** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to [**PhaseStatus**](PhaseStatus.md) |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**OriginId** | Pointer to **string** |  | [optional] 
**CurrentTask** | Pointer to [**Task**](Task.md) |  | [optional] 
**DisplayPath** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Done** | Pointer to **bool** |  | [optional] 
**Defunct** | Pointer to **bool** |  | [optional] 
**Updatable** | Pointer to **bool** |  | [optional] 
**Aborted** | Pointer to **bool** |  | [optional] 
**Planned** | Pointer to **bool** |  | [optional] 
**Failed** | Pointer to **bool** |  | [optional] 
**Failing** | Pointer to **bool** |  | [optional] 
**ReleaseOwner** | Pointer to **string** |  | [optional] 
**AllGates** | Pointer to [**[]GateTask**](GateTask.md) |  | [optional] 
**AllTasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**Children** | Pointer to [**[]PlanItem**](PlanItem.md) |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 
**Original** | Pointer to **bool** |  | [optional] 
**PhaseCopied** | Pointer to **bool** |  | [optional] 
**AncestorId** | Pointer to **string** |  | [optional] 
**LatestCopy** | Pointer to **bool** |  | [optional] 

## Methods

### NewPhase

`func NewPhase() *Phase`

NewPhase instantiates a new Phase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhaseWithDefaults

`func NewPhaseWithDefaults() *Phase`

NewPhaseWithDefaults instantiates a new Phase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Phase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Phase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Phase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Phase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Phase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Phase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Phase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Phase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLocked

`func (o *Phase) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Phase) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Phase) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Phase) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetTitle

`func (o *Phase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Phase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Phase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Phase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Phase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Phase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Phase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Phase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *Phase) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Phase) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Phase) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Phase) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetScheduledStartDate

`func (o *Phase) GetScheduledStartDate() time.Time`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *Phase) GetScheduledStartDateOk() (*time.Time, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *Phase) SetScheduledStartDate(v time.Time)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *Phase) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetDueDate

`func (o *Phase) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Phase) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Phase) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Phase) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *Phase) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Phase) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Phase) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Phase) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Phase) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Phase) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Phase) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Phase) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPlannedDuration

`func (o *Phase) GetPlannedDuration() int32`

GetPlannedDuration returns the PlannedDuration field if non-nil, zero value otherwise.

### GetPlannedDurationOk

`func (o *Phase) GetPlannedDurationOk() (*int32, bool)`

GetPlannedDurationOk returns a tuple with the PlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDuration

`func (o *Phase) SetPlannedDuration(v int32)`

SetPlannedDuration sets PlannedDuration field to given value.

### HasPlannedDuration

`func (o *Phase) HasPlannedDuration() bool`

HasPlannedDuration returns a boolean if a field has been set.

### GetFlagStatus

`func (o *Phase) GetFlagStatus() FlagStatus`

GetFlagStatus returns the FlagStatus field if non-nil, zero value otherwise.

### GetFlagStatusOk

`func (o *Phase) GetFlagStatusOk() (*FlagStatus, bool)`

GetFlagStatusOk returns a tuple with the FlagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagStatus

`func (o *Phase) SetFlagStatus(v FlagStatus)`

SetFlagStatus sets FlagStatus field to given value.

### HasFlagStatus

`func (o *Phase) HasFlagStatus() bool`

HasFlagStatus returns a boolean if a field has been set.

### GetFlagComment

`func (o *Phase) GetFlagComment() string`

GetFlagComment returns the FlagComment field if non-nil, zero value otherwise.

### GetFlagCommentOk

`func (o *Phase) GetFlagCommentOk() (*string, bool)`

GetFlagCommentOk returns a tuple with the FlagComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagComment

`func (o *Phase) SetFlagComment(v string)`

SetFlagComment sets FlagComment field to given value.

### HasFlagComment

`func (o *Phase) HasFlagComment() bool`

HasFlagComment returns a boolean if a field has been set.

### GetOverdueNotified

`func (o *Phase) GetOverdueNotified() bool`

GetOverdueNotified returns the OverdueNotified field if non-nil, zero value otherwise.

### GetOverdueNotifiedOk

`func (o *Phase) GetOverdueNotifiedOk() (*bool, bool)`

GetOverdueNotifiedOk returns a tuple with the OverdueNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueNotified

`func (o *Phase) SetOverdueNotified(v bool)`

SetOverdueNotified sets OverdueNotified field to given value.

### HasOverdueNotified

`func (o *Phase) HasOverdueNotified() bool`

HasOverdueNotified returns a boolean if a field has been set.

### GetFlagged

`func (o *Phase) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *Phase) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *Phase) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *Phase) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetStartOrScheduledDate

`func (o *Phase) GetStartOrScheduledDate() time.Time`

GetStartOrScheduledDate returns the StartOrScheduledDate field if non-nil, zero value otherwise.

### GetStartOrScheduledDateOk

`func (o *Phase) GetStartOrScheduledDateOk() (*time.Time, bool)`

GetStartOrScheduledDateOk returns a tuple with the StartOrScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOrScheduledDate

`func (o *Phase) SetStartOrScheduledDate(v time.Time)`

SetStartOrScheduledDate sets StartOrScheduledDate field to given value.

### HasStartOrScheduledDate

`func (o *Phase) HasStartOrScheduledDate() bool`

HasStartOrScheduledDate returns a boolean if a field has been set.

### GetEndOrDueDate

`func (o *Phase) GetEndOrDueDate() time.Time`

GetEndOrDueDate returns the EndOrDueDate field if non-nil, zero value otherwise.

### GetEndOrDueDateOk

`func (o *Phase) GetEndOrDueDateOk() (*time.Time, bool)`

GetEndOrDueDateOk returns a tuple with the EndOrDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOrDueDate

`func (o *Phase) SetEndOrDueDate(v time.Time)`

SetEndOrDueDate sets EndOrDueDate field to given value.

### HasEndOrDueDate

`func (o *Phase) HasEndOrDueDate() bool`

HasEndOrDueDate returns a boolean if a field has been set.

### GetOverdue

`func (o *Phase) GetOverdue() bool`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *Phase) GetOverdueOk() (*bool, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *Phase) SetOverdue(v bool)`

SetOverdue sets Overdue field to given value.

### HasOverdue

`func (o *Phase) HasOverdue() bool`

HasOverdue returns a boolean if a field has been set.

### GetOrCalculateDueDate

`func (o *Phase) GetOrCalculateDueDate() time.Time`

GetOrCalculateDueDate returns the OrCalculateDueDate field if non-nil, zero value otherwise.

### GetOrCalculateDueDateOk

`func (o *Phase) GetOrCalculateDueDateOk() (*time.Time, bool)`

GetOrCalculateDueDateOk returns a tuple with the OrCalculateDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrCalculateDueDate

`func (o *Phase) SetOrCalculateDueDate(v time.Time)`

SetOrCalculateDueDate sets OrCalculateDueDate field to given value.

### HasOrCalculateDueDate

`func (o *Phase) HasOrCalculateDueDate() bool`

HasOrCalculateDueDate returns a boolean if a field has been set.

### SetOrCalculateDueDateNil

`func (o *Phase) SetOrCalculateDueDateNil(b bool)`

 SetOrCalculateDueDateNil sets the value for OrCalculateDueDate to be an explicit nil

### UnsetOrCalculateDueDate
`func (o *Phase) UnsetOrCalculateDueDate()`

UnsetOrCalculateDueDate ensures that no value is present for OrCalculateDueDate, not even an explicit nil
### GetComputedPlannedDuration

`func (o *Phase) GetComputedPlannedDuration() map[string]interface{}`

GetComputedPlannedDuration returns the ComputedPlannedDuration field if non-nil, zero value otherwise.

### GetComputedPlannedDurationOk

`func (o *Phase) GetComputedPlannedDurationOk() (*map[string]interface{}, bool)`

GetComputedPlannedDurationOk returns a tuple with the ComputedPlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPlannedDuration

`func (o *Phase) SetComputedPlannedDuration(v map[string]interface{})`

SetComputedPlannedDuration sets ComputedPlannedDuration field to given value.

### HasComputedPlannedDuration

`func (o *Phase) HasComputedPlannedDuration() bool`

HasComputedPlannedDuration returns a boolean if a field has been set.

### GetActualDuration

`func (o *Phase) GetActualDuration() map[string]interface{}`

GetActualDuration returns the ActualDuration field if non-nil, zero value otherwise.

### GetActualDurationOk

`func (o *Phase) GetActualDurationOk() (*map[string]interface{}, bool)`

GetActualDurationOk returns a tuple with the ActualDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDuration

`func (o *Phase) SetActualDuration(v map[string]interface{})`

SetActualDuration sets ActualDuration field to given value.

### HasActualDuration

`func (o *Phase) HasActualDuration() bool`

HasActualDuration returns a boolean if a field has been set.

### GetReleaseUid

`func (o *Phase) GetReleaseUid() int32`

GetReleaseUid returns the ReleaseUid field if non-nil, zero value otherwise.

### GetReleaseUidOk

`func (o *Phase) GetReleaseUidOk() (*int32, bool)`

GetReleaseUidOk returns a tuple with the ReleaseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseUid

`func (o *Phase) SetReleaseUid(v int32)`

SetReleaseUid sets ReleaseUid field to given value.

### HasReleaseUid

`func (o *Phase) HasReleaseUid() bool`

HasReleaseUid returns a boolean if a field has been set.

### GetTasks

`func (o *Phase) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Phase) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Phase) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Phase) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetRelease

`func (o *Phase) GetRelease() interface{}`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *Phase) GetReleaseOk() (*interface{}, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *Phase) SetRelease(v interface{})`

SetRelease sets Release field to given value.

### HasRelease

`func (o *Phase) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### SetReleaseNil

`func (o *Phase) SetReleaseNil(b bool)`

 SetReleaseNil sets the value for Release to be an explicit nil

### UnsetRelease
`func (o *Phase) UnsetRelease()`

UnsetRelease ensures that no value is present for Release, not even an explicit nil
### GetStatus

`func (o *Phase) GetStatus() PhaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Phase) GetStatusOk() (*PhaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Phase) SetStatus(v PhaseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Phase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetColor

`func (o *Phase) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Phase) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Phase) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Phase) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetOriginId

`func (o *Phase) GetOriginId() string`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *Phase) GetOriginIdOk() (*string, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *Phase) SetOriginId(v string)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *Phase) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetCurrentTask

`func (o *Phase) GetCurrentTask() Task`

GetCurrentTask returns the CurrentTask field if non-nil, zero value otherwise.

### GetCurrentTaskOk

`func (o *Phase) GetCurrentTaskOk() (*Task, bool)`

GetCurrentTaskOk returns a tuple with the CurrentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTask

`func (o *Phase) SetCurrentTask(v Task)`

SetCurrentTask sets CurrentTask field to given value.

### HasCurrentTask

`func (o *Phase) HasCurrentTask() bool`

HasCurrentTask returns a boolean if a field has been set.

### GetDisplayPath

`func (o *Phase) GetDisplayPath() string`

GetDisplayPath returns the DisplayPath field if non-nil, zero value otherwise.

### GetDisplayPathOk

`func (o *Phase) GetDisplayPathOk() (*string, bool)`

GetDisplayPathOk returns a tuple with the DisplayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPath

`func (o *Phase) SetDisplayPath(v string)`

SetDisplayPath sets DisplayPath field to given value.

### HasDisplayPath

`func (o *Phase) HasDisplayPath() bool`

HasDisplayPath returns a boolean if a field has been set.

### GetActive

`func (o *Phase) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Phase) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Phase) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Phase) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDone

`func (o *Phase) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *Phase) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *Phase) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *Phase) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetDefunct

`func (o *Phase) GetDefunct() bool`

GetDefunct returns the Defunct field if non-nil, zero value otherwise.

### GetDefunctOk

`func (o *Phase) GetDefunctOk() (*bool, bool)`

GetDefunctOk returns a tuple with the Defunct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunct

`func (o *Phase) SetDefunct(v bool)`

SetDefunct sets Defunct field to given value.

### HasDefunct

`func (o *Phase) HasDefunct() bool`

HasDefunct returns a boolean if a field has been set.

### GetUpdatable

`func (o *Phase) GetUpdatable() bool`

GetUpdatable returns the Updatable field if non-nil, zero value otherwise.

### GetUpdatableOk

`func (o *Phase) GetUpdatableOk() (*bool, bool)`

GetUpdatableOk returns a tuple with the Updatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatable

`func (o *Phase) SetUpdatable(v bool)`

SetUpdatable sets Updatable field to given value.

### HasUpdatable

`func (o *Phase) HasUpdatable() bool`

HasUpdatable returns a boolean if a field has been set.

### GetAborted

`func (o *Phase) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *Phase) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *Phase) SetAborted(v bool)`

SetAborted sets Aborted field to given value.

### HasAborted

`func (o *Phase) HasAborted() bool`

HasAborted returns a boolean if a field has been set.

### GetPlanned

`func (o *Phase) GetPlanned() bool`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *Phase) GetPlannedOk() (*bool, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *Phase) SetPlanned(v bool)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *Phase) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.

### GetFailed

`func (o *Phase) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *Phase) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *Phase) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *Phase) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetFailing

`func (o *Phase) GetFailing() bool`

GetFailing returns the Failing field if non-nil, zero value otherwise.

### GetFailingOk

`func (o *Phase) GetFailingOk() (*bool, bool)`

GetFailingOk returns a tuple with the Failing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailing

`func (o *Phase) SetFailing(v bool)`

SetFailing sets Failing field to given value.

### HasFailing

`func (o *Phase) HasFailing() bool`

HasFailing returns a boolean if a field has been set.

### GetReleaseOwner

`func (o *Phase) GetReleaseOwner() string`

GetReleaseOwner returns the ReleaseOwner field if non-nil, zero value otherwise.

### GetReleaseOwnerOk

`func (o *Phase) GetReleaseOwnerOk() (*string, bool)`

GetReleaseOwnerOk returns a tuple with the ReleaseOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseOwner

`func (o *Phase) SetReleaseOwner(v string)`

SetReleaseOwner sets ReleaseOwner field to given value.

### HasReleaseOwner

`func (o *Phase) HasReleaseOwner() bool`

HasReleaseOwner returns a boolean if a field has been set.

### GetAllGates

`func (o *Phase) GetAllGates() []GateTask`

GetAllGates returns the AllGates field if non-nil, zero value otherwise.

### GetAllGatesOk

`func (o *Phase) GetAllGatesOk() (*[]GateTask, bool)`

GetAllGatesOk returns a tuple with the AllGates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGates

`func (o *Phase) SetAllGates(v []GateTask)`

SetAllGates sets AllGates field to given value.

### HasAllGates

`func (o *Phase) HasAllGates() bool`

HasAllGates returns a boolean if a field has been set.

### GetAllTasks

`func (o *Phase) GetAllTasks() []Task`

GetAllTasks returns the AllTasks field if non-nil, zero value otherwise.

### GetAllTasksOk

`func (o *Phase) GetAllTasksOk() (*[]Task, bool)`

GetAllTasksOk returns a tuple with the AllTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTasks

`func (o *Phase) SetAllTasks(v []Task)`

SetAllTasks sets AllTasks field to given value.

### HasAllTasks

`func (o *Phase) HasAllTasks() bool`

HasAllTasks returns a boolean if a field has been set.

### GetChildren

`func (o *Phase) GetChildren() []PlanItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Phase) GetChildrenOk() (*[]PlanItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Phase) SetChildren(v []PlanItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Phase) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetVariableUsages

`func (o *Phase) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *Phase) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *Phase) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *Phase) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.

### GetOriginal

`func (o *Phase) GetOriginal() bool`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *Phase) GetOriginalOk() (*bool, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *Phase) SetOriginal(v bool)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *Phase) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetPhaseCopied

`func (o *Phase) GetPhaseCopied() bool`

GetPhaseCopied returns the PhaseCopied field if non-nil, zero value otherwise.

### GetPhaseCopiedOk

`func (o *Phase) GetPhaseCopiedOk() (*bool, bool)`

GetPhaseCopiedOk returns a tuple with the PhaseCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseCopied

`func (o *Phase) SetPhaseCopied(v bool)`

SetPhaseCopied sets PhaseCopied field to given value.

### HasPhaseCopied

`func (o *Phase) HasPhaseCopied() bool`

HasPhaseCopied returns a boolean if a field has been set.

### GetAncestorId

`func (o *Phase) GetAncestorId() string`

GetAncestorId returns the AncestorId field if non-nil, zero value otherwise.

### GetAncestorIdOk

`func (o *Phase) GetAncestorIdOk() (*string, bool)`

GetAncestorIdOk returns a tuple with the AncestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorId

`func (o *Phase) SetAncestorId(v string)`

SetAncestorId sets AncestorId field to given value.

### HasAncestorId

`func (o *Phase) HasAncestorId() bool`

HasAncestorId returns a boolean if a field has been set.

### GetLatestCopy

`func (o *Phase) GetLatestCopy() bool`

GetLatestCopy returns the LatestCopy field if non-nil, zero value otherwise.

### GetLatestCopyOk

`func (o *Phase) GetLatestCopyOk() (*bool, bool)`

GetLatestCopyOk returns a tuple with the LatestCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCopy

`func (o *Phase) SetLatestCopy(v bool)`

SetLatestCopy sets LatestCopy field to given value.

### HasLatestCopy

`func (o *Phase) HasLatestCopy() bool`

HasLatestCopy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


