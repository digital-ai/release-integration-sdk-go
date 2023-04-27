# UserInputTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledStartDate** | Pointer to **string** |  | [optional] 
**FlagStatus** | Pointer to [**FlagStatus**](FlagStatus.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**PlannedDuration** | Pointer to **int32** |  | [optional] 
**FlagComment** | Pointer to **string** |  | [optional] 
**OverdueNotified** | Pointer to **bool** |  | [optional] 
**Flagged** | Pointer to **bool** |  | [optional] 
**StartOrScheduledDate** | Pointer to **string** |  | [optional] 
**EndOrDueDate** | Pointer to **string** |  | [optional] 
**Overdue** | Pointer to **bool** |  | [optional] 
**OrCalculateDueDate** | Pointer to **NullableString** |  | [optional] 
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
**OriginalScheduledStartDate** | Pointer to **string** |  | [optional] 
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
**UnboundRequiredVariables** | Pointer to **[]string** |  | [optional] 
**Automated** | Pointer to **bool** |  | [optional] 
**TaskType** | Pointer to **map[string]interface{}** |  | [optional] 
**DueSoon** | Pointer to **bool** |  | [optional] 
**ElapsedDurationFraction** | Pointer to **float64** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**ReferencedVariables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 

## Methods

### NewUserInputTask

`func NewUserInputTask() *UserInputTask`

NewUserInputTask instantiates a new UserInputTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInputTaskWithDefaults

`func NewUserInputTaskWithDefaults() *UserInputTask`

NewUserInputTaskWithDefaults instantiates a new UserInputTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledStartDate

`func (o *UserInputTask) GetScheduledStartDate() string`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *UserInputTask) GetScheduledStartDateOk() (*string, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *UserInputTask) SetScheduledStartDate(v string)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *UserInputTask) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetFlagStatus

`func (o *UserInputTask) GetFlagStatus() FlagStatus`

GetFlagStatus returns the FlagStatus field if non-nil, zero value otherwise.

### GetFlagStatusOk

`func (o *UserInputTask) GetFlagStatusOk() (*FlagStatus, bool)`

GetFlagStatusOk returns a tuple with the FlagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagStatus

`func (o *UserInputTask) SetFlagStatus(v FlagStatus)`

SetFlagStatus sets FlagStatus field to given value.

### HasFlagStatus

`func (o *UserInputTask) HasFlagStatus() bool`

HasFlagStatus returns a boolean if a field has been set.

### GetTitle

`func (o *UserInputTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserInputTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserInputTask) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserInputTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UserInputTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserInputTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserInputTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserInputTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *UserInputTask) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UserInputTask) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UserInputTask) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UserInputTask) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDueDate

`func (o *UserInputTask) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *UserInputTask) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *UserInputTask) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *UserInputTask) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *UserInputTask) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UserInputTask) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UserInputTask) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UserInputTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UserInputTask) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UserInputTask) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UserInputTask) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UserInputTask) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPlannedDuration

`func (o *UserInputTask) GetPlannedDuration() int32`

GetPlannedDuration returns the PlannedDuration field if non-nil, zero value otherwise.

### GetPlannedDurationOk

`func (o *UserInputTask) GetPlannedDurationOk() (*int32, bool)`

GetPlannedDurationOk returns a tuple with the PlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDuration

`func (o *UserInputTask) SetPlannedDuration(v int32)`

SetPlannedDuration sets PlannedDuration field to given value.

### HasPlannedDuration

`func (o *UserInputTask) HasPlannedDuration() bool`

HasPlannedDuration returns a boolean if a field has been set.

### GetFlagComment

`func (o *UserInputTask) GetFlagComment() string`

GetFlagComment returns the FlagComment field if non-nil, zero value otherwise.

### GetFlagCommentOk

`func (o *UserInputTask) GetFlagCommentOk() (*string, bool)`

GetFlagCommentOk returns a tuple with the FlagComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagComment

`func (o *UserInputTask) SetFlagComment(v string)`

SetFlagComment sets FlagComment field to given value.

### HasFlagComment

`func (o *UserInputTask) HasFlagComment() bool`

HasFlagComment returns a boolean if a field has been set.

### GetOverdueNotified

`func (o *UserInputTask) GetOverdueNotified() bool`

GetOverdueNotified returns the OverdueNotified field if non-nil, zero value otherwise.

### GetOverdueNotifiedOk

`func (o *UserInputTask) GetOverdueNotifiedOk() (*bool, bool)`

GetOverdueNotifiedOk returns a tuple with the OverdueNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueNotified

`func (o *UserInputTask) SetOverdueNotified(v bool)`

SetOverdueNotified sets OverdueNotified field to given value.

### HasOverdueNotified

`func (o *UserInputTask) HasOverdueNotified() bool`

HasOverdueNotified returns a boolean if a field has been set.

### GetFlagged

`func (o *UserInputTask) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *UserInputTask) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *UserInputTask) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *UserInputTask) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetStartOrScheduledDate

`func (o *UserInputTask) GetStartOrScheduledDate() string`

GetStartOrScheduledDate returns the StartOrScheduledDate field if non-nil, zero value otherwise.

### GetStartOrScheduledDateOk

`func (o *UserInputTask) GetStartOrScheduledDateOk() (*string, bool)`

GetStartOrScheduledDateOk returns a tuple with the StartOrScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOrScheduledDate

`func (o *UserInputTask) SetStartOrScheduledDate(v string)`

SetStartOrScheduledDate sets StartOrScheduledDate field to given value.

### HasStartOrScheduledDate

`func (o *UserInputTask) HasStartOrScheduledDate() bool`

HasStartOrScheduledDate returns a boolean if a field has been set.

### GetEndOrDueDate

`func (o *UserInputTask) GetEndOrDueDate() string`

GetEndOrDueDate returns the EndOrDueDate field if non-nil, zero value otherwise.

### GetEndOrDueDateOk

`func (o *UserInputTask) GetEndOrDueDateOk() (*string, bool)`

GetEndOrDueDateOk returns a tuple with the EndOrDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOrDueDate

`func (o *UserInputTask) SetEndOrDueDate(v string)`

SetEndOrDueDate sets EndOrDueDate field to given value.

### HasEndOrDueDate

`func (o *UserInputTask) HasEndOrDueDate() bool`

HasEndOrDueDate returns a boolean if a field has been set.

### GetOverdue

`func (o *UserInputTask) GetOverdue() bool`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *UserInputTask) GetOverdueOk() (*bool, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *UserInputTask) SetOverdue(v bool)`

