# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewComment

`func NewComment() *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Comment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Comment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Comment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Comment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Comment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *Comment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Comment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Comment) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Comment) HasText() bool`

HasText returns a boolean if a field has been set.

### GetAuthor

`func (o *Comment) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Comment) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Comment) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Comment) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDate

`func (o *Comment) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Comment) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Comment) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Comment) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *Comment) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Comment) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Comment) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Comment) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


