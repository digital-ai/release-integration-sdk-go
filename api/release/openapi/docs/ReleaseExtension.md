# ReleaseExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 

## Methods

### NewReleaseExtension

`func NewReleaseExtension() *ReleaseExtension`

NewReleaseExtension instantiates a new ReleaseExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseExtensionWithDefaults

`func NewReleaseExtensionWithDefaults() *ReleaseExtension`

NewReleaseExtensionWithDefaults instantiates a new ReleaseExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseExtension) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseExtension) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseExtension) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReleaseExtension) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ReleaseExtension) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReleaseExtension) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReleaseExtension) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReleaseExtension) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariableUsages

`func (o *ReleaseExtension) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *ReleaseExtension) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *ReleaseExtension) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *ReleaseExtension) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


