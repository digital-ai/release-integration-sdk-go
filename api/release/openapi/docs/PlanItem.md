# PlanItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
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
**Children** | Pointer to [**[]PlanItem**](PlanItem.md) |  | [optional] 
**Overdue** | Pointer to **bool** |  | [optional] 
**Done** | Pointer to **bool** |  | [optional] 
**Release** | Pointer to [**Release**](Release.md) |  | [optional] 
**ReleaseUid** | Pointer to **int32** |  | [optional] 
**Updatable** | Pointer to **bool** |  | [optional] 
**DisplayPath** | Pointer to **string** |  | [optional] 
**Aborted** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 
**OrCalculateDueDate** | Pointer to **NullableTime** |  | [optional] 
**ComputedPlannedDuration** | Pointer to **map[string]interface{}** |  | [optional] 
**ActualDuration** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPlanItem

`func NewPlanItem() *PlanItem`

NewPlanItem instantiates a new PlanItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanItemWithDefaults

`func NewPlanItemWithDefaults() *PlanItem`

NewPlanItemWithDefaults instantiates a new PlanItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlanItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PlanItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlanItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlanItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlanItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *PlanItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlanItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlanItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlanItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PlanItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PlanItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PlanItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PlanItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PlanItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetScheduledStartDate

`func (o *PlanItem) GetScheduledStartDate() time.Time`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *PlanItem) GetScheduledStartDateOk() (*time.Time, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *PlanItem) SetScheduledStartDate(v time.Time)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *PlanItem) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetDueDate

`func (o *PlanItem) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *PlanItem) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *PlanItem) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *PlanItem) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *PlanItem) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PlanItem) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PlanItem) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PlanItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PlanItem) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PlanItem) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PlanItem) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PlanItem) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPlannedDuration

`func (o *PlanItem) GetPlannedDuration() int32`

GetPlannedDuration returns the PlannedDuration field if non-nil, zero value otherwise.

### GetPlannedDurationOk

`func (o *PlanItem) GetPlannedDurationOk() (*int32, bool)`

GetPlannedDurationOk returns a tuple with the PlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDuration

`func (o *PlanItem) SetPlannedDuration(v int32)`

SetPlannedDuration sets PlannedDuration field to given value.

### HasPlannedDuration

`func (o *PlanItem) HasPlannedDuration() bool`

HasPlannedDuration returns a boolean if a field has been set.

### GetFlagStatus

`func (o *PlanItem) GetFlagStatus() FlagStatus`

GetFlagStatus returns the FlagStatus field if non-nil, zero value otherwise.

### GetFlagStatusOk

`func (o *PlanItem) GetFlagStatusOk() (*FlagStatus, bool)`

GetFlagStatusOk returns a tuple with the FlagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagStatus

`func (o *PlanItem) SetFlagStatus(v FlagStatus)`

SetFlagStatus sets FlagStatus field to given value.

### HasFlagStatus

`func (o *PlanItem) HasFlagStatus() bool`

HasFlagStatus returns a boolean if a field has been set.

### GetFlagComment

`func (o *PlanItem) GetFlagComment() string`

GetFlagComment returns the FlagComment field if non-nil, zero value otherwise.

### GetFlagCommentOk

`func (o *PlanItem) GetFlagCommentOk() (*string, bool)`

GetFlagCommentOk returns a tuple with the FlagComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagComment

`func (o *PlanItem) SetFlagComment(v string)`

SetFlagComment sets FlagComment field to given value.

### HasFlagComment

`func (o *PlanItem) HasFlagComment() bool`

HasFlagComment returns a boolean if a field has been set.

### GetOverdueNotified

`func (o *PlanItem) GetOverdueNotified() bool`

GetOverdueNotified returns the OverdueNotified field if non-nil, zero value otherwise.

### GetOverdueNotifiedOk

`func (o *PlanItem) GetOverdueNotifiedOk() (*bool, bool)`

GetOverdueNotifiedOk returns a tuple with the OverdueNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueNotified

`func (o *PlanItem) SetOverdueNotified(v bool)`

SetOverdueNotified sets OverdueNotified field to given value.

### HasOverdueNotified

`func (o *PlanItem) HasOverdueNotified() bool`

HasOverdueNotified returns a boolean if a field has been set.

### GetFlagged

`func (o *PlanItem) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *PlanItem) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *PlanItem) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *PlanItem) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetStartOrScheduledDate

`func (o *PlanItem) GetStartOrScheduledDate() time.Time`

GetStartOrScheduledDate returns the StartOrScheduledDate field if non-nil, zero value otherwise.

### GetStartOrScheduledDateOk

`func (o *PlanItem) GetStartOrScheduledDateOk() (*time.Time, bool)`

GetStartOrScheduledDateOk returns a tuple with the StartOrScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOrScheduledDate

`func (o *PlanItem) SetStartOrScheduledDate(v time.Time)`

SetStartOrScheduledDate sets StartOrScheduledDate field to given value.

### HasStartOrScheduledDate

`func (o *PlanItem) HasStartOrScheduledDate() bool`

HasStartOrScheduledDate returns a boolean if a field has been set.

### GetEndOrDueDate

`func (o *PlanItem) GetEndOrDueDate() time.Time`

GetEndOrDueDate returns the EndOrDueDate field if non-nil, zero value otherwise.

### GetEndOrDueDateOk

`func (o *PlanItem) GetEndOrDueDateOk() (*time.Time, bool)`

GetEndOrDueDateOk returns a tuple with the EndOrDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOrDueDate

`func (o *PlanItem) SetEndOrDueDate(v time.Time)`

SetEndOrDueDate sets EndOrDueDate field to given value.

### HasEndOrDueDate

`func (o *PlanItem) HasEndOrDueDate() bool`

