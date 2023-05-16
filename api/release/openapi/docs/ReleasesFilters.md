# ReleasesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TaskTags** | Pointer to **[]string** |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 
**From** | Pointer to **time.Time** |  | [optional] 
**To** | Pointer to **time.Time** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Planned** | Pointer to **bool** |  | [optional] 
**InProgress** | Pointer to **bool** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**Failing** | Pointer to **bool** |  | [optional] 
**Failed** | Pointer to **bool** |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**Aborted** | Pointer to **bool** |  | [optional] 
**OnlyMine** | Pointer to **bool** |  | [optional] 
**OnlyFlagged** | Pointer to **bool** |  | [optional] 
**OnlyArchived** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**OrderBy** | Pointer to [**ReleaseOrderMode**](ReleaseOrderMode.md) |  | [optional] 
**OrderDirection** | Pointer to [**ReleaseOrderDirection**](ReleaseOrderDirection.md) |  | [optional] 
**RiskStatusWithThresholds** | Pointer to [**RiskStatusWithThresholds**](RiskStatusWithThresholds.md) |  | [optional] 
**QueryStartDate** | Pointer to **time.Time** |  | [optional] 
**QueryEndDate** | Pointer to **time.Time** |  | [optional] 
**Statuses** | Pointer to [**[]ReleaseStatus**](ReleaseStatus.md) |  | [optional] 

## Methods

### NewReleasesFilters

`func NewReleasesFilters() *ReleasesFilters`

NewReleasesFilters instantiates a new ReleasesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasesFiltersWithDefaults

`func NewReleasesFiltersWithDefaults() *ReleasesFilters`

NewReleasesFiltersWithDefaults instantiates a new ReleasesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ReleasesFilters) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleasesFilters) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleasesFilters) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleasesFilters) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTags

`func (o *ReleasesFilters) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReleasesFilters) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReleasesFilters) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReleasesFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskTags

`func (o *ReleasesFilters) GetTaskTags() []string`

GetTaskTags returns the TaskTags field if non-nil, zero value otherwise.

### GetTaskTagsOk

`func (o *ReleasesFilters) GetTaskTagsOk() (*[]string, bool)`

GetTaskTagsOk returns a tuple with the TaskTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTags

`func (o *ReleasesFilters) SetTaskTags(v []string)`

SetTaskTags sets TaskTags field to given value.

### HasTaskTags

`func (o *ReleasesFilters) HasTaskTags() bool`

HasTaskTags returns a boolean if a field has been set.

### GetTimeFrame

`func (o *ReleasesFilters) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *ReleasesFilters) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *ReleasesFilters) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *ReleasesFilters) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetFrom

`func (o *ReleasesFilters) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReleasesFilters) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReleasesFilters) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ReleasesFilters) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *ReleasesFilters) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ReleasesFilters) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ReleasesFilters) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *ReleasesFilters) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetActive

`func (o *ReleasesFilters) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ReleasesFilters) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ReleasesFilters) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ReleasesFilters) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPlanned

`func (o *ReleasesFilters) GetPlanned() bool`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *ReleasesFilters) GetPlannedOk() (*bool, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *ReleasesFilters) SetPlanned(v bool)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *ReleasesFilters) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.

### GetInProgress

`func (o *ReleasesFilters) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *ReleasesFilters) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *ReleasesFilters) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *ReleasesFilters) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetPaused

`func (o *ReleasesFilters) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ReleasesFilters) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ReleasesFilters) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ReleasesFilters) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetFailing

`func (o *ReleasesFilters) GetFailing() bool`

GetFailing returns the Failing field if non-nil, zero value otherwise.

### GetFailingOk

`func (o *ReleasesFilters) GetFailingOk() (*bool, bool)`

GetFailingOk returns a tuple with the Failing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailing

`func (o *ReleasesFilters) SetFailing(v bool)`

SetFailing sets Failing field to given value.

### HasFailing

`func (o *ReleasesFilters) HasFailing() bool`

HasFailing returns a boolean if a field has been set.

### GetFailed

`func (o *ReleasesFilters) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ReleasesFilters) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ReleasesFilters) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ReleasesFilters) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetInactive

`func (o *ReleasesFilters) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *ReleasesFilters) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *ReleasesFilters) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *ReleasesFilters) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetCompleted

