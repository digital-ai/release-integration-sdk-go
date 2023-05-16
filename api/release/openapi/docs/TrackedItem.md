# TrackedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ReleaseIds** | Pointer to **[]string** |  | [optional] 
**Descoped** | Pointer to **bool** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ModifiedDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTrackedItem

`func NewTrackedItem() *TrackedItem`

NewTrackedItem instantiates a new TrackedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackedItemWithDefaults

`func NewTrackedItemWithDefaults() *TrackedItem`

NewTrackedItemWithDefaults instantiates a new TrackedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrackedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackedItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackedItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TrackedItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrackedItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrackedItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TrackedItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *TrackedItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrackedItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrackedItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrackedItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetReleaseIds

`func (o *TrackedItem) GetReleaseIds() []string`

GetReleaseIds returns the ReleaseIds field if non-nil, zero value otherwise.

### GetReleaseIdsOk

`func (o *TrackedItem) GetReleaseIdsOk() (*[]string, bool)`

GetReleaseIdsOk returns a tuple with the ReleaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseIds

`func (o *TrackedItem) SetReleaseIds(v []string)`

SetReleaseIds sets ReleaseIds field to given value.

### HasReleaseIds

`func (o *TrackedItem) HasReleaseIds() bool`

HasReleaseIds returns a boolean if a field has been set.

### GetDescoped

`func (o *TrackedItem) GetDescoped() bool`

GetDescoped returns the Descoped field if non-nil, zero value otherwise.

### GetDescopedOk

`func (o *TrackedItem) GetDescopedOk() (*bool, bool)`

GetDescopedOk returns a tuple with the Descoped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescoped

`func (o *TrackedItem) SetDescoped(v bool)`

SetDescoped sets Descoped field to given value.

### HasDescoped

`func (o *TrackedItem) HasDescoped() bool`

HasDescoped returns a boolean if a field has been set.

### GetCreatedDate

`func (o *TrackedItem) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TrackedItem) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TrackedItem) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TrackedItem) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *TrackedItem) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TrackedItem) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TrackedItem) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TrackedItem) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


