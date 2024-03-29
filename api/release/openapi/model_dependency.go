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

// checks if the Dependency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dependency{}

// Dependency struct for Dependency
type Dependency struct {
	Id                     *string                            `json:"id,omitempty"`
	Type                   *string                            `json:"type,omitempty"`
	GateTask               *GateTask                          `json:"gateTask,omitempty"`
	Target                 *PlanItem                          `json:"target,omitempty"`
	TargetId               *string                            `json:"targetId,omitempty"`
	ArchivedTargetTitle    *string                            `json:"archivedTargetTitle,omitempty"`
	ArchivedTargetId       *string                            `json:"archivedTargetId,omitempty"`
	ArchivedAsResolved     *bool                              `json:"archivedAsResolved,omitempty"`
	Metadata               *map[string]map[string]interface{} `json:"$metadata,omitempty"`
	Archived               *bool                              `json:"archived,omitempty"`
	Done                   *bool                              `json:"done,omitempty"`
	Aborted                *bool                              `json:"aborted,omitempty"`
	TargetDisplayPath      *string                            `json:"targetDisplayPath,omitempty"`
	TargetTitle            *string                            `json:"targetTitle,omitempty"`
	ValidAllowedPlanItemId *bool                              `json:"validAllowedPlanItemId,omitempty"`
}

// NewDependency instantiates a new Dependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependency() *Dependency {
	this := Dependency{}
	return &this
}

// NewDependencyWithDefaults instantiates a new Dependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependencyWithDefaults() *Dependency {
	this := Dependency{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dependency) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dependency) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Dependency) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Dependency) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Dependency) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Dependency) SetType(v string) {
	o.Type = &v
}

// GetGateTask returns the GateTask field value if set, zero value otherwise.
func (o *Dependency) GetGateTask() GateTask {
	if o == nil || IsNil(o.GateTask) {
		var ret GateTask
		return ret
	}
	return *o.GateTask
}

// GetGateTaskOk returns a tuple with the GateTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetGateTaskOk() (*GateTask, bool) {
	if o == nil || IsNil(o.GateTask) {
		return nil, false
	}
	return o.GateTask, true
}

// HasGateTask returns a boolean if a field has been set.
func (o *Dependency) HasGateTask() bool {
	if o != nil && !IsNil(o.GateTask) {
		return true
	}

	return false
}

// SetGateTask gets a reference to the given GateTask and assigns it to the GateTask field.
func (o *Dependency) SetGateTask(v GateTask) {
	o.GateTask = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Dependency) GetTarget() PlanItem {
	if o == nil || IsNil(o.Target) {
		var ret PlanItem
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetTargetOk() (*PlanItem, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Dependency) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given PlanItem and assigns it to the Target field.
func (o *Dependency) SetTarget(v PlanItem) {
	o.Target = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *Dependency) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *Dependency) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *Dependency) SetTargetId(v string) {
	o.TargetId = &v
}

// GetArchivedTargetTitle returns the ArchivedTargetTitle field value if set, zero value otherwise.
func (o *Dependency) GetArchivedTargetTitle() string {
	if o == nil || IsNil(o.ArchivedTargetTitle) {
		var ret string
		return ret
	}
	return *o.ArchivedTargetTitle
}

// GetArchivedTargetTitleOk returns a tuple with the ArchivedTargetTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetArchivedTargetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedTargetTitle) {
		return nil, false
	}
	return o.ArchivedTargetTitle, true
}

// HasArchivedTargetTitle returns a boolean if a field has been set.
func (o *Dependency) HasArchivedTargetTitle() bool {
	if o != nil && !IsNil(o.ArchivedTargetTitle) {
		return true
	}

	return false
}

// SetArchivedTargetTitle gets a reference to the given string and assigns it to the ArchivedTargetTitle field.
func (o *Dependency) SetArchivedTargetTitle(v string) {
	o.ArchivedTargetTitle = &v
}