SetOverdue sets Overdue field to given value.

### HasOverdue

`func (o *UserInputTask) HasOverdue() bool`

HasOverdue returns a boolean if a field has been set.

### GetOrCalculateDueDate

`func (o *UserInputTask) GetOrCalculateDueDate() string`

GetOrCalculateDueDate returns the OrCalculateDueDate field if non-nil, zero value otherwise.

### GetOrCalculateDueDateOk

`func (o *UserInputTask) GetOrCalculateDueDateOk() (*string, bool)`

GetOrCalculateDueDateOk returns a tuple with the OrCalculateDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrCalculateDueDate

`func (o *UserInputTask) SetOrCalculateDueDate(v string)`

SetOrCalculateDueDate sets OrCalculateDueDate field to given value.

### HasOrCalculateDueDate

`func (o *UserInputTask) HasOrCalculateDueDate() bool`

HasOrCalculateDueDate returns a boolean if a field has been set.

### SetOrCalculateDueDateNil

`func (o *UserInputTask) SetOrCalculateDueDateNil(b bool)`

 SetOrCalculateDueDateNil sets the value for OrCalculateDueDate to be an explicit nil

### UnsetOrCalculateDueDate
`func (o *UserInputTask) UnsetOrCalculateDueDate()`

UnsetOrCalculateDueDate ensures that no value is present for OrCalculateDueDate, not even an explicit nil
### GetComputedPlannedDuration

`func (o *UserInputTask) GetComputedPlannedDuration() map[string]interface{}`

GetComputedPlannedDuration returns the ComputedPlannedDuration field if non-nil, zero value otherwise.

### GetComputedPlannedDurationOk

`func (o *UserInputTask) GetComputedPlannedDurationOk() (*map[string]interface{}, bool)`

GetComputedPlannedDurationOk returns a tuple with the ComputedPlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPlannedDuration

`func (o *UserInputTask) SetComputedPlannedDuration(v map[string]interface{})`

SetComputedPlannedDuration sets ComputedPlannedDuration field to given value.

### HasComputedPlannedDuration

`func (o *UserInputTask) HasComputedPlannedDuration() bool`

HasComputedPlannedDuration returns a boolean if a field has been set.

### GetActualDuration

`func (o *UserInputTask) GetActualDuration() map[string]interface{}`

GetActualDuration returns the ActualDuration field if non-nil, zero value otherwise.

### GetActualDurationOk

`func (o *UserInputTask) GetActualDurationOk() (*map[string]interface{}, bool)`

GetActualDurationOk returns a tuple with the ActualDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDuration

`func (o *UserInputTask) SetActualDuration(v map[string]interface{})`

SetActualDuration sets ActualDuration field to given value.

### HasActualDuration

`func (o *UserInputTask) HasActualDuration() bool`

HasActualDuration returns a boolean if a field has been set.

### GetReleaseUid

`func (o *UserInputTask) GetReleaseUid() int32`

GetReleaseUid returns the ReleaseUid field if non-nil, zero value otherwise.

### GetReleaseUidOk

`func (o *UserInputTask) GetReleaseUidOk() (*int32, bool)`

GetReleaseUidOk returns a tuple with the ReleaseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseUid

`func (o *UserInputTask) SetReleaseUid(v int32)`

SetReleaseUid sets ReleaseUid field to given value.

### HasReleaseUid

`func (o *UserInputTask) HasReleaseUid() bool`

HasReleaseUid returns a boolean if a field has been set.

### GetCiUid

`func (o *UserInputTask) GetCiUid() int32`

GetCiUid returns the CiUid field if non-nil, zero value otherwise.

### GetCiUidOk

`func (o *UserInputTask) GetCiUidOk() (*int32, bool)`

GetCiUidOk returns a tuple with the CiUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiUid

`func (o *UserInputTask) SetCiUid(v int32)`

SetCiUid sets CiUid field to given value.

### HasCiUid

`func (o *UserInputTask) HasCiUid() bool`

HasCiUid returns a boolean if a field has been set.

### GetComments

`func (o *UserInputTask) GetComments() []Comment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *UserInputTask) GetCommentsOk() (*[]Comment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *UserInputTask) SetComments(v []Comment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *UserInputTask) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetContainer

`func (o *UserInputTask) GetContainer() TaskContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *UserInputTask) GetContainerOk() (*TaskContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *UserInputTask) SetContainer(v TaskContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *UserInputTask) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetFacets

`func (o *UserInputTask) GetFacets() []Facet`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *UserInputTask) GetFacetsOk() (*[]Facet, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *UserInputTask) SetFacets(v []Facet)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *UserInputTask) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### GetAttachments

`func (o *UserInputTask) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *UserInputTask) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *UserInputTask) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *UserInputTask) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetStatus

`func (o *UserInputTask) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserInputTask) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserInputTask) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserInputTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTeam

