# ReservationFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentTitle** | Pointer to **string** |  | [optional] 
**Stages** | Pointer to **[]string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Applications** | Pointer to **[]string** |  | [optional] 
**From** | Pointer to **time.Time** |  | [optional] 
**To** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewReservationFilters

`func NewReservationFilters() *ReservationFilters`

NewReservationFilters instantiates a new ReservationFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationFiltersWithDefaults

`func NewReservationFiltersWithDefaults() *ReservationFilters`

NewReservationFiltersWithDefaults instantiates a new ReservationFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentTitle

`func (o *ReservationFilters) GetEnvironmentTitle() string`

GetEnvironmentTitle returns the EnvironmentTitle field if non-nil, zero value otherwise.

### GetEnvironmentTitleOk

`func (o *ReservationFilters) GetEnvironmentTitleOk() (*string, bool)`

GetEnvironmentTitleOk returns a tuple with the EnvironmentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTitle

`func (o *ReservationFilters) SetEnvironmentTitle(v string)`

SetEnvironmentTitle sets EnvironmentTitle field to given value.

### HasEnvironmentTitle

`func (o *ReservationFilters) HasEnvironmentTitle() bool`

HasEnvironmentTitle returns a boolean if a field has been set.

### GetStages

`func (o *ReservationFilters) GetStages() []string`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *ReservationFilters) GetStagesOk() (*[]string, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *ReservationFilters) SetStages(v []string)`

SetStages sets Stages field to given value.

### HasStages

`func (o *ReservationFilters) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetLabels

`func (o *ReservationFilters) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ReservationFilters) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ReservationFilters) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ReservationFilters) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetApplications

`func (o *ReservationFilters) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ReservationFilters) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ReservationFilters) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ReservationFilters) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetFrom

`func (o *ReservationFilters) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReservationFilters) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReservationFilters) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ReservationFilters) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *ReservationFilters) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ReservationFilters) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ReservationFilters) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *ReservationFilters) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


