# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**INDEXED_PROPERTY_PATTERN** | Pointer to **string** |  | [optional] 
**PropertyName** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Indexed** | Pointer to **bool** |  | [optional] 

## Methods

### NewProperty

`func NewProperty() *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetINDEXED_PROPERTY_PATTERN

`func (o *Property) GetINDEXED_PROPERTY_PATTERN() string`

GetINDEXED_PROPERTY_PATTERN returns the INDEXED_PROPERTY_PATTERN field if non-nil, zero value otherwise.

### GetINDEXED_PROPERTY_PATTERNOk

`func (o *Property) GetINDEXED_PROPERTY_PATTERNOk() (*string, bool)`

GetINDEXED_PROPERTY_PATTERNOk returns a tuple with the INDEXED_PROPERTY_PATTERN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINDEXED_PROPERTY_PATTERN

`func (o *Property) SetINDEXED_PROPERTY_PATTERN(v string)`

SetINDEXED_PROPERTY_PATTERN sets INDEXED_PROPERTY_PATTERN field to given value.

### HasINDEXED_PROPERTY_PATTERN

`func (o *Property) HasINDEXED_PROPERTY_PATTERN() bool`

HasINDEXED_PROPERTY_PATTERN returns a boolean if a field has been set.

### GetPropertyName

`func (o *Property) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *Property) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *Property) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *Property) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetIndex

`func (o *Property) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Property) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Property) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Property) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIndexed

`func (o *Property) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *Property) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *Property) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *Property) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


