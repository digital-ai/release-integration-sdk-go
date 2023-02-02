# ReleaseGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ReleaseGroupStatus**](ReleaseGroupStatus.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**RiskScore** | Pointer to **int32** |  | [optional] 
**ReleaseIds** | Pointer to **[]string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Updatable** | Pointer to **bool** |  | [optional] 

## Methods

### NewReleaseGroup

`func NewReleaseGroup() *ReleaseGroup`

NewReleaseGroup instantiates a new ReleaseGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseGroupWithDefaults

`func NewReleaseGroupWithDefaults() *ReleaseGroup`

NewReleaseGroupWithDefaults instantiates a new ReleaseGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ReleaseGroup) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReleaseGroup) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReleaseGroup) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ReleaseGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTitle

`func (o *ReleaseGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseGroup) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseGroup) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *ReleaseGroup) GetStatus() ReleaseGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReleaseGroup) GetStatusOk() (*ReleaseGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReleaseGroup) SetStatus(v ReleaseGroupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReleaseGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *ReleaseGroup) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ReleaseGroup) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ReleaseGroup) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ReleaseGroup) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ReleaseGroup) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ReleaseGroup) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ReleaseGroup) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ReleaseGroup) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRiskScore

`func (o *ReleaseGroup) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *ReleaseGroup) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *ReleaseGroup) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *ReleaseGroup) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetReleaseIds

`func (o *ReleaseGroup) GetReleaseIds() []string`

GetReleaseIds returns the ReleaseIds field if non-nil, zero value otherwise.

### GetReleaseIdsOk

`func (o *ReleaseGroup) GetReleaseIdsOk() (*[]string, bool)`

GetReleaseIdsOk returns a tuple with the ReleaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseIds

`func (o *ReleaseGroup) SetReleaseIds(v []string)`

SetReleaseIds sets ReleaseIds field to given value.

### HasReleaseIds

`func (o *ReleaseGroup) HasReleaseIds() bool`

HasReleaseIds returns a boolean if a field has been set.

### GetProgress

`func (o *ReleaseGroup) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ReleaseGroup) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ReleaseGroup) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ReleaseGroup) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetFolderId

`func (o *ReleaseGroup) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ReleaseGroup) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ReleaseGroup) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ReleaseGroup) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetUpdatable

`func (o *ReleaseGroup) GetUpdatable() bool`

GetUpdatable returns the Updatable field if non-nil, zero value otherwise.

### GetUpdatableOk

`func (o *ReleaseGroup) GetUpdatableOk() (*bool, bool)`

GetUpdatableOk returns a tuple with the Updatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatable

`func (o *ReleaseGroup) SetUpdatable(v bool)`

SetUpdatable sets Updatable field to given value.

### HasUpdatable

`func (o *ReleaseGroup) HasUpdatable() bool`

HasUpdatable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


