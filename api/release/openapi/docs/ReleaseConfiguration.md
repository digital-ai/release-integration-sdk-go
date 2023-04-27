# ReleaseConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**VariableMapping** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewReleaseConfiguration

`func NewReleaseConfiguration() *ReleaseConfiguration`

NewReleaseConfiguration instantiates a new ReleaseConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseConfigurationWithDefaults

`func NewReleaseConfigurationWithDefaults() *ReleaseConfiguration`

NewReleaseConfigurationWithDefaults instantiates a new ReleaseConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *ReleaseConfiguration) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ReleaseConfiguration) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ReleaseConfiguration) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ReleaseConfiguration) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTitle

`func (o *ReleaseConfiguration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseConfiguration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseConfiguration) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseConfiguration) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVariableMapping

`func (o *ReleaseConfiguration) GetVariableMapping() map[string]string`

GetVariableMapping returns the VariableMapping field if non-nil, zero value otherwise.

### GetVariableMappingOk

`func (o *ReleaseConfiguration) GetVariableMappingOk() (*map[string]string, bool)`

GetVariableMappingOk returns a tuple with the VariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMapping

`func (o *ReleaseConfiguration) SetVariableMapping(v map[string]string)`

SetVariableMapping sets VariableMapping field to given value.

### HasVariableMapping

`func (o *ReleaseConfiguration) HasVariableMapping() bool`

HasVariableMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


