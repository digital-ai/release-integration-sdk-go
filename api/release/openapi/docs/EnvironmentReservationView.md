# EnvironmentReservationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to [**EnvironmentView**](EnvironmentView.md) |  | [optional] 
**Applications** | Pointer to [**[]BaseApplicationView**](BaseApplicationView.md) |  | [optional] 

## Methods

### NewEnvironmentReservationView

`func NewEnvironmentReservationView() *EnvironmentReservationView`

NewEnvironmentReservationView instantiates a new EnvironmentReservationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReservationViewWithDefaults

`func NewEnvironmentReservationViewWithDefaults() *EnvironmentReservationView`

NewEnvironmentReservationViewWithDefaults instantiates a new EnvironmentReservationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentReservationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentReservationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentReservationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentReservationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartDate

`func (o *EnvironmentReservationView) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EnvironmentReservationView) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EnvironmentReservationView) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EnvironmentReservationView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *EnvironmentReservationView) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnvironmentReservationView) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnvironmentReservationView) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EnvironmentReservationView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNote

`func (o *EnvironmentReservationView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *EnvironmentReservationView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *EnvironmentReservationView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *EnvironmentReservationView) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetEnvironment

`func (o *EnvironmentReservationView) GetEnvironment() EnvironmentView`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentReservationView) GetEnvironmentOk() (*EnvironmentView, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentReservationView) SetEnvironment(v EnvironmentView)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentReservationView) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetApplications

`func (o *EnvironmentReservationView) GetApplications() []BaseApplicationView`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *EnvironmentReservationView) GetApplicationsOk() (*[]BaseApplicationView, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *EnvironmentReservationView) SetApplications(v []BaseApplicationView)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *EnvironmentReservationView) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