HasEndOrDueDate returns a boolean if a field has been set.

### GetChildren

`func (o *PlanItem) GetChildren() []PlanItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PlanItem) GetChildrenOk() (*[]PlanItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PlanItem) SetChildren(v []PlanItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *PlanItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetOverdue

`func (o *PlanItem) GetOverdue() bool`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *PlanItem) GetOverdueOk() (*bool, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *PlanItem) SetOverdue(v bool)`

SetOverdue sets Overdue field to given value.

### HasOverdue

`func (o *PlanItem) HasOverdue() bool`

HasOverdue returns a boolean if a field has been set.

### GetDone

`func (o *PlanItem) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *PlanItem) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *PlanItem) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *PlanItem) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetRelease

`func (o *PlanItem) GetRelease() Release`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PlanItem) GetReleaseOk() (*Release, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PlanItem) SetRelease(v Release)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PlanItem) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetReleaseUid

`func (o *PlanItem) GetReleaseUid() int32`

GetReleaseUid returns the ReleaseUid field if non-nil, zero value otherwise.

### GetReleaseUidOk

`func (o *PlanItem) GetReleaseUidOk() (*int32, bool)`

GetReleaseUidOk returns a tuple with the ReleaseUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseUid

`func (o *PlanItem) SetReleaseUid(v int32)`

SetReleaseUid sets ReleaseUid field to given value.

### HasReleaseUid

`func (o *PlanItem) HasReleaseUid() bool`

HasReleaseUid returns a boolean if a field has been set.

### GetUpdatable

`func (o *PlanItem) GetUpdatable() bool`

GetUpdatable returns the Updatable field if non-nil, zero value otherwise.

### GetUpdatableOk

`func (o *PlanItem) GetUpdatableOk() (*bool, bool)`

GetUpdatableOk returns a tuple with the Updatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatable

`func (o *PlanItem) SetUpdatable(v bool)`

SetUpdatable sets Updatable field to given value.

### HasUpdatable

`func (o *PlanItem) HasUpdatable() bool`

HasUpdatable returns a boolean if a field has been set.

### GetDisplayPath

`func (o *PlanItem) GetDisplayPath() string`

GetDisplayPath returns the DisplayPath field if non-nil, zero value otherwise.

### GetDisplayPathOk

`func (o *PlanItem) GetDisplayPathOk() (*string, bool)`

GetDisplayPathOk returns a tuple with the DisplayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPath

`func (o *PlanItem) SetDisplayPath(v string)`

SetDisplayPath sets DisplayPath field to given value.

### HasDisplayPath

`func (o *PlanItem) HasDisplayPath() bool`

HasDisplayPath returns a boolean if a field has been set.

### GetAborted

`func (o *PlanItem) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *PlanItem) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *PlanItem) SetAborted(v bool)`

SetAborted sets Aborted field to given value.

### HasAborted

`func (o *PlanItem) HasAborted() bool`

HasAborted returns a boolean if a field has been set.

### GetActive

`func (o *PlanItem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PlanItem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PlanItem) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PlanItem) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVariableUsages

`func (o *PlanItem) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *PlanItem) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *PlanItem) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *PlanItem) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.

### GetOrCalculateDueDate

`func (o *PlanItem) GetOrCalculateDueDate() time.Time`

GetOrCalculateDueDate returns the OrCalculateDueDate field if non-nil, zero value otherwise.

### GetOrCalculateDueDateOk

`func (o *PlanItem) GetOrCalculateDueDateOk() (*time.Time, bool)`

GetOrCalculateDueDateOk returns a tuple with the OrCalculateDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrCalculateDueDate

`func (o *PlanItem) SetOrCalculateDueDate(v time.Time)`

SetOrCalculateDueDate sets OrCalculateDueDate field to given value.

### HasOrCalculateDueDate

`func (o *PlanItem) HasOrCalculateDueDate() bool`

HasOrCalculateDueDate returns a boolean if a field has been set.

### SetOrCalculateDueDateNil

`func (o *PlanItem) SetOrCalculateDueDateNil(b bool)`

 SetOrCalculateDueDateNil sets the value for OrCalculateDueDate to be an explicit nil

### UnsetOrCalculateDueDate
`func (o *PlanItem) UnsetOrCalculateDueDate()`

UnsetOrCalculateDueDate ensures that no value is present for OrCalculateDueDate, not even an explicit nil
### GetComputedPlannedDuration

`func (o *PlanItem) GetComputedPlannedDuration() map[string]interface{}`

GetComputedPlannedDuration returns the ComputedPlannedDuration field if non-nil, zero value otherwise.

### GetComputedPlannedDurationOk

`func (o *PlanItem) GetComputedPlannedDurationOk() (*map[string]interface{}, bool)`

GetComputedPlannedDurationOk returns a tuple with the ComputedPlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPlannedDuration

`func (o *PlanItem) SetComputedPlannedDuration(v map[string]interface{})`

SetComputedPlannedDuration sets ComputedPlannedDuration field to given value.

### HasComputedPlannedDuration

`func (o *PlanItem) HasComputedPlannedDuration() bool`

HasComputedPlannedDuration returns a boolean if a field has been set.

### GetActualDuration

`func (o *PlanItem) GetActualDuration() map[string]interface{}`

GetActualDuration returns the ActualDuration field if non-nil, zero value otherwise.

### GetActualDurationOk

`func (o *PlanItem) GetActualDurationOk() (*map[string]interface{}, bool)`

GetActualDurationOk returns a tuple with the ActualDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDuration

`func (o *PlanItem) SetActualDuration(v map[string]interface{})`

SetActualDuration sets ActualDuration field to given value.

### HasActualDuration

`func (o *PlanItem) HasActualDuration() bool`

HasActualDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


