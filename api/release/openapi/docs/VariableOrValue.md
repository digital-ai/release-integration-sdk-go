# VariableOrValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variable** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewVariableOrValue

`func NewVariableOrValue() *VariableOrValue`

NewVariableOrValue instantiates a new VariableOrValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableOrValueWithDefaults

`func NewVariableOrValueWithDefaults() *VariableOrValue`

NewVariableOrValueWithDefaults instantiates a new VariableOrValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariable

`func (o *VariableOrValue) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *VariableOrValue) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *VariableOrValue) SetVariable(v string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *VariableOrValue) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### GetValue

`func (o *VariableOrValue) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableOrValue) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableOrValue) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableOrValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VariableOrValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VariableOrValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


