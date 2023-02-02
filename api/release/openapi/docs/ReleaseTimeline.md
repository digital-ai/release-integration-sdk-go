# ReleaseTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ScheduledStartDate** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**PlannedStartDate** | Pointer to **string** |  | [optional] 
**PlannedEndDate** | Pointer to **string** |  | [optional] 
**Phases** | Pointer to [**[]PhaseTimeline**](PhaseTimeline.md) |  | [optional] 
**RiskScore** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**ReleaseStatus**](ReleaseStatus.md) |  | [optional] 

## Methods

### NewReleaseTimeline

`func NewReleaseTimeline() *ReleaseTimeline`

NewReleaseTimeline instantiates a new ReleaseTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseTimelineWithDefaults

`func NewReleaseTimelineWithDefaults() *ReleaseTimeline`

NewReleaseTimelineWithDefaults instantiates a new ReleaseTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseTimeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseTimeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseTimeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReleaseTimeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ReleaseTimeline) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseTimeline) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseTimeline) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseTimeline) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetScheduledStartDate

`func (o *ReleaseTimeline) GetScheduledStartDate() string`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *ReleaseTimeline) GetScheduledStartDateOk() (*string, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *ReleaseTimeline) SetScheduledStartDate(v string)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *ReleaseTimeline) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetDueDate

`func (o *ReleaseTimeline) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ReleaseTimeline) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ReleaseTimeline) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ReleaseTimeline) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *ReleaseTimeline) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ReleaseTimeline) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ReleaseTimeline) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ReleaseTimeline) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ReleaseTimeline) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ReleaseTimeline) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ReleaseTimeline) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ReleaseTimeline) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPlannedStartDate

`func (o *ReleaseTimeline) GetPlannedStartDate() string`

GetPlannedStartDate returns the PlannedStartDate field if non-nil, zero value otherwise.

### GetPlannedStartDateOk

`func (o *ReleaseTimeline) GetPlannedStartDateOk() (*string, bool)`

GetPlannedStartDateOk returns a tuple with the PlannedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartDate

`func (o *ReleaseTimeline) SetPlannedStartDate(v string)`

SetPlannedStartDate sets PlannedStartDate field to given value.

### HasPlannedStartDate

`func (o *ReleaseTimeline) HasPlannedStartDate() bool`

HasPlannedStartDate returns a boolean if a field has been set.

### GetPlannedEndDate

`func (o *ReleaseTimeline) GetPlannedEndDate() string`

GetPlannedEndDate returns the PlannedEndDate field if non-nil, zero value otherwise.

### GetPlannedEndDateOk

`func (o *ReleaseTimeline) GetPlannedEndDateOk() (*string, bool)`

GetPlannedEndDateOk returns a tuple with the PlannedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEndDate

`func (o *ReleaseTimeline) SetPlannedEndDate(v string)`

SetPlannedEndDate sets PlannedEndDate field to given value.

### HasPlannedEndDate

`func (o *ReleaseTimeline) HasPlannedEndDate() bool`

HasPlannedEndDate returns a boolean if a field has been set.

### GetPhases

`func (o *ReleaseTimeline) GetPhases() []PhaseTimeline`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *ReleaseTimeline) GetPhasesOk() (*[]PhaseTimeline, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *ReleaseTimeline) SetPhases(v []PhaseTimeline)`

SetPhases sets Phases field to given value.

### HasPhases

`func (o *ReleaseTimeline) HasPhases() bool`

HasPhases returns a boolean if a field has been set.

### GetRiskScore

`func (o *ReleaseTimeline) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *ReleaseTimeline) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *ReleaseTimeline) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *ReleaseTimeline) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetStatus

`func (o *ReleaseTimeline) GetStatus() ReleaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReleaseTimeline) GetStatusOk() (*ReleaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReleaseTimeline) SetStatus(v ReleaseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReleaseTimeline) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


