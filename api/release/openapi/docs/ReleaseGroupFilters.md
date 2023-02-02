# ReleaseGroupFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Statuses** | Pointer to [**[]ReleaseGroupStatus**](ReleaseGroupStatus.md) |  | [optional] 

## Methods

### NewReleaseGroupFilters

`func NewReleaseGroupFilters() *ReleaseGroupFilters`

NewReleaseGroupFilters instantiates a new ReleaseGroupFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseGroupFiltersWithDefaults

`func NewReleaseGroupFiltersWithDefaults() *ReleaseGroupFilters`

NewReleaseGroupFiltersWithDefaults instantiates a new ReleaseGroupFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ReleaseGroupFilters) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseGroupFilters) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseGroupFilters) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseGroupFilters) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFolderId

`func (o *ReleaseGroupFilters) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ReleaseGroupFilters) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ReleaseGroupFilters) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ReleaseGroupFilters) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetStatuses

`func (o *ReleaseGroupFilters) GetStatuses() []ReleaseGroupStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ReleaseGroupFilters) GetStatusesOk() (*[]ReleaseGroupStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ReleaseGroupFilters) SetStatuses(v []ReleaseGroupStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ReleaseGroupFilters) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


