# Transition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to [**Stage**](Stage.md) |  | [optional] 
**Conditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 
**Automated** | Pointer to **bool** |  | [optional] 
**AllConditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 
**LeafConditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 
**RootCondition** | Pointer to [**Condition**](Condition.md) |  | [optional] 

## Methods

### NewTransition

`func NewTransition() *Transition`

NewTransition instantiates a new Transition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionWithDefaults

`func NewTransitionWithDefaults() *Transition`

NewTransitionWithDefaults instantiates a new Transition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Transition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Transition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Transition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Transition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStage

`func (o *Transition) GetStage() Stage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Transition) GetStageOk() (*Stage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Transition) SetStage(v Stage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *Transition) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetConditions

`func (o *Transition) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Transition) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Transition) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Transition) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetAutomated

`func (o *Transition) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *Transition) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *Transition) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *Transition) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetAllConditions

`func (o *Transition) GetAllConditions() []Condition`

GetAllConditions returns the AllConditions field if non-nil, zero value otherwise.

### GetAllConditionsOk

`func (o *Transition) GetAllConditionsOk() (*[]Condition, bool)`

GetAllConditionsOk returns a tuple with the AllConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllConditions

`func (o *Transition) SetAllConditions(v []Condition)`

SetAllConditions sets AllConditions field to given value.

### HasAllConditions

`func (o *Transition) HasAllConditions() bool`

HasAllConditions returns a boolean if a field has been set.

### GetLeafConditions

`func (o *Transition) GetLeafConditions() []Condition`

GetLeafConditions returns the LeafConditions field if non-nil, zero value otherwise.

### GetLeafConditionsOk

`func (o *Transition) GetLeafConditionsOk() (*[]Condition, bool)`

GetLeafConditionsOk returns a tuple with the LeafConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafConditions

`func (o *Transition) SetLeafConditions(v []Condition)`

SetLeafConditions sets LeafConditions field to given value.

### HasLeafConditions

`func (o *Transition) HasLeafConditions() bool`

HasLeafConditions returns a boolean if a field has been set.

### GetRootCondition

`func (o *Transition) GetRootCondition() Condition`

GetRootCondition returns the RootCondition field if non-nil, zero value otherwise.

### GetRootConditionOk

`func (o *Transition) GetRootConditionOk() (*Condition, bool)`

GetRootConditionOk returns a tuple with the RootCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCondition

`func (o *Transition) SetRootCondition(v Condition)`

SetRootCondition sets RootCondition field to given value.

### HasRootCondition

`func (o *Transition) HasRootCondition() bool`

HasRootCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


