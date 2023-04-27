# StartRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseTitle** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ReleaseVariables** | Pointer to **map[string]string** |  | [optional] 
**ReleasePasswordVariables** | Pointer to **map[string]string** |  | [optional] 
**ScheduledStartDate** | Pointer to **string** |  | [optional] 
**AutoStart** | Pointer to **bool** |  | [optional] 
**StartedFromTaskId** | Pointer to **string** |  | [optional] 
**ReleaseOwner** | Pointer to **string** |  | [optional] 

## Methods

### NewStartRelease

`func NewStartRelease() *StartRelease`

NewStartRelease instantiates a new StartRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartReleaseWithDefaults

`func NewStartReleaseWithDefaults() *StartRelease`

NewStartReleaseWithDefaults instantiates a new StartRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseTitle

`func (o *StartRelease) GetReleaseTitle() string`

GetReleaseTitle returns the ReleaseTitle field if non-nil, zero value otherwise.

### GetReleaseTitleOk

`func (o *StartRelease) GetReleaseTitleOk() (*string, bool)`

GetReleaseTitleOk returns a tuple with the ReleaseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTitle

`func (o *StartRelease) SetReleaseTitle(v string)`

SetReleaseTitle sets ReleaseTitle field to given value.

### HasReleaseTitle

`func (o *StartRelease) HasReleaseTitle() bool`

HasReleaseTitle returns a boolean if a field has been set.

### GetFolderId

`func (o *StartRelease) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *StartRelease) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *StartRelease) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *StartRelease) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetVariables

`func (o *StartRelease) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *StartRelease) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *StartRelease) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *StartRelease) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetReleaseVariables

`func (o *StartRelease) GetReleaseVariables() map[string]string`

GetReleaseVariables returns the ReleaseVariables field if non-nil, zero value otherwise.

### GetReleaseVariablesOk

`func (o *StartRelease) GetReleaseVariablesOk() (*map[string]string, bool)`

GetReleaseVariablesOk returns a tuple with the ReleaseVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVariables

`func (o *StartRelease) SetReleaseVariables(v map[string]string)`

SetReleaseVariables sets ReleaseVariables field to given value.

### HasReleaseVariables

`func (o *StartRelease) HasReleaseVariables() bool`

HasReleaseVariables returns a boolean if a field has been set.

### GetReleasePasswordVariables

`func (o *StartRelease) GetReleasePasswordVariables() map[string]string`

GetReleasePasswordVariables returns the ReleasePasswordVariables field if non-nil, zero value otherwise.

### GetReleasePasswordVariablesOk

`func (o *StartRelease) GetReleasePasswordVariablesOk() (*map[string]string, bool)`

GetReleasePasswordVariablesOk returns a tuple with the ReleasePasswordVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePasswordVariables

`func (o *StartRelease) SetReleasePasswordVariables(v map[string]string)`

SetReleasePasswordVariables sets ReleasePasswordVariables field to given value.

### HasReleasePasswordVariables

`func (o *StartRelease) HasReleasePasswordVariables() bool`

HasReleasePasswordVariables returns a boolean if a field has been set.

### GetScheduledStartDate

`func (o *StartRelease) GetScheduledStartDate() string`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *StartRelease) GetScheduledStartDateOk() (*string, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *StartRelease) SetScheduledStartDate(v string)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *StartRelease) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetAutoStart

`func (o *StartRelease) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *StartRelease) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *StartRelease) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.

### HasAutoStart

`func (o *StartRelease) HasAutoStart() bool`

HasAutoStart returns a boolean if a field has been set.

### GetStartedFromTaskId

`func (o *StartRelease) GetStartedFromTaskId() string`

GetStartedFromTaskId returns the StartedFromTaskId field if non-nil, zero value otherwise.

### GetStartedFromTaskIdOk

`func (o *StartRelease) GetStartedFromTaskIdOk() (*string, bool)`

GetStartedFromTaskIdOk returns a tuple with the StartedFromTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedFromTaskId

`func (o *StartRelease) SetStartedFromTaskId(v string)`

SetStartedFromTaskId sets StartedFromTaskId field to given value.

### HasStartedFromTaskId

`func (o *StartRelease) HasStartedFromTaskId() bool`

HasStartedFromTaskId returns a boolean if a field has been set.

### GetReleaseOwner

`func (o *StartRelease) GetReleaseOwner() string`

GetReleaseOwner returns the ReleaseOwner field if non-nil, zero value otherwise.

### GetReleaseOwnerOk

`func (o *StartRelease) GetReleaseOwnerOk() (*string, bool)`

GetReleaseOwnerOk returns a tuple with the ReleaseOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseOwner

`func (o *StartRelease) SetReleaseOwner(v string)`

SetReleaseOwner sets ReleaseOwner field to given value.

### HasReleaseOwner

`func (o *StartRelease) HasReleaseOwner() bool`

HasReleaseOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


