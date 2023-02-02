# TeamView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]TeamMemberView**](TeamMemberView.md) |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**SystemTeam** | Pointer to **bool** |  | [optional] 

## Methods

### NewTeamView

`func NewTeamView() *TeamView`

NewTeamView instantiates a new TeamView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamViewWithDefaults

`func NewTeamViewWithDefaults() *TeamView`

NewTeamViewWithDefaults instantiates a new TeamView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamName

`func (o *TeamView) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *TeamView) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *TeamView) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *TeamView) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetMembers

`func (o *TeamView) GetMembers() []TeamMemberView`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamView) GetMembersOk() (*[]TeamMemberView, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamView) SetMembers(v []TeamMemberView)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamView) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetPermissions

`func (o *TeamView) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TeamView) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TeamView) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TeamView) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSystemTeam

`func (o *TeamView) GetSystemTeam() bool`

GetSystemTeam returns the SystemTeam field if non-nil, zero value otherwise.

### GetSystemTeamOk

`func (o *TeamView) GetSystemTeamOk() (*bool, bool)`

GetSystemTeamOk returns a tuple with the SystemTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTeam

`func (o *TeamView) SetSystemTeam(v bool)`

SetSystemTeam sets SystemTeam field to given value.

### HasSystemTeam

`func (o *TeamView) HasSystemTeam() bool`

HasSystemTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