`func (o *UserInputTask) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *UserInputTask) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *UserInputTask) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *UserInputTask) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetWatchers

`func (o *UserInputTask) GetWatchers() []string`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *UserInputTask) GetWatchersOk() (*[]string, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *UserInputTask) SetWatchers(v []string)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *UserInputTask) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.

### GetWaitForScheduledStartDate

`func (o *UserInputTask) GetWaitForScheduledStartDate() bool`

GetWaitForScheduledStartDate returns the WaitForScheduledStartDate field if non-nil, zero value otherwise.

### GetWaitForScheduledStartDateOk

`func (o *UserInputTask) GetWaitForScheduledStartDateOk() (*bool, bool)`

GetWaitForScheduledStartDateOk returns a tuple with the WaitForScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForScheduledStartDate

`func (o *UserInputTask) SetWaitForScheduledStartDate(v bool)`

SetWaitForScheduledStartDate sets WaitForScheduledStartDate field to given value.

### HasWaitForScheduledStartDate

`func (o *UserInputTask) HasWaitForScheduledStartDate() bool`

HasWaitForScheduledStartDate returns a boolean if a field has been set.

### GetDelayDuringBlackout

`func (o *UserInputTask) GetDelayDuringBlackout() bool`

GetDelayDuringBlackout returns the DelayDuringBlackout field if non-nil, zero value otherwise.

### GetDelayDuringBlackoutOk

`func (o *UserInputTask) GetDelayDuringBlackoutOk() (*bool, bool)`

GetDelayDuringBlackoutOk returns a tuple with the DelayDuringBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDuringBlackout

`func (o *UserInputTask) SetDelayDuringBlackout(v bool)`

SetDelayDuringBlackout sets DelayDuringBlackout field to given value.

### HasDelayDuringBlackout

`func (o *UserInputTask) HasDelayDuringBlackout() bool`

HasDelayDuringBlackout returns a boolean if a field has been set.

### GetPostponedDueToBlackout

`func (o *UserInputTask) GetPostponedDueToBlackout() bool`

GetPostponedDueToBlackout returns the PostponedDueToBlackout field if non-nil, zero value otherwise.

### GetPostponedDueToBlackoutOk

`func (o *UserInputTask) GetPostponedDueToBlackoutOk() (*bool, bool)`

GetPostponedDueToBlackoutOk returns a tuple with the PostponedDueToBlackout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostponedDueToBlackout

`func (o *UserInputTask) SetPostponedDueToBlackout(v bool)`

SetPostponedDueToBlackout sets PostponedDueToBlackout field to given value.

### HasPostponedDueToBlackout

`func (o *UserInputTask) HasPostponedDueToBlackout() bool`

HasPostponedDueToBlackout returns a boolean if a field has been set.

### GetPostponedUntilEnvironmentsAreReserved

`func (o *UserInputTask) GetPostponedUntilEnvironmentsAreReserved() bool`

GetPostponedUntilEnvironmentsAreReserved returns the PostponedUntilEnvironmentsAreReserved field if non-nil, zero value otherwise.

### GetPostponedUntilEnvironmentsAreReservedOk

`func (o *UserInputTask) GetPostponedUntilEnvironmentsAreReservedOk() (*bool, bool)`

GetPostponedUntilEnvironmentsAreReservedOk returns a tuple with the PostponedUntilEnvironmentsAreReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostponedUntilEnvironmentsAreReserved

`func (o *UserInputTask) SetPostponedUntilEnvironmentsAreReserved(v bool)`

SetPostponedUntilEnvironmentsAreReserved sets PostponedUntilEnvironmentsAreReserved field to given value.

### HasPostponedUntilEnvironmentsAreReserved

`func (o *UserInputTask) HasPostponedUntilEnvironmentsAreReserved() bool`

HasPostponedUntilEnvironmentsAreReserved returns a boolean if a field has been set.

### GetOriginalScheduledStartDate

`func (o *UserInputTask) GetOriginalScheduledStartDate() string`

GetOriginalScheduledStartDate returns the OriginalScheduledStartDate field if non-nil, zero value otherwise.

### GetOriginalScheduledStartDateOk

`func (o *UserInputTask) GetOriginalScheduledStartDateOk() (*string, bool)`

GetOriginalScheduledStartDateOk returns a tuple with the OriginalScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalScheduledStartDate

`func (o *UserInputTask) SetOriginalScheduledStartDate(v string)`

SetOriginalScheduledStartDate sets OriginalScheduledStartDate field to given value.

### HasOriginalScheduledStartDate

`func (o *UserInputTask) HasOriginalScheduledStartDate() bool`

HasOriginalScheduledStartDate returns a boolean if a field has been set.

### GetHasBeenFlagged

`func (o *UserInputTask) GetHasBeenFlagged() bool`

GetHasBeenFlagged returns the HasBeenFlagged field if non-nil, zero value otherwise.

### GetHasBeenFlaggedOk

`func (o *UserInputTask) GetHasBeenFlaggedOk() (*bool, bool)`

GetHasBeenFlaggedOk returns a tuple with the HasBeenFlagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenFlagged

`func (o *UserInputTask) SetHasBeenFlagged(v bool)`

SetHasBeenFlagged sets HasBeenFlagged field to given value.

### HasHasBeenFlagged

`func (o *UserInputTask) HasHasBeenFlagged() bool`

HasHasBeenFlagged returns a boolean if a field has been set.

### GetHasBeenDelayed

`func (o *UserInputTask) GetHasBeenDelayed() bool`

GetHasBeenDelayed returns the HasBeenDelayed field if non-nil, zero value otherwise.

### GetHasBeenDelayedOk

`func (o *UserInputTask) GetHasBeenDelayedOk() (*bool, bool)`

GetHasBeenDelayedOk returns a tuple with the HasBeenDelayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenDelayed

`func (o *UserInputTask) SetHasBeenDelayed(v bool)`

SetHasBeenDelayed sets HasBeenDelayed field to given value.

### HasHasBeenDelayed

`func (o *UserInputTask) HasHasBeenDelayed() bool`

HasHasBeenDelayed returns a boolean if a field has been set.

### GetPrecondition

`func (o *UserInputTask) GetPrecondition() string`

GetPrecondition returns the Precondition field if non-nil, zero value otherwise.

### GetPreconditionOk

`func (o *UserInputTask) GetPreconditionOk() (*string, bool)`

GetPreconditionOk returns a tuple with the Precondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecondition

`func (o *UserInputTask) SetPrecondition(v string)`

SetPrecondition sets Precondition field to given value.

### HasPrecondition

`func (o *UserInputTask) HasPrecondition() bool`

HasPrecondition returns a boolean if a field has been set.

### GetFailureHandler

`func (o *UserInputTask) GetFailureHandler() string`

GetFailureHandler returns the FailureHandler field if non-nil, zero value otherwise.

### GetFailureHandlerOk

`func (o *UserInputTask) GetFailureHandlerOk() (*string, bool)`

GetFailureHandlerOk returns a tuple with the FailureHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureHandler

`func (o *UserInputTask) SetFailureHandler(v string)`

SetFailureHandler sets FailureHandler field to given value.

### HasFailureHandler

`func (o *UserInputTask) HasFailureHandler() bool`

HasFailureHandler returns a boolean if a field has been set.

### GetTaskFailureHandlerEnabled

`func (o *UserInputTask) GetTaskFailureHandlerEnabled() bool`

GetTaskFailureHandlerEnabled returns the TaskFailureHandlerEnabled field if non-nil, zero value otherwise.

### GetTaskFailureHandlerEnabledOk

`func (o *UserInputTask) GetTaskFailureHandlerEnabledOk() (*bool, bool)`

GetTaskFailureHandlerEnabledOk returns a tuple with the TaskFailureHandlerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFailureHandlerEnabled

`func (o *UserInputTask) SetTaskFailureHandlerEnabled(v bool)`

SetTaskFailureHandlerEnabled sets TaskFailureHandlerEnabled field to given value.

### HasTaskFailureHandlerEnabled

`func (o *UserInputTask) HasTaskFailureHandlerEnabled() bool`

HasTaskFailureHandlerEnabled returns a boolean if a field has been set.

### GetTaskRecoverOp

`func (o *UserInputTask) GetTaskRecoverOp() TaskRecoverOp`

GetTaskRecoverOp returns the TaskRecoverOp field if non-nil, zero value otherwise.

### GetTaskRecoverOpOk

`func (o *UserInputTask) GetTaskRecoverOpOk() (*TaskRecoverOp, bool)`

GetTaskRecoverOpOk returns a tuple with the TaskRecoverOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRecoverOp

`func (o *UserInputTask) SetTaskRecoverOp(v TaskRecoverOp)`

SetTaskRecoverOp sets TaskRecoverOp field to given value.

### HasTaskRecoverOp

`func (o *UserInputTask) HasTaskRecoverOp() bool`

HasTaskRecoverOp returns a boolean if a field has been set.

### GetFailuresCount

`func (o *UserInputTask) GetFailuresCount() int32`

GetFailuresCount returns the FailuresCount field if non-nil, zero value otherwise.

### GetFailuresCountOk

`func (o *UserInputTask) GetFailuresCountOk() (*int32, bool)`

GetFailuresCountOk returns a tuple with the FailuresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailuresCount

`func (o *UserInputTask) SetFailuresCount(v int32)`

SetFailuresCount sets FailuresCount field to given value.

### HasFailuresCount

`func (o *UserInputTask) HasFailuresCount() bool`

HasFailuresCount returns a boolean if a field has been set.

### GetExecutionId

`func (o *UserInputTask) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *UserInputTask) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *UserInputTask) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *UserInputTask) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetVariableMapping

