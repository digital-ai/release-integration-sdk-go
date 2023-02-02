# ReleaseCountResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**ByStatus** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewReleaseCountResult

`func NewReleaseCountResult() *ReleaseCountResult`

NewReleaseCountResult instantiates a new ReleaseCountResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseCountResultWithDefaults

`func NewReleaseCountResultWithDefaults() *ReleaseCountResult`

NewReleaseCountResultWithDefaults instantiates a new ReleaseCountResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ReleaseCountResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ReleaseCountResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ReleaseCountResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ReleaseCountResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *ReleaseCountResult) GetByStatus() map[string]int32`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *ReleaseCountResult) GetByStatusOk() (*map[string]int32, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *ReleaseCountResult) SetByStatus(v map[string]int32)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *ReleaseCountResult) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


