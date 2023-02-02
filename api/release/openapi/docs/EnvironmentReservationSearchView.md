# EnvironmentReservationSearchView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to [**EnvironmentStageView**](EnvironmentStageView.md) |  | [optional] 
**Labels** | Pointer to [**[]EnvironmentLabelView**](EnvironmentLabelView.md) |  | [optional] 
**Reservations** | Pointer to [**[]ReservationSearchView**](ReservationSearchView.md) |  | [optional] 

## Methods

### NewEnvironmentReservationSearchView

`func NewEnvironmentReservationSearchView() *EnvironmentReservationSearchView`

NewEnvironmentReservationSearchView instantiates a new EnvironmentReservationSearchView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReservationSearchViewWithDefaults

`func NewEnvironmentReservationSearchViewWithDefaults() *EnvironmentReservationSearchView`

NewEnvironmentReservationSearchViewWithDefaults instantiates a new EnvironmentReservationSearchView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentReservationSearchView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentReservationSearchView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentReservationSearchView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentReservationSearchView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *EnvironmentReservationSearchView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EnvironmentReservationSearchView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EnvironmentReservationSearchView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EnvironmentReservationSearchView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentReservationSearchView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentReservationSearchView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentReservationSearchView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentReservationSearchView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStage

`func (o *EnvironmentReservationSearchView) GetStage() EnvironmentStageView`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *EnvironmentReservationSearchView) GetStageOk() (*EnvironmentStageView, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *EnvironmentReservationSearchView) SetStage(v EnvironmentStageView)`

SetStage sets Stage field to given value.

### HasStage

`func (o *EnvironmentReservationSearchView) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetLabels

`func (o *EnvironmentReservationSearchView) GetLabels() []EnvironmentLabelView`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EnvironmentReservationSearchView) GetLabelsOk() (*[]EnvironmentLabelView, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EnvironmentReservationSearchView) SetLabels(v []EnvironmentLabelView)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EnvironmentReservationSearchView) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetReservations

`func (o *EnvironmentReservationSearchView) GetReservations() []ReservationSearchView`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *EnvironmentReservationSearchView) GetReservationsOk() (*[]ReservationSearchView, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *EnvironmentReservationSearchView) SetReservations(v []ReservationSearchView)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *EnvironmentReservationSearchView) HasReservations() bool`

HasReservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


