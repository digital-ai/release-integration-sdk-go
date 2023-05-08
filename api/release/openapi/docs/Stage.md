# Stage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StageStatus**](StageStatus.md) |  | [optional] 
**Items** | Pointer to [**[]StageTrackedItem**](StageTrackedItem.md) |  | [optional] 
**Transition** | Pointer to [**Transition**](Transition.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Team** | Pointer to **string** |  | [optional] 
**Open** | Pointer to **bool** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 

## Methods

### NewStage

`func NewStage() *Stage`

NewStage instantiates a new Stage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageWithDefaults

`func NewStageWithDefaults() *Stage`

NewStageWithDefaults instantiates a new Stage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Stage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Stage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Stage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Stage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Stage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Stage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Stage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Stage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *Stage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Stage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Stage) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Stage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *Stage) GetStatus() StageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Stage) GetStatusOk() (*StageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Stage) SetStatus(v StageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Stage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetItems

`func (o *Stage) GetItems() []StageTrackedItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Stage) GetItemsOk() (*[]StageTrackedItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Stage) SetItems(v []StageTrackedItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Stage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTransition

`func (o *Stage) GetTransition() Transition`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *Stage) GetTransitionOk() (*Transition, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *Stage) SetTransition(v Transition)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *Stage) HasTransition() bool`

HasTransition returns a boolean if a field has been set.

### GetOwner

`func (o *Stage) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Stage) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Stage) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Stage) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTeam

`func (o *Stage) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *Stage) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *Stage) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *Stage) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetOpen

`func (o *Stage) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Stage) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Stage) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *Stage) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetClosed

`func (o *Stage) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *Stage) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *Stage) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *Stage) HasClosed() bool`

HasClosed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


