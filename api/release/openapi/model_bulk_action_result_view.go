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

// checks if the BulkActionResultView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkActionResultView{}

// BulkActionResultView struct for BulkActionResultView
type BulkActionResultView struct {
	UpdatedIds []string `json:"updatedIds,omitempty"`
}

// NewBulkActionResultView instantiates a new BulkActionResultView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkActionResultView() *BulkActionResultView {
	this := BulkActionResultView{}
	return &this
}

// NewBulkActionResultViewWithDefaults instantiates a new BulkActionResultView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkActionResultViewWithDefaults() *BulkActionResultView {
	this := BulkActionResultView{}
	return &this
}

// GetUpdatedIds returns the UpdatedIds field value if set, zero value otherwise.
func (o *BulkActionResultView) GetUpdatedIds() []string {
	if o == nil || isNil(o.UpdatedIds) {
		var ret []string
		return ret
	}
	return o.UpdatedIds
}

// GetUpdatedIdsOk returns a tuple with the UpdatedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkActionResultView) GetUpdatedIdsOk() ([]string, bool) {
	if o == nil || isNil(o.UpdatedIds) {
		return nil, false
	}
	return o.UpdatedIds, true
}

// HasUpdatedIds returns a boolean if a field has been set.
func (o *BulkActionResultView) HasUpdatedIds() bool {
	if o != nil && !isNil(o.UpdatedIds) {
		return true
	}

	return false
}

// SetUpdatedIds gets a reference to the given []string and assigns it to the UpdatedIds field.
func (o *BulkActionResultView) SetUpdatedIds(v []string) {
	o.UpdatedIds = v
}

func (o BulkActionResultView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkActionResultView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UpdatedIds) {
		toSerialize["updatedIds"] = o.UpdatedIds
	}
	return toSerialize, nil
}

type NullableBulkActionResultView struct {
	value *BulkActionResultView
	isSet bool
}

func (v NullableBulkActionResultView) Get() *BulkActionResultView {
	return v.value
}

func (v *NullableBulkActionResultView) Set(val *BulkActionResultView) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkActionResultView) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkActionResultView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkActionResultView(val *BulkActionResultView) *NullableBulkActionResultView {
	return &NullableBulkActionResultView{value: val, isSet: true}
}

func (v NullableBulkActionResultView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkActionResultView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

