# ApplicationFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Environments** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApplicationFilters

`func NewApplicationFilters() *ApplicationFilters`

NewApplicationFilters instantiates a new ApplicationFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFiltersWithDefaults

`func NewApplicationFiltersWithDefaults() *ApplicationFilters`

NewApplicationFiltersWithDefaults instantiates a new ApplicationFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ApplicationFilters) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApplicationFilters) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApplicationFilters) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApplicationFilters) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEnvironments

`func (o *ApplicationFilters) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ApplicationFilters) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ApplicationFilters) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ApplicationFilters) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


