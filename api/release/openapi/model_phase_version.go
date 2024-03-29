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

// PhaseVersion the model 'PhaseVersion'
type PhaseVersion string

// List of PhaseVersion
const (
	PHASEVERSION_LATEST   PhaseVersion = "LATEST"
	PHASEVERSION_ORIGINAL PhaseVersion = "ORIGINAL"
	PHASEVERSION_ALL      PhaseVersion = "ALL"
)

// All allowed values of PhaseVersion enum
var AllowedPhaseVersionEnumValues = []PhaseVersion{
	"LATEST",
	"ORIGINAL",
	"ALL",
}

func (v *PhaseVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PhaseVersion(value)
	for _, existing := range AllowedPhaseVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PhaseVersion", value)
}

// NewPhaseVersionFromValue returns a pointer to a valid PhaseVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPhaseVersionFromValue(v string) (*PhaseVersion, error) {
	ev := PhaseVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PhaseVersion: valid values are %v", v, AllowedPhaseVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PhaseVersion) IsValid() bool {
	for _, existing := range AllowedPhaseVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PhaseVersion value
func (v PhaseVersion) Ptr() *PhaseVersion {
	return &v
}

type NullablePhaseVersion struct {
	value *PhaseVersion
	isSet bool
}

func (v NullablePhaseVersion) Get() *PhaseVersion {
	return v.value
}

func (v *NullablePhaseVersion) Set(val *PhaseVersion) {
	v.value = val
	v.isSet = true
}

func (v NullablePhaseVersion) IsSet() bool {
	return v.isSet
}

func (v *NullablePhaseVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhaseVersion(val *PhaseVersion) *NullablePhaseVersion {
	return &NullablePhaseVersion{value: val, isSet: true}
}

func (v NullablePhaseVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhaseVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
