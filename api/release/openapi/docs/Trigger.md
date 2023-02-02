# Trigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** |  | [optional] 
**AbortScript** | Pointer to **string** |  | [optional] 
**CiUid** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TriggerState** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**AllowParallelExecution** | Pointer to **bool** |  | [optional] 
**LastRunDate** | Pointer to **string** |  | [optional] 
**LastRunStatus** | Pointer to [**TriggerExecutionStatus**](TriggerExecutionStatus.md) |  | [optional] 
**InternalProperties** | Pointer to **[]string** |  | [optional] 
**ContainerId** | Pointer to **string** |  | [optional] 

## Methods

### NewTrigger

`func NewTrigger() *Trigger`

NewTrigger instantiates a new Trigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWithDefaults

`func NewTriggerWithDefaults() *Trigger`

NewTriggerWithDefaults instantiates a new Trigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *Trigger) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *Trigger) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *Trigger) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *Trigger) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetAbortScript

`func (o *Trigger) GetAbortScript() string`

GetAbortScript returns the AbortScript field if non-nil, zero value otherwise.

### GetAbortScriptOk

`func (o *Trigger) GetAbortScriptOk() (*string, bool)`

GetAbortScriptOk returns a tuple with the AbortScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortScript

`func (o *Trigger) SetAbortScript(v string)`

SetAbortScript sets AbortScript field to given value.

### HasAbortScript

`func (o *Trigger) HasAbortScript() bool`

HasAbortScript returns a boolean if a field has been set.

### GetCiUid

`func (o *Trigger) GetCiUid() int32`

GetCiUid returns the CiUid field if non-nil, zero value otherwise.

### GetCiUidOk

`func (o *Trigger) GetCiUidOk() (*int32, bool)`

GetCiUidOk returns a tuple with the CiUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiUid

`func (o *Trigger) SetCiUid(v int32)`

SetCiUid sets CiUid field to given value.

### HasCiUid

`func (o *Trigger) HasCiUid() bool`

HasCiUid returns a boolean if a field has been set.

### GetTitle

`func (o *Trigger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Trigger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Trigger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Trigger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Trigger) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Trigger) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Trigger) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Trigger) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Trigger) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Trigger) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Trigger) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Trigger) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTriggerState

`func (o *Trigger) GetTriggerState() string`

GetTriggerState returns the TriggerState field if non-nil, zero value otherwise.

### GetTriggerStateOk

`func (o *Trigger) GetTriggerStateOk() (*string, bool)`

GetTriggerStateOk returns a tuple with the TriggerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerState

`func (o *Trigger) SetTriggerState(v string)`

SetTriggerState sets TriggerState field to given value.

### HasTriggerState

`func (o *Trigger) HasTriggerState() bool`

HasTriggerState returns a boolean if a field has been set.

### GetFolderId

`func (o *Trigger) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Trigger) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Trigger) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Trigger) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetAllowParallelExecution

`func (o *Trigger) GetAllowParallelExecution() bool`

GetAllowParallelExecution returns the AllowParallelExecution field if non-nil, zero value otherwise.

### GetAllowParallelExecutionOk

`func (o *Trigger) GetAllowParallelExecutionOk() (*bool, bool)`

GetAllowParallelExecutionOk returns a tuple with the AllowParallelExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowParallelExecution

`func (o *Trigger) SetAllowParallelExecution(v bool)`

SetAllowParallelExecution sets AllowParallelExecution field to given value.

### HasAllowParallelExecution

`func (o *Trigger) HasAllowParallelExecution() bool`

HasAllowParallelExecution returns a boolean if a field has been set.

### GetLastRunDate

`func (o *Trigger) GetLastRunDate() string`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *Trigger) GetLastRunDateOk() (*string, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *Trigger) SetLastRunDate(v string)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *Trigger) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### GetLastRunStatus

`func (o *Trigger) GetLastRunStatus() TriggerExecutionStatus`

GetLastRunStatus returns the LastRunStatus field if non-nil, zero value otherwise.

### GetLastRunStatusOk

`func (o *Trigger) GetLastRunStatusOk() (*TriggerExecutionStatus, bool)`

GetLastRunStatusOk returns a tuple with the LastRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStatus

`func (o *Trigger) SetLastRunStatus(v TriggerExecutionStatus)`

SetLastRunStatus sets LastRunStatus field to given value.

### HasLastRunStatus

`func (o *Trigger) HasLastRunStatus() bool`

HasLastRunStatus returns a boolean if a field has been set.

### GetInternalProperties

`func (o *Trigger) GetInternalProperties() []string`

GetInternalProperties returns the InternalProperties field if non-nil, zero value otherwise.

### GetInternalPropertiesOk

`func (o *Trigger) GetInternalPropertiesOk() (*[]string, bool)`

GetInternalPropertiesOk returns a tuple with the InternalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalProperties

`func (o *Trigger) SetInternalProperties(v []string)`

SetInternalProperties sets InternalProperties field to given value.

### HasInternalProperties

`func (o *Trigger) HasInternalProperties() bool`

HasInternalProperties returns a boolean if a field has been set.

### GetContainerId

`func (o *Trigger) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *Trigger) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *Trigger) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *Trigger) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