`func (o *UserInputTask) GetVariableMapping() map[string]string`

GetVariableMapping returns the VariableMapping field if non-nil, zero value otherwise.

### GetVariableMappingOk

`func (o *UserInputTask) GetVariableMappingOk() (*map[string]string, bool)`

GetVariableMappingOk returns a tuple with the VariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMapping

`func (o *UserInputTask) SetVariableMapping(v map[string]string)`

SetVariableMapping sets VariableMapping field to given value.

### HasVariableMapping

`func (o *UserInputTask) HasVariableMapping() bool`

HasVariableMapping returns a boolean if a field has been set.

### GetExternalVariableMapping

`func (o *UserInputTask) GetExternalVariableMapping() map[string]string`

GetExternalVariableMapping returns the ExternalVariableMapping field if non-nil, zero value otherwise.

### GetExternalVariableMappingOk

`func (o *UserInputTask) GetExternalVariableMappingOk() (*map[string]string, bool)`

GetExternalVariableMappingOk returns a tuple with the ExternalVariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVariableMapping

`func (o *UserInputTask) SetExternalVariableMapping(v map[string]string)`

SetExternalVariableMapping sets ExternalVariableMapping field to given value.

### HasExternalVariableMapping

`func (o *UserInputTask) HasExternalVariableMapping() bool`

HasExternalVariableMapping returns a boolean if a field has been set.

### GetMaxCommentSize

`func (o *UserInputTask) GetMaxCommentSize() int32`

GetMaxCommentSize returns the MaxCommentSize field if non-nil, zero value otherwise.

### GetMaxCommentSizeOk

`func (o *UserInputTask) GetMaxCommentSizeOk() (*int32, bool)`

GetMaxCommentSizeOk returns a tuple with the MaxCommentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCommentSize

`func (o *UserInputTask) SetMaxCommentSize(v int32)`

SetMaxCommentSize sets MaxCommentSize field to given value.

### HasMaxCommentSize

`func (o *UserInputTask) HasMaxCommentSize() bool`

HasMaxCommentSize returns a boolean if a field has been set.

### GetTags

