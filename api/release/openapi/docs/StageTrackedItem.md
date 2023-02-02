# StageTrackedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackedItemId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TrackedItemStatus**](TrackedItemStatus.md) |  | [optional] 
**ReleaseIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStageTrackedItem

`func NewStageTrackedItem() *StageTrackedItem`

NewStageTrackedItem instantiates a new StageTrackedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageTrackedItemWithDefaults

`func NewStageTrackedItemWithDefaults() *StageTrackedItem`

NewStageTrackedItemWithDefaults instantiates a new StageTrackedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackedItemId

`func (o *StageTrackedItem) GetTrackedItemId() string`

GetTrackedItemId returns the TrackedItemId field if non-nil, zero value otherwise.

### GetTrackedItemIdOk

`func (o *StageTrackedItem) GetTrackedItemIdOk() (*string, bool)`

GetTrackedItemIdOk returns a tuple with the TrackedItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedItemId

`func (o *StageTrackedItem) SetTrackedItemId(v string)`

SetTrackedItemId sets TrackedItemId field to given value.

### HasTrackedItemId

`func (o *StageTrackedItem) HasTrackedItemId() bool`

HasTrackedItemId returns a boolean if a field has been set.

### GetStatus

`func (o *StageTrackedItem) GetStatus() TrackedItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StageTrackedItem) GetStatusOk() (*TrackedItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StageTrackedItem) SetStatus(v TrackedItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StageTrackedItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReleaseIds

`func (o *StageTrackedItem) GetReleaseIds() []string`

GetReleaseIds returns the ReleaseIds field if non-nil, zero value otherwise.

### GetReleaseIdsOk

`func (o *StageTrackedItem) GetReleaseIdsOk() (*[]string, bool)`

GetReleaseIdsOk returns a tuple with the ReleaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseIds

`func (o *StageTrackedItem) SetReleaseIds(v []string)`

SetReleaseIds sets ReleaseIds field to given value.

### HasReleaseIds

`func (o *StageTrackedItem) HasReleaseIds() bool`

HasReleaseIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


