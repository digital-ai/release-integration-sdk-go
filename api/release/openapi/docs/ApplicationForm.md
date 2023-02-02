# ApplicationForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**EnvironmentIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApplicationForm

`func NewApplicationForm() *ApplicationForm`

NewApplicationForm instantiates a new ApplicationForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFormWithDefaults

`func NewApplicationFormWithDefaults() *ApplicationForm`

NewApplicationFormWithDefaults instantiates a new ApplicationForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ApplicationForm) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApplicationForm) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApplicationForm) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApplicationForm) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEnvironmentIds

`func (o *ApplicationForm) GetEnvironmentIds() []string`

GetEnvironmentIds returns the EnvironmentIds field if non-nil, zero value otherwise.

### GetEnvironmentIdsOk

`func (o *ApplicationForm) GetEnvironmentIdsOk() (*[]string, bool)`

GetEnvironmentIdsOk returns a tuple with the EnvironmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentIds

`func (o *ApplicationForm) SetEnvironmentIds(v []string)`

SetEnvironmentIds sets EnvironmentIds field to given value.

### HasEnvironmentIds

`func (o *ApplicationForm) HasEnvironmentIds() bool`

HasEnvironmentIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


