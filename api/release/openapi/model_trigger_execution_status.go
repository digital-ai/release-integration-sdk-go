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

// TriggerExecutionStatus the model 'TriggerExecutionStatus'
type TriggerExecutionStatus string

// List of TriggerExecutionStatus
const (
	SUCCESS TriggerExecutionStatus = "SUCCESS"
	FAILURE TriggerExecutionStatus = "FAILURE"
)

// All allowed values of TriggerExecutionStatus enum
var AllowedTriggerExecutionStatusEnumValues = []TriggerExecutionStatus{
	"SUCCESS",
	"FAILURE",
}

func (v *TriggerExecutionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TriggerExecutionStatus(value)
	for _, existing := range AllowedTriggerExecutionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TriggerExecutionStatus", value)
}

// NewTriggerExecutionStatusFromValue returns a pointer to a valid TriggerExecutionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggerExecutionStatusFromValue(v string) (*TriggerExecutionStatus, error) {
	ev := TriggerExecutionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TriggerExecutionStatus: valid values are %v", v, AllowedTriggerExecutionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TriggerExecutionStatus) IsValid() bool {
	for _, existing := range AllowedTriggerExecutionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerExecutionStatus value
func (v TriggerExecutionStatus) Ptr() *TriggerExecutionStatus {
	return &v
}

type NullableTriggerExecutionStatus struct {
	value *TriggerExecutionStatus
	isSet bool
}

func (v NullableTriggerExecutionStatus) Get() *TriggerExecutionStatus {
	return v.value
}

func (v *NullableTriggerExecutionStatus) Set(val *TriggerExecutionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerExecutionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerExecutionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerExecutionStatus(val *TriggerExecutionStatus) *NullableTriggerExecutionStatus {
	return &NullableTriggerExecutionStatus{value: val, isSet: true}
}

func (v NullableTriggerExecutionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerExecutionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
