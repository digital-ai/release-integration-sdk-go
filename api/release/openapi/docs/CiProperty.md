# CiProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wrapped** | Pointer to [**CiProperty**](CiProperty.md) |  | [optional] 
**LastProperty** | Pointer to [**Property**](Property.md) |  | [optional] 
**Parent** | Pointer to **map[string]interface{}** |  | [optional] 
**Exists** | Pointer to **bool** |  | [optional] 
**PropertyName** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**ParentCi** | Pointer to **map[string]interface{}** |  | [optional] 
**Descriptor** | Pointer to **map[string]interface{}** |  | [optional] 
**Kind** | Pointer to **map[string]interface{}** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **bool** |  | [optional] 
**Indexed** | Pointer to **bool** |  | [optional] 

## Methods

### NewCiProperty

`func NewCiProperty() *CiProperty`

NewCiProperty instantiates a new CiProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiPropertyWithDefaults

`func NewCiPropertyWithDefaults() *CiProperty`

NewCiPropertyWithDefaults instantiates a new CiProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWrapped

`func (o *CiProperty) GetWrapped() CiProperty`

GetWrapped returns the Wrapped field if non-nil, zero value otherwise.

### GetWrappedOk

`func (o *CiProperty) GetWrappedOk() (*CiProperty, bool)`

GetWrappedOk returns a tuple with the Wrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapped

`func (o *CiProperty) SetWrapped(v CiProperty)`

SetWrapped sets Wrapped field to given value.

### HasWrapped

`func (o *CiProperty) HasWrapped() bool`

HasWrapped returns a boolean if a field has been set.

### GetLastProperty

`func (o *CiProperty) GetLastProperty() Property`

GetLastProperty returns the LastProperty field if non-nil, zero value otherwise.

### GetLastPropertyOk

`func (o *CiProperty) GetLastPropertyOk() (*Property, bool)`

GetLastPropertyOk returns a tuple with the LastProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProperty

`func (o *CiProperty) SetLastProperty(v Property)`

SetLastProperty sets LastProperty field to given value.

### HasLastProperty

`func (o *CiProperty) HasLastProperty() bool`

HasLastProperty returns a boolean if a field has been set.

### GetParent

`func (o *CiProperty) GetParent() map[string]interface{}`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CiProperty) GetParentOk() (*map[string]interface{}, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CiProperty) SetParent(v map[string]interface{})`

SetParent sets Parent field to given value.

### HasParent

`func (o *CiProperty) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetExists

`func (o *CiProperty) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *CiProperty) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *CiProperty) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *CiProperty) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetPropertyName

`func (o *CiProperty) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *CiProperty) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *CiProperty) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *CiProperty) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetValue

`func (o *CiProperty) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CiProperty) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CiProperty) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CiProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetParentCi

`func (o *CiProperty) GetParentCi() map[string]interface{}`

GetParentCi returns the ParentCi field if non-nil, zero value otherwise.

### GetParentCiOk

`func (o *CiProperty) GetParentCiOk() (*map[string]interface{}, bool)`

GetParentCiOk returns a tuple with the ParentCi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCi

`func (o *CiProperty) SetParentCi(v map[string]interface{})`

SetParentCi sets ParentCi field to given value.

### HasParentCi

`func (o *CiProperty) HasParentCi() bool`

HasParentCi returns a boolean if a field has been set.

### GetDescriptor

`func (o *CiProperty) GetDescriptor() map[string]interface{}`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *CiProperty) GetDescriptorOk() (*map[string]interface{}, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *CiProperty) SetDescriptor(v map[string]interface{})`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *CiProperty) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetKind

`func (o *CiProperty) GetKind() map[string]interface{}`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CiProperty) GetKindOk() (*map[string]interface{}, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CiProperty) SetKind(v map[string]interface{})`

SetKind sets Kind field to given value.

### HasKind

`func (o *CiProperty) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCategory

`func (o *CiProperty) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CiProperty) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CiProperty) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CiProperty) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPassword

`func (o *CiProperty) GetPassword() bool`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CiProperty) GetPasswordOk() (*bool, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CiProperty) SetPassword(v bool)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CiProperty) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetIndexed

`func (o *CiProperty) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *CiProperty) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *CiProperty) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *CiProperty) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


