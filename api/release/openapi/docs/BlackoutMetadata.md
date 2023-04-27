# BlackoutMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Periods** | Pointer to [**[]BlackoutPeriod**](BlackoutPeriod.md) |  | [optional] 

## Methods

### NewBlackoutMetadata

`func NewBlackoutMetadata() *BlackoutMetadata`

NewBlackoutMetadata instantiates a new BlackoutMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlackoutMetadataWithDefaults

`func NewBlackoutMetadataWithDefaults() *BlackoutMetadata`

NewBlackoutMetadataWithDefaults instantiates a new BlackoutMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriods

`func (o *BlackoutMetadata) GetPeriods() []BlackoutPeriod`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *BlackoutMetadata) GetPeriodsOk() (*[]BlackoutPeriod, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *BlackoutMetadata) SetPeriods(v []BlackoutPeriod)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *BlackoutMetadata) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


