# GateTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ScheduledStartDate** | Pointer to **time.Time** |  | [optional] 
**FlagStatus** | Pointer to [**FlagStatus**](FlagStatus.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**PlannedDuration** | Pointer to **int32** |  | [optional] 
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
**CiUid** | Pointer to **int32** |  | [optional] 
**Comments** | Pointer to [**[]Comment**](Comment.md) |  | [optional] 
**Container** | Pointer to [**TaskContainer**](TaskContainer.md) |  | [optional] 
**Facets** | Pointer to [**[]Facet**](Facet.md) |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**Status** | Pointer to [**TaskStatus**](TaskStatus.md) |  | [optional] 
**Team** | Pointer to **string** |  | [optional] 
**Watchers** | Pointer to **[]string** |  | [optional] 
**WaitForScheduledStartDate** | Pointer to **bool** |  | [optional] 
**DelayDuringBlackout** | Pointer to **bool** |  | [optional] 
**PostponedDueToBlackout** | Pointer to **bool** |  | [optional] 
**PostponedUntilEnvironmentsAreReserved** | Pointer to **bool** |  | [optional] 
**OriginalScheduledStartDate** | Pointer to **time.Time** |  | [optional] 
**HasBeenFlagged** | Pointer to **bool** |  | [optional] 
**HasBeenDelayed** | Pointer to **bool** |  | [optional] 
**Precondition** | Pointer to **string** |  | [optional] 
**FailureHandler** | Pointer to **string** |  | [optional] 
**TaskFailureHandlerEnabled** | Pointer to **bool** |  | [optional] 
**TaskRecoverOp** | Pointer to [**TaskRecoverOp**](TaskRecoverOp.md) |  | [optional] 
**FailuresCount** | Pointer to **int32** |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**VariableMapping** | Pointer to **map[string]string** |  | [optional] 
**ExternalVariableMapping** | Pointer to **map[string]string** |  | [optional] 
**MaxCommentSize** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ConfigurationUri** | Pointer to **string** |  | [optional] 
**DueSoonNotified** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**CheckAttributes** | Pointer to **bool** |  | [optional] 
**AbortScript** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to [**Phase**](Phase.md) |  | [optional] 
**BlackoutMetadata** | Pointer to [**BlackoutMetadata**](BlackoutMetadata.md) |  | [optional] 
**FlaggedCount** | Pointer to **int32** |  | [optional] 
**DelayedCount** | Pointer to **int32** |  | [optional] 
**Done** | Pointer to **bool** |  | [optional] 
**DoneInAdvance** | Pointer to **bool** |  | [optional] 
**Defunct** | Pointer to **bool** |  | [optional] 
**Updatable** | Pointer to **bool** |  | [optional] 
**Aborted** | Pointer to **bool** |  | [optional] 
**NotYetReached** | Pointer to **bool** |  | [optional] 
**Planned** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**InProgress** | Pointer to **bool** |  | [optional] 
**Pending** | Pointer to **bool** |  | [optional] 
**WaitingForInput** | Pointer to **bool** |  | [optional] 
**Failed** | Pointer to **bool** |  | [optional] 
**Failing** | Pointer to **bool** |  | [optional] 
**CompletedInAdvance** | Pointer to **bool** |  | [optional] 
**Skipped** | Pointer to **bool** |  | [optional] 
**SkippedInAdvance** | Pointer to **bool** |  | [optional] 
**PreconditionInProgress** | Pointer to **bool** |  | [optional] 
**FailureHandlerInProgress** | Pointer to **bool** |  | [optional] 
**AbortScriptInProgress** | Pointer to **bool** |  | [optional] 
**FacetInProgress** | Pointer to **bool** |  | [optional] 
**Movable** | Pointer to **bool** |  | [optional] 
**Gate** | Pointer to **bool** |  | [optional] 
**TaskGroup** | Pointer to **bool** |  | [optional] 
**ParallelGroup** | Pointer to **bool** |  | [optional] 
**PreconditionEnabled** | Pointer to **bool** |  | [optional] 
**FailureHandlerEnabled** | Pointer to **bool** |  | [optional] 
**Release** | Pointer to [**Release**](Release.md) |  | [optional] 
**DisplayPath** | Pointer to **string** |  | [optional] 
**ReleaseOwner** | Pointer to **string** |  | [optional] 
**AllTasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**Children** | Pointer to [**[]PlanItem**](PlanItem.md) |  | [optional] 
**InputVariables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**ReferencedVariables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**UnboundRequiredVariables** | Pointer to **[]string** |  | [optional] 
**Automated** | Pointer to **bool** |  | [optional] 
**TaskType** | Pointer to **map[string]interface{}** |  | [optional] 
**DueSoon** | Pointer to **bool** |  | [optional] 
**ElapsedDurationFraction** | Pointer to **float64** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]GateCondition**](GateCondition.md) |  | [optional] 
**Dependencies** | Pointer to [**[]Dependency**](Dependency.md) |  | [optional] 
**Open** | Pointer to **bool** |  | [optional] 
**OpenInAdvance** | Pointer to **bool** |  | [optional] 
**Completable** | Pointer to **bool** |  | [optional] 
**AbortedDependencyTitles** | Pointer to **string** |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 

## Methods

### NewGateTask

`func NewGateTask() *GateTask`

NewGateTask instantiates a new GateTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGateTaskWithDefaults

`func NewGateTaskWithDefaults() *GateTask`

NewGateTaskWithDefaults instantiates a new GateTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GateTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GateTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GateTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GateTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GateTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GateTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GateTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GateTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScheduledStartDate

`func (o *GateTask) GetScheduledStartDate() time.Time`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *GateTask) GetScheduledStartDateOk() (*time.Time, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *GateTask) SetScheduledStartDate(v time.Time)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *GateTask) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetFlagStatus

`func (o *GateTask) GetFlagStatus() FlagStatus`

GetFlagStatus returns the FlagStatus field if non-nil, zero value otherwise.

### GetFlagStatusOk

`func (o *GateTask) GetFlagStatusOk() (*FlagStatus, bool)`

GetFlagStatusOk returns a tuple with the FlagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagStatus

`func (o *GateTask) SetFlagStatus(v FlagStatus)`

SetFlagStatus sets FlagStatus field to given value.

### HasFlagStatus

`func (o *GateTask) HasFlagStatus() bool`

HasFlagStatus returns a boolean if a field has been set.

### GetTitle

