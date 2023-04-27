# DeliveryPatternFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**StrictTitleMatch** | Pointer to **bool** |  | [optional] 
**TrackedItemTitle** | Pointer to **string** |  | [optional] 
**StrictTrackedItemTitleMatch** | Pointer to **bool** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Statuses** | Pointer to [**[]DeliveryStatus**](DeliveryStatus.md) |  | [optional] 

## Methods

### NewDeliveryPatternFilters

`func NewDeliveryPatternFilters() *DeliveryPatternFilters`

NewDeliveryPatternFilters instantiates a new DeliveryPatternFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryPatternFiltersWithDefaults

`func NewDeliveryPatternFiltersWithDefaults() *DeliveryPatternFilters`

NewDeliveryPatternFiltersWithDefaults instantiates a new DeliveryPatternFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DeliveryPatternFilters) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeliveryPatternFilters) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeliveryPatternFilters) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DeliveryPatternFilters) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStrictTitleMatch

`func (o *DeliveryPatternFilters) GetStrictTitleMatch() bool`

GetStrictTitleMatch returns the StrictTitleMatch field if non-nil, zero value otherwise.

### GetStrictTitleMatchOk

`func (o *DeliveryPatternFilters) GetStrictTitleMatchOk() (*bool, bool)`

GetStrictTitleMatchOk returns a tuple with the StrictTitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictTitleMatch

`func (o *DeliveryPatternFilters) SetStrictTitleMatch(v bool)`

SetStrictTitleMatch sets StrictTitleMatch field to given value.

### HasStrictTitleMatch

`func (o *DeliveryPatternFilters) HasStrictTitleMatch() bool`

HasStrictTitleMatch returns a boolean if a field has been set.

### GetTrackedItemTitle

`func (o *DeliveryPatternFilters) GetTrackedItemTitle() string`

GetTrackedItemTitle returns the TrackedItemTitle field if non-nil, zero value otherwise.

### GetTrackedItemTitleOk

`func (o *DeliveryPatternFilters) GetTrackedItemTitleOk() (*string, bool)`

GetTrackedItemTitleOk returns a tuple with the TrackedItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedItemTitle

`func (o *DeliveryPatternFilters) SetTrackedItemTitle(v string)`

SetTrackedItemTitle sets TrackedItemTitle field to given value.

### HasTrackedItemTitle

`func (o *DeliveryPatternFilters) HasTrackedItemTitle() bool`

HasTrackedItemTitle returns a boolean if a field has been set.

### GetStrictTrackedItemTitleMatch

`func (o *DeliveryPatternFilters) GetStrictTrackedItemTitleMatch() bool`

GetStrictTrackedItemTitleMatch returns the StrictTrackedItemTitleMatch field if non-nil, zero value otherwise.

### GetStrictTrackedItemTitleMatchOk

`func (o *DeliveryPatternFilters) GetStrictTrackedItemTitleMatchOk() (*bool, bool)`

GetStrictTrackedItemTitleMatchOk returns a tuple with the StrictTrackedItemTitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictTrackedItemTitleMatch

`func (o *DeliveryPatternFilters) SetStrictTrackedItemTitleMatch(v bool)`

SetStrictTrackedItemTitleMatch sets StrictTrackedItemTitleMatch field to given value.

### HasStrictTrackedItemTitleMatch

`func (o *DeliveryPatternFilters) HasStrictTrackedItemTitleMatch() bool`

HasStrictTrackedItemTitleMatch returns a boolean if a field has been set.

### GetFolderId

`func (o *DeliveryPatternFilters) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DeliveryPatternFilters) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DeliveryPatternFilters) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DeliveryPatternFilters) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetStatuses

`func (o *DeliveryPatternFilters) GetStatuses() []DeliveryStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *DeliveryPatternFilters) GetStatusesOk() (*[]DeliveryStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *DeliveryPatternFilters) SetStatuses(v []DeliveryStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *DeliveryPatternFilters) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


