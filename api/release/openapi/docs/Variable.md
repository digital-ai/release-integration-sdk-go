# Variable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**RequiresValue** | Pointer to **bool** |  | [optional] 
**ShowOnReleaseStart** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ValueProvider** | Pointer to [**ValueProviderConfiguration**](ValueProviderConfiguration.md) |  | [optional] 
**Inherited** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**EmptyValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ValueEmpty** | Pointer to **bool** |  | [optional] 
**UntypedValue** | Pointer to **map[string]interface{}** |  | [optional] 
**Password** | Pointer to **bool** |  | [optional] 
**ValueAsString** | Pointer to **string** |  | [optional] 
**EmptyValueAsString** | Pointer to **string** |  | [optional] 

## Methods

### NewVariable

`func NewVariable() *Variable`

NewVariable instantiates a new Variable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableWithDefaults

`func NewVariableWithDefaults() *Variable`

NewVariableWithDefaults instantiates a new Variable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Variable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Variable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Variable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Variable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Variable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Variable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Variable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Variable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFolderId

`func (o *Variable) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Variable) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Variable) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Variable) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTitle

`func (o *Variable) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Variable) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Variable) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Variable) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetKey

`func (o *Variable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Variable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Variable) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Variable) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRequiresValue

`func (o *Variable) GetRequiresValue() bool`

GetRequiresValue returns the RequiresValue field if non-nil, zero value otherwise.

### GetRequiresValueOk

`func (o *Variable) GetRequiresValueOk() (*bool, bool)`

GetRequiresValueOk returns a tuple with the RequiresValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresValue

`func (o *Variable) SetRequiresValue(v bool)`

SetRequiresValue sets RequiresValue field to given value.

### HasRequiresValue

`func (o *Variable) HasRequiresValue() bool`

HasRequiresValue returns a boolean if a field has been set.

### GetShowOnReleaseStart

`func (o *Variable) GetShowOnReleaseStart() bool`

GetShowOnReleaseStart returns the ShowOnReleaseStart field if non-nil, zero value otherwise.

### GetShowOnReleaseStartOk

`func (o *Variable) GetShowOnReleaseStartOk() (*bool, bool)`

GetShowOnReleaseStartOk returns a tuple with the ShowOnReleaseStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnReleaseStart

`func (o *Variable) SetShowOnReleaseStart(v bool)`

SetShowOnReleaseStart sets ShowOnReleaseStart field to given value.

### HasShowOnReleaseStart

`func (o *Variable) HasShowOnReleaseStart() bool`

HasShowOnReleaseStart returns a boolean if a field has been set.

### GetLabel

`func (o *Variable) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Variable) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Variable) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Variable) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *Variable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Variable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Variable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Variable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueProvider

`func (o *Variable) GetValueProvider() ValueProviderConfiguration`

GetValueProvider returns the ValueProvider field if non-nil, zero value otherwise.

### GetValueProviderOk

`func (o *Variable) GetValueProviderOk() (*ValueProviderConfiguration, bool)`

GetValueProviderOk returns a tuple with the ValueProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueProvider

`func (o *Variable) SetValueProvider(v ValueProviderConfiguration)`

SetValueProvider sets ValueProvider field to given value.

### HasValueProvider

`func (o *Variable) HasValueProvider() bool`

HasValueProvider returns a boolean if a field has been set.

### GetInherited

`func (o *Variable) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *Variable) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *Variable) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *Variable) HasInherited() bool`

HasInherited returns a boolean if a field has been set.

### GetValue

`func (o *Variable) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Variable) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Variable) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *Variable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Variable) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Variable) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetEmptyValue

`func (o *Variable) GetEmptyValue() map[string]interface{}`

GetEmptyValue returns the EmptyValue field if non-nil, zero value otherwise.

### GetEmptyValueOk

`func (o *Variable) GetEmptyValueOk() (*map[string]interface{}, bool)`

GetEmptyValueOk returns a tuple with the EmptyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyValue

`func (o *Variable) SetEmptyValue(v map[string]interface{})`

SetEmptyValue sets EmptyValue field to given value.

### HasEmptyValue

`func (o *Variable) HasEmptyValue() bool`

HasEmptyValue returns a boolean if a field has been set.

### GetValueEmpty

`func (o *Variable) GetValueEmpty() bool`

GetValueEmpty returns the ValueEmpty field if non-nil, zero value otherwise.

### GetValueEmptyOk

`func (o *Variable) GetValueEmptyOk() (*bool, bool)`

GetValueEmptyOk returns a tuple with the ValueEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueEmpty

`func (o *Variable) SetValueEmpty(v bool)`

SetValueEmpty sets ValueEmpty field to given value.

### HasValueEmpty

`func (o *Variable) HasValueEmpty() bool`

HasValueEmpty returns a boolean if a field has been set.

### GetUntypedValue

`func (o *Variable) GetUntypedValue() map[string]interface{}`

GetUntypedValue returns the UntypedValue field if non-nil, zero value otherwise.

### GetUntypedValueOk

`func (o *Variable) GetUntypedValueOk() (*map[string]interface{}, bool)`

GetUntypedValueOk returns a tuple with the UntypedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntypedValue

`func (o *Variable) SetUntypedValue(v map[string]interface{})`

SetUntypedValue sets UntypedValue field to given value.

### HasUntypedValue

`func (o *Variable) HasUntypedValue() bool`

HasUntypedValue returns a boolean if a field has been set.

### GetPassword

`func (o *Variable) GetPassword() bool`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Variable) GetPasswordOk() (*bool, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Variable) SetPassword(v bool)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Variable) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetValueAsString

`func (o *Variable) GetValueAsString() string`

GetValueAsString returns the ValueAsString field if non-nil, zero value otherwise.

### GetValueAsStringOk

`func (o *Variable) GetValueAsStringOk() (*string, bool)`

GetValueAsStringOk returns a tuple with the ValueAsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAsString

`func (o *Variable) SetValueAsString(v string)`

SetValueAsString sets ValueAsString field to given value.

### HasValueAsString

`func (o *Variable) HasValueAsString() bool`

HasValueAsString returns a boolean if a field has been set.

### GetEmptyValueAsString

`func (o *Variable) GetEmptyValueAsString() string`

GetEmptyValueAsString returns the EmptyValueAsString field if non-nil, zero value otherwise.

### GetEmptyValueAsStringOk

`func (o *Variable) GetEmptyValueAsStringOk() (*string, bool)`

GetEmptyValueAsStringOk returns a tuple with the EmptyValueAsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyValueAsString

`func (o *Variable) SetEmptyValueAsString(v string)`

SetEmptyValueAsString sets EmptyValueAsString field to given value.

### HasEmptyValueAsString

`func (o *Variable) HasEmptyValueAsString() bool`

HasEmptyValueAsString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


