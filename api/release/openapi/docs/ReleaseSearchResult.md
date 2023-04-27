# ReleaseSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Releases** | Pointer to [**[]Release**](Release.md) |  | [optional] 

## Methods

### NewReleaseSearchResult

`func NewReleaseSearchResult() *ReleaseSearchResult`

NewReleaseSearchResult instantiates a new ReleaseSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseSearchResultWithDefaults

`func NewReleaseSearchResultWithDefaults() *ReleaseSearchResult`

NewReleaseSearchResultWithDefaults instantiates a new ReleaseSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ReleaseSearchResult) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ReleaseSearchResult) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ReleaseSearchResult) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ReleaseSearchResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *ReleaseSearchResult) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReleaseSearchResult) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReleaseSearchResult) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ReleaseSearchResult) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetReleases

`func (o *ReleaseSearchResult) GetReleases() []Release`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *ReleaseSearchResult) GetReleasesOk() (*[]Release, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *ReleaseSearchResult) SetReleases(v []Release)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *ReleaseSearchResult) HasReleases() bool`

HasReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


