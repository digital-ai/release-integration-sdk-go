# ReservationSearchView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Applications** | Pointer to [**[]BaseApplicationView**](BaseApplicationView.md) |  | [optional] 

## Methods

### NewReservationSearchView

`func NewReservationSearchView() *ReservationSearchView`

NewReservationSearchView instantiates a new ReservationSearchView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationSearchViewWithDefaults

`func NewReservationSearchViewWithDefaults() *ReservationSearchView`

NewReservationSearchViewWithDefaults instantiates a new ReservationSearchView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReservationSearchView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReservationSearchView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReservationSearchView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReservationSearchView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartDate

`func (o *ReservationSearchView) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ReservationSearchView) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ReservationSearchView) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ReservationSearchView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ReservationSearchView) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ReservationSearchView) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ReservationSearchView) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ReservationSearchView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNote

`func (o *ReservationSearchView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ReservationSearchView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ReservationSearchView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ReservationSearchView) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetApplications

`func (o *ReservationSearchView) GetApplications() []BaseApplicationView`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ReservationSearchView) GetApplicationsOk() (*[]BaseApplicationView, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ReservationSearchView) SetApplications(v []BaseApplicationView)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ReservationSearchView) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


