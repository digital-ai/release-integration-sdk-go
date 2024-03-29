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

// checks if the EnvironmentFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentFilters{}

// EnvironmentFilters struct for EnvironmentFilters
type EnvironmentFilters struct {
	Title  *string  `json:"title,omitempty"`
	Stage  *string  `json:"stage,omitempty"`
	Labels []string `json:"labels,omitempty"`
}

// NewEnvironmentFilters instantiates a new EnvironmentFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentFilters() *EnvironmentFilters {
	this := EnvironmentFilters{}
	return &this
}

// NewEnvironmentFiltersWithDefaults instantiates a new EnvironmentFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentFiltersWithDefaults() *EnvironmentFilters {
	this := EnvironmentFilters{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EnvironmentFilters) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentFilters) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EnvironmentFilters) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *EnvironmentFilters) SetTitle(v string) {
	o.Title = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *EnvironmentFilters) GetStage() string {
	if o == nil || IsNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentFilters) GetStageOk() (*string, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *EnvironmentFilters) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *EnvironmentFilters) SetStage(v string) {
	o.Stage = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *EnvironmentFilters) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentFilters) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *EnvironmentFilters) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *EnvironmentFilters) SetLabels(v []string) {
	o.Labels = v
}

func (o EnvironmentFilters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableEnvironmentFilters struct {
	value *EnvironmentFilters
	isSet bool
}

func (v NullableEnvironmentFilters) Get() *EnvironmentFilters {
	return v.value
}

func (v *NullableEnvironmentFilters) Set(val *EnvironmentFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentFilters(val *EnvironmentFilters) *NullableEnvironmentFilters {
	return &NullableEnvironmentFilters{value: val, isSet: true}
}

func (v NullableEnvironmentFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
