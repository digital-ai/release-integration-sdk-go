/*
Digital.ai Release API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Delivery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Delivery{}

// Delivery struct for Delivery
type Delivery struct {
	Metadata                        *map[string]map[string]interface{} `json:"$metadata,omitempty"`
	Title                           *string                            `json:"title,omitempty"`
	Description                     *string                            `json:"description,omitempty"`
	Status                          *DeliveryStatus                    `json:"status,omitempty"`
	StartDate                       *time.Time                         `json:"startDate,omitempty"`
	EndDate                         *time.Time                         `json:"endDate,omitempty"`
	PlannedDuration                 *int32                             `json:"plannedDuration,omitempty"`
	ReleaseIds                      []string                           `json:"releaseIds,omitempty"`
	FolderId                        *string                            `json:"folderId,omitempty"`
	OriginPatternId                 *string                            `json:"originPatternId,omitempty"`
	Stages                          []Stage                            `json:"stages,omitempty"`
	TrackedItems                    []TrackedItem                      `json:"trackedItems,omitempty"`
	Subscribers                     []Subscriber                       `json:"subscribers,omitempty"`
	AutoComplete                    *bool                              `json:"autoComplete,omitempty"`
	Template                        *bool                              `json:"template,omitempty"`
	Transitions                     []Transition                       `json:"transitions,omitempty"`
	StagesBeforeFirstOpenTransition []Stage                            `json:"stagesBeforeFirstOpenTransition,omitempty"`
	Updatable                       *bool                              `json:"updatable,omitempty"`
}

// NewDelivery instantiates a new Delivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelivery() *Delivery {
	this := Delivery{}
	return &this
}

// NewDeliveryWithDefaults instantiates a new Delivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryWithDefaults() *Delivery {
	this := Delivery{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Delivery) GetMetadata() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetMetadataOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Delivery) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *Delivery) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Delivery) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Delivery) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Delivery) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Delivery) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Delivery) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Delivery) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Delivery) GetStatus() DeliveryStatus {
	if o == nil || IsNil(o.Status) {
		var ret DeliveryStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetStatusOk() (*DeliveryStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Delivery) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeliveryStatus and assigns it to the Status field.
func (o *Delivery) SetStatus(v DeliveryStatus) {
	o.Status = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Delivery) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Delivery) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Delivery) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Delivery) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Delivery) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *Delivery) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetPlannedDuration returns the PlannedDuration field value if set, zero value otherwise.
func (o *Delivery) GetPlannedDuration() int32 {
	if o == nil || IsNil(o.PlannedDuration) {
		var ret int32
		return ret
	}
	return *o.PlannedDuration
}

// GetPlannedDurationOk returns a tuple with the PlannedDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetPlannedDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.PlannedDuration) {
		return nil, false
	}
	return o.PlannedDuration, true
}

// HasPlannedDuration returns a boolean if a field has been set.
func (o *Delivery) HasPlannedDuration() bool {
	if o != nil && !IsNil(o.PlannedDuration) {
		return true
	}

	return false
}

// SetPlannedDuration gets a reference to the given int32 and assigns it to the PlannedDuration field.
func (o *Delivery) SetPlannedDuration(v int32) {
	o.PlannedDuration = &v
}

// GetReleaseIds returns the ReleaseIds field value if set, zero value otherwise.
func (o *Delivery) GetReleaseIds() []string {
	if o == nil || IsNil(o.ReleaseIds) {
		var ret []string
		return ret
	}
	return o.ReleaseIds
}

// GetReleaseIdsOk returns a tuple with the ReleaseIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetReleaseIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ReleaseIds) {
		return nil, false
	}
	return o.ReleaseIds, true
}

// HasReleaseIds returns a boolean if a field has been set.
func (o *Delivery) HasReleaseIds() bool {
	if o != nil && !IsNil(o.ReleaseIds) {
		return true
	}

	return false
}

// SetReleaseIds gets a reference to the given []string and assigns it to the ReleaseIds field.
func (o *Delivery) SetReleaseIds(v []string) {
	o.ReleaseIds = v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *Delivery) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *Delivery) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *Delivery) SetFolderId(v string) {
	o.FolderId = &v
}

// GetOriginPatternId returns the OriginPatternId field value if set, zero value otherwise.
func (o *Delivery) GetOriginPatternId() string {
	if o == nil || IsNil(o.OriginPatternId) {
		var ret string
		return ret
	}
	return *o.OriginPatternId
}

// GetOriginPatternIdOk returns a tuple with the OriginPatternId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetOriginPatternIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginPatternId) {
		return nil, false
	}
	return o.OriginPatternId, true
}

// HasOriginPatternId returns a boolean if a field has been set.
func (o *Delivery) HasOriginPatternId() bool {
	if o != nil && !IsNil(o.OriginPatternId) {
		return true
	}

	return false
}

// SetOriginPatternId gets a reference to the given string and assigns it to the OriginPatternId field.
func (o *Delivery) SetOriginPatternId(v string) {
	o.OriginPatternId = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *Delivery) GetStages() []Stage {
	if o == nil || IsNil(o.Stages) {
		var ret []Stage
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetStagesOk() ([]Stage, bool) {
	if o == nil || IsNil(o.Stages) {
		return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *Delivery) HasStages() bool {
	if o != nil && !IsNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []Stage and assigns it to the Stages field.
func (o *Delivery) SetStages(v []Stage) {
	o.Stages = v
}

// GetTrackedItems returns the TrackedItems field value if set, zero value otherwise.
func (o *Delivery) GetTrackedItems() []TrackedItem {
	if o == nil || IsNil(o.TrackedItems) {
		var ret []TrackedItem
		return ret
	}
	return o.TrackedItems
}

// GetTrackedItemsOk returns a tuple with the TrackedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetTrackedItemsOk() ([]TrackedItem, bool) {
	if o == nil || IsNil(o.TrackedItems) {
		return nil, false
	}
	return o.TrackedItems, true
}

// HasTrackedItems returns a boolean if a field has been set.
func (o *Delivery) HasTrackedItems() bool {
	if o != nil && !IsNil(o.TrackedItems) {
		return true
	}

	return false
}

// SetTrackedItems gets a reference to the given []TrackedItem and assigns it to the TrackedItems field.
func (o *Delivery) SetTrackedItems(v []TrackedItem) {
	o.TrackedItems = v
}

// GetSubscribers returns the Subscribers field value if set, zero value otherwise.
func (o *Delivery) GetSubscribers() []Subscriber {
	if o == nil || IsNil(o.Subscribers) {
		var ret []Subscriber
		return ret
	}
	return o.Subscribers
}

// GetSubscribersOk returns a tuple with the Subscribers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetSubscribersOk() ([]Subscriber, bool) {
	if o == nil || IsNil(o.Subscribers) {
		return nil, false
	}
	return o.Subscribers, true
}

// HasSubscribers returns a boolean if a field has been set.
func (o *Delivery) HasSubscribers() bool {
	if o != nil && !IsNil(o.Subscribers) {
		return true
	}

	return false
}

// SetSubscribers gets a reference to the given []Subscriber and assigns it to the Subscribers field.
func (o *Delivery) SetSubscribers(v []Subscriber) {
	o.Subscribers = v
}

// GetAutoComplete returns the AutoComplete field value if set, zero value otherwise.
func (o *Delivery) GetAutoComplete() bool {
	if o == nil || IsNil(o.AutoComplete) {
		var ret bool
		return ret
	}
	return *o.AutoComplete
}

// GetAutoCompleteOk returns a tuple with the AutoComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetAutoCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoComplete) {
		return nil, false
	}
	return o.AutoComplete, true
}

// HasAutoComplete returns a boolean if a field has been set.
func (o *Delivery) HasAutoComplete() bool {
	if o != nil && !IsNil(o.AutoComplete) {
		return true
	}

	return false
}

// SetAutoComplete gets a reference to the given bool and assigns it to the AutoComplete field.
func (o *Delivery) SetAutoComplete(v bool) {
	o.AutoComplete = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *Delivery) GetTemplate() bool {
	if o == nil || IsNil(o.Template) {
		var ret bool
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetTemplateOk() (*bool, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *Delivery) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given bool and assigns it to the Template field.
func (o *Delivery) SetTemplate(v bool) {
	o.Template = &v
}

// GetTransitions returns the Transitions field value if set, zero value otherwise.
func (o *Delivery) GetTransitions() []Transition {
	if o == nil || IsNil(o.Transitions) {
		var ret []Transition
		return ret
	}
	return o.Transitions
}

// GetTransitionsOk returns a tuple with the Transitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetTransitionsOk() ([]Transition, bool) {
	if o == nil || IsNil(o.Transitions) {
		return nil, false
	}
	return o.Transitions, true
}

// HasTransitions returns a boolean if a field has been set.
func (o *Delivery) HasTransitions() bool {
	if o != nil && !IsNil(o.Transitions) {
		return true
	}

	return false
}

// SetTransitions gets a reference to the given []Transition and assigns it to the Transitions field.
func (o *Delivery) SetTransitions(v []Transition) {
	o.Transitions = v
}

// GetStagesBeforeFirstOpenTransition returns the StagesBeforeFirstOpenTransition field value if set, zero value otherwise.
func (o *Delivery) GetStagesBeforeFirstOpenTransition() []Stage {
	if o == nil || IsNil(o.StagesBeforeFirstOpenTransition) {
		var ret []Stage
		return ret
	}
	return o.StagesBeforeFirstOpenTransition
}

// GetStagesBeforeFirstOpenTransitionOk returns a tuple with the StagesBeforeFirstOpenTransition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetStagesBeforeFirstOpenTransitionOk() ([]Stage, bool) {
	if o == nil || IsNil(o.StagesBeforeFirstOpenTransition) {
		return nil, false
	}
	return o.StagesBeforeFirstOpenTransition, true
}

// HasStagesBeforeFirstOpenTransition returns a boolean if a field has been set.
func (o *Delivery) HasStagesBeforeFirstOpenTransition() bool {
	if o != nil && !IsNil(o.StagesBeforeFirstOpenTransition) {
		return true
	}

	return false
}

// SetStagesBeforeFirstOpenTransition gets a reference to the given []Stage and assigns it to the StagesBeforeFirstOpenTransition field.
func (o *Delivery) SetStagesBeforeFirstOpenTransition(v []Stage) {
	o.StagesBeforeFirstOpenTransition = v
}

// GetUpdatable returns the Updatable field value if set, zero value otherwise.
func (o *Delivery) GetUpdatable() bool {
	if o == nil || IsNil(o.Updatable) {
		var ret bool
		return ret
	}
	return *o.Updatable
}

// GetUpdatableOk returns a tuple with the Updatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Delivery) GetUpdatableOk() (*bool, bool) {
	if o == nil || IsNil(o.Updatable) {
		return nil, false
	}
	return o.Updatable, true
}

// HasUpdatable returns a boolean if a field has been set.
func (o *Delivery) HasUpdatable() bool {
	if o != nil && !IsNil(o.Updatable) {
		return true
	}

	return false
}

// SetUpdatable gets a reference to the given bool and assigns it to the Updatable field.
func (o *Delivery) SetUpdatable(v bool) {
	o.Updatable = &v
}

func (o Delivery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Delivery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["$metadata"] = o.Metadata
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.PlannedDuration) {
		toSerialize["plannedDuration"] = o.PlannedDuration
	}
	if !IsNil(o.ReleaseIds) {
		toSerialize["releaseIds"] = o.ReleaseIds
	}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !IsNil(o.OriginPatternId) {
		toSerialize["originPatternId"] = o.OriginPatternId
	}
	if !IsNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	if !IsNil(o.TrackedItems) {
		toSerialize["trackedItems"] = o.TrackedItems
	}
	if !IsNil(o.Subscribers) {
		toSerialize["subscribers"] = o.Subscribers
	}
	if !IsNil(o.AutoComplete) {
		toSerialize["autoComplete"] = o.AutoComplete
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Transitions) {
		toSerialize["transitions"] = o.Transitions
	}
	if !IsNil(o.StagesBeforeFirstOpenTransition) {
		toSerialize["stagesBeforeFirstOpenTransition"] = o.StagesBeforeFirstOpenTransition
	}
	if !IsNil(o.Updatable) {
		toSerialize["updatable"] = o.Updatable
	}
	return toSerialize, nil
}

type NullableDelivery struct {
	value *Delivery
	isSet bool
}

func (v NullableDelivery) Get() *Delivery {
	return v.value
}

func (v *NullableDelivery) Set(val *Delivery) {
	v.value = val
	v.isSet = true
}

func (v NullableDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelivery(val *Delivery) *NullableDelivery {
	return &NullableDelivery{value: val, isSet: true}
}

func (v NullableDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
