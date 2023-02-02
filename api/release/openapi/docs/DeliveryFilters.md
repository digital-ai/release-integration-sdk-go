# DeliveryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**StrictTitleMatch** | Pointer to **bool** |  | [optional] 
**TrackedItemTitle** | Pointer to **string** |  | [optional] 
**StrictTrackedItemTitleMatch** | Pointer to **bool** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**OriginPatternId** | Pointer to **string** |  | [optional] 
**Statuses** | Pointer to [**[]DeliveryStatus**](DeliveryStatus.md) |  | [optional] 

## Methods

### NewDeliveryFilters

`func NewDeliveryFilters() *DeliveryFilters`

NewDeliveryFilters instantiates a new DeliveryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryFiltersWithDefaults

`func NewDeliveryFiltersWithDefaults() *DeliveryFilters`

NewDeliveryFiltersWithDefaults instantiates a new DeliveryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DeliveryFilters) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeliveryFilters) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeliveryFilters) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DeliveryFilters) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStrictTitleMatch

`func (o *DeliveryFilters) GetStrictTitleMatch() bool`

GetStrictTitleMatch returns the StrictTitleMatch field if non-nil, zero value otherwise.

### GetStrictTitleMatchOk

`func (o *DeliveryFilters) GetStrictTitleMatchOk() (*bool, bool)`

GetStrictTitleMatchOk returns a tuple with the StrictTitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictTitleMatch

`func (o *DeliveryFilters) SetStrictTitleMatch(v bool)`

SetStrictTitleMatch sets StrictTitleMatch field to given value.

### HasStrictTitleMatch

`func (o *DeliveryFilters) HasStrictTitleMatch() bool`

HasStrictTitleMatch returns a boolean if a field has been set.

### GetTrackedItemTitle

`func (o *DeliveryFilters) GetTrackedItemTitle() string`

GetTrackedItemTitle returns the TrackedItemTitle field if non-nil, zero value otherwise.

### GetTrackedItemTitleOk

`func (o *DeliveryFilters) GetTrackedItemTitleOk() (*string, bool)`

GetTrackedItemTitleOk returns a tuple with the TrackedItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedItemTitle

`func (o *DeliveryFilters) SetTrackedItemTitle(v string)`

SetTrackedItemTitle sets TrackedItemTitle field to given value.

### HasTrackedItemTitle

`func (o *DeliveryFilters) HasTrackedItemTitle() bool`

HasTrackedItemTitle returns a boolean if a field has been set.

### GetStrictTrackedItemTitleMatch

`func (o *DeliveryFilters) GetStrictTrackedItemTitleMatch() bool`

GetStrictTrackedItemTitleMatch returns the StrictTrackedItemTitleMatch field if non-nil, zero value otherwise.

### GetStrictTrackedItemTitleMatchOk

`func (o *DeliveryFilters) GetStrictTrackedItemTitleMatchOk() (*bool, bool)`

GetStrictTrackedItemTitleMatchOk returns a tuple with the StrictTrackedItemTitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictTrackedItemTitleMatch

`func (o *DeliveryFilters) SetStrictTrackedItemTitleMatch(v bool)`

SetStrictTrackedItemTitleMatch sets StrictTrackedItemTitleMatch field to given value.

### HasStrictTrackedItemTitleMatch

`func (o *DeliveryFilters) HasStrictTrackedItemTitleMatch() bool`

HasStrictTrackedItemTitleMatch returns a boolean if a field has been set.

### GetFolderId

`func (o *DeliveryFilters) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DeliveryFilters) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DeliveryFilters) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DeliveryFilters) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetOriginPatternId

`func (o *DeliveryFilters) GetOriginPatternId() string`

GetOriginPatternId returns the OriginPatternId field if non-nil, zero value otherwise.

### GetOriginPatternIdOk

`func (o *DeliveryFilters) GetOriginPatternIdOk() (*string, bool)`

GetOriginPatternIdOk returns a tuple with the OriginPatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPatternId

`func (o *DeliveryFilters) SetOriginPatternId(v string)`

SetOriginPatternId sets OriginPatternId field to given value.

### HasOriginPatternId

`func (o *DeliveryFilters) HasOriginPatternId() bool`

HasOriginPatternId returns a boolean if a field has been set.

### GetStatuses

`func (o *DeliveryFilters) GetStatuses() []DeliveryStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *DeliveryFilters) GetStatusesOk() (*[]DeliveryStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *DeliveryFilters) SetStatuses(v []DeliveryStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *DeliveryFilters) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


