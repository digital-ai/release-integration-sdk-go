# ReleaseTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**AbortScript** | Pointer to **string** |  | [optional] 
**CiUid** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TriggerState** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**AllowParallelExecution** | Pointer to **bool** |  | [optional] 
**LastRunDate** | Pointer to **time.Time** |  | [optional] 
**LastRunStatus** | Pointer to [**TriggerExecutionStatus**](TriggerExecutionStatus.md) |  | [optional] 
**PollType** | Pointer to [**PollType**](PollType.md) |  | [optional] 
**Periodicity** | Pointer to **string** |  | [optional] 
**InitialFire** | Pointer to **bool** |  | [optional] 
**ReleaseTitle** | Pointer to **string** |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ReleaseFolder** | Pointer to **string** |  | [optional] 
**InternalProperties** | Pointer to **[]string** |  | [optional] 
**TemplateVariables** | Pointer to **map[string]string** |  | [optional] 
**TemplatePasswordVariables** | Pointer to **map[string]string** |  | [optional] 
**TriggerStateFromResults** | Pointer to **string** |  | [optional] 
**ScriptVariableNames** | Pointer to **[]string** |  | [optional] 
**ScriptVariablesFromResults** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**StringScriptVariableValues** | Pointer to **map[string]string** |  | [optional] 
**ScriptVariableValues** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**VariablesByKeys** | Pointer to [**map[string]Variable**](Variable.md) |  | [optional] 
**ContainerId** | Pointer to **string** |  | [optional] 

## Methods

### NewReleaseTrigger

`func NewReleaseTrigger() *ReleaseTrigger`

NewReleaseTrigger instantiates a new ReleaseTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseTriggerWithDefaults

`func NewReleaseTriggerWithDefaults() *ReleaseTrigger`

NewReleaseTriggerWithDefaults instantiates a new ReleaseTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseTrigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReleaseTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ReleaseTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReleaseTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReleaseTrigger) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReleaseTrigger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScript

`func (o *ReleaseTrigger) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *ReleaseTrigger) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *ReleaseTrigger) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *ReleaseTrigger) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetAbortScript

`func (o *ReleaseTrigger) GetAbortScript() string`

GetAbortScript returns the AbortScript field if non-nil, zero value otherwise.

### GetAbortScriptOk

`func (o *ReleaseTrigger) GetAbortScriptOk() (*string, bool)`

GetAbortScriptOk returns a tuple with the AbortScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortScript

`func (o *ReleaseTrigger) SetAbortScript(v string)`

SetAbortScript sets AbortScript field to given value.

### HasAbortScript

`func (o *ReleaseTrigger) HasAbortScript() bool`

HasAbortScript returns a boolean if a field has been set.

### GetCiUid

`func (o *ReleaseTrigger) GetCiUid() int32`

GetCiUid returns the CiUid field if non-nil, zero value otherwise.

### GetCiUidOk

`func (o *ReleaseTrigger) GetCiUidOk() (*int32, bool)`

GetCiUidOk returns a tuple with the CiUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiUid

`func (o *ReleaseTrigger) SetCiUid(v int32)`

SetCiUid sets CiUid field to given value.

### HasCiUid

`func (o *ReleaseTrigger) HasCiUid() bool`

HasCiUid returns a boolean if a field has been set.

### GetTitle

`func (o *ReleaseTrigger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseTrigger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseTrigger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseTrigger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ReleaseTrigger) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReleaseTrigger) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReleaseTrigger) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReleaseTrigger) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ReleaseTrigger) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReleaseTrigger) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReleaseTrigger) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ReleaseTrigger) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTriggerState

`func (o *ReleaseTrigger) GetTriggerState() string`

GetTriggerState returns the TriggerState field if non-nil, zero value otherwise.

### GetTriggerStateOk

`func (o *ReleaseTrigger) GetTriggerStateOk() (*string, bool)`

GetTriggerStateOk returns a tuple with the TriggerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerState

`func (o *ReleaseTrigger) SetTriggerState(v string)`

SetTriggerState sets TriggerState field to given value.

### HasTriggerState

`func (o *ReleaseTrigger) HasTriggerState() bool`

