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

// checks if the GlobalVariables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalVariables{}

// GlobalVariables struct for GlobalVariables
type GlobalVariables struct {
	Variables []Variable `json:"variables,omitempty"`
	VariablesByKeys *map[string]Variable `json:"variablesByKeys,omitempty"`
	StringVariableValues *map[string]string `json:"stringVariableValues,omitempty"`
	PasswordVariableValues *map[string]string `json:"passwordVariableValues,omitempty"`
}

// NewGlobalVariables instantiates a new GlobalVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalVariables() *GlobalVariables {
	this := GlobalVariables{}
	return &this
}

// NewGlobalVariablesWithDefaults instantiates a new GlobalVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalVariablesWithDefaults() *GlobalVariables {
	this := GlobalVariables{}
	return &this
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *GlobalVariables) GetVariables() []Variable {
	if o == nil || isNil(o.Variables) {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalVariables) GetVariablesOk() ([]Variable, bool) {
	if o == nil || isNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *GlobalVariables) HasVariables() bool {
	if o != nil && !isNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *GlobalVariables) SetVariables(v []Variable) {
	o.Variables = v
}

// GetVariablesByKeys returns the VariablesByKeys field value if set, zero value otherwise.
func (o *GlobalVariables) GetVariablesByKeys() map[string]Variable {
	if o == nil || isNil(o.VariablesByKeys) {
		var ret map[string]Variable
		return ret
	}
	return *o.VariablesByKeys
}

// GetVariablesByKeysOk returns a tuple with the VariablesByKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalVariables) GetVariablesByKeysOk() (*map[string]Variable, bool) {
	if o == nil || isNil(o.VariablesByKeys) {
		return nil, false
	}
	return o.VariablesByKeys, true
}

// HasVariablesByKeys returns a boolean if a field has been set.
func (o *GlobalVariables) HasVariablesByKeys() bool {
	if o != nil && !isNil(o.VariablesByKeys) {
		return true
	}

	return false
}

// SetVariablesByKeys gets a reference to the given map[string]Variable and assigns it to the VariablesByKeys field.
func (o *GlobalVariables) SetVariablesByKeys(v map[string]Variable) {
	o.VariablesByKeys = &v
}

// GetStringVariableValues returns the StringVariableValues field value if set, zero value otherwise.
func (o *GlobalVariables) GetStringVariableValues() map[string]string {
	if o == nil || isNil(o.StringVariableValues) {
		var ret map[string]string
		return ret
	}
	return *o.StringVariableValues
}

// GetStringVariableValuesOk returns a tuple with the StringVariableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalVariables) GetStringVariableValuesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.StringVariableValues) {
		return nil, false
	}
	return o.StringVariableValues, true
}

// HasStringVariableValues returns a boolean if a field has been set.
func (o *GlobalVariables) HasStringVariableValues() bool {
	if o != nil && !isNil(o.StringVariableValues) {
		return true
	}

	return false
}

// SetStringVariableValues gets a reference to the given map[string]string and assigns it to the StringVariableValues field.
func (o *GlobalVariables) SetStringVariableValues(v map[string]string) {
	o.StringVariableValues = &v
}

// GetPasswordVariableValues returns the PasswordVariableValues field value if set, zero value otherwise.
func (o *GlobalVariables) GetPasswordVariableValues() map[string]string {
	if o == nil || isNil(o.PasswordVariableValues) {
		var ret map[string]string
		return ret
	}
	return *o.PasswordVariableValues
}

// GetPasswordVariableValuesOk returns a tuple with the PasswordVariableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalVariables) GetPasswordVariableValuesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.PasswordVariableValues) {
		return nil, false
	}
	return o.PasswordVariableValues, true
}

// HasPasswordVariableValues returns a boolean if a field has been set.
func (o *GlobalVariables) HasPasswordVariableValues() bool {
	if o != nil && !isNil(o.PasswordVariableValues) {
		return true
	}

	return false
}

// SetPasswordVariableValues gets a reference to the given map[string]string and assigns it to the PasswordVariableValues field.
func (o *GlobalVariables) SetPasswordVariableValues(v map[string]string) {
	o.PasswordVariableValues = &v
}

func (o GlobalVariables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalVariables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !isNil(o.VariablesByKeys) {
		toSerialize["variablesByKeys"] = o.VariablesByKeys
	}
	if !isNil(o.StringVariableValues) {
		toSerialize["stringVariableValues"] = o.StringVariableValues
	}
	if !isNil(o.PasswordVariableValues) {
		toSerialize["passwordVariableValues"] = o.PasswordVariableValues
	}
	return toSerialize, nil
}

type NullableGlobalVariables struct {
	value *GlobalVariables
	isSet bool
}

func (v NullableGlobalVariables) Get() *GlobalVariables {
	return v.value
}

func (v *NullableGlobalVariables) Set(val *GlobalVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalVariables(val *GlobalVariables) *NullableGlobalVariables {
	return &NullableGlobalVariables{value: val, isSet: true}
}

func (v NullableGlobalVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