`func (o *GateTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GateTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GateTask) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GateTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *GateTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GateTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GateTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GateTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *GateTask) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GateTask) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GateTask) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GateTask) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDueDate

`func (o *GateTask) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *GateTask) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *GateTask) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *GateTask) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *GateTask) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GateTask) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GateTask) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GateTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GateTask) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GateTask) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GateTask) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GateTask) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPlannedDuration

`func (o *GateTask) GetPlannedDuration() int32`

GetPlannedDuration returns the PlannedDuration field if non-nil, zero value otherwise.

### GetPlannedDurationOk

`func (o *GateTask) GetPlannedDurationOk() (*int32, bool)`

GetPlannedDurationOk returns a tuple with the PlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDuration

`func (o *GateTask) SetPlannedDuration(v int32)`

SetPlannedDuration sets PlannedDuration field to given value.

### HasPlannedDuration

`func (o *GateTask) HasPlannedDuration() bool`

HasPlannedDuration returns a boolean if a field has been set.

### GetFlagComment

`func (o *GateTask) GetFlagComment() string`

GetFlagComment returns the FlagComment field if non-nil, zero value otherwise.

### GetFlagCommentOk

`func (o *GateTask) GetFlagCommentOk() (*string, bool)`

GetFlagCommentOk returns a tuple with the FlagComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagComment

`func (o *GateTask) SetFlagComment(v string)`

SetFlagComment sets FlagComment field to given value.

### HasFlagComment

`func (o *GateTask) HasFlagComment() bool`

HasFlagComment returns a boolean if a field has been set.

### GetOverdueNotified

`func (o *GateTask) GetOverdueNotified() bool`

GetOverdueNotified returns the OverdueNotified field if non-nil, zero value otherwise.

### GetOverdueNotifiedOk

`func (o *GateTask) GetOverdueNotifiedOk() (*bool, bool)`

GetOverdueNotifiedOk returns a tuple with the OverdueNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueNotified

`func (o *GateTask) SetOverdueNotified(v bool)`

SetOverdueNotified sets OverdueNotified field to given value.

### HasOverdueNotified

`func (o *GateTask) HasOverdueNotified() bool`

HasOverdueNotified returns a boolean if a field has been set.

### GetFlagged

`func (o *GateTask) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *GateTask) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *GateTask) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *GateTask) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetStartOrScheduledDate

`func (o *GateTask) GetStartOrScheduledDate() time.Time`

GetStartOrScheduledDate returns the StartOrScheduledDate field if non-nil, zero value otherwise.

### GetStartOrScheduledDateOk

`func (o *GateTask) GetStartOrScheduledDateOk() (*time.Time, bool)`

GetStartOrScheduledDateOk returns a tuple with the StartOrScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOrScheduledDate

`func (o *GateTask) SetStartOrScheduledDate(v time.Time)`

SetStartOrScheduledDate sets StartOrScheduledDate field to given value.

### HasStartOrScheduledDate

`func (o *GateTask) HasStartOrScheduledDate() bool`

HasStartOrScheduledDate returns a boolean if a field has been set.

### GetEndOrDueDate

`func (o *GateTask) GetEndOrDueDate() time.Time`

GetEndOrDueDate returns the EndOrDueDate field if non-nil, zero value otherwise.

### GetEndOrDueDateOk

`func (o *GateTask) GetEndOrDueDateOk() (*time.Time, bool)`

GetEndOrDueDateOk returns a tuple with the EndOrDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOrDueDate

`func (o *GateTask) SetEndOrDueDate(v time.Time)`

SetEndOrDueDate sets EndOrDueDate field to given value.

### HasEndOrDueDate

`func (o *GateTask) HasEndOrDueDate() bool`

HasEndOrDueDate returns a boolean if a field has been set.

### GetOverdue

`func (o *GateTask) GetOverdue() bool`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *GateTask) GetOverdueOk() (*bool, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *GateTask) SetOverdue(v bool)`

SetOverdue sets Overdue field to given value.

### HasOverdue

`func (o *GateTask) HasOverdue() bool`

HasOverdue returns a boolean if a field has been set.

### GetOrCalculateDueDate

`func (o *GateTask) GetOrCalculateDueDate() time.Time`

GetOrCalculateDueDate returns the OrCalculateDueDate field if non-nil, zero value otherwise.

### GetOrCalculateDueDateOk

`func (o *GateTask) GetOrCalculateDueDateOk() (*time.Time, bool)`

GetOrCalculateDueDateOk returns a tuple with the OrCalculateDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrCalculateDueDate

`func (o *GateTask) SetOrCalculateDueDate(v time.Time)`

SetOrCalculateDueDate sets OrCalculateDueDate field to given value.

### HasOrCalculateDueDate

`func (o *GateTask) HasOrCalculateDueDate() bool`

HasOrCalculateDueDate returns a boolean if a field has been set.

### SetOrCalculateDueDateNil

`func (o *GateTask) SetOrCalculateDueDateNil(b bool)`

 SetOrCalculateDueDateNil sets the value for OrCalculateDueDate to be an explicit nil

### UnsetOrCalculateDueDate
`func (o *GateTask) UnsetOrCalculateDueDate()`

UnsetOrCalculateDueDate ensures that no value is present for OrCalculateDueDate, not even an explicit nil
### GetComputedPlannedDuration

`func (o *GateTask) GetComputedPlannedDuration() map[string]interface{}`

GetComputedPlannedDuration returns the ComputedPlannedDuration field if non-nil, zero value otherwise.

### GetComputedPlannedDurationOk

`func (o *GateTask) GetComputedPlannedDurationOk() (*map[string]interface{}, bool)`

GetComputedPlannedDurationOk returns a tuple with the ComputedPlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPlannedDuration

`func (o *GateTask) SetComputedPlannedDuration(v map[string]interface{})`

SetComputedPlannedDuration sets ComputedPlannedDuration field to given value.

### HasComputedPlannedDuration

`func (o *GateTask) HasComputedPlannedDuration() bool`

HasComputedPlannedDuration returns a boolean if a field has been set.

### GetActualDuration

`func (o *GateTask) GetActualDuration() map[string]interface{}`

GetActualDuration returns the ActualDuration field if non-nil, zero value otherwise.

### GetActualDurationOk

`func (o *GateTask) GetActualDurationOk() (*map[string]interface{}, bool)`

GetActualDurationOk returns a tuple with the ActualDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDuration

`func (o *GateTask) SetActualDuration(v map[string]interface{})`

SetActualDuration sets ActualDuration field to given value.

### HasActualDuration

`func (o *GateTask) HasActualDuration() bool`

