# Folder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**[]Folder**](Folder.md) |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**AllVariables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**FolderVariables** | Pointer to [**FolderVariables**](FolderVariables.md) |  | [optional] 

## Methods

### NewFolder

`func NewFolder() *Folder`

NewFolder instantiates a new Folder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderWithDefaults

`func NewFolderWithDefaults() *Folder`

NewFolderWithDefaults instantiates a new Folder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Folder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Folder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Folder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Folder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Folder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Folder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Folder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Folder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *Folder) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Folder) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Folder) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Folder) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetChildren

`func (o *Folder) GetChildren() []Folder`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Folder) GetChildrenOk() (*[]Folder, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Folder) SetChildren(v []Folder)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Folder) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetMetadata

`func (o *Folder) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Folder) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Folder) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Folder) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAllVariables

`func (o *Folder) GetAllVariables() []Variable`

GetAllVariables returns the AllVariables field if non-nil, zero value otherwise.

### GetAllVariablesOk

`func (o *Folder) GetAllVariablesOk() (*[]Variable, bool)`

GetAllVariablesOk returns a tuple with the AllVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllVariables

`func (o *Folder) SetAllVariables(v []Variable)`

SetAllVariables sets AllVariables field to given value.

### HasAllVariables

`func (o *Folder) HasAllVariables() bool`

HasAllVariables returns a boolean if a field has been set.

### GetFolderVariables

`func (o *Folder) GetFolderVariables() FolderVariables`

GetFolderVariables returns the FolderVariables field if non-nil, zero value otherwise.

### GetFolderVariablesOk

`func (o *Folder) GetFolderVariablesOk() (*FolderVariables, bool)`

GetFolderVariablesOk returns a tuple with the FolderVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderVariables

`func (o *Folder) SetFolderVariables(v FolderVariables)`

SetFolderVariables sets FolderVariables field to given value.

### HasFolderVariables

`func (o *Folder) HasFolderVariables() bool`

HasFolderVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