HasTriggerState returns a boolean if a field has been set.

### GetFolderId

`func (o *ReleaseTrigger) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ReleaseTrigger) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ReleaseTrigger) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ReleaseTrigger) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetAllowParallelExecution

`func (o *ReleaseTrigger) GetAllowParallelExecution() bool`

GetAllowParallelExecution returns the AllowParallelExecution field if non-nil, zero value otherwise.

### GetAllowParallelExecutionOk

`func (o *ReleaseTrigger) GetAllowParallelExecutionOk() (*bool, bool)`

GetAllowParallelExecutionOk returns a tuple with the AllowParallelExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowParallelExecution

`func (o *ReleaseTrigger) SetAllowParallelExecution(v bool)`

SetAllowParallelExecution sets AllowParallelExecution field to given value.

### HasAllowParallelExecution

`func (o *ReleaseTrigger) HasAllowParallelExecution() bool`

HasAllowParallelExecution returns a boolean if a field has been set.

### GetLastRunDate

`func (o *ReleaseTrigger) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *ReleaseTrigger) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *ReleaseTrigger) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *ReleaseTrigger) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### GetLastRunStatus

`func (o *ReleaseTrigger) GetLastRunStatus() TriggerExecutionStatus`

GetLastRunStatus returns the LastRunStatus field if non-nil, zero value otherwise.

### GetLastRunStatusOk

`func (o *ReleaseTrigger) GetLastRunStatusOk() (*TriggerExecutionStatus, bool)`

GetLastRunStatusOk returns a tuple with the LastRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStatus

`func (o *ReleaseTrigger) SetLastRunStatus(v TriggerExecutionStatus)`

SetLastRunStatus sets LastRunStatus field to given value.

### HasLastRunStatus

`func (o *ReleaseTrigger) HasLastRunStatus() bool`

HasLastRunStatus returns a boolean if a field has been set.

### GetPollType

`func (o *ReleaseTrigger) GetPollType() PollType`

GetPollType returns the PollType field if non-nil, zero value otherwise.

### GetPollTypeOk

`func (o *ReleaseTrigger) GetPollTypeOk() (*PollType, bool)`

GetPollTypeOk returns a tuple with the PollType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollType

`func (o *ReleaseTrigger) SetPollType(v PollType)`

SetPollType sets PollType field to given value.

### HasPollType

`func (o *ReleaseTrigger) HasPollType() bool`

HasPollType returns a boolean if a field has been set.

### GetPeriodicity

`func (o *ReleaseTrigger) GetPeriodicity() string`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *ReleaseTrigger) GetPeriodicityOk() (*string, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *ReleaseTrigger) SetPeriodicity(v string)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *ReleaseTrigger) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### GetInitialFire

`func (o *ReleaseTrigger) GetInitialFire() bool`

GetInitialFire returns the InitialFire field if non-nil, zero value otherwise.

### GetInitialFireOk

`func (o *ReleaseTrigger) GetInitialFireOk() (*bool, bool)`

GetInitialFireOk returns a tuple with the InitialFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFire

`func (o *ReleaseTrigger) SetInitialFire(v bool)`

SetInitialFire sets InitialFire field to given value.

### HasInitialFire

`func (o *ReleaseTrigger) HasInitialFire() bool`

HasInitialFire returns a boolean if a field has been set.

### GetReleaseTitle

`func (o *ReleaseTrigger) GetReleaseTitle() string`

GetReleaseTitle returns the ReleaseTitle field if non-nil, zero value otherwise.

### GetReleaseTitleOk

`func (o *ReleaseTrigger) GetReleaseTitleOk() (*string, bool)`

GetReleaseTitleOk returns a tuple with the ReleaseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTitle

`func (o *ReleaseTrigger) SetReleaseTitle(v string)`

SetReleaseTitle sets ReleaseTitle field to given value.

### HasReleaseTitle

`func (o *ReleaseTrigger) HasReleaseTitle() bool`

HasReleaseTitle returns a boolean if a field has been set.

### GetExecutionId

`func (o *ReleaseTrigger) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ReleaseTrigger) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ReleaseTrigger) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ReleaseTrigger) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetVariables

`func (o *ReleaseTrigger) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ReleaseTrigger) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ReleaseTrigger) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ReleaseTrigger) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetTemplate

`func (o *ReleaseTrigger) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ReleaseTrigger) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ReleaseTrigger) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ReleaseTrigger) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTags

`func (o *ReleaseTrigger) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReleaseTrigger) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReleaseTrigger) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReleaseTrigger) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetReleaseFolder

`func (o *ReleaseTrigger) GetReleaseFolder() string`

GetReleaseFolder returns the ReleaseFolder field if non-nil, zero value otherwise.

### GetReleaseFolderOk

`func (o *ReleaseTrigger) GetReleaseFolderOk() (*string, bool)`

GetReleaseFolderOk returns a tuple with the ReleaseFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseFolder

`func (o *ReleaseTrigger) SetReleaseFolder(v string)`

SetReleaseFolder sets ReleaseFolder field to given value.

### HasReleaseFolder

`func (o *ReleaseTrigger) HasReleaseFolder() bool`

HasReleaseFolder returns a boolean if a field has been set.

### GetInternalProperties

`func (o *ReleaseTrigger) GetInternalProperties() []string`

GetInternalProperties returns the InternalProperties field if non-nil, zero value otherwise.

### GetInternalPropertiesOk

`func (o *ReleaseTrigger) GetInternalPropertiesOk() (*[]string, bool)`

GetInternalPropertiesOk returns a tuple with the InternalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalProperties

`func (o *ReleaseTrigger) SetInternalProperties(v []string)`

SetInternalProperties sets InternalProperties field to given value.

### HasInternalProperties

`func (o *ReleaseTrigger) HasInternalProperties() bool`

HasInternalProperties returns a boolean if a field has been set.

### GetTemplateVariables

`func (o *ReleaseTrigger) GetTemplateVariables() map[string]string`

GetTemplateVariables returns the TemplateVariables field if non-nil, zero value otherwise.

### GetTemplateVariablesOk

`func (o *ReleaseTrigger) GetTemplateVariablesOk() (*map[string]string, bool)`

GetTemplateVariablesOk returns a tuple with the TemplateVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVariables

`func (o *ReleaseTrigger) SetTemplateVariables(v map[string]string)`

SetTemplateVariables sets TemplateVariables field to given value.

### HasTemplateVariables

`func (o *ReleaseTrigger) HasTemplateVariables() bool`

HasTemplateVariables returns a boolean if a field has been set.

### GetTemplatePasswordVariables

`func (o *ReleaseTrigger) GetTemplatePasswordVariables() map[string]string`

GetTemplatePasswordVariables returns the TemplatePasswordVariables field if non-nil, zero value otherwise.

### GetTemplatePasswordVariablesOk

`func (o *ReleaseTrigger) GetTemplatePasswordVariablesOk() (*map[string]string, bool)`

GetTemplatePasswordVariablesOk returns a tuple with the TemplatePasswordVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePasswordVariables

`func (o *ReleaseTrigger) SetTemplatePasswordVariables(v map[string]string)`

SetTemplatePasswordVariables sets TemplatePasswordVariables field to given value.

### HasTemplatePasswordVariables

`func (o *ReleaseTrigger) HasTemplatePasswordVariables() bool`

HasTemplatePasswordVariables returns a boolean if a field has been set.

### GetTriggerStateFromResults

`func (o *ReleaseTrigger) GetTriggerStateFromResults() string`

GetTriggerStateFromResults returns the TriggerStateFromResults field if non-nil, zero value otherwise.

### GetTriggerStateFromResultsOk

`func (o *ReleaseTrigger) GetTriggerStateFromResultsOk() (*string, bool)`

GetTriggerStateFromResultsOk returns a tuple with the TriggerStateFromResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerStateFromResults

`func (o *ReleaseTrigger) SetTriggerStateFromResults(v string)`

SetTriggerStateFromResults sets TriggerStateFromResults field to given value.

### HasTriggerStateFromResults

`func (o *ReleaseTrigger) HasTriggerStateFromResults() bool`

HasTriggerStateFromResults returns a boolean if a field has been set.

### GetScriptVariableNames

`func (o *ReleaseTrigger) GetScriptVariableNames() []string`