HasActualDuration returns a boolean if a field has been set.

### GetReleaseUid

`func (o *GateTask) GetReleaseUid() int32`

GetReleaseUid returns the ReleaseUid field if non-nil, zero value otherwise.

### GetReleaseUidOk

`func (o *GateTask) GetReleaseUidOk() (*int32, bool)`

GetReleaseUidOk returns a tuple with the ReleaseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseUid

`func (o *GateTask) SetReleaseUid(v int32)`

SetReleaseUid sets ReleaseUid field to given value.

### HasReleaseUid

`func (o *GateTask) HasReleaseUid() bool`

HasReleaseUid returns a boolean if a field has been set.

### GetCiUid

`func (o *GateTask) GetCiUid() int32`

GetCiUid returns the CiUid field if non-nil, zero value otherwise.

### GetCiUidOk

`func (o *GateTask) GetCiUidOk() (*int32, bool)`

GetCiUidOk returns a tuple with the CiUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiUid

`func (o *GateTask) SetCiUid(v int32)`

SetCiUid sets CiUid field to given value.

### HasCiUid

`func (o *GateTask) HasCiUid() bool`

HasCiUid returns a boolean if a field has been set.

### GetComments

`func (o *GateTask) GetComments() []Comment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GateTask) GetCommentsOk() (*[]Comment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GateTask) SetComments(v []Comment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *GateTask) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetContainer

`func (o *GateTask) GetContainer() TaskContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *GateTask) GetContainerOk() (*TaskContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *GateTask) SetContainer(v TaskContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *GateTask) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetFacets

`func (o *GateTask) GetFacets() []Facet`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *GateTask) GetFacetsOk() (*[]Facet, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *GateTask) SetFacets(v []Facet)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *GateTask) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### GetAttachments

`func (o *GateTask) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GateTask) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GateTask) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GateTask) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetStatus

`func (o *GateTask) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GateTask) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GateTask) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GateTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTeam

