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

// checks if the SystemMessageSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemMessageSettings{}

// SystemMessageSettings struct for SystemMessageSettings
type SystemMessageSettings struct {
	FolderId *string `json:"folderId,omitempty"`
	Title *string `json:"title,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Message *string `json:"message,omitempty"`
	Automated *bool `json:"automated,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	EndDate *string `json:"endDate,omitempty"`
}

// NewSystemMessageSettings instantiates a new SystemMessageSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemMessageSettings() *SystemMessageSettings {
	this := SystemMessageSettings{}
	return &this
}

// NewSystemMessageSettingsWithDefaults instantiates a new SystemMessageSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemMessageSettingsWithDefaults() *SystemMessageSettings {
	this := SystemMessageSettings{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *SystemMessageSettings) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMessageSettings) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *SystemMessageSettings) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *SystemMessageSettings) SetFolderId(v string) {
	o.FolderId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SystemMessageSettings) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMessageSettings) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SystemMessageSettings) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SystemMessageSettings) SetTitle(v string) {
	o.Title = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SystemMessageSettings) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMessageSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SystemMessageSettings) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SystemMessageSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SystemMessageSettings) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMessageSettings) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SystemMessageSettings) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SystemMessageSettings) SetMessage(v string) {
	o.Message = &v
}

// GetAutomated returns the Automated field value if set, zero value otherwise.
func (o *SystemMessageSettings) GetAutomated() bool {
	if o == nil || isNil(o.Automated) {
		var ret bool
		return ret
	}
	return *o.Automated
}

// GetAutomatedOk returns a tuple with the Automated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMessageSettings) GetAutomatedOk() (*bool, bool) {
	if o == nil || isNil(o.Automated) {
		return nil, false
	}
	return o.Automated, true
}

// HasAutomated returns a boolean if a field has been set.
func (o *SystemMessageSettings) HasAutomated() bool {
	if o != nil && !isNil(o.Automated) {
		return true
	}

	return false
}

// SetAutomated gets a reference to the given bool and assigns it to the Automated field.
func (o *SystemMessageSettings) SetAutomated(v bool) {
	o.Automated = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SystemMessageSettings) GetStartDate() string {
	if o == nil || isNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMessageSettings) GetStartDateOk() (*string, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SystemMessageSettings) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SystemMessageSettings) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SystemMessageSettings) GetEndDate() string {
	if o == nil || isNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMessageSettings) GetEndDateOk() (*string, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SystemMessageSettings) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SystemMessageSettings) SetEndDate(v string) {
	o.EndDate = &v
}

func (o SystemMessageSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemMessageSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Automated) {
		toSerialize["automated"] = o.Automated
	}
	if !isNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !isNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableSystemMessageSettings struct {
	value *SystemMessageSettings
	isSet bool
}

func (v NullableSystemMessageSettings) Get() *SystemMessageSettings {
	return v.value
}

func (v *NullableSystemMessageSettings) Set(val *SystemMessageSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemMessageSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemMessageSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemMessageSettings(val *SystemMessageSettings) *NullableSystemMessageSettings {
	return &NullableSystemMessageSettings{value: val, isSet: true}
}

func (v NullableSystemMessageSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemMessageSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

