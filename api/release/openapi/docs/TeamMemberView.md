# TeamMemberView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**MemberType**](MemberType.md) |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 

## Methods

### NewTeamMemberView

`func NewTeamMemberView() *TeamMemberView`

NewTeamMemberView instantiates a new TeamMemberView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberViewWithDefaults

`func NewTeamMemberViewWithDefaults() *TeamMemberView`

NewTeamMemberViewWithDefaults instantiates a new TeamMemberView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamMemberView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamMemberView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamMemberView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamMemberView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFullName

`func (o *TeamMemberView) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *TeamMemberView) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *TeamMemberView) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *TeamMemberView) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetType

`func (o *TeamMemberView) GetType() MemberType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamMemberView) GetTypeOk() (*MemberType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamMemberView) SetType(v MemberType)`

SetType sets Type field to given value.

### HasType

`func (o *TeamMemberView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoleId

`func (o *TeamMemberView) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *TeamMemberView) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *TeamMemberView) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *TeamMemberView) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


