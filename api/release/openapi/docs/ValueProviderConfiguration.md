# ValueProviderConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Variable** | Pointer to [**Variable**](Variable.md) |  | [optional] 

## Methods

### NewValueProviderConfiguration

`func NewValueProviderConfiguration() *ValueProviderConfiguration`

NewValueProviderConfiguration instantiates a new ValueProviderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueProviderConfigurationWithDefaults

`func NewValueProviderConfigurationWithDefaults() *ValueProviderConfiguration`

NewValueProviderConfigurationWithDefaults instantiates a new ValueProviderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ValueProviderConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValueProviderConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValueProviderConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ValueProviderConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ValueProviderConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValueProviderConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValueProviderConfiguration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValueProviderConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariable

`func (o *ValueProviderConfiguration) GetVariable() Variable`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *ValueProviderConfiguration) GetVariableOk() (*Variable, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *ValueProviderConfiguration) SetVariable(v Variable)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *ValueProviderConfiguration) HasVariable() bool`

HasVariable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


