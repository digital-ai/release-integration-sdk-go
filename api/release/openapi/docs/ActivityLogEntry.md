# ActivityLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**ActivityType** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**EventTime** | Pointer to **time.Time** |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**DataId** | Pointer to **string** |  | [optional] 

## Methods

### NewActivityLogEntry

`func NewActivityLogEntry() *ActivityLogEntry`

NewActivityLogEntry instantiates a new ActivityLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogEntryWithDefaults

`func NewActivityLogEntryWithDefaults() *ActivityLogEntry`

NewActivityLogEntryWithDefaults instantiates a new ActivityLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLogEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ActivityLogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityLogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityLogEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityLogEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *ActivityLogEntry) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ActivityLogEntry) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ActivityLogEntry) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ActivityLogEntry) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetActivityType

`func (o *ActivityLogEntry) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *ActivityLogEntry) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *ActivityLogEntry) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *ActivityLogEntry) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetMessage

`func (o *ActivityLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActivityLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActivityLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ActivityLogEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEventTime

`func (o *ActivityLogEntry) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *ActivityLogEntry) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *ActivityLogEntry) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *ActivityLogEntry) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetTargetType

`func (o *ActivityLogEntry) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ActivityLogEntry) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ActivityLogEntry) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ActivityLogEntry) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargetId

`func (o *ActivityLogEntry) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ActivityLogEntry) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ActivityLogEntry) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ActivityLogEntry) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetDataId

`func (o *ActivityLogEntry) GetDataId() string`

GetDataId returns the DataId field if non-nil, zero value otherwise.

### GetDataIdOk

`func (o *ActivityLogEntry) GetDataIdOk() (*string, bool)`

GetDataIdOk returns a tuple with the DataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataId

`func (o *ActivityLogEntry) SetDataId(v string)`

SetDataId sets DataId field to given value.

### HasDataId

`func (o *ActivityLogEntry) HasDataId() bool`

HasDataId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


