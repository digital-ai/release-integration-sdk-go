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

// checks if the BaseApplicationView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseApplicationView{}

// BaseApplicationView struct for BaseApplicationView
type BaseApplicationView struct {
	Id    *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewBaseApplicationView instantiates a new BaseApplicationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseApplicationView() *BaseApplicationView {
	this := BaseApplicationView{}
	return &this
}

// NewBaseApplicationViewWithDefaults instantiates a new BaseApplicationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseApplicationViewWithDefaults() *BaseApplicationView {
	this := BaseApplicationView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseApplicationView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseApplicationView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseApplicationView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseApplicationView) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BaseApplicationView) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseApplicationView) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BaseApplicationView) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BaseApplicationView) SetTitle(v string) {
	o.Title = &v
}

func (o BaseApplicationView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseApplicationView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableBaseApplicationView struct {
	value *BaseApplicationView
	isSet bool
}

func (v NullableBaseApplicationView) Get() *BaseApplicationView {
	return v.value
}

func (v *NullableBaseApplicationView) Set(val *BaseApplicationView) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseApplicationView) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseApplicationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseApplicationView(val *BaseApplicationView) *NullableBaseApplicationView {
	return &NullableBaseApplicationView{value: val, isSet: true}
}

func (v NullableBaseApplicationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseApplicationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