`func (o *UserInputTask) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserInputTask) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserInputTask) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UserInputTask) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfigurationUri

`func (o *UserInputTask) GetConfigurationUri() string`

GetConfigurationUri returns the ConfigurationUri field if non-nil, zero value otherwise.

### GetConfigurationUriOk

`func (o *UserInputTask) GetConfigurationUriOk() (*string, bool)`

GetConfigurationUriOk returns a tuple with the ConfigurationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUri

`func (o *UserInputTask) SetConfigurationUri(v string)`

SetConfigurationUri sets ConfigurationUri field to given value.

### HasConfigurationUri

`func (o *UserInputTask) HasConfigurationUri() bool`

HasConfigurationUri returns a boolean if a field has been set.

### GetDueSoonNotified

`func (o *UserInputTask) GetDueSoonNotified() bool`

GetDueSoonNotified returns the DueSoonNotified field if non-nil, zero value otherwise.

### GetDueSoonNotifiedOk

`func (o *UserInputTask) GetDueSoonNotifiedOk() (*bool, bool)`

GetDueSoonNotifiedOk returns a tuple with the DueSoonNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueSoonNotified

`func (o *UserInputTask) SetDueSoonNotified(v bool)`

SetDueSoonNotified sets DueSoonNotified field to given value.

### HasDueSoonNotified

`func (o *UserInputTask) HasDueSoonNotified() bool`

HasDueSoonNotified returns a boolean if a field has been set.

### GetLocked

`func (o *UserInputTask) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UserInputTask) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UserInputTask) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UserInputTask) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetCheckAttributes

`func (o *UserInputTask) GetCheckAttributes() bool`

GetCheckAttributes returns the CheckAttributes field if non-nil, zero value otherwise.

### GetCheckAttributesOk

`func (o *UserInputTask) GetCheckAttributesOk() (*bool, bool)`

GetCheckAttributesOk returns a tuple with the CheckAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAttributes

`func (o *UserInputTask) SetCheckAttributes(v bool)`

SetCheckAttributes sets CheckAttributes field to given value.

### HasCheckAttributes

`func (o *UserInputTask) HasCheckAttributes() bool`

HasCheckAttributes returns a boolean if a field has been set.

### GetAbortScript

`func (o *UserInputTask) GetAbortScript() string`

GetAbortScript returns the AbortScript field if non-nil, zero value otherwise.

### GetAbortScriptOk

`func (o *UserInputTask) GetAbortScriptOk() (*string, bool)`

GetAbortScriptOk returns a tuple with the AbortScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortScript

`func (o *UserInputTask) SetAbortScript(v string)`

SetAbortScript sets AbortScript field to given value.

### HasAbortScript

`func (o *UserInputTask) HasAbortScript() bool`

HasAbortScript returns a boolean if a field has been set.

### GetPhase

`func (o *UserInputTask) GetPhase() Phase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *UserInputTask) GetPhaseOk() (*Phase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *UserInputTask) SetPhase(v Phase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *UserInputTask) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetBlackoutMetadata

`func (o *UserInputTask) GetBlackoutMetadata() BlackoutMetadata`

GetBlackoutMetadata returns the BlackoutMetadata field if non-nil, zero value otherwise.

### GetBlackoutMetadataOk

`func (o *UserInputTask) GetBlackoutMetadataOk() (*BlackoutMetadata, bool)`

GetBlackoutMetadataOk returns a tuple with the BlackoutMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutMetadata

`func (o *UserInputTask) SetBlackoutMetadata(v BlackoutMetadata)`

SetBlackoutMetadata sets BlackoutMetadata field to given value.

### HasBlackoutMetadata

`func (o *UserInputTask) HasBlackoutMetadata() bool`

HasBlackoutMetadata returns a boolean if a field has been set.

### GetFlaggedCount

`func (o *UserInputTask) GetFlaggedCount() int32`

GetFlaggedCount returns the FlaggedCount field if non-nil, zero value otherwise.

### GetFlaggedCountOk

`func (o *UserInputTask) GetFlaggedCountOk() (*int32, bool)`

GetFlaggedCountOk returns a tuple with the FlaggedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedCount

`func (o *UserInputTask) SetFlaggedCount(v int32)`

SetFlaggedCount sets FlaggedCount field to given value.

### HasFlaggedCount

`func (o *UserInputTask) HasFlaggedCount() bool`

HasFlaggedCount returns a boolean if a field has been set.

### GetDelayedCount

`func (o *UserInputTask) GetDelayedCount() int32`

GetDelayedCount returns the DelayedCount field if non-nil, zero value otherwise.

### GetDelayedCountOk

`func (o *UserInputTask) GetDelayedCountOk() (*int32, bool)`

GetDelayedCountOk returns a tuple with the DelayedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedCount

`func (o *UserInputTask) SetDelayedCount(v int32)`

SetDelayedCount sets DelayedCount field to given value.

### HasDelayedCount

`func (o *UserInputTask) HasDelayedCount() bool`

HasDelayedCount returns a boolean if a field has been set.

### GetDone

`func (o *UserInputTask) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *UserInputTask) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *UserInputTask) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *UserInputTask) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetDoneInAdvance

`func (o *UserInputTask) GetDoneInAdvance() bool`

GetDoneInAdvance returns the DoneInAdvance field if non-nil, zero value otherwise.

### GetDoneInAdvanceOk

`func (o *UserInputTask) GetDoneInAdvanceOk() (*bool, bool)`

GetDoneInAdvanceOk returns a tuple with the DoneInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneInAdvance

`func (o *UserInputTask) SetDoneInAdvance(v bool)`

SetDoneInAdvance sets DoneInAdvance field to given value.

### HasDoneInAdvance

`func (o *UserInputTask) HasDoneInAdvance() bool`

HasDoneInAdvance returns a boolean if a field has been set.

### GetDefunct

`func (o *UserInputTask) GetDefunct() bool`

GetDefunct returns the Defunct field if non-nil, zero value otherwise.

### GetDefunctOk

`func (o *UserInputTask) GetDefunctOk() (*bool, bool)`

GetDefunctOk returns a tuple with the Defunct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunct

`func (o *UserInputTask) SetDefunct(v bool)`

SetDefunct sets Defunct field to given value.

### HasDefunct

`func (o *UserInputTask) HasDefunct() bool`

HasDefunct returns a boolean if a field has been set.

### GetUpdatable

`func (o *UserInputTask) GetUpdatable() bool`

GetUpdatable returns the Updatable field if non-nil, zero value otherwise.

