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

// ReleaseOrderMode the model 'ReleaseOrderMode'
type ReleaseOrderMode string

// List of ReleaseOrderMode
const (
	RELEASEORDERMODE_RISK       ReleaseOrderMode = "risk"
	RELEASEORDERMODE_START_DATE ReleaseOrderMode = "start_date"
	RELEASEORDERMODE_END_DATE   ReleaseOrderMode = "end_date"
	RELEASEORDERMODE_TITLE      ReleaseOrderMode = "title"
)

// All allowed values of ReleaseOrderMode enum
var AllowedReleaseOrderModeEnumValues = []ReleaseOrderMode{
	"risk",
	"start_date",
	"end_date",
	"title",
}

func (v *ReleaseOrderMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReleaseOrderMode(value)
	for _, existing := range AllowedReleaseOrderModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReleaseOrderMode", value)
}

// NewReleaseOrderModeFromValue returns a pointer to a valid ReleaseOrderMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReleaseOrderModeFromValue(v string) (*ReleaseOrderMode, error) {
	ev := ReleaseOrderMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReleaseOrderMode: valid values are %v", v, AllowedReleaseOrderModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReleaseOrderMode) IsValid() bool {
	for _, existing := range AllowedReleaseOrderModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReleaseOrderMode value
func (v ReleaseOrderMode) Ptr() *ReleaseOrderMode {
	return &v
}

type NullableReleaseOrderMode struct {
	value *ReleaseOrderMode
	isSet bool
}

func (v NullableReleaseOrderMode) Get() *ReleaseOrderMode {
	return v.value
}

func (v *NullableReleaseOrderMode) Set(val *ReleaseOrderMode) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseOrderMode) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseOrderMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseOrderMode(val *ReleaseOrderMode) *NullableReleaseOrderMode {
	return &NullableReleaseOrderMode{value: val, isSet: true}
}

func (v NullableReleaseOrderMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseOrderMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
