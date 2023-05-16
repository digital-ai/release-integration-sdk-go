# CreateRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseTitle** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ReleaseVariables** | Pointer to **map[string]string** |  | [optional] 
**ReleasePasswordVariables** | Pointer to **map[string]string** |  | [optional] 
**ScheduledStartDate** | Pointer to **time.Time** |  | [optional] 
**AutoStart** | Pointer to **bool** |  | [optional] 
**StartedFromTaskId** | Pointer to **string** |  | [optional] 
**ReleaseOwner** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRelease

`func NewCreateRelease() *CreateRelease`

NewCreateRelease instantiates a new CreateRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReleaseWithDefaults

`func NewCreateReleaseWithDefaults() *CreateRelease`

NewCreateReleaseWithDefaults instantiates a new CreateRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseTitle

`func (o *CreateRelease) GetReleaseTitle() string`

GetReleaseTitle returns the ReleaseTitle field if non-nil, zero value otherwise.

### GetReleaseTitleOk

`func (o *CreateRelease) GetReleaseTitleOk() (*string, bool)`

GetReleaseTitleOk returns a tuple with the ReleaseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTitle

`func (o *CreateRelease) SetReleaseTitle(v string)`

SetReleaseTitle sets ReleaseTitle field to given value.

### HasReleaseTitle

`func (o *CreateRelease) HasReleaseTitle() bool`

HasReleaseTitle returns a boolean if a field has been set.

### GetFolderId

`func (o *CreateRelease) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CreateRelease) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CreateRelease) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CreateRelease) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetVariables

`func (o *CreateRelease) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *CreateRelease) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *CreateRelease) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *CreateRelease) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetReleaseVariables

`func (o *CreateRelease) GetReleaseVariables() map[string]string`

GetReleaseVariables returns the ReleaseVariables field if non-nil, zero value otherwise.

### GetReleaseVariablesOk

`func (o *CreateRelease) GetReleaseVariablesOk() (*map[string]string, bool)`

GetReleaseVariablesOk returns a tuple with the ReleaseVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVariables

`func (o *CreateRelease) SetReleaseVariables(v map[string]string)`

SetReleaseVariables sets ReleaseVariables field to given value.

### HasReleaseVariables

`func (o *CreateRelease) HasReleaseVariables() bool`

HasReleaseVariables returns a boolean if a field has been set.

### GetReleasePasswordVariables

`func (o *CreateRelease) GetReleasePasswordVariables() map[string]string`

GetReleasePasswordVariables returns the ReleasePasswordVariables field if non-nil, zero value otherwise.

### GetReleasePasswordVariablesOk

`func (o *CreateRelease) GetReleasePasswordVariablesOk() (*map[string]string, bool)`

GetReleasePasswordVariablesOk returns a tuple with the ReleasePasswordVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePasswordVariables

`func (o *CreateRelease) SetReleasePasswordVariables(v map[string]string)`

SetReleasePasswordVariables sets ReleasePasswordVariables field to given value.

### HasReleasePasswordVariables

`func (o *CreateRelease) HasReleasePasswordVariables() bool`

HasReleasePasswordVariables returns a boolean if a field has been set.

### GetScheduledStartDate

`func (o *CreateRelease) GetScheduledStartDate() time.Time`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *CreateRelease) GetScheduledStartDateOk() (*time.Time, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *CreateRelease) SetScheduledStartDate(v time.Time)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *CreateRelease) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetAutoStart

`func (o *CreateRelease) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *CreateRelease) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *CreateRelease) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.

### HasAutoStart

`func (o *CreateRelease) HasAutoStart() bool`

HasAutoStart returns a boolean if a field has been set.

### GetStartedFromTaskId

`func (o *CreateRelease) GetStartedFromTaskId() string`

GetStartedFromTaskId returns the StartedFromTaskId field if non-nil, zero value otherwise.

### GetStartedFromTaskIdOk

`func (o *CreateRelease) GetStartedFromTaskIdOk() (*string, bool)`

GetStartedFromTaskIdOk returns a tuple with the StartedFromTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedFromTaskId

`func (o *CreateRelease) SetStartedFromTaskId(v string)`

SetStartedFromTaskId sets StartedFromTaskId field to given value.

### HasStartedFromTaskId

`func (o *CreateRelease) HasStartedFromTaskId() bool`

HasStartedFromTaskId returns a boolean if a field has been set.

### GetReleaseOwner

`func (o *CreateRelease) GetReleaseOwner() string`

GetReleaseOwner returns the ReleaseOwner field if non-nil, zero value otherwise.

### GetReleaseOwnerOk

`func (o *CreateRelease) GetReleaseOwnerOk() (*string, bool)`

GetReleaseOwnerOk returns a tuple with the ReleaseOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseOwner

`func (o *CreateRelease) SetReleaseOwner(v string)`

SetReleaseOwner sets ReleaseOwner field to given value.

### HasReleaseOwner

`func (o *CreateRelease) HasReleaseOwner() bool`

HasReleaseOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


