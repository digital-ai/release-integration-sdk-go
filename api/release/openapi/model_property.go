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

// checks if the Property type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Property{}

// Property struct for Property
type Property struct {
	INDEXED_PROPERTY_PATTERN *string `json:"INDEXED_PROPERTY_PATTERN,omitempty"`
	PropertyName *string `json:"propertyName,omitempty"`
	Index *int32 `json:"index,omitempty"`
	Indexed *bool `json:"indexed,omitempty"`
}

// NewProperty instantiates a new Property object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProperty() *Property {
	this := Property{}
	return &this
}

// NewPropertyWithDefaults instantiates a new Property object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyWithDefaults() *Property {
	this := Property{}
	return &this
}

// GetINDEXED_PROPERTY_PATTERN returns the INDEXED_PROPERTY_PATTERN field value if set, zero value otherwise.
func (o *Property) GetINDEXED_PROPERTY_PATTERN() string {
	if o == nil || isNil(o.INDEXED_PROPERTY_PATTERN) {
		var ret string
		return ret
	}
	return *o.INDEXED_PROPERTY_PATTERN
}

// GetINDEXED_PROPERTY_PATTERNOk returns a tuple with the INDEXED_PROPERTY_PATTERN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetINDEXED_PROPERTY_PATTERNOk() (*string, bool) {
	if o == nil || isNil(o.INDEXED_PROPERTY_PATTERN) {
		return nil, false
	}
	return o.INDEXED_PROPERTY_PATTERN, true
}

// HasINDEXED_PROPERTY_PATTERN returns a boolean if a field has been set.
func (o *Property) HasINDEXED_PROPERTY_PATTERN() bool {
	if o != nil && !isNil(o.INDEXED_PROPERTY_PATTERN) {
		return true
	}

	return false
}

// SetINDEXED_PROPERTY_PATTERN gets a reference to the given string and assigns it to the INDEXED_PROPERTY_PATTERN field.
func (o *Property) SetINDEXED_PROPERTY_PATTERN(v string) {
	o.INDEXED_PROPERTY_PATTERN = &v
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *Property) GetPropertyName() string {
	if o == nil || isNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetPropertyNameOk() (*string, bool) {
	if o == nil || isNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *Property) HasPropertyName() bool {
	if o != nil && !isNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *Property) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *Property) GetIndex() int32 {
	if o == nil || isNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetIndexOk() (*int32, bool) {
	if o == nil || isNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *Property) HasIndex() bool {
	if o != nil && !isNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *Property) SetIndex(v int32) {
	o.Index = &v
}

// GetIndexed returns the Indexed field value if set, zero value otherwise.
func (o *Property) GetIndexed() bool {
	if o == nil || isNil(o.Indexed) {
		var ret bool
		return ret
	}
	return *o.Indexed
}

// GetIndexedOk returns a tuple with the Indexed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetIndexedOk() (*bool, bool) {
	if o == nil || isNil(o.Indexed) {
		return nil, false
	}
	return o.Indexed, true
}

// HasIndexed returns a boolean if a field has been set.
func (o *Property) HasIndexed() bool {
	if o != nil && !isNil(o.Indexed) {
		return true
	}

	return false
}

// SetIndexed gets a reference to the given bool and assigns it to the Indexed field.
func (o *Property) SetIndexed(v bool) {
	o.Indexed = &v
}

func (o Property) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Property) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.INDEXED_PROPERTY_PATTERN) {
		toSerialize["INDEXED_PROPERTY_PATTERN"] = o.INDEXED_PROPERTY_PATTERN
	}
	if !isNil(o.PropertyName) {
		toSerialize["propertyName"] = o.PropertyName
	}
	if !isNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !isNil(o.Indexed) {
		toSerialize["indexed"] = o.Indexed
	}
	return toSerialize, nil
}

type NullableProperty struct {
	value *Property
	isSet bool
}

func (v NullableProperty) Get() *Property {
	return v.value
}

func (v *NullableProperty) Set(val *Property) {
	v.value = val
	v.isSet = true
}

func (v NullableProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProperty(val *Property) *NullableProperty {
	return &NullableProperty{value: val, isSet: true}
}

func (v NullableProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

