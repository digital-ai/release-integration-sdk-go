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

// checks if the CompleteTransition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteTransition{}

// CompleteTransition struct for CompleteTransition
type CompleteTransition struct {
	TransitionItems []string `json:"transitionItems,omitempty"`
	CloseStages *bool `json:"closeStages,omitempty"`
}

// NewCompleteTransition instantiates a new CompleteTransition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteTransition() *CompleteTransition {
	this := CompleteTransition{}
	return &this
}

// NewCompleteTransitionWithDefaults instantiates a new CompleteTransition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteTransitionWithDefaults() *CompleteTransition {
	this := CompleteTransition{}
	return &this
}

// GetTransitionItems returns the TransitionItems field value if set, zero value otherwise.
func (o *CompleteTransition) GetTransitionItems() []string {
	if o == nil || isNil(o.TransitionItems) {
		var ret []string
		return ret
	}
	return o.TransitionItems
}

// GetTransitionItemsOk returns a tuple with the TransitionItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteTransition) GetTransitionItemsOk() ([]string, bool) {
	if o == nil || isNil(o.TransitionItems) {
		return nil, false
	}
	return o.TransitionItems, true
}

// HasTransitionItems returns a boolean if a field has been set.
func (o *CompleteTransition) HasTransitionItems() bool {
	if o != nil && !isNil(o.TransitionItems) {
		return true
	}

	return false
}

// SetTransitionItems gets a reference to the given []string and assigns it to the TransitionItems field.
func (o *CompleteTransition) SetTransitionItems(v []string) {
	o.TransitionItems = v
}

// GetCloseStages returns the CloseStages field value if set, zero value otherwise.
func (o *CompleteTransition) GetCloseStages() bool {
	if o == nil || isNil(o.CloseStages) {
		var ret bool
		return ret
	}
	return *o.CloseStages
}

// GetCloseStagesOk returns a tuple with the CloseStages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteTransition) GetCloseStagesOk() (*bool, bool) {
	if o == nil || isNil(o.CloseStages) {
		return nil, false
	}
	return o.CloseStages, true
}

// HasCloseStages returns a boolean if a field has been set.
func (o *CompleteTransition) HasCloseStages() bool {
	if o != nil && !isNil(o.CloseStages) {
		return true
	}

	return false
}

// SetCloseStages gets a reference to the given bool and assigns it to the CloseStages field.
func (o *CompleteTransition) SetCloseStages(v bool) {
	o.CloseStages = &v
}

func (o CompleteTransition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteTransition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TransitionItems) {
		toSerialize["transitionItems"] = o.TransitionItems
	}
	if !isNil(o.CloseStages) {
		toSerialize["closeStages"] = o.CloseStages
	}
	return toSerialize, nil
}

type NullableCompleteTransition struct {
	value *CompleteTransition
	isSet bool
}

func (v NullableCompleteTransition) Get() *CompleteTransition {
	return v.value
}

func (v *NullableCompleteTransition) Set(val *CompleteTransition) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteTransition) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteTransition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteTransition(val *CompleteTransition) *NullableCompleteTransition {
	return &NullableCompleteTransition{value: val, isSet: true}
}

func (v NullableCompleteTransition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteTransition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