### GetUpdatableOk

`func (o *UserInputTask) GetUpdatableOk() (*bool, bool)`

GetUpdatableOk returns a tuple with the Updatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatable

`func (o *UserInputTask) SetUpdatable(v bool)`

SetUpdatable sets Updatable field to given value.

### HasUpdatable

`func (o *UserInputTask) HasUpdatable() bool`

HasUpdatable returns a boolean if a field has been set.

### GetAborted

`func (o *UserInputTask) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *UserInputTask) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *UserInputTask) SetAborted(v bool)`

SetAborted sets Aborted field to given value.

### HasAborted

`func (o *UserInputTask) HasAborted() bool`

HasAborted returns a boolean if a field has been set.

### GetNotYetReached

`func (o *UserInputTask) GetNotYetReached() bool`

GetNotYetReached returns the NotYetReached field if non-nil, zero value otherwise.

### GetNotYetReachedOk

`func (o *UserInputTask) GetNotYetReachedOk() (*bool, bool)`

GetNotYetReachedOk returns a tuple with the NotYetReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotYetReached

`func (o *UserInputTask) SetNotYetReached(v bool)`

SetNotYetReached sets NotYetReached field to given value.

### HasNotYetReached

`func (o *UserInputTask) HasNotYetReached() bool`

HasNotYetReached returns a boolean if a field has been set.

### GetPlanned

`func (o *UserInputTask) GetPlanned() bool`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *UserInputTask) GetPlannedOk() (*bool, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *UserInputTask) SetPlanned(v bool)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *UserInputTask) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.

### GetActive

`func (o *UserInputTask) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserInputTask) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserInputTask) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UserInputTask) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetInProgress

`func (o *UserInputTask) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *UserInputTask) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *UserInputTask) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *UserInputTask) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetPending

`func (o *UserInputTask) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *UserInputTask) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *UserInputTask) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *UserInputTask) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetWaitingForInput

`func (o *UserInputTask) GetWaitingForInput() bool`

GetWaitingForInput returns the WaitingForInput field if non-nil, zero value otherwise.

### GetWaitingForInputOk

`func (o *UserInputTask) GetWaitingForInputOk() (*bool, bool)`

GetWaitingForInputOk returns a tuple with the WaitingForInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingForInput

`func (o *UserInputTask) SetWaitingForInput(v bool)`

SetWaitingForInput sets WaitingForInput field to given value.

### HasWaitingForInput

`func (o *UserInputTask) HasWaitingForInput() bool`

HasWaitingForInput returns a boolean if a field has been set.

### GetFailed

`func (o *UserInputTask) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *UserInputTask) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *UserInputTask) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *UserInputTask) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetFailing

`func (o *UserInputTask) GetFailing() bool`

GetFailing returns the Failing field if non-nil, zero value otherwise.

### GetFailingOk

`func (o *UserInputTask) GetFailingOk() (*bool, bool)`

GetFailingOk returns a tuple with the Failing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailing

`func (o *UserInputTask) SetFailing(v bool)`

SetFailing sets Failing field to given value.

### HasFailing

`func (o *UserInputTask) HasFailing() bool`

HasFailing returns a boolean if a field has been set.

### GetCompletedInAdvance

`func (o *UserInputTask) GetCompletedInAdvance() bool`

GetCompletedInAdvance returns the CompletedInAdvance field if non-nil, zero value otherwise.

### GetCompletedInAdvanceOk

`func (o *UserInputTask) GetCompletedInAdvanceOk() (*bool, bool)`

GetCompletedInAdvanceOk returns a tuple with the CompletedInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedInAdvance

`func (o *UserInputTask) SetCompletedInAdvance(v bool)`

SetCompletedInAdvance sets CompletedInAdvance field to given value.

### HasCompletedInAdvance

`func (o *UserInputTask) HasCompletedInAdvance() bool`

HasCompletedInAdvance returns a boolean if a field has been set.

### GetSkipped

`func (o *UserInputTask) GetSkipped() bool`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *UserInputTask) GetSkippedOk() (*bool, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *UserInputTask) SetSkipped(v bool)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *UserInputTask) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSkippedInAdvance

`func (o *UserInputTask) GetSkippedInAdvance() bool`

GetSkippedInAdvance returns the SkippedInAdvance field if non-nil, zero value otherwise.

### GetSkippedInAdvanceOk

`func (o *UserInputTask) GetSkippedInAdvanceOk() (*bool, bool)`

GetSkippedInAdvanceOk returns a tuple with the SkippedInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedInAdvance

`func (o *UserInputTask) SetSkippedInAdvance(v bool)`

SetSkippedInAdvance sets SkippedInAdvance field to given value.

### HasSkippedInAdvance

`func (o *UserInputTask) HasSkippedInAdvance() bool`

HasSkippedInAdvance returns a boolean if a field has been set.

### GetPreconditionInProgress

`func (o *UserInputTask) GetPreconditionInProgress() bool`

GetPreconditionInProgress returns the PreconditionInProgress field if non-nil, zero value otherwise.

### GetPreconditionInProgressOk

`func (o *UserInputTask) GetPreconditionInProgressOk() (*bool, bool)`

GetPreconditionInProgressOk returns a tuple with the PreconditionInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionInProgress

`func (o *UserInputTask) SetPreconditionInProgress(v bool)`

SetPreconditionInProgress sets PreconditionInProgress field to given value.

### HasPreconditionInProgress

`func (o *UserInputTask) HasPreconditionInProgress() bool`

HasPreconditionInProgress returns a boolean if a field has been set.

### GetFailureHandlerInProgress

`func (o *UserInputTask) GetFailureHandlerInProgress() bool`

GetFailureHandlerInProgress returns the FailureHandlerInProgress field if non-nil, zero value otherwise.

### GetFailureHandlerInProgressOk

`func (o *UserInputTask) GetFailureHandlerInProgressOk() (*bool, bool)`