GetScriptVariableNames returns the ScriptVariableNames field if non-nil, zero value otherwise.

### GetScriptVariableNamesOk

`func (o *ReleaseTrigger) GetScriptVariableNamesOk() (*[]string, bool)`

GetScriptVariableNamesOk returns a tuple with the ScriptVariableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptVariableNames

`func (o *ReleaseTrigger) SetScriptVariableNames(v []string)`

SetScriptVariableNames sets ScriptVariableNames field to given value.

### HasScriptVariableNames

`func (o *ReleaseTrigger) HasScriptVariableNames() bool`

HasScriptVariableNames returns a boolean if a field has been set.

### GetScriptVariablesFromResults

`func (o *ReleaseTrigger) GetScriptVariablesFromResults() map[string]map[string]interface{}`

GetScriptVariablesFromResults returns the ScriptVariablesFromResults field if non-nil, zero value otherwise.

### GetScriptVariablesFromResultsOk

`func (o *ReleaseTrigger) GetScriptVariablesFromResultsOk() (*map[string]map[string]interface{}, bool)`

GetScriptVariablesFromResultsOk returns a tuple with the ScriptVariablesFromResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptVariablesFromResults

`func (o *ReleaseTrigger) SetScriptVariablesFromResults(v map[string]map[string]interface{})`

SetScriptVariablesFromResults sets ScriptVariablesFromResults field to given value.

### HasScriptVariablesFromResults

`func (o *ReleaseTrigger) HasScriptVariablesFromResults() bool`

HasScriptVariablesFromResults returns a boolean if a field has been set.

### GetStringScriptVariableValues

`func (o *ReleaseTrigger) GetStringScriptVariableValues() map[string]string`

GetStringScriptVariableValues returns the StringScriptVariableValues field if non-nil, zero value otherwise.

### GetStringScriptVariableValuesOk

`func (o *ReleaseTrigger) GetStringScriptVariableValuesOk() (*map[string]string, bool)`

GetStringScriptVariableValuesOk returns a tuple with the StringScriptVariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringScriptVariableValues

`func (o *ReleaseTrigger) SetStringScriptVariableValues(v map[string]string)`

SetStringScriptVariableValues sets StringScriptVariableValues field to given value.

### HasStringScriptVariableValues

`func (o *ReleaseTrigger) HasStringScriptVariableValues() bool`

HasStringScriptVariableValues returns a boolean if a field has been set.

### GetScriptVariableValues

`func (o *ReleaseTrigger) GetScriptVariableValues() map[string]map[string]interface{}`

GetScriptVariableValues returns the ScriptVariableValues field if non-nil, zero value otherwise.

### GetScriptVariableValuesOk

`func (o *ReleaseTrigger) GetScriptVariableValuesOk() (*map[string]map[string]interface{}, bool)`

GetScriptVariableValuesOk returns a tuple with the ScriptVariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptVariableValues

`func (o *ReleaseTrigger) SetScriptVariableValues(v map[string]map[string]interface{})`

SetScriptVariableValues sets ScriptVariableValues field to given value.

### HasScriptVariableValues

`func (o *ReleaseTrigger) HasScriptVariableValues() bool`

HasScriptVariableValues returns a boolean if a field has been set.

### GetVariablesByKeys

`func (o *ReleaseTrigger) GetVariablesByKeys() map[string]Variable`

GetVariablesByKeys returns the VariablesByKeys field if non-nil, zero value otherwise.

### GetVariablesByKeysOk

`func (o *ReleaseTrigger) GetVariablesByKeysOk() (*map[string]Variable, bool)`

GetVariablesByKeysOk returns a tuple with the VariablesByKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesByKeys

`func (o *ReleaseTrigger) SetVariablesByKeys(v map[string]Variable)`

SetVariablesByKeys sets VariablesByKeys field to given value.

### HasVariablesByKeys

`func (o *ReleaseTrigger) HasVariablesByKeys() bool`

HasVariablesByKeys returns a boolean if a field has been set.

### GetContainerId

`func (o *ReleaseTrigger) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ReleaseTrigger) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ReleaseTrigger) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ReleaseTrigger) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


