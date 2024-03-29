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

// checks if the ApplicationFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationFilters{}

// ApplicationFilters struct for ApplicationFilters
type ApplicationFilters struct {
	Title        *string  `json:"title,omitempty"`
	Environments []string `json:"environments,omitempty"`
}

// NewApplicationFilters instantiates a new ApplicationFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFilters() *ApplicationFilters {
	this := ApplicationFilters{}
	return &this
}

// NewApplicationFiltersWithDefaults instantiates a new ApplicationFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFiltersWithDefaults() *ApplicationFilters {
	this := ApplicationFilters{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApplicationFilters) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFilters) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApplicationFilters) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApplicationFilters) SetTitle(v string) {
	o.Title = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *ApplicationFilters) GetEnvironments() []string {
	if o == nil || IsNil(o.Environments) {
		var ret []string
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFilters) GetEnvironmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *ApplicationFilters) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []string and assigns it to the Environments field.
func (o *ApplicationFilters) SetEnvironments(v []string) {
	o.Environments = v
}

func (o ApplicationFilters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	return toSerialize, nil
}

type NullableApplicationFilters struct {
	value *ApplicationFilters
	isSet bool
}

func (v NullableApplicationFilters) Get() *ApplicationFilters {
	return v.value
}

func (v *NullableApplicationFilters) Set(val *ApplicationFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFilters(val *ApplicationFilters) *NullableApplicationFilters {
	return &NullableApplicationFilters{value: val, isSet: true}
}

func (v NullableApplicationFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
