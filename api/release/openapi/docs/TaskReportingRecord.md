# TaskReportingRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**FacetScope**](FacetScope.md) |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**ConfigurationUri** | Pointer to **string** |  | [optional] 
**VariableMapping** | Pointer to **map[string]string** |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 
**PropertiesWithVariables** | Pointer to **[]interface{}** |  | [optional] 
**ServerUrl** | Pointer to **string** |  | [optional] 
**ServerUser** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **string** |  | [optional] 
**RetryAttemptNumber** | Pointer to **int32** |  | [optional] 
**CreatedViaApi** | Pointer to **bool** |  | [optional] 

## Methods

### NewTaskReportingRecord

`func NewTaskReportingRecord() *TaskReportingRecord`

NewTaskReportingRecord instantiates a new TaskReportingRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskReportingRecordWithDefaults

`func NewTaskReportingRecordWithDefaults() *TaskReportingRecord`

NewTaskReportingRecordWithDefaults instantiates a new TaskReportingRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *TaskReportingRecord) GetScope() FacetScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TaskReportingRecord) GetScopeOk() (*FacetScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TaskReportingRecord) SetScope(v FacetScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TaskReportingRecord) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTargetId

`func (o *TaskReportingRecord) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TaskReportingRecord) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TaskReportingRecord) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TaskReportingRecord) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetConfigurationUri

`func (o *TaskReportingRecord) GetConfigurationUri() string`

GetConfigurationUri returns the ConfigurationUri field if non-nil, zero value otherwise.

### GetConfigurationUriOk

`func (o *TaskReportingRecord) GetConfigurationUriOk() (*string, bool)`

GetConfigurationUriOk returns a tuple with the ConfigurationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUri

`func (o *TaskReportingRecord) SetConfigurationUri(v string)`

SetConfigurationUri sets ConfigurationUri field to given value.

### HasConfigurationUri

`func (o *TaskReportingRecord) HasConfigurationUri() bool`

HasConfigurationUri returns a boolean if a field has been set.

### GetVariableMapping

`func (o *TaskReportingRecord) GetVariableMapping() map[string]string`

GetVariableMapping returns the VariableMapping field if non-nil, zero value otherwise.

### GetVariableMappingOk

`func (o *TaskReportingRecord) GetVariableMappingOk() (*map[string]string, bool)`

GetVariableMappingOk returns a tuple with the VariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMapping

`func (o *TaskReportingRecord) SetVariableMapping(v map[string]string)`

SetVariableMapping sets VariableMapping field to given value.

### HasVariableMapping

`func (o *TaskReportingRecord) HasVariableMapping() bool`

HasVariableMapping returns a boolean if a field has been set.

### GetVariableUsages

`func (o *TaskReportingRecord) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *TaskReportingRecord) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *TaskReportingRecord) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *TaskReportingRecord) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.

### GetPropertiesWithVariables

`func (o *TaskReportingRecord) GetPropertiesWithVariables() []interface{}`

GetPropertiesWithVariables returns the PropertiesWithVariables field if non-nil, zero value otherwise.

### GetPropertiesWithVariablesOk

`func (o *TaskReportingRecord) GetPropertiesWithVariablesOk() (*[]interface{}, bool)`

GetPropertiesWithVariablesOk returns a tuple with the PropertiesWithVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesWithVariables

`func (o *TaskReportingRecord) SetPropertiesWithVariables(v []interface{})`

SetPropertiesWithVariables sets PropertiesWithVariables field to given value.

### HasPropertiesWithVariables

`func (o *TaskReportingRecord) HasPropertiesWithVariables() bool`

HasPropertiesWithVariables returns a boolean if a field has been set.

### GetServerUrl

`func (o *TaskReportingRecord) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *TaskReportingRecord) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *TaskReportingRecord) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.

### HasServerUrl

`func (o *TaskReportingRecord) HasServerUrl() bool`

HasServerUrl returns a boolean if a field has been set.

### GetServerUser

`func (o *TaskReportingRecord) GetServerUser() string`

GetServerUser returns the ServerUser field if non-nil, zero value otherwise.

### GetServerUserOk

`func (o *TaskReportingRecord) GetServerUserOk() (*string, bool)`

GetServerUserOk returns a tuple with the ServerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUser

`func (o *TaskReportingRecord) SetServerUser(v string)`

SetServerUser sets ServerUser field to given value.

### HasServerUser

`func (o *TaskReportingRecord) HasServerUser() bool`

HasServerUser returns a boolean if a field has been set.

### GetCreationDate

`func (o *TaskReportingRecord) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *TaskReportingRecord) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *TaskReportingRecord) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *TaskReportingRecord) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetRetryAttemptNumber

`func (o *TaskReportingRecord) GetRetryAttemptNumber() int32`

GetRetryAttemptNumber returns the RetryAttemptNumber field if non-nil, zero value otherwise.

### GetRetryAttemptNumberOk

`func (o *TaskReportingRecord) GetRetryAttemptNumberOk() (*int32, bool)`

GetRetryAttemptNumberOk returns a tuple with the RetryAttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAttemptNumber

`func (o *TaskReportingRecord) SetRetryAttemptNumber(v int32)`

SetRetryAttemptNumber sets RetryAttemptNumber field to given value.

### HasRetryAttemptNumber

`func (o *TaskReportingRecord) HasRetryAttemptNumber() bool`

HasRetryAttemptNumber returns a boolean if a field has been set.

### GetCreatedViaApi

`func (o *TaskReportingRecord) GetCreatedViaApi() bool`

GetCreatedViaApi returns the CreatedViaApi field if non-nil, zero value otherwise.

### GetCreatedViaApiOk

`func (o *TaskReportingRecord) GetCreatedViaApiOk() (*bool, bool)`

GetCreatedViaApiOk returns a tuple with the CreatedViaApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedViaApi

`func (o *TaskReportingRecord) SetCreatedViaApi(v bool)`

SetCreatedViaApi sets CreatedViaApi field to given value.

### HasCreatedViaApi

`func (o *TaskReportingRecord) HasCreatedViaApi() bool`

HasCreatedViaApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


