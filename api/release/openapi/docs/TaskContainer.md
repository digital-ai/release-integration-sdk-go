# TaskContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskContainer

`func NewTaskContainer() *TaskContainer`

NewTaskContainer instantiates a new TaskContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskContainerWithDefaults

`func NewTaskContainerWithDefaults() *TaskContainer`

NewTaskContainerWithDefaults instantiates a new TaskContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskContainer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TaskContainer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskContainer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskContainer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskContainer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTasks

`func (o *TaskContainer) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TaskContainer) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TaskContainer) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TaskContainer) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetLocked

`func (o *TaskContainer) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *TaskContainer) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *TaskContainer) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *TaskContainer) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetTitle

`func (o *TaskContainer) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskContainer) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskContainer) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskContainer) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