// GetArchivedTargetId returns the ArchivedTargetId field value if set, zero value otherwise.
func (o *Dependency) GetArchivedTargetId() string {
	if o == nil || IsNil(o.ArchivedTargetId) {
		var ret string
		return ret
	}
	return *o.ArchivedTargetId
}

// GetArchivedTargetIdOk returns a tuple with the ArchivedTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetArchivedTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedTargetId) {
		return nil, false
	}
	return o.ArchivedTargetId, true
}

// HasArchivedTargetId returns a boolean if a field has been set.
func (o *Dependency) HasArchivedTargetId() bool {
	if o != nil && !IsNil(o.ArchivedTargetId) {
		return true
	}

	return false
}

// SetArchivedTargetId gets a reference to the given string and assigns it to the ArchivedTargetId field.
func (o *Dependency) SetArchivedTargetId(v string) {
	o.ArchivedTargetId = &v
}

// GetArchivedAsResolved returns the ArchivedAsResolved field value if set, zero value otherwise.
func (o *Dependency) GetArchivedAsResolved() bool {
	if o == nil || IsNil(o.ArchivedAsResolved) {
		var ret bool
		return ret
	}
	return *o.ArchivedAsResolved
}

// GetArchivedAsResolvedOk returns a tuple with the ArchivedAsResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetArchivedAsResolvedOk() (*bool, bool) {
	if o == nil || IsNil(o.ArchivedAsResolved) {
		return nil, false
	}
	return o.ArchivedAsResolved, true
}

// HasArchivedAsResolved returns a boolean if a field has been set.
func (o *Dependency) HasArchivedAsResolved() bool {
	if o != nil && !IsNil(o.ArchivedAsResolved) {
		return true
	}

	return false
}

// SetArchivedAsResolved gets a reference to the given bool and assigns it to the ArchivedAsResolved field.
func (o *Dependency) SetArchivedAsResolved(v bool) {
	o.ArchivedAsResolved = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Dependency) GetMetadata() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetMetadataOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Dependency) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *Dependency) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *Dependency) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *Dependency) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *Dependency) SetArchived(v bool) {
	o.Archived = &v
}

// GetDone returns the Done field value if set, zero value otherwise.
func (o *Dependency) GetDone() bool {
	if o == nil || IsNil(o.Done) {
		var ret bool
		return ret
	}
	return *o.Done
}

// GetDoneOk returns a tuple with the Done field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetDoneOk() (*bool, bool) {
	if o == nil || IsNil(o.Done) {
		return nil, false
	}
	return o.Done, true
}

// HasDone returns a boolean if a field has been set.
func (o *Dependency) HasDone() bool {
	if o != nil && !IsNil(o.Done) {
		return true
	}

	return false
}

// SetDone gets a reference to the given bool and assigns it to the Done field.
func (o *Dependency) SetDone(v bool) {
	o.Done = &v
}

// GetAborted returns the Aborted field value if set, zero value otherwise.
func (o *Dependency) GetAborted() bool {
	if o == nil || IsNil(o.Aborted) {
		var ret bool
		return ret
	}
	return *o.Aborted
}

// GetAbortedOk returns a tuple with the Aborted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetAbortedOk() (*bool, bool) {
	if o == nil || IsNil(o.Aborted) {
		return nil, false
	}
	return o.Aborted, true
}

// HasAborted returns a boolean if a field has been set.
func (o *Dependency) HasAborted() bool {
	if o != nil && !IsNil(o.Aborted) {
		return true
	}

	return false
}

// SetAborted gets a reference to the given bool and assigns it to the Aborted field.
func (o *Dependency) SetAborted(v bool) {
	o.Aborted = &v
}

// GetTargetDisplayPath returns the TargetDisplayPath field value if set, zero value otherwise.
func (o *Dependency) GetTargetDisplayPath() string {
	if o == nil || IsNil(o.TargetDisplayPath) {
		var ret string
		return ret
	}
	return *o.TargetDisplayPath
}

// GetTargetDisplayPathOk returns a tuple with the TargetDisplayPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetTargetDisplayPathOk() (*string, bool) {
	if o == nil || IsNil(o.TargetDisplayPath) {
		return nil, false
	}
	return o.TargetDisplayPath, true
}