`func (o *GateTask) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GateTask) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GateTask) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GateTask) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetWatchers

`func (o *GateTask) GetWatchers() []string`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *GateTask) GetWatchersOk() (*[]string, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *GateTask) SetWatchers(v []string)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *GateTask) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.

### GetWaitForScheduledStartDate

`func (o *GateTask) GetWaitForScheduledStartDate() bool`

GetWaitForScheduledStartDate returns the WaitForScheduledStartDate field if non-nil, zero value otherwise.

### GetWaitForScheduledStartDateOk

`func (o *GateTask) GetWaitForScheduledStartDateOk() (*bool, bool)`

GetWaitForScheduledStartDateOk returns a tuple with the WaitForScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForScheduledStartDate

`func (o *GateTask) SetWaitForScheduledStartDate(v bool)`

SetWaitForScheduledStartDate sets WaitForScheduledStartDate field to given value.

### HasWaitForScheduledStartDate

`func (o *GateTask) HasWaitForScheduledStartDate() bool`

HasWaitForScheduledStartDate returns a boolean if a field has been set.

### GetDelayDuringBlackout

`func (o *GateTask) GetDelayDuringBlackout() bool`

GetDelayDuringBlackout returns the DelayDuringBlackout field if non-nil, zero value otherwise.

### GetDelayDuringBlackoutOk

`func (o *GateTask) GetDelayDuringBlackoutOk() (*bool, bool)`

GetDelayDuringBlackoutOk returns a tuple with the DelayDuringBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDuringBlackout

`func (o *GateTask) SetDelayDuringBlackout(v bool)`

SetDelayDuringBlackout sets DelayDuringBlackout field to given value.

### HasDelayDuringBlackout

`func (o *GateTask) HasDelayDuringBlackout() bool`

HasDelayDuringBlackout returns a boolean if a field has been set.

### GetPostponedDueToBlackout

`func (o *GateTask) GetPostponedDueToBlackout() bool`

GetPostponedDueToBlackout returns the PostponedDueToBlackout field if non-nil, zero value otherwise.

### GetPostponedDueToBlackoutOk

`func (o *GateTask) GetPostponedDueToBlackoutOk() (*bool, bool)`

GetPostponedDueToBlackoutOk returns a tuple with the PostponedDueToBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostponedDueToBlackout

`func (o *GateTask) SetPostponedDueToBlackout(v bool)`

SetPostponedDueToBlackout sets PostponedDueToBlackout field to given value.

### HasPostponedDueToBlackout

`func (o *GateTask) HasPostponedDueToBlackout() bool`

HasPostponedDueToBlackout returns a boolean if a field has been set.

### GetPostponedUntilEnvironmentsAreReserved

`func (o *GateTask) GetPostponedUntilEnvironmentsAreReserved() bool`

GetPostponedUntilEnvironmentsAreReserved returns the PostponedUntilEnvironmentsAreReserved field if non-nil, zero value otherwise.

### GetPostponedUntilEnvironmentsAreReservedOk

`func (o *GateTask) GetPostponedUntilEnvironmentsAreReservedOk() (*bool, bool)`

GetPostponedUntilEnvironmentsAreReservedOk returns a tuple with the PostponedUntilEnvironmentsAreReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostponedUntilEnvironmentsAreReserved

`func (o *GateTask) SetPostponedUntilEnvironmentsAreReserved(v bool)`

SetPostponedUntilEnvironmentsAreReserved sets PostponedUntilEnvironmentsAreReserved field to given value.

### HasPostponedUntilEnvironmentsAreReserved

`func (o *GateTask) HasPostponedUntilEnvironmentsAreReserved() bool`

HasPostponedUntilEnvironmentsAreReserved returns a boolean if a field has been set.

### GetOriginalScheduledStartDate

`func (o *GateTask) GetOriginalScheduledStartDate() time.Time`

GetOriginalScheduledStartDate returns the OriginalScheduledStartDate field if non-nil, zero value otherwise.

### GetOriginalScheduledStartDateOk

`func (o *GateTask) GetOriginalScheduledStartDateOk() (*time.Time, bool)`

GetOriginalScheduledStartDateOk returns a tuple with the OriginalScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalScheduledStartDate

`func (o *GateTask) SetOriginalScheduledStartDate(v time.Time)`

SetOriginalScheduledStartDate sets OriginalScheduledStartDate field to given value.

### HasOriginalScheduledStartDate

`func (o *GateTask) HasOriginalScheduledStartDate() bool`

HasOriginalScheduledStartDate returns a boolean if a field has been set.

### GetHasBeenFlagged

`func (o *GateTask) GetHasBeenFlagged() bool`

GetHasBeenFlagged returns the HasBeenFlagged field if non-nil, zero value otherwise.

### GetHasBeenFlaggedOk

`func (o *GateTask) GetHasBeenFlaggedOk() (*bool, bool)`

GetHasBeenFlaggedOk returns a tuple with the HasBeenFlagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenFlagged

`func (o *GateTask) SetHasBeenFlagged(v bool)`

SetHasBeenFlagged sets HasBeenFlagged field to given value.

### HasHasBeenFlagged

`func (o *GateTask) HasHasBeenFlagged() bool`

HasHasBeenFlagged returns a boolean if a field has been set.

### GetHasBeenDelayed

`func (o *GateTask) GetHasBeenDelayed() bool`

GetHasBeenDelayed returns the HasBeenDelayed field if non-nil, zero value otherwise.

### GetHasBeenDelayedOk

`func (o *GateTask) GetHasBeenDelayedOk() (*bool, bool)`

GetHasBeenDelayedOk returns a tuple with the HasBeenDelayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenDelayed

`func (o *GateTask) SetHasBeenDelayed(v bool)`

SetHasBeenDelayed sets HasBeenDelayed field to given value.

### HasHasBeenDelayed

`func (o *GateTask) HasHasBeenDelayed() bool`

HasHasBeenDelayed returns a boolean if a field has been set.

### GetPrecondition

`func (o *GateTask) GetPrecondition() string`

GetPrecondition returns the Precondition field if non-nil, zero value otherwise.

### GetPreconditionOk

`func (o *GateTask) GetPreconditionOk() (*string, bool)`

GetPreconditionOk returns a tuple with the Precondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecondition

`func (o *GateTask) SetPrecondition(v string)`

SetPrecondition sets Precondition field to given value.

### HasPrecondition

`func (o *GateTask) HasPrecondition() bool`

HasPrecondition returns a boolean if a field has been set.

### GetFailureHandler

`func (o *GateTask) GetFailureHandler() string`

GetFailureHandler returns the FailureHandler field if non-nil, zero value otherwise.

### GetFailureHandlerOk

`func (o *GateTask) GetFailureHandlerOk() (*string, bool)`

GetFailureHandlerOk returns a tuple with the FailureHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureHandler

`func (o *GateTask) SetFailureHandler(v string)`

SetFailureHandler sets FailureHandler field to given value.

### HasFailureHandler

`func (o *GateTask) HasFailureHandler() bool`

HasFailureHandler returns a boolean if a field has been set.

### GetTaskFailureHandlerEnabled

`func (o *GateTask) GetTaskFailureHandlerEnabled() bool`

GetTaskFailureHandlerEnabled returns the TaskFailureHandlerEnabled field if non-nil, zero value otherwise.

### GetTaskFailureHandlerEnabledOk

`func (o *GateTask) GetTaskFailureHandlerEnabledOk() (*bool, bool)`

GetTaskFailureHandlerEnabledOk returns a tuple with the TaskFailureHandlerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFailureHandlerEnabled

`func (o *GateTask) SetTaskFailureHandlerEnabled(v bool)`

SetTaskFailureHandlerEnabled sets TaskFailureHandlerEnabled field to given value.

### HasTaskFailureHandlerEnabled

`func (o *GateTask) HasTaskFailureHandlerEnabled() bool`

HasTaskFailureHandlerEnabled returns a boolean if a field has been set.

### GetTaskRecoverOp

`func (o *GateTask) GetTaskRecoverOp() TaskRecoverOp`

GetTaskRecoverOp returns the TaskRecoverOp field if non-nil, zero value otherwise.

### GetTaskRecoverOpOk

`func (o *GateTask) GetTaskRecoverOpOk() (*TaskRecoverOp, bool)`

GetTaskRecoverOpOk returns a tuple with the TaskRecoverOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRecoverOp

`func (o *GateTask) SetTaskRecoverOp(v TaskRecoverOp)`

SetTaskRecoverOp sets TaskRecoverOp field to given value.

### HasTaskRecoverOp

`func (o *GateTask) HasTaskRecoverOp() bool`

HasTaskRecoverOp returns a boolean if a field has been set.

### GetFailuresCount

`func (o *GateTask) GetFailuresCount() int32`

GetFailuresCount returns the FailuresCount field if non-nil, zero value otherwise.

### GetFailuresCountOk

`func (o *GateTask) GetFailuresCountOk() (*int32, bool)`

GetFailuresCountOk returns a tuple with the FailuresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailuresCount

`func (o *GateTask) SetFailuresCount(v int32)`

SetFailuresCount sets FailuresCount field to given value.

### HasFailuresCount

`func (o *GateTask) HasFailuresCount() bool`

HasFailuresCount returns a boolean if a field has been set.

### GetExecutionId

`func (o *GateTask) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *GateTask) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *GateTask) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *GateTask) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetVariableMapping

`func (o *GateTask) GetVariableMapping() map[string]string`

GetVariableMapping returns the VariableMapping field if non-nil, zero value otherwise.

### GetVariableMappingOk

`func (o *GateTask) GetVariableMappingOk() (*map[string]string, bool)`

GetVariableMappingOk returns a tuple with the VariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMapping

`func (o *GateTask) SetVariableMapping(v map[string]string)`

SetVariableMapping sets VariableMapping field to given value.

### HasVariableMapping

`func (o *GateTask) HasVariableMapping() bool`

HasVariableMapping returns a boolean if a field has been set.

### GetExternalVariableMapping

`func (o *GateTask) GetExternalVariableMapping() map[string]string`

GetExternalVariableMapping returns the ExternalVariableMapping field if non-nil, zero value otherwise.

### GetExternalVariableMappingOk

`func (o *GateTask) GetExternalVariableMappingOk() (*map[string]string, bool)`

GetExternalVariableMappingOk returns a tuple with the ExternalVariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVariableMapping

`func (o *GateTask) SetExternalVariableMapping(v map[string]string)`

SetExternalVariableMapping sets ExternalVariableMapping field to given value.

### HasExternalVariableMapping

`func (o *GateTask) HasExternalVariableMapping() bool`

HasExternalVariableMapping returns a boolean if a field has been set.

### GetMaxCommentSize

`func (o *GateTask) GetMaxCommentSize() int32`

GetMaxCommentSize returns the MaxCommentSize field if non-nil, zero value otherwise.

### GetMaxCommentSizeOk

`func (o *GateTask) GetMaxCommentSizeOk() (*int32, bool)`

