# ReleaseGroupTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**RiskScore** | Pointer to **int32** |  | [optional] 
**Releases** | Pointer to [**[]ReleaseTimeline**](ReleaseTimeline.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewReleaseGroupTimeline

`func NewReleaseGroupTimeline() *ReleaseGroupTimeline`

NewReleaseGroupTimeline instantiates a new ReleaseGroupTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseGroupTimelineWithDefaults

`func NewReleaseGroupTimelineWithDefaults() *ReleaseGroupTimeline`

NewReleaseGroupTimelineWithDefaults instantiates a new ReleaseGroupTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseGroupTimeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseGroupTimeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseGroupTimeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReleaseGroupTimeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ReleaseGroupTimeline) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseGroupTimeline) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseGroupTimeline) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseGroupTimeline) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRiskScore

`func (o *ReleaseGroupTimeline) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *ReleaseGroupTimeline) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *ReleaseGroupTimeline) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *ReleaseGroupTimeline) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetReleases

`func (o *ReleaseGroupTimeline) GetReleases() []ReleaseTimeline`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *ReleaseGroupTimeline) GetReleasesOk() (*[]ReleaseTimeline, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *ReleaseGroupTimeline) SetReleases(v []ReleaseTimeline)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *ReleaseGroupTimeline) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### GetStartDate

`func (o *ReleaseGroupTimeline) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ReleaseGroupTimeline) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ReleaseGroupTimeline) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ReleaseGroupTimeline) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ReleaseGroupTimeline) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ReleaseGroupTimeline) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ReleaseGroupTimeline) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ReleaseGroupTimeline) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


