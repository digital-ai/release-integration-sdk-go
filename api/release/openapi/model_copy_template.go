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

// checks if the CopyTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyTemplate{}

// CopyTemplate struct for CopyTemplate
type CopyTemplate struct {
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewCopyTemplate instantiates a new CopyTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyTemplate() *CopyTemplate {
	this := CopyTemplate{}
	return &this
}

// NewCopyTemplateWithDefaults instantiates a new CopyTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyTemplateWithDefaults() *CopyTemplate {
	this := CopyTemplate{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CopyTemplate) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyTemplate) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CopyTemplate) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CopyTemplate) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CopyTemplate) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CopyTemplate) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CopyTemplate) SetDescription(v string) {
	o.Description = &v
}

func (o CopyTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableCopyTemplate struct {
	value *CopyTemplate
	isSet bool
}

func (v NullableCopyTemplate) Get() *CopyTemplate {
	return v.value
}

func (v *NullableCopyTemplate) Set(val *CopyTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyTemplate(val *CopyTemplate) *NullableCopyTemplate {
	return &NullableCopyTemplate{value: val, isSet: true}
}

func (v NullableCopyTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