GetMaxCommentSizeOk returns a tuple with the MaxCommentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCommentSize

`func (o *GateTask) SetMaxCommentSize(v int32)`

SetMaxCommentSize sets MaxCommentSize field to given value.

### HasMaxCommentSize

`func (o *GateTask) HasMaxCommentSize() bool`

HasMaxCommentSize returns a boolean if a field has been set.

### GetTags

`func (o *GateTask) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GateTask) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GateTask) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GateTask) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfigurationUri

`func (o *GateTask) GetConfigurationUri() string`

GetConfigurationUri returns the ConfigurationUri field if non-nil, zero value otherwise.

### GetConfigurationUriOk

`func (o *GateTask) GetConfigurationUriOk() (*string, bool)`

GetConfigurationUriOk returns a tuple with the ConfigurationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUri

`func (o *GateTask) SetConfigurationUri(v string)`

SetConfigurationUri sets ConfigurationUri field to given value.

### HasConfigurationUri

`func (o *GateTask) HasConfigurationUri() bool`

HasConfigurationUri returns a boolean if a field has been set.

### GetDueSoonNotified

`func (o *GateTask) GetDueSoonNotified() bool`

GetDueSoonNotified returns the DueSoonNotified field if non-nil, zero value otherwise.

### GetDueSoonNotifiedOk

`func (o *GateTask) GetDueSoonNotifiedOk() (*bool, bool)`

GetDueSoonNotifiedOk returns a tuple with the DueSoonNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueSoonNotified

`func (o *GateTask) SetDueSoonNotified(v bool)`

SetDueSoonNotified sets DueSoonNotified field to given value.

### HasDueSoonNotified

`func (o *GateTask) HasDueSoonNotified() bool`

HasDueSoonNotified returns a boolean if a field has been set.

### GetLocked

`func (o *GateTask) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *GateTask) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *GateTask) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *GateTask) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetCheckAttributes

`func (o *GateTask) GetCheckAttributes() bool`

GetCheckAttributes returns the CheckAttributes field if non-nil, zero value otherwise.

### GetCheckAttributesOk

`func (o *GateTask) GetCheckAttributesOk() (*bool, bool)`

GetCheckAttributesOk returns a tuple with the CheckAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAttributes

`func (o *GateTask) SetCheckAttributes(v bool)`

SetCheckAttributes sets CheckAttributes field to given value.

### HasCheckAttributes

`func (o *GateTask) HasCheckAttributes() bool`

HasCheckAttributes returns a boolean if a field has been set.

### GetAbortScript

`func (o *GateTask) GetAbortScript() string`

GetAbortScript returns the AbortScript field if non-nil, zero value otherwise.

### GetAbortScriptOk

`func (o *GateTask) GetAbortScriptOk() (*string, bool)`

GetAbortScriptOk returns a tuple with the AbortScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortScript

`func (o *GateTask) SetAbortScript(v string)`

SetAbortScript sets AbortScript field to given value.

### HasAbortScript

`func (o *GateTask) HasAbortScript() bool`

HasAbortScript returns a boolean if a field has been set.

### GetPhase

`func (o *GateTask) GetPhase() Phase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *GateTask) GetPhaseOk() (*Phase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *GateTask) SetPhase(v Phase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *GateTask) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetBlackoutMetadata

`func (o *GateTask) GetBlackoutMetadata() BlackoutMetadata`

GetBlackoutMetadata returns the BlackoutMetadata field if non-nil, zero value otherwise.

### GetBlackoutMetadataOk

`func (o *GateTask) GetBlackoutMetadataOk() (*BlackoutMetadata, bool)`

GetBlackoutMetadataOk returns a tuple with the BlackoutMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutMetadata

`func (o *GateTask) SetBlackoutMetadata(v BlackoutMetadata)`

SetBlackoutMetadata sets BlackoutMetadata field to given value.

### HasBlackoutMetadata

`func (o *GateTask) HasBlackoutMetadata() bool`

HasBlackoutMetadata returns a boolean if a field has been set.

### GetFlaggedCount

`func (o *GateTask) GetFlaggedCount() int32`

GetFlaggedCount returns the FlaggedCount field if non-nil, zero value otherwise.

### GetFlaggedCountOk

`func (o *GateTask) GetFlaggedCountOk() (*int32, bool)`

GetFlaggedCountOk returns a tuple with the FlaggedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedCount

`func (o *GateTask) SetFlaggedCount(v int32)`

SetFlaggedCount sets FlaggedCount field to given value.

### HasFlaggedCount

`func (o *GateTask) HasFlaggedCount() bool`

HasFlaggedCount returns a boolean if a field has been set.

### GetDelayedCount

`func (o *GateTask) GetDelayedCount() int32`

GetDelayedCount returns the DelayedCount field if non-nil, zero value otherwise.

### GetDelayedCountOk

`func (o *GateTask) GetDelayedCountOk() (*int32, bool)`

GetDelayedCountOk returns a tuple with the DelayedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedCount

`func (o *GateTask) SetDelayedCount(v int32)`

SetDelayedCount sets DelayedCount field to given value.

### HasDelayedCount

`func (o *GateTask) HasDelayedCount() bool`

HasDelayedCount returns a boolean if a field has been set.

### GetDone

`func (o *GateTask) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *GateTask) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *GateTask) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *GateTask) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetDoneInAdvance

`func (o *GateTask) GetDoneInAdvance() bool`

GetDoneInAdvance returns the DoneInAdvance field if non-nil, zero value otherwise.

### GetDoneInAdvanceOk

`func (o *GateTask) GetDoneInAdvanceOk() (*bool, bool)`

GetDoneInAdvanceOk returns a tuple with the DoneInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneInAdvance

`func (o *GateTask) SetDoneInAdvance(v bool)`

SetDoneInAdvance sets DoneInAdvance field to given value.

### HasDoneInAdvance

`func (o *GateTask) HasDoneInAdvance() bool`

HasDoneInAdvance returns a boolean if a field has been set.

### GetDefunct

`func (o *GateTask) GetDefunct() bool`

GetDefunct returns the Defunct field if non-nil, zero value otherwise.

### GetDefunctOk

`func (o *GateTask) GetDefunctOk() (*bool, bool)`

GetDefunctOk returns a tuple with the Defunct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunct

`func (o *GateTask) SetDefunct(v bool)`

SetDefunct sets Defunct field to given value.

### HasDefunct

`func (o *GateTask) HasDefunct() bool`

HasDefunct returns a boolean if a field has been set.

### GetUpdatable

