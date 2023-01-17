# Variable1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RequiresValue** | Pointer to **bool** |  | [optional] 
**ShowOnReleaseStart** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Multiline** | Pointer to **bool** |  | [optional] 
**Inherited** | Pointer to **bool** |  | [optional] 
**PreventInterpolation** | Pointer to **bool** |  | [optional] 
**ExternalVariableValue** | Pointer to [**ExternalVariableValue**](Documents/xebialabs/go-remote-runner-wrapper/api/release/docs/ExternalVariableValue.md) |  | [optional] 
**ValueProvider** | Pointer to [**ValueProviderConfiguration**](Documents/xebialabs/go-remote-runner-wrapper/api/release/docs/ValueProviderConfiguration.md) |  | [optional] 

## Methods

### NewVariable1

`func NewVariable1() *Variable1`

NewVariable1 instantiates a new Variable1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariable1WithDefaults

`func NewVariable1WithDefaults() *Variable1`

NewVariable1WithDefaults instantiates a new Variable1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Variable1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Variable1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Variable1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Variable1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *Variable1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Variable1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Variable1) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Variable1) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *Variable1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Variable1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Variable1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Variable1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequiresValue

`func (o *Variable1) GetRequiresValue() bool`

GetRequiresValue returns the RequiresValue field if non-nil, zero value otherwise.

### GetRequiresValueOk

`func (o *Variable1) GetRequiresValueOk() (*bool, bool)`

GetRequiresValueOk returns a tuple with the RequiresValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresValue

`func (o *Variable1) SetRequiresValue(v bool)`

SetRequiresValue sets RequiresValue field to given value.

### HasRequiresValue

`func (o *Variable1) HasRequiresValue() bool`

HasRequiresValue returns a boolean if a field has been set.

### GetShowOnReleaseStart

`func (o *Variable1) GetShowOnReleaseStart() bool`

GetShowOnReleaseStart returns the ShowOnReleaseStart field if non-nil, zero value otherwise.

### GetShowOnReleaseStartOk

`func (o *Variable1) GetShowOnReleaseStartOk() (*bool, bool)`

GetShowOnReleaseStartOk returns a tuple with the ShowOnReleaseStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnReleaseStart

`func (o *Variable1) SetShowOnReleaseStart(v bool)`

SetShowOnReleaseStart sets ShowOnReleaseStart field to given value.

### HasShowOnReleaseStart

`func (o *Variable1) HasShowOnReleaseStart() bool`

HasShowOnReleaseStart returns a boolean if a field has been set.

### GetValue

`func (o *Variable1) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Variable1) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Variable1) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *Variable1) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *Variable1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Variable1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Variable1) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Variable1) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *Variable1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Variable1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Variable1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Variable1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiline

`func (o *Variable1) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *Variable1) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *Variable1) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.

### HasMultiline

`func (o *Variable1) HasMultiline() bool`

HasMultiline returns a boolean if a field has been set.

### GetInherited

`func (o *Variable1) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *Variable1) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *Variable1) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *Variable1) HasInherited() bool`

HasInherited returns a boolean if a field has been set.

### GetPreventInterpolation

`func (o *Variable1) GetPreventInterpolation() bool`

GetPreventInterpolation returns the PreventInterpolation field if non-nil, zero value otherwise.

### GetPreventInterpolationOk

`func (o *Variable1) GetPreventInterpolationOk() (*bool, bool)`

GetPreventInterpolationOk returns a tuple with the PreventInterpolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventInterpolation

`func (o *Variable1) SetPreventInterpolation(v bool)`

SetPreventInterpolation sets PreventInterpolation field to given value.

### HasPreventInterpolation

`func (o *Variable1) HasPreventInterpolation() bool`

HasPreventInterpolation returns a boolean if a field has been set.

### GetExternalVariableValue

`func (o *Variable1) GetExternalVariableValue() ExternalVariableValue`

GetExternalVariableValue returns the ExternalVariableValue field if non-nil, zero value otherwise.

### GetExternalVariableValueOk

`func (o *Variable1) GetExternalVariableValueOk() (*ExternalVariableValue, bool)`

GetExternalVariableValueOk returns a tuple with the ExternalVariableValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVariableValue

`func (o *Variable1) SetExternalVariableValue(v ExternalVariableValue)`

SetExternalVariableValue sets ExternalVariableValue field to given value.

### HasExternalVariableValue

`func (o *Variable1) HasExternalVariableValue() bool`

HasExternalVariableValue returns a boolean if a field has been set.

### GetValueProvider

`func (o *Variable1) GetValueProvider() ValueProviderConfiguration`

GetValueProvider returns the ValueProvider field if non-nil, zero value otherwise.

### GetValueProviderOk

`func (o *Variable1) GetValueProviderOk() (*ValueProviderConfiguration, bool)`

GetValueProviderOk returns a tuple with the ValueProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueProvider

`func (o *Variable1) SetValueProvider(v ValueProviderConfiguration)`

SetValueProvider sets ValueProvider field to given value.

### HasValueProvider

`func (o *Variable1) HasValueProvider() bool`

HasValueProvider returns a boolean if a field has been set.


[[Back to Model list]](Documents/xebialabs/go-remote-runner-wrapper/api/release/README.mdbialabs/go-remote-runner-wrapper/api/release/README.md#documentation-for-models) [[Back to API list]](Documents/xebialabs/go-remote-runner-wrapper/api/release/README.mdbialabs/go-remote-runner-wrapper/api/release/README.md#documentation-for-api-endpoints) [[Back to README]](Documents/xebialabs/go-remote-runner-wrapper/api/release/README.mdbialabs/go-remote-runner-wrapper/api/release/README.md)


