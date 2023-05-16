# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Satisfied** | Pointer to **bool** |  | [optional] 
**SatisfiedDate** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**InputProperties** | Pointer to **[]interface{}** |  | [optional] 
**Leaf** | Pointer to **bool** |  | [optional] 
**AllConditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 
**LeafConditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 

## Methods

### NewCondition

`func NewCondition() *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Condition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Condition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Condition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Condition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Condition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Condition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Condition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Condition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSatisfied

`func (o *Condition) GetSatisfied() bool`

GetSatisfied returns the Satisfied field if non-nil, zero value otherwise.

### GetSatisfiedOk

`func (o *Condition) GetSatisfiedOk() (*bool, bool)`

GetSatisfiedOk returns a tuple with the Satisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfied

`func (o *Condition) SetSatisfied(v bool)`

SetSatisfied sets Satisfied field to given value.

### HasSatisfied

`func (o *Condition) HasSatisfied() bool`

HasSatisfied returns a boolean if a field has been set.

### GetSatisfiedDate

`func (o *Condition) GetSatisfiedDate() time.Time`

GetSatisfiedDate returns the SatisfiedDate field if non-nil, zero value otherwise.

### GetSatisfiedDateOk

`func (o *Condition) GetSatisfiedDateOk() (*time.Time, bool)`

GetSatisfiedDateOk returns a tuple with the SatisfiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfiedDate

`func (o *Condition) SetSatisfiedDate(v time.Time)`

SetSatisfiedDate sets SatisfiedDate field to given value.

### HasSatisfiedDate

`func (o *Condition) HasSatisfiedDate() bool`

HasSatisfiedDate returns a boolean if a field has been set.

### GetDescription

`func (o *Condition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Condition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Condition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Condition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *Condition) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Condition) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Condition) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Condition) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetInputProperties

`func (o *Condition) GetInputProperties() []interface{}`

GetInputProperties returns the InputProperties field if non-nil, zero value otherwise.

### GetInputPropertiesOk

`func (o *Condition) GetInputPropertiesOk() (*[]interface{}, bool)`

GetInputPropertiesOk returns a tuple with the InputProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputProperties

`func (o *Condition) SetInputProperties(v []interface{})`

SetInputProperties sets InputProperties field to given value.

### HasInputProperties

`func (o *Condition) HasInputProperties() bool`

HasInputProperties returns a boolean if a field has been set.

### GetLeaf

`func (o *Condition) GetLeaf() bool`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *Condition) GetLeafOk() (*bool, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *Condition) SetLeaf(v bool)`

SetLeaf sets Leaf field to given value.

### HasLeaf

`func (o *Condition) HasLeaf() bool`

HasLeaf returns a boolean if a field has been set.

### GetAllConditions

`func (o *Condition) GetAllConditions() []Condition`

GetAllConditions returns the AllConditions field if non-nil, zero value otherwise.

### GetAllConditionsOk

`func (o *Condition) GetAllConditionsOk() (*[]Condition, bool)`

GetAllConditionsOk returns a tuple with the AllConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllConditions

`func (o *Condition) SetAllConditions(v []Condition)`

SetAllConditions sets AllConditions field to given value.

### HasAllConditions

`func (o *Condition) HasAllConditions() bool`

HasAllConditions returns a boolean if a field has been set.

### GetLeafConditions

`func (o *Condition) GetLeafConditions() []Condition`

GetLeafConditions returns the LeafConditions field if non-nil, zero value otherwise.

### GetLeafConditionsOk

`func (o *Condition) GetLeafConditionsOk() (*[]Condition, bool)`

GetLeafConditionsOk returns a tuple with the LeafConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafConditions

`func (o *Condition) SetLeafConditions(v []Condition)`

SetLeafConditions sets LeafConditions field to given value.

### HasLeafConditions

`func (o *Condition) HasLeafConditions() bool`

HasLeafConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


