# Delivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**DeliveryStatus**](DeliveryStatus.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**PlannedDuration** | Pointer to **int32** |  | [optional] 
**ReleaseIds** | Pointer to **[]string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**OriginPatternId** | Pointer to **string** |  | [optional] 
**Stages** | Pointer to [**[]Stage**](Stage.md) |  | [optional] 
**TrackedItems** | Pointer to [**[]TrackedItem**](TrackedItem.md) |  | [optional] 
**Subscribers** | Pointer to [**[]Subscriber**](Subscriber.md) |  | [optional] 
**AutoComplete** | Pointer to **bool** |  | [optional] 
**Template** | Pointer to **bool** |  | [optional] 
**Transitions** | Pointer to [**[]Transition**](Transition.md) |  | [optional] 
**StagesBeforeFirstOpenTransition** | Pointer to [**[]Stage**](Stage.md) |  | [optional] 
**Updatable** | Pointer to **bool** |  | [optional] 

## Methods

### NewDelivery

`func NewDelivery() *Delivery`

NewDelivery instantiates a new Delivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryWithDefaults

`func NewDeliveryWithDefaults() *Delivery`

NewDeliveryWithDefaults instantiates a new Delivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Delivery) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Delivery) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Delivery) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Delivery) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTitle

`func (o *Delivery) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Delivery) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Delivery) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Delivery) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Delivery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Delivery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Delivery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Delivery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *Delivery) GetStatus() DeliveryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Delivery) GetStatusOk() (*DeliveryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Delivery) SetStatus(v DeliveryStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Delivery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *Delivery) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Delivery) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Delivery) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Delivery) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Delivery) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Delivery) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Delivery) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Delivery) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPlannedDuration

`func (o *Delivery) GetPlannedDuration() int32`

GetPlannedDuration returns the PlannedDuration field if non-nil, zero value otherwise.

### GetPlannedDurationOk

`func (o *Delivery) GetPlannedDurationOk() (*int32, bool)`

GetPlannedDurationOk returns a tuple with the PlannedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDuration

`func (o *Delivery) SetPlannedDuration(v int32)`

SetPlannedDuration sets PlannedDuration field to given value.

### HasPlannedDuration

`func (o *Delivery) HasPlannedDuration() bool`

HasPlannedDuration returns a boolean if a field has been set.

### GetReleaseIds

`func (o *Delivery) GetReleaseIds() []string`

GetReleaseIds returns the ReleaseIds field if non-nil, zero value otherwise.

### GetReleaseIdsOk

`func (o *Delivery) GetReleaseIdsOk() (*[]string, bool)`

GetReleaseIdsOk returns a tuple with the ReleaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseIds

`func (o *Delivery) SetReleaseIds(v []string)`

SetReleaseIds sets ReleaseIds field to given value.

### HasReleaseIds

`func (o *Delivery) HasReleaseIds() bool`

HasReleaseIds returns a boolean if a field has been set.

### GetFolderId

`func (o *Delivery) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Delivery) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Delivery) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Delivery) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetOriginPatternId

`func (o *Delivery) GetOriginPatternId() string`

GetOriginPatternId returns the OriginPatternId field if non-nil, zero value otherwise.

### GetOriginPatternIdOk

`func (o *Delivery) GetOriginPatternIdOk() (*string, bool)`

GetOriginPatternIdOk returns a tuple with the OriginPatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPatternId

`func (o *Delivery) SetOriginPatternId(v string)`

SetOriginPatternId sets OriginPatternId field to given value.

### HasOriginPatternId

`func (o *Delivery) HasOriginPatternId() bool`

HasOriginPatternId returns a boolean if a field has been set.

### GetStages

`func (o *Delivery) GetStages() []Stage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *Delivery) GetStagesOk() (*[]Stage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *Delivery) SetStages(v []Stage)`

SetStages sets Stages field to given value.

### HasStages

`func (o *Delivery) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetTrackedItems

`func (o *Delivery) GetTrackedItems() []TrackedItem`

GetTrackedItems returns the TrackedItems field if non-nil, zero value otherwise.

### GetTrackedItemsOk

`func (o *Delivery) GetTrackedItemsOk() (*[]TrackedItem, bool)`

GetTrackedItemsOk returns a tuple with the TrackedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedItems

`func (o *Delivery) SetTrackedItems(v []TrackedItem)`

SetTrackedItems sets TrackedItems field to given value.

### HasTrackedItems

`func (o *Delivery) HasTrackedItems() bool`

HasTrackedItems returns a boolean if a field has been set.

### GetSubscribers

`func (o *Delivery) GetSubscribers() []Subscriber`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *Delivery) GetSubscribersOk() (*[]Subscriber, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *Delivery) SetSubscribers(v []Subscriber)`

SetSubscribers sets Subscribers field to given value.

### HasSubscribers

`func (o *Delivery) HasSubscribers() bool`

HasSubscribers returns a boolean if a field has been set.

### GetAutoComplete

`func (o *Delivery) GetAutoComplete() bool`

GetAutoComplete returns the AutoComplete field if non-nil, zero value otherwise.

### GetAutoCompleteOk

`func (o *Delivery) GetAutoCompleteOk() (*bool, bool)`

GetAutoCompleteOk returns a tuple with the AutoComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoComplete

`func (o *Delivery) SetAutoComplete(v bool)`

SetAutoComplete sets AutoComplete field to given value.

### HasAutoComplete

`func (o *Delivery) HasAutoComplete() bool`

HasAutoComplete returns a boolean if a field has been set.

### GetTemplate

`func (o *Delivery) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Delivery) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Delivery) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Delivery) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTransitions

`func (o *Delivery) GetTransitions() []Transition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *Delivery) GetTransitionsOk() (*[]Transition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *Delivery) SetTransitions(v []Transition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *Delivery) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetStagesBeforeFirstOpenTransition

`func (o *Delivery) GetStagesBeforeFirstOpenTransition() []Stage`

GetStagesBeforeFirstOpenTransition returns the StagesBeforeFirstOpenTransition field if non-nil, zero value otherwise.

### GetStagesBeforeFirstOpenTransitionOk

`func (o *Delivery) GetStagesBeforeFirstOpenTransitionOk() (*[]Stage, bool)`

GetStagesBeforeFirstOpenTransitionOk returns a tuple with the StagesBeforeFirstOpenTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagesBeforeFirstOpenTransition

`func (o *Delivery) SetStagesBeforeFirstOpenTransition(v []Stage)`

SetStagesBeforeFirstOpenTransition sets StagesBeforeFirstOpenTransition field to given value.

### HasStagesBeforeFirstOpenTransition

`func (o *Delivery) HasStagesBeforeFirstOpenTransition() bool`

HasStagesBeforeFirstOpenTransition returns a boolean if a field has been set.

### GetUpdatable

`func (o *Delivery) GetUpdatable() bool`

GetUpdatable returns the Updatable field if non-nil, zero value otherwise.

### GetUpdatableOk

`func (o *Delivery) GetUpdatableOk() (*bool, bool)`

GetUpdatableOk returns a tuple with the Updatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatable

`func (o *Delivery) SetUpdatable(v bool)`

SetUpdatable sets Updatable field to given value.

### HasUpdatable

`func (o *Delivery) HasUpdatable() bool`

HasUpdatable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


