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

// checks if the RiskStatusWithThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskStatusWithThresholds{}

// RiskStatusWithThresholds struct for RiskStatusWithThresholds
type RiskStatusWithThresholds struct {
	RiskStatus *RiskStatus `json:"riskStatus,omitempty"`
	AttentionNeededFrom *int32 `json:"attentionNeededFrom,omitempty"`
	AtRiskFrom *int32 `json:"atRiskFrom,omitempty"`
}

// NewRiskStatusWithThresholds instantiates a new RiskStatusWithThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskStatusWithThresholds() *RiskStatusWithThresholds {
	this := RiskStatusWithThresholds{}
	return &this
}

// NewRiskStatusWithThresholdsWithDefaults instantiates a new RiskStatusWithThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskStatusWithThresholdsWithDefaults() *RiskStatusWithThresholds {
	this := RiskStatusWithThresholds{}
	return &this
}

// GetRiskStatus returns the RiskStatus field value if set, zero value otherwise.
func (o *RiskStatusWithThresholds) GetRiskStatus() RiskStatus {
	if o == nil || isNil(o.RiskStatus) {
		var ret RiskStatus
		return ret
	}
	return *o.RiskStatus
}

// GetRiskStatusOk returns a tuple with the RiskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskStatusWithThresholds) GetRiskStatusOk() (*RiskStatus, bool) {
	if o == nil || isNil(o.RiskStatus) {
		return nil, false
	}
	return o.RiskStatus, true
}

// HasRiskStatus returns a boolean if a field has been set.
func (o *RiskStatusWithThresholds) HasRiskStatus() bool {
	if o != nil && !isNil(o.RiskStatus) {
		return true
	}

	return false
}

// SetRiskStatus gets a reference to the given RiskStatus and assigns it to the RiskStatus field.
func (o *RiskStatusWithThresholds) SetRiskStatus(v RiskStatus) {
	o.RiskStatus = &v
}

// GetAttentionNeededFrom returns the AttentionNeededFrom field value if set, zero value otherwise.
func (o *RiskStatusWithThresholds) GetAttentionNeededFrom() int32 {
	if o == nil || isNil(o.AttentionNeededFrom) {
		var ret int32
		return ret
	}
	return *o.AttentionNeededFrom
}

// GetAttentionNeededFromOk returns a tuple with the AttentionNeededFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskStatusWithThresholds) GetAttentionNeededFromOk() (*int32, bool) {
	if o == nil || isNil(o.AttentionNeededFrom) {
		return nil, false
	}
	return o.AttentionNeededFrom, true
}

// HasAttentionNeededFrom returns a boolean if a field has been set.
func (o *RiskStatusWithThresholds) HasAttentionNeededFrom() bool {
	if o != nil && !isNil(o.AttentionNeededFrom) {
		return true
	}

	return false
}

// SetAttentionNeededFrom gets a reference to the given int32 and assigns it to the AttentionNeededFrom field.
func (o *RiskStatusWithThresholds) SetAttentionNeededFrom(v int32) {
	o.AttentionNeededFrom = &v
}

// GetAtRiskFrom returns the AtRiskFrom field value if set, zero value otherwise.
func (o *RiskStatusWithThresholds) GetAtRiskFrom() int32 {
	if o == nil || isNil(o.AtRiskFrom) {
		var ret int32
		return ret
	}
	return *o.AtRiskFrom
}

// GetAtRiskFromOk returns a tuple with the AtRiskFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskStatusWithThresholds) GetAtRiskFromOk() (*int32, bool) {
	if o == nil || isNil(o.AtRiskFrom) {
		return nil, false
	}
	return o.AtRiskFrom, true
}

// HasAtRiskFrom returns a boolean if a field has been set.
func (o *RiskStatusWithThresholds) HasAtRiskFrom() bool {
	if o != nil && !isNil(o.AtRiskFrom) {
		return true
	}

	return false
}

// SetAtRiskFrom gets a reference to the given int32 and assigns it to the AtRiskFrom field.
func (o *RiskStatusWithThresholds) SetAtRiskFrom(v int32) {
	o.AtRiskFrom = &v
}

func (o RiskStatusWithThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskStatusWithThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RiskStatus) {
		toSerialize["riskStatus"] = o.RiskStatus
	}
	if !isNil(o.AttentionNeededFrom) {
		toSerialize["attentionNeededFrom"] = o.AttentionNeededFrom
	}
	if !isNil(o.AtRiskFrom) {
		toSerialize["atRiskFrom"] = o.AtRiskFrom
	}
	return toSerialize, nil
}

type NullableRiskStatusWithThresholds struct {
	value *RiskStatusWithThresholds
	isSet bool
}

func (v NullableRiskStatusWithThresholds) Get() *RiskStatusWithThresholds {
	return v.value
}

func (v *NullableRiskStatusWithThresholds) Set(val *RiskStatusWithThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskStatusWithThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskStatusWithThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskStatusWithThresholds(val *RiskStatusWithThresholds) *NullableRiskStatusWithThresholds {
	return &NullableRiskStatusWithThresholds{value: val, isSet: true}
}

func (v NullableRiskStatusWithThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskStatusWithThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

