/*
Digital.ai Release API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Variable struct for Variable
type Variable struct {
	FolderId *string `json:"folderId,omitempty"`
	Title *string `json:"title,omitempty"`
	Key *string `json:"key,omitempty"`
	RequiresValue *bool `json:"requiresValue,omitempty"`
	ShowOnReleaseStart *bool `json:"showOnReleaseStart,omitempty"`
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	ValueProvider *ValueProviderConfiguration `json:"valueProvider,omitempty"`
	Inherited *bool `json:"inherited,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
	EmptyValue map[string]interface{} `json:"emptyValue,omitempty"`
	ValueEmpty *bool `json:"valueEmpty,omitempty"`
	UntypedValue map[string]interface{} `json:"untypedValue,omitempty"`
	Password *bool `json:"password,omitempty"`
	ValueAsString *string `json:"valueAsString,omitempty"`
	EmptyValueAsString *string `json:"emptyValueAsString,omitempty"`
}

// NewVariable instantiates a new Variable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariable() *Variable {
	this := Variable{}
	return &this
}

// NewVariableWithDefaults instantiates a new Variable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableWithDefaults() *Variable {
	this := Variable{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *Variable) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *Variable) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *Variable) SetFolderId(v string) {
	o.FolderId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Variable) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Variable) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Variable) SetTitle(v string) {
	o.Title = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Variable) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Variable) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Variable) SetKey(v string) {
	o.Key = &v
}

// GetRequiresValue returns the RequiresValue field value if set, zero value otherwise.
func (o *Variable) GetRequiresValue() bool {
	if o == nil || isNil(o.RequiresValue) {
		var ret bool
		return ret
	}
	return *o.RequiresValue
}

// GetRequiresValueOk returns a tuple with the RequiresValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetRequiresValueOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresValue) {
    return nil, false
	}
	return o.RequiresValue, true
}

// HasRequiresValue returns a boolean if a field has been set.
func (o *Variable) HasRequiresValue() bool {
	if o != nil && !isNil(o.RequiresValue) {
		return true
	}

	return false
}

// SetRequiresValue gets a reference to the given bool and assigns it to the RequiresValue field.
func (o *Variable) SetRequiresValue(v bool) {
	o.RequiresValue = &v
}

// GetShowOnReleaseStart returns the ShowOnReleaseStart field value if set, zero value otherwise.
func (o *Variable) GetShowOnReleaseStart() bool {
	if o == nil || isNil(o.ShowOnReleaseStart) {
		var ret bool
		return ret
	}
	return *o.ShowOnReleaseStart
}

// GetShowOnReleaseStartOk returns a tuple with the ShowOnReleaseStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetShowOnReleaseStartOk() (*bool, bool) {
	if o == nil || isNil(o.ShowOnReleaseStart) {
    return nil, false
	}
	return o.ShowOnReleaseStart, true
}

// HasShowOnReleaseStart returns a boolean if a field has been set.
func (o *Variable) HasShowOnReleaseStart() bool {
	if o != nil && !isNil(o.ShowOnReleaseStart) {
		return true
	}

	return false
}

// SetShowOnReleaseStart gets a reference to the given bool and assigns it to the ShowOnReleaseStart field.
func (o *Variable) SetShowOnReleaseStart(v bool) {
	o.ShowOnReleaseStart = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Variable) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Variable) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Variable) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Variable) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Variable) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Variable) SetDescription(v string) {
	o.Description = &v
}

// GetValueProvider returns the ValueProvider field value if set, zero value otherwise.
func (o *Variable) GetValueProvider() ValueProviderConfiguration {
	if o == nil || isNil(o.ValueProvider) {
		var ret ValueProviderConfiguration
		return ret
	}
	return *o.ValueProvider
}

// GetValueProviderOk returns a tuple with the ValueProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetValueProviderOk() (*ValueProviderConfiguration, bool) {
	if o == nil || isNil(o.ValueProvider) {
    return nil, false
	}
	return o.ValueProvider, true
}

// HasValueProvider returns a boolean if a field has been set.
func (o *Variable) HasValueProvider() bool {
	if o != nil && !isNil(o.ValueProvider) {
		return true
	}

	return false
}

// SetValueProvider gets a reference to the given ValueProviderConfiguration and assigns it to the ValueProvider field.
func (o *Variable) SetValueProvider(v ValueProviderConfiguration) {
	o.ValueProvider = &v
}

// GetInherited returns the Inherited field value if set, zero value otherwise.
func (o *Variable) GetInherited() bool {
	if o == nil || isNil(o.Inherited) {
		var ret bool
		return ret
	}
	return *o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetInheritedOk() (*bool, bool) {
	if o == nil || isNil(o.Inherited) {
    return nil, false
	}
	return o.Inherited, true
}

// HasInherited returns a boolean if a field has been set.
func (o *Variable) HasInherited() bool {
	if o != nil && !isNil(o.Inherited) {
		return true
	}

	return false
}

// SetInherited gets a reference to the given bool and assigns it to the Inherited field.
func (o *Variable) SetInherited(v bool) {
	o.Inherited = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Variable) GetValue() map[string]interface{} {
	if o == nil || isNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Value) {
    return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Variable) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *Variable) SetValue(v map[string]interface{}) {
	o.Value = v
}

// GetEmptyValue returns the EmptyValue field value if set, zero value otherwise.
func (o *Variable) GetEmptyValue() map[string]interface{} {
	if o == nil || isNil(o.EmptyValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.EmptyValue
}

// GetEmptyValueOk returns a tuple with the EmptyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetEmptyValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.EmptyValue) {
    return map[string]interface{}{}, false
	}
	return o.EmptyValue, true
}

// HasEmptyValue returns a boolean if a field has been set.
func (o *Variable) HasEmptyValue() bool {
	if o != nil && !isNil(o.EmptyValue) {
		return true
	}

	return false
}

// SetEmptyValue gets a reference to the given map[string]interface{} and assigns it to the EmptyValue field.
func (o *Variable) SetEmptyValue(v map[string]interface{}) {
	o.EmptyValue = v
}

// GetValueEmpty returns the ValueEmpty field value if set, zero value otherwise.
func (o *Variable) GetValueEmpty() bool {
	if o == nil || isNil(o.ValueEmpty) {
		var ret bool
		return ret
	}
	return *o.ValueEmpty
}

// GetValueEmptyOk returns a tuple with the ValueEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetValueEmptyOk() (*bool, bool) {
	if o == nil || isNil(o.ValueEmpty) {
    return nil, false
	}
	return o.ValueEmpty, true
}

// HasValueEmpty returns a boolean if a field has been set.
func (o *Variable) HasValueEmpty() bool {
	if o != nil && !isNil(o.ValueEmpty) {
		return true
	}

	return false
}

// SetValueEmpty gets a reference to the given bool and assigns it to the ValueEmpty field.
func (o *Variable) SetValueEmpty(v bool) {
	o.ValueEmpty = &v
}

// GetUntypedValue returns the UntypedValue field value if set, zero value otherwise.
func (o *Variable) GetUntypedValue() map[string]interface{} {
	if o == nil || isNil(o.UntypedValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.UntypedValue
}

// GetUntypedValueOk returns a tuple with the UntypedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetUntypedValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.UntypedValue) {
    return map[string]interface{}{}, false
	}
	return o.UntypedValue, true
}

// HasUntypedValue returns a boolean if a field has been set.
func (o *Variable) HasUntypedValue() bool {
	if o != nil && !isNil(o.UntypedValue) {
		return true
	}

	return false
}

// SetUntypedValue gets a reference to the given map[string]interface{} and assigns it to the UntypedValue field.
func (o *Variable) SetUntypedValue(v map[string]interface{}) {
	o.UntypedValue = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Variable) GetPassword() bool {
	if o == nil || isNil(o.Password) {
		var ret bool
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Variable) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given bool and assigns it to the Password field.
func (o *Variable) SetPassword(v bool) {
	o.Password = &v
}

// GetValueAsString returns the ValueAsString field value if set, zero value otherwise.
func (o *Variable) GetValueAsString() string {
	if o == nil || isNil(o.ValueAsString) {
		var ret string
		return ret
	}
	return *o.ValueAsString
}

// GetValueAsStringOk returns a tuple with the ValueAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetValueAsStringOk() (*string, bool) {
	if o == nil || isNil(o.ValueAsString) {
    return nil, false
	}
	return o.ValueAsString, true
}

// HasValueAsString returns a boolean if a field has been set.
func (o *Variable) HasValueAsString() bool {
	if o != nil && !isNil(o.ValueAsString) {
		return true
	}

	return false
}

// SetValueAsString gets a reference to the given string and assigns it to the ValueAsString field.
func (o *Variable) SetValueAsString(v string) {
	o.ValueAsString = &v
}

// GetEmptyValueAsString returns the EmptyValueAsString field value if set, zero value otherwise.
func (o *Variable) GetEmptyValueAsString() string {
	if o == nil || isNil(o.EmptyValueAsString) {
		var ret string
		return ret
	}
	return *o.EmptyValueAsString
}

// GetEmptyValueAsStringOk returns a tuple with the EmptyValueAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetEmptyValueAsStringOk() (*string, bool) {
	if o == nil || isNil(o.EmptyValueAsString) {
    return nil, false
	}
	return o.EmptyValueAsString, true
}

// HasEmptyValueAsString returns a boolean if a field has been set.
func (o *Variable) HasEmptyValueAsString() bool {
	if o != nil && !isNil(o.EmptyValueAsString) {
		return true
	}

	return false
}

// SetEmptyValueAsString gets a reference to the given string and assigns it to the EmptyValueAsString field.
func (o *Variable) SetEmptyValueAsString(v string) {
	o.EmptyValueAsString = &v
}

func (o Variable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.RequiresValue) {
		toSerialize["requiresValue"] = o.RequiresValue
	}
	if !isNil(o.ShowOnReleaseStart) {
		toSerialize["showOnReleaseStart"] = o.ShowOnReleaseStart
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ValueProvider) {
		toSerialize["valueProvider"] = o.ValueProvider
	}
	if !isNil(o.Inherited) {
		toSerialize["inherited"] = o.Inherited
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.EmptyValue) {
		toSerialize["emptyValue"] = o.EmptyValue
	}
	if !isNil(o.ValueEmpty) {
		toSerialize["valueEmpty"] = o.ValueEmpty
	}
	if !isNil(o.UntypedValue) {
		toSerialize["untypedValue"] = o.UntypedValue
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.ValueAsString) {
		toSerialize["valueAsString"] = o.ValueAsString
	}
	if !isNil(o.EmptyValueAsString) {
		toSerialize["emptyValueAsString"] = o.EmptyValueAsString
	}
	return json.Marshal(toSerialize)
}

type NullableVariable struct {
	value *Variable
	isSet bool
}

func (v NullableVariable) Get() *Variable {
	return v.value
}

func (v *NullableVariable) Set(val *Variable) {
	v.value = val
	v.isSet = true
}

func (v NullableVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariable(val *Variable) *NullableVariable {
	return &NullableVariable{value: val, isSet: true}
}

func (v NullableVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