`func (o *GateTask) GetUpdatable() bool`

GetUpdatable returns the Updatable field if non-nil, zero value otherwise.

### GetUpdatableOk

`func (o *GateTask) GetUpdatableOk() (*bool, bool)`

GetUpdatableOk returns a tuple with the Updatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatable

`func (o *GateTask) SetUpdatable(v bool)`

SetUpdatable sets Updatable field to given value.

### HasUpdatable

`func (o *GateTask) HasUpdatable() bool`

HasUpdatable returns a boolean if a field has been set.

### GetAborted

`func (o *GateTask) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *GateTask) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *GateTask) SetAborted(v bool)`

SetAborted sets Aborted field to given value.

### HasAborted

`func (o *GateTask) HasAborted() bool`

HasAborted returns a boolean if a field has been set.

### GetNotYetReached

`func (o *GateTask) GetNotYetReached() bool`

GetNotYetReached returns the NotYetReached field if non-nil, zero value otherwise.

### GetNotYetReachedOk

`func (o *GateTask) GetNotYetReachedOk() (*bool, bool)`

GetNotYetReachedOk returns a tuple with the NotYetReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotYetReached

`func (o *GateTask) SetNotYetReached(v bool)`

SetNotYetReached sets NotYetReached field to given value.

### HasNotYetReached

`func (o *GateTask) HasNotYetReached() bool`

HasNotYetReached returns a boolean if a field has been set.

### GetPlanned

`func (o *GateTask) GetPlanned() bool`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *GateTask) GetPlannedOk() (*bool, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *GateTask) SetPlanned(v bool)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *GateTask) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.

### GetActive

`func (o *GateTask) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GateTask) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GateTask) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GateTask) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetInProgress

`func (o *GateTask) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *GateTask) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *GateTask) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *GateTask) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetPending

`func (o *GateTask) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GateTask) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GateTask) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *GateTask) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetWaitingForInput

`func (o *GateTask) GetWaitingForInput() bool`

GetWaitingForInput returns the WaitingForInput field if non-nil, zero value otherwise.

### GetWaitingForInputOk

`func (o *GateTask) GetWaitingForInputOk() (*bool, bool)`

GetWaitingForInputOk returns a tuple with the WaitingForInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingForInput

`func (o *GateTask) SetWaitingForInput(v bool)`

SetWaitingForInput sets WaitingForInput field to given value.

### HasWaitingForInput

`func (o *GateTask) HasWaitingForInput() bool`

HasWaitingForInput returns a boolean if a field has been set.

### GetFailed

`func (o *GateTask) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *GateTask) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *GateTask) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *GateTask) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetFailing

`func (o *GateTask) GetFailing() bool`

GetFailing returns the Failing field if non-nil, zero value otherwise.

### GetFailingOk

`func (o *GateTask) GetFailingOk() (*bool, bool)`

GetFailingOk returns a tuple with the Failing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailing

`func (o *GateTask) SetFailing(v bool)`

SetFailing sets Failing field to given value.

### HasFailing

`func (o *GateTask) HasFailing() bool`

HasFailing returns a boolean if a field has been set.

### GetCompletedInAdvance

`func (o *GateTask) GetCompletedInAdvance() bool`

GetCompletedInAdvance returns the CompletedInAdvance field if non-nil, zero value otherwise.

### GetCompletedInAdvanceOk

`func (o *GateTask) GetCompletedInAdvanceOk() (*bool, bool)`

GetCompletedInAdvanceOk returns a tuple with the CompletedInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedInAdvance

`func (o *GateTask) SetCompletedInAdvance(v bool)`

SetCompletedInAdvance sets CompletedInAdvance field to given value.

### HasCompletedInAdvance

`func (o *GateTask) HasCompletedInAdvance() bool`

HasCompletedInAdvance returns a boolean if a field has been set.

### GetSkipped

`func (o *GateTask) GetSkipped() bool`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *GateTask) GetSkippedOk() (*bool, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *GateTask) SetSkipped(v bool)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *GateTask) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSkippedInAdvance

`func (o *GateTask) GetSkippedInAdvance() bool`

GetSkippedInAdvance returns the SkippedInAdvance field if non-nil, zero value otherwise.

### GetSkippedInAdvanceOk

`func (o *GateTask) GetSkippedInAdvanceOk() (*bool, bool)`

GetSkippedInAdvanceOk returns a tuple with the SkippedInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedInAdvance

`func (o *GateTask) SetSkippedInAdvance(v bool)`

SetSkippedInAdvance sets SkippedInAdvance field to given value.

### HasSkippedInAdvance

`func (o *GateTask) HasSkippedInAdvance() bool`

HasSkippedInAdvance returns a boolean if a field has been set.

### GetPreconditionInProgress

`func (o *GateTask) GetPreconditionInProgress() bool`

GetPreconditionInProgress returns the PreconditionInProgress field if non-nil, zero value otherwise.

### GetPreconditionInProgressOk

`func (o *GateTask) GetPreconditionInProgressOk() (*bool, bool)`

GetPreconditionInProgressOk returns a tuple with the PreconditionInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionInProgress

`func (o *GateTask) SetPreconditionInProgress(v bool)`

SetPreconditionInProgress sets PreconditionInProgress field to given value.

### HasPreconditionInProgress

`func (o *GateTask) HasPreconditionInProgress() bool`

HasPreconditionInProgress returns a boolean if a field has been set.

### GetFailureHandlerInProgress

`func (o *GateTask) GetFailureHandlerInProgress() bool`

GetFailureHandlerInProgress returns the FailureHandlerInProgress field if non-nil, zero value otherwise.

### GetFailureHandlerInProgressOk

`func (o *GateTask) GetFailureHandlerInProgressOk() (*bool, bool)`

GetFailureHandlerInProgressOk returns a tuple with the FailureHandlerInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureHandlerInProgress

`func (o *GateTask) SetFailureHandlerInProgress(v bool)`

SetFailureHandlerInProgress sets FailureHandlerInProgress field to given value.

### HasFailureHandlerInProgress

`func (o *GateTask) HasFailureHandlerInProgress() bool`

HasFailureHandlerInProgress returns a boolean if a field has been set.

### GetAbortScriptInProgress

`func (o *GateTask) GetAbortScriptInProgress() bool`

GetAbortScriptInProgress returns the AbortScriptInProgress field if non-nil, zero value otherwise.

### GetAbortScriptInProgressOk

`func (o *GateTask) GetAbortScriptInProgressOk() (*bool, bool)`