GetFailureHandlerInProgressOk returns a tuple with the FailureHandlerInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureHandlerInProgress

`func (o *UserInputTask) SetFailureHandlerInProgress(v bool)`

SetFailureHandlerInProgress sets FailureHandlerInProgress field to given value.

### HasFailureHandlerInProgress

`func (o *UserInputTask) HasFailureHandlerInProgress() bool`

HasFailureHandlerInProgress returns a boolean if a field has been set.

### GetAbortScriptInProgress

`func (o *UserInputTask) GetAbortScriptInProgress() bool`

GetAbortScriptInProgress returns the AbortScriptInProgress field if non-nil, zero value otherwise.

### GetAbortScriptInProgressOk

`func (o *UserInputTask) GetAbortScriptInProgressOk() (*bool, bool)`

GetAbortScriptInProgressOk returns a tuple with the AbortScriptInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortScriptInProgress

`func (o *UserInputTask) SetAbortScriptInProgress(v bool)`

SetAbortScriptInProgress sets AbortScriptInProgress field to given value.

### HasAbortScriptInProgress

`func (o *UserInputTask) HasAbortScriptInProgress() bool`

HasAbortScriptInProgress returns a boolean if a field has been set.

### GetFacetInProgress

`func (o *UserInputTask) GetFacetInProgress() bool`

GetFacetInProgress returns the FacetInProgress field if non-nil, zero value otherwise.

### GetFacetInProgressOk

`func (o *UserInputTask) GetFacetInProgressOk() (*bool, bool)`

GetFacetInProgressOk returns a tuple with the FacetInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetInProgress

`func (o *UserInputTask) SetFacetInProgress(v bool)`

SetFacetInProgress sets FacetInProgress field to given value.

### HasFacetInProgress

`func (o *UserInputTask) HasFacetInProgress() bool`

HasFacetInProgress returns a boolean if a field has been set.

### GetMovable

`func (o *UserInputTask) GetMovable() bool`

GetMovable returns the Movable field if non-nil, zero value otherwise.

### GetMovableOk

`func (o *UserInputTask) GetMovableOk() (*bool, bool)`

GetMovableOk returns a tuple with the Movable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovable

`func (o *UserInputTask) SetMovable(v bool)`

SetMovable sets Movable field to given value.

### HasMovable

`func (o *UserInputTask) HasMovable() bool`

HasMovable returns a boolean if a field has been set.

### GetGate

`func (o *UserInputTask) GetGate() bool`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *UserInputTask) GetGateOk() (*bool, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *UserInputTask) SetGate(v bool)`

SetGate sets Gate field to given value.

### HasGate

`func (o *UserInputTask) HasGate() bool`

HasGate returns a boolean if a field has been set.

### GetTaskGroup

`func (o *UserInputTask) GetTaskGroup() bool`

GetTaskGroup returns the TaskGroup field if non-nil, zero value otherwise.

### GetTaskGroupOk

`func (o *UserInputTask) GetTaskGroupOk() (*bool, bool)`

GetTaskGroupOk returns a tuple with the TaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroup

`func (o *UserInputTask) SetTaskGroup(v bool)`

SetTaskGroup sets TaskGroup field to given value.

### HasTaskGroup

`func (o *UserInputTask) HasTaskGroup() bool`

HasTaskGroup returns a boolean if a field has been set.

### GetParallelGroup

`func (o *UserInputTask) GetParallelGroup() bool`

GetParallelGroup returns the ParallelGroup field if non-nil, zero value otherwise.

### GetParallelGroupOk

`func (o *UserInputTask) GetParallelGroupOk() (*bool, bool)`

GetParallelGroupOk returns a tuple with the ParallelGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelGroup

`func (o *UserInputTask) SetParallelGroup(v bool)`

SetParallelGroup sets ParallelGroup field to given value.

### HasParallelGroup

`func (o *UserInputTask) HasParallelGroup() bool`

HasParallelGroup returns a boolean if a field has been set.

### GetPreconditionEnabled

`func (o *UserInputTask) GetPreconditionEnabled() bool`

GetPreconditionEnabled returns the PreconditionEnabled field if non-nil, zero value otherwise.

### GetPreconditionEnabledOk

`func (o *UserInputTask) GetPreconditionEnabledOk() (*bool, bool)`

GetPreconditionEnabledOk returns a tuple with the PreconditionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionEnabled

`func (o *UserInputTask) SetPreconditionEnabled(v bool)`

SetPreconditionEnabled sets PreconditionEnabled field to given value.

### HasPreconditionEnabled

`func (o *UserInputTask) HasPreconditionEnabled() bool`

HasPreconditionEnabled returns a boolean if a field has been set.

### GetFailureHandlerEnabled

`func (o *UserInputTask) GetFailureHandlerEnabled() bool`

GetFailureHandlerEnabled returns the FailureHandlerEnabled field if non-nil, zero value otherwise.

### GetFailureHandlerEnabledOk

`func (o *UserInputTask) GetFailureHandlerEnabledOk() (*bool, bool)`

GetFailureHandlerEnabledOk returns a tuple with the FailureHandlerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureHandlerEnabled

`func (o *UserInputTask) SetFailureHandlerEnabled(v bool)`

SetFailureHandlerEnabled sets FailureHandlerEnabled field to given value.

### HasFailureHandlerEnabled

`func (o *UserInputTask) HasFailureHandlerEnabled() bool`

HasFailureHandlerEnabled returns a boolean if a field has been set.

### GetRelease

`func (o *UserInputTask) GetRelease() Release`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *UserInputTask) GetReleaseOk() (*Release, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *UserInputTask) SetRelease(v Release)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *UserInputTask) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetDisplayPath

`func (o *UserInputTask) GetDisplayPath() string`

GetDisplayPath returns the DisplayPath field if non-nil, zero value otherwise.

### GetDisplayPathOk

`func (o *UserInputTask) GetDisplayPathOk() (*string, bool)`

GetDisplayPathOk returns a tuple with the DisplayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPath

`func (o *UserInputTask) SetDisplayPath(v string)`

SetDisplayPath sets DisplayPath field to given value.

### HasDisplayPath

`func (o *UserInputTask) HasDisplayPath() bool`

HasDisplayPath returns a boolean if a field has been set.

### GetReleaseOwner

`func (o *UserInputTask) GetReleaseOwner() string`

GetReleaseOwner returns the ReleaseOwner field if non-nil, zero value otherwise.

### GetReleaseOwnerOk

`func (o *UserInputTask) GetReleaseOwnerOk() (*string, bool)`

GetReleaseOwnerOk returns a tuple with the ReleaseOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseOwner

`func (o *UserInputTask) SetReleaseOwner(v string)`

SetReleaseOwner sets ReleaseOwner field to given value.

### HasReleaseOwner

`func (o *UserInputTask) HasReleaseOwner() bool`

HasReleaseOwner returns a boolean if a field has been set.

### GetAllTasks

`func (o *UserInputTask) GetAllTasks() []Task`

GetAllTasks returns the AllTasks field if non-nil, zero value otherwise.

### GetAllTasksOk

`func (o *UserInputTask) GetAllTasksOk() (*[]Task, bool)`

GetAllTasksOk returns a tuple with the AllTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTasks

`func (o *UserInputTask) SetAllTasks(v []Task)`

SetAllTasks sets AllTasks field to given value.

### HasAllTasks

`func (o *UserInputTask) HasAllTasks() bool`

HasAllTasks returns a boolean if a field has been set.

### GetChildren

`func (o *UserInputTask) GetChildren() []PlanItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *UserInputTask) GetChildrenOk() (*[]PlanItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *UserInputTask) SetChildren(v []PlanItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *UserInputTask) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetInputVariables

`func (o *UserInputTask) GetInputVariables() []Variable`

GetInputVariables returns the InputVariables field if non-nil, zero value otherwise.

### GetInputVariablesOk

`func (o *UserInputTask) GetInputVariablesOk() (*[]Variable, bool)`

GetInputVariablesOk returns a tuple with the InputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVariables

`func (o *UserInputTask) SetInputVariables(v []Variable)`

SetInputVariables sets InputVariables field to given value.

### HasInputVariables

`func (o *UserInputTask) HasInputVariables() bool`

HasInputVariables returns a boolean if a field has been set.

### GetUnboundRequiredVariables

`func (o *UserInputTask) GetUnboundRequiredVariables() []string`

GetUnboundRequiredVariables returns the UnboundRequiredVariables field if non-nil, zero value otherwise.

### GetUnboundRequiredVariablesOk

`func (o *UserInputTask) GetUnboundRequiredVariablesOk() (*[]string, bool)`

GetUnboundRequiredVariablesOk returns a tuple with the UnboundRequiredVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnboundRequiredVariables

`func (o *UserInputTask) SetUnboundRequiredVariables(v []string)`

SetUnboundRequiredVariables sets UnboundRequiredVariables field to given value.

### HasUnboundRequiredVariables

`func (o *UserInputTask) HasUnboundRequiredVariables() bool`

HasUnboundRequiredVariables returns a boolean if a field has been set.

### GetAutomated

`func (o *UserInputTask) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *UserInputTask) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *UserInputTask) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *UserInputTask) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetTaskType

`func (o *UserInputTask) GetTaskType() map[string]interface{}`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *UserInputTask) GetTaskTypeOk() (*map[string]interface{}, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *UserInputTask) SetTaskType(v map[string]interface{})`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *UserInputTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetDueSoon

`func (o *UserInputTask) GetDueSoon() bool`

GetDueSoon returns the DueSoon field if non-nil, zero value otherwise.

### GetDueSoonOk

`func (o *UserInputTask) GetDueSoonOk() (*bool, bool)`

GetDueSoonOk returns a tuple with the DueSoon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueSoon

`func (o *UserInputTask) SetDueSoon(v bool)`

SetDueSoon sets DueSoon field to given value.

### HasDueSoon

`func (o *UserInputTask) HasDueSoon() bool`

HasDueSoon returns a boolean if a field has been set.

### GetElapsedDurationFraction

`func (o *UserInputTask) GetElapsedDurationFraction() float64`

GetElapsedDurationFraction returns the ElapsedDurationFraction field if non-nil, zero value otherwise.

### GetElapsedDurationFractionOk

`func (o *UserInputTask) GetElapsedDurationFractionOk() (*float64, bool)`

GetElapsedDurationFractionOk returns a tuple with the ElapsedDurationFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedDurationFraction

`func (o *UserInputTask) SetElapsedDurationFraction(v float64)`

SetElapsedDurationFraction sets ElapsedDurationFraction field to given value.

### HasElapsedDurationFraction

`func (o *UserInputTask) HasElapsedDurationFraction() bool`

HasElapsedDurationFraction returns a boolean if a field has been set.

### GetUrl

`func (o *UserInputTask) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserInputTask) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserInputTask) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserInputTask) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVariables

`func (o *UserInputTask) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *UserInputTask) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *UserInputTask) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *UserInputTask) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetReferencedVariables

`func (o *UserInputTask) GetReferencedVariables() []Variable`

GetReferencedVariables returns the ReferencedVariables field if non-nil, zero value otherwise.

### GetReferencedVariablesOk

`func (o *UserInputTask) GetReferencedVariablesOk() (*[]Variable, bool)`

GetReferencedVariablesOk returns a tuple with the ReferencedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedVariables

`func (o *UserInputTask) SetReferencedVariables(v []Variable)`

SetReferencedVariables sets ReferencedVariables field to given value.

### HasReferencedVariables

`func (o *UserInputTask) HasReferencedVariables() bool`

HasReferencedVariables returns a boolean if a field has been set.

### GetVariableUsages

`func (o *UserInputTask) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *UserInputTask) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *UserInputTask) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *UserInputTask) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


