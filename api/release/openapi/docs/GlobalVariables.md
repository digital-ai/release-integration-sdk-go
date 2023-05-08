# GlobalVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**VariablesByKeys** | Pointer to [**map[string]Variable**](Variable.md) |  | [optional] 
**StringVariableValues** | Pointer to **map[string]string** |  | [optional] 
**PasswordVariableValues** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGlobalVariables

`func NewGlobalVariables() *GlobalVariables`

NewGlobalVariables instantiates a new GlobalVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalVariablesWithDefaults

`func NewGlobalVariablesWithDefaults() *GlobalVariables`

NewGlobalVariablesWithDefaults instantiates a new GlobalVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalVariables) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalVariables) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalVariables) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalVariables) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GlobalVariables) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalVariables) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalVariables) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GlobalVariables) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariables

`func (o *GlobalVariables) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *GlobalVariables) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *GlobalVariables) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *GlobalVariables) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVariablesByKeys

`func (o *GlobalVariables) GetVariablesByKeys() map[string]Variable`

GetVariablesByKeys returns the VariablesByKeys field if non-nil, zero value otherwise.

### GetVariablesByKeysOk

`func (o *GlobalVariables) GetVariablesByKeysOk() (*map[string]Variable, bool)`

GetVariablesByKeysOk returns a tuple with the VariablesByKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesByKeys

`func (o *GlobalVariables) SetVariablesByKeys(v map[string]Variable)`

SetVariablesByKeys sets VariablesByKeys field to given value.

### HasVariablesByKeys

`func (o *GlobalVariables) HasVariablesByKeys() bool`

HasVariablesByKeys returns a boolean if a field has been set.

### GetStringVariableValues

`func (o *GlobalVariables) GetStringVariableValues() map[string]string`

GetStringVariableValues returns the StringVariableValues field if non-nil, zero value otherwise.

### GetStringVariableValuesOk

`func (o *GlobalVariables) GetStringVariableValuesOk() (*map[string]string, bool)`

GetStringVariableValuesOk returns a tuple with the StringVariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringVariableValues

`func (o *GlobalVariables) SetStringVariableValues(v map[string]string)`

SetStringVariableValues sets StringVariableValues field to given value.

### HasStringVariableValues

`func (o *GlobalVariables) HasStringVariableValues() bool`

HasStringVariableValues returns a boolean if a field has been set.

### GetPasswordVariableValues

`func (o *GlobalVariables) GetPasswordVariableValues() map[string]string`

GetPasswordVariableValues returns the PasswordVariableValues field if non-nil, zero value otherwise.

### GetPasswordVariableValuesOk

`func (o *GlobalVariables) GetPasswordVariableValuesOk() (*map[string]string, bool)`

GetPasswordVariableValuesOk returns a tuple with the PasswordVariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordVariableValues

`func (o *GlobalVariables) SetPasswordVariableValues(v map[string]string)`

SetPasswordVariableValues sets PasswordVariableValues field to given value.

### HasPasswordVariableValues

`func (o *GlobalVariables) HasPasswordVariableValues() bool`

HasPasswordVariableValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