GetAbortScriptInProgressOk returns a tuple with the AbortScriptInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortScriptInProgress

`func (o *GateTask) SetAbortScriptInProgress(v bool)`

SetAbortScriptInProgress sets AbortScriptInProgress field to given value.

### HasAbortScriptInProgress

`func (o *GateTask) HasAbortScriptInProgress() bool`

HasAbortScriptInProgress returns a boolean if a field has been set.

### GetFacetInProgress

`func (o *GateTask) GetFacetInProgress() bool`

GetFacetInProgress returns the FacetInProgress field if non-nil, zero value otherwise.

### GetFacetInProgressOk

`func (o *GateTask) GetFacetInProgressOk() (*bool, bool)`

GetFacetInProgressOk returns a tuple with the FacetInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetInProgress

`func (o *GateTask) SetFacetInProgress(v bool)`

SetFacetInProgress sets FacetInProgress field to given value.

### HasFacetInProgress

`func (o *GateTask) HasFacetInProgress() bool`

HasFacetInProgress returns a boolean if a field has been set.

### GetMovable

`func (o *GateTask) GetMovable() bool`

GetMovable returns the Movable field if non-nil, zero value otherwise.

### GetMovableOk

`func (o *GateTask) GetMovableOk() (*bool, bool)`

GetMovableOk returns a tuple with the Movable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovable

`func (o *GateTask) SetMovable(v bool)`

SetMovable sets Movable field to given value.

### HasMovable

`func (o *GateTask) HasMovable() bool`

HasMovable returns a boolean if a field has been set.

### GetGate

`func (o *GateTask) GetGate() bool`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *GateTask) GetGateOk() (*bool, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *GateTask) SetGate(v bool)`

SetGate sets Gate field to given value.

### HasGate

`func (o *GateTask) HasGate() bool`

HasGate returns a boolean if a field has been set.

### GetTaskGroup

`func (o *GateTask) GetTaskGroup() bool`

GetTaskGroup returns the TaskGroup field if non-nil, zero value otherwise.

### GetTaskGroupOk

`func (o *GateTask) GetTaskGroupOk() (*bool, bool)`

GetTaskGroupOk returns a tuple with the TaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroup

`func (o *GateTask) SetTaskGroup(v bool)`

SetTaskGroup sets TaskGroup field to given value.

### HasTaskGroup

`func (o *GateTask) HasTaskGroup() bool`

HasTaskGroup returns a boolean if a field has been set.

### GetParallelGroup

`func (o *GateTask) GetParallelGroup() bool`

GetParallelGroup returns the ParallelGroup field if non-nil, zero value otherwise.

### GetParallelGroupOk

`func (o *GateTask) GetParallelGroupOk() (*bool, bool)`

GetParallelGroupOk returns a tuple with the ParallelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelGroup

`func (o *GateTask) SetParallelGroup(v bool)`

SetParallelGroup sets ParallelGroup field to given value.

### HasParallelGroup

`func (o *GateTask) HasParallelGroup() bool`

HasParallelGroup returns a boolean if a field has been set.

### GetPreconditionEnabled

`func (o *GateTask) GetPreconditionEnabled() bool`

GetPreconditionEnabled returns the PreconditionEnabled field if non-nil, zero value otherwise.

### GetPreconditionEnabledOk

`func (o *GateTask) GetPreconditionEnabledOk() (*bool, bool)`

GetPreconditionEnabledOk returns a tuple with the PreconditionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionEnabled

`func (o *GateTask) SetPreconditionEnabled(v bool)`

SetPreconditionEnabled sets PreconditionEnabled field to given value.

### HasPreconditionEnabled

`func (o *GateTask) HasPreconditionEnabled() bool`

HasPreconditionEnabled returns a boolean if a field has been set.

### GetFailureHandlerEnabled

`func (o *GateTask) GetFailureHandlerEnabled() bool`

GetFailureHandlerEnabled returns the FailureHandlerEnabled field if non-nil, zero value otherwise.

### GetFailureHandlerEnabledOk

`func (o *GateTask) GetFailureHandlerEnabledOk() (*bool, bool)`

GetFailureHandlerEnabledOk returns a tuple with the FailureHandlerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureHandlerEnabled

`func (o *GateTask) SetFailureHandlerEnabled(v bool)`

SetFailureHandlerEnabled sets FailureHandlerEnabled field to given value.

### HasFailureHandlerEnabled

`func (o *GateTask) HasFailureHandlerEnabled() bool`

HasFailureHandlerEnabled returns a boolean if a field has been set.

### GetRelease

`func (o *GateTask) GetRelease() Release`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *GateTask) GetReleaseOk() (*Release, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *GateTask) SetRelease(v Release)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *GateTask) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetDisplayPath

`func (o *GateTask) GetDisplayPath() string`

GetDisplayPath returns the DisplayPath field if non-nil, zero value otherwise.

### GetDisplayPathOk

`func (o *GateTask) GetDisplayPathOk() (*string, bool)`

GetDisplayPathOk returns a tuple with the DisplayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPath

`func (o *GateTask) SetDisplayPath(v string)`

SetDisplayPath sets DisplayPath field to given value.

### HasDisplayPath

`func (o *GateTask) HasDisplayPath() bool`

HasDisplayPath returns a boolean if a field has been set.

### GetReleaseOwner

`func (o *GateTask) GetReleaseOwner() string`

GetReleaseOwner returns the ReleaseOwner field if non-nil, zero value otherwise.

### GetReleaseOwnerOk

`func (o *GateTask) GetReleaseOwnerOk() (*string, bool)`

GetReleaseOwnerOk returns a tuple with the ReleaseOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseOwner

`func (o *GateTask) SetReleaseOwner(v string)`

SetReleaseOwner sets ReleaseOwner field to given value.

### HasReleaseOwner

`func (o *GateTask) HasReleaseOwner() bool`

HasReleaseOwner returns a boolean if a field has been set.

### GetAllTasks

`func (o *GateTask) GetAllTasks() []Task`

GetAllTasks returns the AllTasks field if non-nil, zero value otherwise.

### GetAllTasksOk

`func (o *GateTask) GetAllTasksOk() (*[]Task, bool)`

GetAllTasksOk returns a tuple with the AllTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTasks

`func (o *GateTask) SetAllTasks(v []Task)`

SetAllTasks sets AllTasks field to given value.

### HasAllTasks

`func (o *GateTask) HasAllTasks() bool`

HasAllTasks returns a boolean if a field has been set.

### GetChildren

`func (o *GateTask) GetChildren() []PlanItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *GateTask) GetChildrenOk() (*[]PlanItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *GateTask) SetChildren(v []PlanItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *GateTask) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetInputVariables

`func (o *GateTask) GetInputVariables() []Variable`

GetInputVariables returns the InputVariables field if non-nil, zero value otherwise.

### GetInputVariablesOk

`func (o *GateTask) GetInputVariablesOk() (*[]Variable, bool)`

GetInputVariablesOk returns a tuple with the InputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVariables

`func (o *GateTask) SetInputVariables(v []Variable)`

SetInputVariables sets InputVariables field to given value.

### HasInputVariables

`func (o *GateTask) HasInputVariables() bool`

HasInputVariables returns a boolean if a field has been set.

### GetReferencedVariables

`func (o *GateTask) GetReferencedVariables() []Variable`

GetReferencedVariables returns the ReferencedVariables field if non-nil, zero value otherwise.

### GetReferencedVariablesOk

`func (o *GateTask) GetReferencedVariablesOk() (*[]Variable, bool)`

GetReferencedVariablesOk returns a tuple with the ReferencedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedVariables

`func (o *GateTask) SetReferencedVariables(v []Variable)`

SetReferencedVariables sets ReferencedVariables field to given value.

### HasReferencedVariables

`func (o *GateTask) HasReferencedVariables() bool`

HasReferencedVariables returns a boolean if a field has been set.

### GetUnboundRequiredVariables

`func (o *GateTask) GetUnboundRequiredVariables() []string`

GetUnboundRequiredVariables returns the UnboundRequiredVariables field if non-nil, zero value otherwise.

### GetUnboundRequiredVariablesOk

`func (o *GateTask) GetUnboundRequiredVariablesOk() (*[]string, bool)`

GetUnboundRequiredVariablesOk returns a tuple with the UnboundRequiredVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnboundRequiredVariables

`func (o *GateTask) SetUnboundRequiredVariables(v []string)`

SetUnboundRequiredVariables sets UnboundRequiredVariables field to given value.

### HasUnboundRequiredVariables

`func (o *GateTask) HasUnboundRequiredVariables() bool`

HasUnboundRequiredVariables returns a boolean if a field has been set.

### GetAutomated

`func (o *GateTask) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *GateTask) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *GateTask) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *GateTask) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetTaskType

`func (o *GateTask) GetTaskType() map[string]interface{}`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *GateTask) GetTaskTypeOk() (*map[string]interface{}, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *GateTask) SetTaskType(v map[string]interface{})`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *GateTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetDueSoon

`func (o *GateTask) GetDueSoon() bool`

GetDueSoon returns the DueSoon field if non-nil, zero value otherwise.

### GetDueSoonOk

`func (o *GateTask) GetDueSoonOk() (*bool, bool)`

GetDueSoonOk returns a tuple with the DueSoon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueSoon

`func (o *GateTask) SetDueSoon(v bool)`

SetDueSoon sets DueSoon field to given value.

### HasDueSoon

`func (o *GateTask) HasDueSoon() bool`

HasDueSoon returns a boolean if a field has been set.

### GetElapsedDurationFraction

`func (o *GateTask) GetElapsedDurationFraction() float64`

GetElapsedDurationFraction returns the ElapsedDurationFraction field if non-nil, zero value otherwise.

### GetElapsedDurationFractionOk

`func (o *GateTask) GetElapsedDurationFractionOk() (*float64, bool)`

GetElapsedDurationFractionOk returns a tuple with the ElapsedDurationFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedDurationFraction

`func (o *GateTask) SetElapsedDurationFraction(v float64)`

SetElapsedDurationFraction sets ElapsedDurationFraction field to given value.

### HasElapsedDurationFraction

`func (o *GateTask) HasElapsedDurationFraction() bool`

HasElapsedDurationFraction returns a boolean if a field has been set.

### GetUrl

`func (o *GateTask) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GateTask) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GateTask) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GateTask) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetConditions

`func (o *GateTask) GetConditions() []GateCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GateTask) GetConditionsOk() (*[]GateCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GateTask) SetConditions(v []GateCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *GateTask) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDependencies

`func (o *GateTask) GetDependencies() []Dependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *GateTask) GetDependenciesOk() (*[]Dependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *GateTask) SetDependencies(v []Dependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *GateTask) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetOpen

`func (o *GateTask) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *GateTask) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *GateTask) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *GateTask) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetOpenInAdvance

`func (o *GateTask) GetOpenInAdvance() bool`

GetOpenInAdvance returns the OpenInAdvance field if non-nil, zero value otherwise.

### GetOpenInAdvanceOk

`func (o *GateTask) GetOpenInAdvanceOk() (*bool, bool)`

GetOpenInAdvanceOk returns a tuple with the OpenInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInAdvance

`func (o *GateTask) SetOpenInAdvance(v bool)`

SetOpenInAdvance sets OpenInAdvance field to given value.

### HasOpenInAdvance

`func (o *GateTask) HasOpenInAdvance() bool`

HasOpenInAdvance returns a boolean if a field has been set.

### GetCompletable

`func (o *GateTask) GetCompletable() bool`

GetCompletable returns the Completable field if non-nil, zero value otherwise.

### GetCompletableOk

`func (o *GateTask) GetCompletableOk() (*bool, bool)`

GetCompletableOk returns a tuple with the Completable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletable

`func (o *GateTask) SetCompletable(v bool)`

SetCompletable sets Completable field to given value.

### HasCompletable

`func (o *GateTask) HasCompletable() bool`

HasCompletable returns a boolean if a field has been set.

### GetAbortedDependencyTitles

`func (o *GateTask) GetAbortedDependencyTitles() string`

GetAbortedDependencyTitles returns the AbortedDependencyTitles field if non-nil, zero value otherwise.

### GetAbortedDependencyTitlesOk

`func (o *GateTask) GetAbortedDependencyTitlesOk() (*string, bool)`

GetAbortedDependencyTitlesOk returns a tuple with the AbortedDependencyTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortedDependencyTitles

`func (o *GateTask) SetAbortedDependencyTitles(v string)`

SetAbortedDependencyTitles sets AbortedDependencyTitles field to given value.

### HasAbortedDependencyTitles

`func (o *GateTask) HasAbortedDependencyTitles() bool`

HasAbortedDependencyTitles returns a boolean if a field has been set.

### GetVariableUsages

`func (o *GateTask) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *GateTask) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *GateTask) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *GateTask) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


