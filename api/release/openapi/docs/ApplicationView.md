# ApplicationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Environments** | Pointer to [**[]BaseEnvironmentView**](BaseEnvironmentView.md) |  | [optional] 

## Methods

### NewApplicationView

`func NewApplicationView() *ApplicationView`

NewApplicationView instantiates a new ApplicationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationViewWithDefaults

`func NewApplicationViewWithDefaults() *ApplicationView`

NewApplicationViewWithDefaults instantiates a new ApplicationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ApplicationView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApplicationView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApplicationView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApplicationView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEnvironments

`func (o *ApplicationView) GetEnvironments() []BaseEnvironmentView`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ApplicationView) GetEnvironmentsOk() (*[]BaseEnvironmentView, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ApplicationView) SetEnvironments(v []BaseEnvironmentView)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ApplicationView) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