`func (o *ReleasesFilters) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *ReleasesFilters) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *ReleasesFilters) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *ReleasesFilters) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetAborted

`func (o *ReleasesFilters) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *ReleasesFilters) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *ReleasesFilters) SetAborted(v bool)`

SetAborted sets Aborted field to given value.

### HasAborted

`func (o *ReleasesFilters) HasAborted() bool`

HasAborted returns a boolean if a field has been set.

### GetOnlyMine

`func (o *ReleasesFilters) GetOnlyMine() bool`

GetOnlyMine returns the OnlyMine field if non-nil, zero value otherwise.

### GetOnlyMineOk

`func (o *ReleasesFilters) GetOnlyMineOk() (*bool, bool)`

GetOnlyMineOk returns a tuple with the OnlyMine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyMine

`func (o *ReleasesFilters) SetOnlyMine(v bool)`

SetOnlyMine sets OnlyMine field to given value.

### HasOnlyMine

`func (o *ReleasesFilters) HasOnlyMine() bool`

HasOnlyMine returns a boolean if a field has been set.

### GetOnlyFlagged

`func (o *ReleasesFilters) GetOnlyFlagged() bool`

GetOnlyFlagged returns the OnlyFlagged field if non-nil, zero value otherwise.

### GetOnlyFlaggedOk

`func (o *ReleasesFilters) GetOnlyFlaggedOk() (*bool, bool)`

GetOnlyFlaggedOk returns a tuple with the OnlyFlagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyFlagged

`func (o *ReleasesFilters) SetOnlyFlagged(v bool)`

SetOnlyFlagged sets OnlyFlagged field to given value.

### HasOnlyFlagged

`func (o *ReleasesFilters) HasOnlyFlagged() bool`

HasOnlyFlagged returns a boolean if a field has been set.

### GetOnlyArchived

`func (o *ReleasesFilters) GetOnlyArchived() bool`

GetOnlyArchived returns the OnlyArchived field if non-nil, zero value otherwise.

### GetOnlyArchivedOk

`func (o *ReleasesFilters) GetOnlyArchivedOk() (*bool, bool)`

GetOnlyArchivedOk returns a tuple with the OnlyArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyArchived

`func (o *ReleasesFilters) SetOnlyArchived(v bool)`

SetOnlyArchived sets OnlyArchived field to given value.

### HasOnlyArchived

`func (o *ReleasesFilters) HasOnlyArchived() bool`

HasOnlyArchived returns a boolean if a field has been set.

### GetParentId

`func (o *ReleasesFilters) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ReleasesFilters) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ReleasesFilters) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ReleasesFilters) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetOrderBy

`func (o *ReleasesFilters) GetOrderBy() ReleaseOrderMode`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ReleasesFilters) GetOrderByOk() (*ReleaseOrderMode, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ReleasesFilters) SetOrderBy(v ReleaseOrderMode)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ReleasesFilters) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDirection

`func (o *ReleasesFilters) GetOrderDirection() ReleaseOrderDirection`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *ReleasesFilters) GetOrderDirectionOk() (*ReleaseOrderDirection, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *ReleasesFilters) SetOrderDirection(v ReleaseOrderDirection)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *ReleasesFilters) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetRiskStatusWithThresholds

`func (o *ReleasesFilters) GetRiskStatusWithThresholds() RiskStatusWithThresholds`

GetRiskStatusWithThresholds returns the RiskStatusWithThresholds field if non-nil, zero value otherwise.

### GetRiskStatusWithThresholdsOk

`func (o *ReleasesFilters) GetRiskStatusWithThresholdsOk() (*RiskStatusWithThresholds, bool)`

GetRiskStatusWithThresholdsOk returns a tuple with the RiskStatusWithThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskStatusWithThresholds

`func (o *ReleasesFilters) SetRiskStatusWithThresholds(v RiskStatusWithThresholds)`

SetRiskStatusWithThresholds sets RiskStatusWithThresholds field to given value.

### HasRiskStatusWithThresholds

`func (o *ReleasesFilters) HasRiskStatusWithThresholds() bool`

HasRiskStatusWithThresholds returns a boolean if a field has been set.

### GetQueryStartDate

`func (o *ReleasesFilters) GetQueryStartDate() time.Time`

GetQueryStartDate returns the QueryStartDate field if non-nil, zero value otherwise.

### GetQueryStartDateOk

`func (o *ReleasesFilters) GetQueryStartDateOk() (*time.Time, bool)`

GetQueryStartDateOk returns a tuple with the QueryStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStartDate

`func (o *ReleasesFilters) SetQueryStartDate(v time.Time)`

SetQueryStartDate sets QueryStartDate field to given value.

### HasQueryStartDate

`func (o *ReleasesFilters) HasQueryStartDate() bool`

HasQueryStartDate returns a boolean if a field has been set.

### GetQueryEndDate

`func (o *ReleasesFilters) GetQueryEndDate() time.Time`

GetQueryEndDate returns the QueryEndDate field if non-nil, zero value otherwise.

### GetQueryEndDateOk

`func (o *ReleasesFilters) GetQueryEndDateOk() (*time.Time, bool)`

GetQueryEndDateOk returns a tuple with the QueryEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEndDate

`func (o *ReleasesFilters) SetQueryEndDate(v time.Time)`

SetQueryEndDate sets QueryEndDate field to given value.

### HasQueryEndDate

`func (o *ReleasesFilters) HasQueryEndDate() bool`

HasQueryEndDate returns a boolean if a field has been set.

### GetStatuses

`func (o *ReleasesFilters) GetStatuses() []ReleaseStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ReleasesFilters) GetStatusesOk() (*[]ReleaseStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ReleasesFilters) SetStatuses(v []ReleaseStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ReleasesFilters) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