// HasTargetDisplayPath returns a boolean if a field has been set.
func (o *Dependency) HasTargetDisplayPath() bool {
	if o != nil && !IsNil(o.TargetDisplayPath) {
		return true
	}

	return false
}

// SetTargetDisplayPath gets a reference to the given string and assigns it to the TargetDisplayPath field.
func (o *Dependency) SetTargetDisplayPath(v string) {
	o.TargetDisplayPath = &v
}

// GetTargetTitle returns the TargetTitle field value if set, zero value otherwise.
func (o *Dependency) GetTargetTitle() string {
	if o == nil || IsNil(o.TargetTitle) {
		var ret string
		return ret
	}
	return *o.TargetTitle
}

// GetTargetTitleOk returns a tuple with the TargetTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetTargetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.TargetTitle) {
		return nil, false
	}
	return o.TargetTitle, true
}

// HasTargetTitle returns a boolean if a field has been set.
func (o *Dependency) HasTargetTitle() bool {
	if o != nil && !IsNil(o.TargetTitle) {
		return true
	}

	return false
}

// SetTargetTitle gets a reference to the given string and assigns it to the TargetTitle field.
func (o *Dependency) SetTargetTitle(v string) {
	o.TargetTitle = &v
}

// GetValidAllowedPlanItemId returns the ValidAllowedPlanItemId field value if set, zero value otherwise.
func (o *Dependency) GetValidAllowedPlanItemId() bool {
	if o == nil || IsNil(o.ValidAllowedPlanItemId) {
		var ret bool
		return ret
	}
	return *o.ValidAllowedPlanItemId
}

// GetValidAllowedPlanItemIdOk returns a tuple with the ValidAllowedPlanItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetValidAllowedPlanItemIdOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidAllowedPlanItemId) {
		return nil, false
	}
	return o.ValidAllowedPlanItemId, true
}

// HasValidAllowedPlanItemId returns a boolean if a field has been set.
func (o *Dependency) HasValidAllowedPlanItemId() bool {
	if o != nil && !IsNil(o.ValidAllowedPlanItemId) {
		return true
	}

	return false
}

// SetValidAllowedPlanItemId gets a reference to the given bool and assigns it to the ValidAllowedPlanItemId field.
func (o *Dependency) SetValidAllowedPlanItemId(v bool) {
	o.ValidAllowedPlanItemId = &v
}

func (o Dependency) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dependency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.GateTask) {
		toSerialize["gateTask"] = o.GateTask
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !IsNil(o.ArchivedTargetTitle) {
		toSerialize["archivedTargetTitle"] = o.ArchivedTargetTitle
	}
	if !IsNil(o.ArchivedTargetId) {
		toSerialize["archivedTargetId"] = o.ArchivedTargetId
	}
	if !IsNil(o.ArchivedAsResolved) {
		toSerialize["archivedAsResolved"] = o.ArchivedAsResolved
	}
	if !IsNil(o.Metadata) {
		toSerialize["$metadata"] = o.Metadata
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.Done) {
		toSerialize["done"] = o.Done
	}
	if !IsNil(o.Aborted) {
		toSerialize["aborted"] = o.Aborted
	}
	if !IsNil(o.TargetDisplayPath) {
		toSerialize["targetDisplayPath"] = o.TargetDisplayPath
	}
	if !IsNil(o.TargetTitle) {
		toSerialize["targetTitle"] = o.TargetTitle
	}
	if !IsNil(o.ValidAllowedPlanItemId) {
		toSerialize["validAllowedPlanItemId"] = o.ValidAllowedPlanItemId
	}
	return toSerialize, nil
}

type NullableDependency struct {
	value *Dependency
	isSet bool
}

func (v NullableDependency) Get() *Dependency {
	return v.value
}

func (v *NullableDependency) Set(val *Dependency) {
	v.value = val
	v.isSet = true
}

func (v NullableDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependency(val *Dependency) *NullableDependency {
	return &NullableDependency{value: val, isSet: true}
}

func (v NullableDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
