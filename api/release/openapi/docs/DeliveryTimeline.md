# DeliveryTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Releases** | Pointer to [**[]ReleaseTimeline**](ReleaseTimeline.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 

## Methods

### NewDeliveryTimeline

`func NewDeliveryTimeline() *DeliveryTimeline`

NewDeliveryTimeline instantiates a new DeliveryTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryTimelineWithDefaults

`func NewDeliveryTimelineWithDefaults() *DeliveryTimeline`

NewDeliveryTimelineWithDefaults instantiates a new DeliveryTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryTimeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryTimeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryTimeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryTimeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *DeliveryTimeline) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeliveryTimeline) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeliveryTimeline) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DeliveryTimeline) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetReleases

`func (o *DeliveryTimeline) GetReleases() []ReleaseTimeline`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *DeliveryTimeline) GetReleasesOk() (*[]ReleaseTimeline, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *DeliveryTimeline) SetReleases(v []ReleaseTimeline)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *DeliveryTimeline) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### GetStartDate

`func (o *DeliveryTimeline) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DeliveryTimeline) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DeliveryTimeline) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DeliveryTimeline) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *DeliveryTimeline) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DeliveryTimeline) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DeliveryTimeline) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DeliveryTimeline) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


