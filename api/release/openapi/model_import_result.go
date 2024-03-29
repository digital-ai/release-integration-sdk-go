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

// checks if the ImportResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportResult{}

// ImportResult struct for ImportResult
type ImportResult struct {
	Id       *string  `json:"id,omitempty"`
	Title    *string  `json:"title,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// NewImportResult instantiates a new ImportResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportResult() *ImportResult {
	this := ImportResult{}
	return &this
}

// NewImportResultWithDefaults instantiates a new ImportResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportResultWithDefaults() *ImportResult {
	this := ImportResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImportResult) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ImportResult) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportResult) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ImportResult) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ImportResult) SetTitle(v string) {
	o.Title = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ImportResult) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportResult) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ImportResult) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *ImportResult) SetWarnings(v []string) {
	o.Warnings = v
}

func (o ImportResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableImportResult struct {
	value *ImportResult
	isSet bool
}

func (v NullableImportResult) Get() *ImportResult {
	return v.value
}

func (v *NullableImportResult) Set(val *ImportResult) {
	v.value = val
	v.isSet = true
}

func (v NullableImportResult) IsSet() bool {
	return v.isSet
}

func (v *NullableImportResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportResult(val *ImportResult) *NullableImportResult {
	return &NullableImportResult{value: val, isSet: true}
}

func (v NullableImportResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
