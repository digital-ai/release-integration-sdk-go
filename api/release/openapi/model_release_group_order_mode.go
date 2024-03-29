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

// ReleaseGroupOrderMode the model 'ReleaseGroupOrderMode'
type ReleaseGroupOrderMode string

// List of ReleaseGroupOrderMode
const (
	RELEASEGROUPORDERMODE_RISK       ReleaseGroupOrderMode = "RISK"
	RELEASEGROUPORDERMODE_START_DATE ReleaseGroupOrderMode = "START_DATE"
	RELEASEGROUPORDERMODE_END_DATE   ReleaseGroupOrderMode = "END_DATE"
)

// All allowed values of ReleaseGroupOrderMode enum
var AllowedReleaseGroupOrderModeEnumValues = []ReleaseGroupOrderMode{
	"RISK",
	"START_DATE",
	"END_DATE",
}

func (v *ReleaseGroupOrderMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReleaseGroupOrderMode(value)
	for _, existing := range AllowedReleaseGroupOrderModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReleaseGroupOrderMode", value)
}

// NewReleaseGroupOrderModeFromValue returns a pointer to a valid ReleaseGroupOrderMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReleaseGroupOrderModeFromValue(v string) (*ReleaseGroupOrderMode, error) {
	ev := ReleaseGroupOrderMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReleaseGroupOrderMode: valid values are %v", v, AllowedReleaseGroupOrderModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReleaseGroupOrderMode) IsValid() bool {
	for _, existing := range AllowedReleaseGroupOrderModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReleaseGroupOrderMode value
func (v ReleaseGroupOrderMode) Ptr() *ReleaseGroupOrderMode {
	return &v
}

type NullableReleaseGroupOrderMode struct {
	value *ReleaseGroupOrderMode
	isSet bool
}

func (v NullableReleaseGroupOrderMode) Get() *ReleaseGroupOrderMode {
	return v.value
}

func (v *NullableReleaseGroupOrderMode) Set(val *ReleaseGroupOrderMode) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseGroupOrderMode) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseGroupOrderMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseGroupOrderMode(val *ReleaseGroupOrderMode) *NullableReleaseGroupOrderMode {
	return &NullableReleaseGroupOrderMode{value: val, isSet: true}
}

func (v NullableReleaseGroupOrderMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseGroupOrderMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
