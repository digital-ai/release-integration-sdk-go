/*
Digital.ai Release API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TaskReportingRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskReportingRecord{}

// TaskReportingRecord struct for TaskReportingRecord
type TaskReportingRecord struct {
	Scope *FacetScope `json:"scope,omitempty"`
	TargetId *string `json:"targetId,omitempty"`
	ConfigurationUri *string `json:"configurationUri,omitempty"`
	VariableMapping *map[string]string `json:"variableMapping,omitempty"`
	VariableUsages []UsagePoint `json:"variableUsages,omitempty"`
	PropertiesWithVariables []interface{} `json:"propertiesWithVariables,omitempty"`
	ServerUrl *string `json:"serverUrl,omitempty"`
	ServerUser *string `json:"serverUser,omitempty"`
	CreationDate *string `json:"creationDate,omitempty"`
	RetryAttemptNumber *int32 `json:"retryAttemptNumber,omitempty"`
	CreatedViaApi *bool `json:"createdViaApi,omitempty"`
}

// NewTaskReportingRecord instantiates a new TaskReportingRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskReportingRecord() *TaskReportingRecord {
	this := TaskReportingRecord{}
	return &this
}

// NewTaskReportingRecordWithDefaults instantiates a new TaskReportingRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskReportingRecordWithDefaults() *TaskReportingRecord {
	this := TaskReportingRecord{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetScope() FacetScope {
	if o == nil || isNil(o.Scope) {
		var ret FacetScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetScopeOk() (*FacetScope, bool) {
	if o == nil || isNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given FacetScope and assigns it to the Scope field.
func (o *TaskReportingRecord) SetScope(v FacetScope) {
	o.Scope = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetTargetId() string {
	if o == nil || isNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetTargetIdOk() (*string, bool) {
	if o == nil || isNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasTargetId() bool {
	if o != nil && !isNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *TaskReportingRecord) SetTargetId(v string) {
	o.TargetId = &v
}

// GetConfigurationUri returns the ConfigurationUri field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetConfigurationUri() string {
	if o == nil || isNil(o.ConfigurationUri) {
		var ret string
		return ret
	}
	return *o.ConfigurationUri
}

// GetConfigurationUriOk returns a tuple with the ConfigurationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetConfigurationUriOk() (*string, bool) {
	if o == nil || isNil(o.ConfigurationUri) {
		return nil, false
	}
	return o.ConfigurationUri, true
}

// HasConfigurationUri returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasConfigurationUri() bool {
	if o != nil && !isNil(o.ConfigurationUri) {
		return true
	}

	return false
}

// SetConfigurationUri gets a reference to the given string and assigns it to the ConfigurationUri field.
func (o *TaskReportingRecord) SetConfigurationUri(v string) {
	o.ConfigurationUri = &v
}

// GetVariableMapping returns the VariableMapping field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetVariableMapping() map[string]string {
	if o == nil || isNil(o.VariableMapping) {
		var ret map[string]string
		return ret
	}
	return *o.VariableMapping
}

// GetVariableMappingOk returns a tuple with the VariableMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetVariableMappingOk() (*map[string]string, bool) {
	if o == nil || isNil(o.VariableMapping) {
		return nil, false
	}
	return o.VariableMapping, true
}

// HasVariableMapping returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasVariableMapping() bool {
	if o != nil && !isNil(o.VariableMapping) {
		return true
	}

	return false
}

// SetVariableMapping gets a reference to the given map[string]string and assigns it to the VariableMapping field.
func (o *TaskReportingRecord) SetVariableMapping(v map[string]string) {
	o.VariableMapping = &v
}

// GetVariableUsages returns the VariableUsages field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetVariableUsages() []UsagePoint {
	if o == nil || isNil(o.VariableUsages) {
		var ret []UsagePoint
		return ret
	}
	return o.VariableUsages
}

// GetVariableUsagesOk returns a tuple with the VariableUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetVariableUsagesOk() ([]UsagePoint, bool) {
	if o == nil || isNil(o.VariableUsages) {
		return nil, false
	}
	return o.VariableUsages, true
}

// HasVariableUsages returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasVariableUsages() bool {
	if o != nil && !isNil(o.VariableUsages) {
		return true
	}

	return false
}

// SetVariableUsages gets a reference to the given []UsagePoint and assigns it to the VariableUsages field.
func (o *TaskReportingRecord) SetVariableUsages(v []UsagePoint) {
	o.VariableUsages = v
}

// GetPropertiesWithVariables returns the PropertiesWithVariables field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetPropertiesWithVariables() []interface{} {
	if o == nil || isNil(o.PropertiesWithVariables) {
		var ret []interface{}
		return ret
	}
	return o.PropertiesWithVariables
}

// GetPropertiesWithVariablesOk returns a tuple with the PropertiesWithVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetPropertiesWithVariablesOk() ([]interface{}, bool) {
	if o == nil || isNil(o.PropertiesWithVariables) {
		return nil, false
	}
	return o.PropertiesWithVariables, true
}

// HasPropertiesWithVariables returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasPropertiesWithVariables() bool {
	if o != nil && !isNil(o.PropertiesWithVariables) {
		return true
	}

	return false
}

// SetPropertiesWithVariables gets a reference to the given []interface{} and assigns it to the PropertiesWithVariables field.
func (o *TaskReportingRecord) SetPropertiesWithVariables(v []interface{}) {
	o.PropertiesWithVariables = v
}

// GetServerUrl returns the ServerUrl field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetServerUrl() string {
	if o == nil || isNil(o.ServerUrl) {
		var ret string
		return ret
	}
	return *o.ServerUrl
}

// GetServerUrlOk returns a tuple with the ServerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetServerUrlOk() (*string, bool) {
	if o == nil || isNil(o.ServerUrl) {
		return nil, false
	}
	return o.ServerUrl, true
}

// HasServerUrl returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasServerUrl() bool {
	if o != nil && !isNil(o.ServerUrl) {
		return true
	}

	return false
}

// SetServerUrl gets a reference to the given string and assigns it to the ServerUrl field.
func (o *TaskReportingRecord) SetServerUrl(v string) {
	o.ServerUrl = &v
}

// GetServerUser returns the ServerUser field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetServerUser() string {
	if o == nil || isNil(o.ServerUser) {
		var ret string
		return ret
	}
	return *o.ServerUser
}

// GetServerUserOk returns a tuple with the ServerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetServerUserOk() (*string, bool) {
	if o == nil || isNil(o.ServerUser) {
		return nil, false
	}
	return o.ServerUser, true
}

// HasServerUser returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasServerUser() bool {
	if o != nil && !isNil(o.ServerUser) {
		return true
	}

	return false
}

// SetServerUser gets a reference to the given string and assigns it to the ServerUser field.
func (o *TaskReportingRecord) SetServerUser(v string) {
	o.ServerUser = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetCreationDate() string {
	if o == nil || isNil(o.CreationDate) {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetCreationDateOk() (*string, bool) {
	if o == nil || isNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasCreationDate() bool {
	if o != nil && !isNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *TaskReportingRecord) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetRetryAttemptNumber returns the RetryAttemptNumber field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetRetryAttemptNumber() int32 {
	if o == nil || isNil(o.RetryAttemptNumber) {
		var ret int32
		return ret
	}
	return *o.RetryAttemptNumber
}

// GetRetryAttemptNumberOk returns a tuple with the RetryAttemptNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetRetryAttemptNumberOk() (*int32, bool) {
	if o == nil || isNil(o.RetryAttemptNumber) {
		return nil, false
	}
	return o.RetryAttemptNumber, true
}

// HasRetryAttemptNumber returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasRetryAttemptNumber() bool {
	if o != nil && !isNil(o.RetryAttemptNumber) {
		return true
	}

	return false
}

// SetRetryAttemptNumber gets a reference to the given int32 and assigns it to the RetryAttemptNumber field.
func (o *TaskReportingRecord) SetRetryAttemptNumber(v int32) {
	o.RetryAttemptNumber = &v
}

// GetCreatedViaApi returns the CreatedViaApi field value if set, zero value otherwise.
func (o *TaskReportingRecord) GetCreatedViaApi() bool {
	if o == nil || isNil(o.CreatedViaApi) {
		var ret bool
		return ret
	}
	return *o.CreatedViaApi
}

// GetCreatedViaApiOk returns a tuple with the CreatedViaApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskReportingRecord) GetCreatedViaApiOk() (*bool, bool) {
	if o == nil || isNil(o.CreatedViaApi) {
		return nil, false
	}
	return o.CreatedViaApi, true
}

// HasCreatedViaApi returns a boolean if a field has been set.
func (o *TaskReportingRecord) HasCreatedViaApi() bool {
	if o != nil && !isNil(o.CreatedViaApi) {
		return true
	}

	return false
}

// SetCreatedViaApi gets a reference to the given bool and assigns it to the CreatedViaApi field.
func (o *TaskReportingRecord) SetCreatedViaApi(v bool) {
	o.CreatedViaApi = &v
}

func (o TaskReportingRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskReportingRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !isNil(o.ConfigurationUri) {
		toSerialize["configurationUri"] = o.ConfigurationUri
	}
	if !isNil(o.VariableMapping) {
		toSerialize["variableMapping"] = o.VariableMapping
	}
	if !isNil(o.VariableUsages) {
		toSerialize["variableUsages"] = o.VariableUsages
	}
	if !isNil(o.PropertiesWithVariables) {
		toSerialize["propertiesWithVariables"] = o.PropertiesWithVariables
	}
	if !isNil(o.ServerUrl) {
		toSerialize["serverUrl"] = o.ServerUrl
	}
	if !isNil(o.ServerUser) {
		toSerialize["serverUser"] = o.ServerUser
	}
	if !isNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !isNil(o.RetryAttemptNumber) {
		toSerialize["retryAttemptNumber"] = o.RetryAttemptNumber
	}
	if !isNil(o.CreatedViaApi) {
		toSerialize["createdViaApi"] = o.CreatedViaApi
	}
	return toSerialize, nil
}

type NullableTaskReportingRecord struct {
	value *TaskReportingRecord
	isSet bool
}

func (v NullableTaskReportingRecord) Get() *TaskReportingRecord {
	return v.value
}

func (v *NullableTaskReportingRecord) Set(val *TaskReportingRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskReportingRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskReportingRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskReportingRecord(val *TaskReportingRecord) *NullableTaskReportingRecord {
	return &NullableTaskReportingRecord{value: val, isSet: true}
}

func (v NullableTaskReportingRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskReportingRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

