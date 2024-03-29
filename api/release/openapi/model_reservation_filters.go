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

// checks if the ReservationFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationFilters{}

// ReservationFilters struct for ReservationFilters
type ReservationFilters struct {
	EnvironmentTitle *string    `json:"environmentTitle,omitempty"`
	Stages           []string   `json:"stages,omitempty"`
	Labels           []string   `json:"labels,omitempty"`
	Applications     []string   `json:"applications,omitempty"`
	From             *time.Time `json:"from,omitempty"`
	To               *time.Time `json:"to,omitempty"`
}

// NewReservationFilters instantiates a new ReservationFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationFilters() *ReservationFilters {
	this := ReservationFilters{}
	return &this
}

// NewReservationFiltersWithDefaults instantiates a new ReservationFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationFiltersWithDefaults() *ReservationFilters {
	this := ReservationFilters{}
	return &this
}

// GetEnvironmentTitle returns the EnvironmentTitle field value if set, zero value otherwise.
func (o *ReservationFilters) GetEnvironmentTitle() string {
	if o == nil || IsNil(o.EnvironmentTitle) {
		var ret string
		return ret
	}
	return *o.EnvironmentTitle
}

// GetEnvironmentTitleOk returns a tuple with the EnvironmentTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFilters) GetEnvironmentTitleOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentTitle) {
		return nil, false
	}
	return o.EnvironmentTitle, true
}

// HasEnvironmentTitle returns a boolean if a field has been set.
func (o *ReservationFilters) HasEnvironmentTitle() bool {
	if o != nil && !IsNil(o.EnvironmentTitle) {
		return true
	}

	return false
}

// SetEnvironmentTitle gets a reference to the given string and assigns it to the EnvironmentTitle field.
func (o *ReservationFilters) SetEnvironmentTitle(v string) {
	o.EnvironmentTitle = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *ReservationFilters) GetStages() []string {
	if o == nil || IsNil(o.Stages) {
		var ret []string
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFilters) GetStagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Stages) {
		return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *ReservationFilters) HasStages() bool {
	if o != nil && !IsNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []string and assigns it to the Stages field.
func (o *ReservationFilters) SetStages(v []string) {
	o.Stages = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ReservationFilters) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFilters) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ReservationFilters) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ReservationFilters) SetLabels(v []string) {
	o.Labels = v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *ReservationFilters) GetApplications() []string {
	if o == nil || IsNil(o.Applications) {
		var ret []string
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFilters) GetApplicationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *ReservationFilters) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []string and assigns it to the Applications field.
func (o *ReservationFilters) SetApplications(v []string) {
	o.Applications = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ReservationFilters) GetFrom() time.Time {
	if o == nil || IsNil(o.From) {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFilters) GetFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ReservationFilters) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *ReservationFilters) SetFrom(v time.Time) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ReservationFilters) GetTo() time.Time {
	if o == nil || IsNil(o.To) {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationFilters) GetToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ReservationFilters) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *ReservationFilters) SetTo(v time.Time) {
	o.To = &v
}

func (o ReservationFilters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentTitle) {
		toSerialize["environmentTitle"] = o.EnvironmentTitle
	}
	if !IsNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableReservationFilters struct {
	value *ReservationFilters
	isSet bool
}

func (v NullableReservationFilters) Get() *ReservationFilters {
	return v.value
}

func (v *NullableReservationFilters) Set(val *ReservationFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationFilters(val *ReservationFilters) *NullableReservationFilters {
	return &NullableReservationFilters{value: val, isSet: true}
}

func (v NullableReservationFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
