/*
Digital.ai Release API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// FlagStatus the model 'FlagStatus'
type FlagStatus string

// List of FlagStatus
const (
	OK_FLAG FlagStatus = "OK_FLAG"
	ATTENTION_NEEDED_FLAG FlagStatus = "ATTENTION_NEEDED_FLAG"
	AT_RISK_FLAG FlagStatus = "AT_RISK_FLAG"
)

// All allowed values of FlagStatus enum
var AllowedFlagStatusEnumValues = []FlagStatus{
	"OK_FLAG",
	"ATTENTION_NEEDED_FLAG",
	"AT_RISK_FLAG",
}

func (v *FlagStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlagStatus(value)
	for _, existing := range AllowedFlagStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlagStatus", value)
}

// NewFlagStatusFromValue returns a pointer to a valid FlagStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlagStatusFromValue(v string) (*FlagStatus, error) {
	ev := FlagStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlagStatus: valid values are %v", v, AllowedFlagStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlagStatus) IsValid() bool {
	for _, existing := range AllowedFlagStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlagStatus value
func (v FlagStatus) Ptr() *FlagStatus {
	return &v
}

type NullableFlagStatus struct {
	value *FlagStatus
	isSet bool
}

func (v NullableFlagStatus) Get() *FlagStatus {
	return v.value
}

func (v *NullableFlagStatus) Set(val *FlagStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagStatus(val *FlagStatus) *NullableFlagStatus {
	return &NullableFlagStatus{value: val, isSet: true}
}

func (v NullableFlagStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
