# PhaseTimeline

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
**Color** | Pointer to **string** |  | [optional] 
**CurrentTask** | Pointer to **string** |  | [optional] 

## Methods

### NewPhaseTimeline

`func NewPhaseTimeline() *PhaseTimeline`

NewPhaseTimeline instantiates a new PhaseTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhaseTimelineWithDefaults

`func NewPhaseTimelineWithDefaults() *PhaseTimeline`

NewPhaseTimelineWithDefaults instantiates a new PhaseTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhaseTimeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhaseTimeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhaseTimeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhaseTimeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *PhaseTimeline) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PhaseTimeline) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PhaseTimeline) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PhaseTimeline) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetScheduledStartDate

`func (o *PhaseTimeline) GetScheduledStartDate() string`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *PhaseTimeline) GetScheduledStartDateOk() (*string, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *PhaseTimeline) SetScheduledStartDate(v string)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *PhaseTimeline) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetDueDate

`func (o *PhaseTimeline) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *PhaseTimeline) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *PhaseTimeline) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *PhaseTimeline) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *PhaseTimeline) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PhaseTimeline) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PhaseTimeline) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PhaseTimeline) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PhaseTimeline) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PhaseTimeline) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PhaseTimeline) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PhaseTimeline) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPlannedStartDate

`func (o *PhaseTimeline) GetPlannedStartDate() string`

GetPlannedStartDate returns the PlannedStartDate field if non-nil, zero value otherwise.

### GetPlannedStartDateOk

`func (o *PhaseTimeline) GetPlannedStartDateOk() (*string, bool)`

GetPlannedStartDateOk returns a tuple with the PlannedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartDate

`func (o *PhaseTimeline) SetPlannedStartDate(v string)`

SetPlannedStartDate sets PlannedStartDate field to given value.

### HasPlannedStartDate

`func (o *PhaseTimeline) HasPlannedStartDate() bool`

HasPlannedStartDate returns a boolean if a field has been set.

### GetPlannedEndDate

`func (o *PhaseTimeline) GetPlannedEndDate() string`

GetPlannedEndDate returns the PlannedEndDate field if non-nil, zero value otherwise.

### GetPlannedEndDateOk

`func (o *PhaseTimeline) GetPlannedEndDateOk() (*string, bool)`

GetPlannedEndDateOk returns a tuple with the PlannedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEndDate

`func (o *PhaseTimeline) SetPlannedEndDate(v string)`

SetPlannedEndDate sets PlannedEndDate field to given value.

### HasPlannedEndDate

`func (o *PhaseTimeline) HasPlannedEndDate() bool`

HasPlannedEndDate returns a boolean if a field has been set.

### GetColor

`func (o *PhaseTimeline) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PhaseTimeline) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PhaseTimeline) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PhaseTimeline) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCurrentTask

`func (o *PhaseTimeline) GetCurrentTask() string`

GetCurrentTask returns the CurrentTask field if non-nil, zero value otherwise.

### GetCurrentTaskOk

`func (o *PhaseTimeline) GetCurrentTaskOk() (*string, bool)`

GetCurrentTaskOk returns a tuple with the CurrentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTask

`func (o *PhaseTimeline) SetCurrentTask(v string)`

SetCurrentTask sets CurrentTask field to given value.

### HasCurrentTask

`func (o *PhaseTimeline) HasCurrentTask() bool`

HasCurrentTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


