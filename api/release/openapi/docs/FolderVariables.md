# FolderVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**StringVariableValues** | Pointer to **map[string]string** |  | [optional] 
**PasswordVariableValues** | Pointer to **map[string]string** |  | [optional] 
**VariablesByKeys** | Pointer to [**map[string]Variable**](Variable.md) |  | [optional] 

## Methods

### NewFolderVariables

`func NewFolderVariables() *FolderVariables`

NewFolderVariables instantiates a new FolderVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderVariablesWithDefaults

`func NewFolderVariablesWithDefaults() *FolderVariables`

NewFolderVariablesWithDefaults instantiates a new FolderVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *FolderVariables) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *FolderVariables) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *FolderVariables) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *FolderVariables) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetStringVariableValues

`func (o *FolderVariables) GetStringVariableValues() map[string]string`

GetStringVariableValues returns the StringVariableValues field if non-nil, zero value otherwise.

### GetStringVariableValuesOk

`func (o *FolderVariables) GetStringVariableValuesOk() (*map[string]string, bool)`

GetStringVariableValuesOk returns a tuple with the StringVariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringVariableValues

`func (o *FolderVariables) SetStringVariableValues(v map[string]string)`

SetStringVariableValues sets StringVariableValues field to given value.

### HasStringVariableValues

`func (o *FolderVariables) HasStringVariableValues() bool`

HasStringVariableValues returns a boolean if a field has been set.

### GetPasswordVariableValues

`func (o *FolderVariables) GetPasswordVariableValues() map[string]string`

GetPasswordVariableValues returns the PasswordVariableValues field if non-nil, zero value otherwise.

### GetPasswordVariableValuesOk

`func (o *FolderVariables) GetPasswordVariableValuesOk() (*map[string]string, bool)`

GetPasswordVariableValuesOk returns a tuple with the PasswordVariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordVariableValues

`func (o *FolderVariables) SetPasswordVariableValues(v map[string]string)`

SetPasswordVariableValues sets PasswordVariableValues field to given value.

### HasPasswordVariableValues

`func (o *FolderVariables) HasPasswordVariableValues() bool`

HasPasswordVariableValues returns a boolean if a field has been set.

### GetVariablesByKeys

`func (o *FolderVariables) GetVariablesByKeys() map[string]Variable`

GetVariablesByKeys returns the VariablesByKeys field if non-nil, zero value otherwise.

### GetVariablesByKeysOk

`func (o *FolderVariables) GetVariablesByKeysOk() (*map[string]Variable, bool)`

GetVariablesByKeysOk returns a tuple with the VariablesByKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesByKeys

`func (o *FolderVariables) SetVariablesByKeys(v map[string]Variable)`

SetVariablesByKeys sets VariablesByKeys field to given value.

### HasVariablesByKeys

`func (o *FolderVariables) HasVariablesByKeys() bool`

HasVariablesByKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


