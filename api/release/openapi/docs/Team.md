# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**Members** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**ReleaseAdminTeam** | Pointer to **bool** |  | [optional] 
**TemplateOwnerTeam** | Pointer to **bool** |  | [optional] 
**FolderOwnerTeam** | Pointer to **bool** |  | [optional] 
**FolderAdminTeam** | Pointer to **bool** |  | [optional] 
**SystemTeam** | Pointer to **bool** |  | [optional] 

## Methods

### NewTeam

`func NewTeam() *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Team) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Team) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Team) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Team) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Team) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Team) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTeamName

`func (o *Team) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *Team) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *Team) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *Team) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetMembers

`func (o *Team) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Team) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Team) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Team) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetRoles

`func (o *Team) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Team) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Team) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Team) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPermissions

`func (o *Team) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Team) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Team) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Team) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetReleaseAdminTeam

`func (o *Team) GetReleaseAdminTeam() bool`

GetReleaseAdminTeam returns the ReleaseAdminTeam field if non-nil, zero value otherwise.

### GetReleaseAdminTeamOk

`func (o *Team) GetReleaseAdminTeamOk() (*bool, bool)`

GetReleaseAdminTeamOk returns a tuple with the ReleaseAdminTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseAdminTeam

`func (o *Team) SetReleaseAdminTeam(v bool)`

SetReleaseAdminTeam sets ReleaseAdminTeam field to given value.

### HasReleaseAdminTeam

`func (o *Team) HasReleaseAdminTeam() bool`

HasReleaseAdminTeam returns a boolean if a field has been set.

### GetTemplateOwnerTeam

`func (o *Team) GetTemplateOwnerTeam() bool`

GetTemplateOwnerTeam returns the TemplateOwnerTeam field if non-nil, zero value otherwise.

### GetTemplateOwnerTeamOk

`func (o *Team) GetTemplateOwnerTeamOk() (*bool, bool)`

GetTemplateOwnerTeamOk returns a tuple with the TemplateOwnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateOwnerTeam

`func (o *Team) SetTemplateOwnerTeam(v bool)`

SetTemplateOwnerTeam sets TemplateOwnerTeam field to given value.

### HasTemplateOwnerTeam

`func (o *Team) HasTemplateOwnerTeam() bool`

HasTemplateOwnerTeam returns a boolean if a field has been set.

### GetFolderOwnerTeam

`func (o *Team) GetFolderOwnerTeam() bool`

GetFolderOwnerTeam returns the FolderOwnerTeam field if non-nil, zero value otherwise.

### GetFolderOwnerTeamOk

`func (o *Team) GetFolderOwnerTeamOk() (*bool, bool)`

GetFolderOwnerTeamOk returns a tuple with the FolderOwnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderOwnerTeam

`func (o *Team) SetFolderOwnerTeam(v bool)`

SetFolderOwnerTeam sets FolderOwnerTeam field to given value.

### HasFolderOwnerTeam

`func (o *Team) HasFolderOwnerTeam() bool`

HasFolderOwnerTeam returns a boolean if a field has been set.

### GetFolderAdminTeam

`func (o *Team) GetFolderAdminTeam() bool`

GetFolderAdminTeam returns the FolderAdminTeam field if non-nil, zero value otherwise.

### GetFolderAdminTeamOk

`func (o *Team) GetFolderAdminTeamOk() (*bool, bool)`

GetFolderAdminTeamOk returns a tuple with the FolderAdminTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderAdminTeam

`func (o *Team) SetFolderAdminTeam(v bool)`

SetFolderAdminTeam sets FolderAdminTeam field to given value.

### HasFolderAdminTeam

`func (o *Team) HasFolderAdminTeam() bool`

HasFolderAdminTeam returns a boolean if a field has been set.

### GetSystemTeam

`func (o *Team) GetSystemTeam() bool`

GetSystemTeam returns the SystemTeam field if non-nil, zero value otherwise.

### GetSystemTeamOk

`func (o *Team) GetSystemTeamOk() (*bool, bool)`

GetSystemTeamOk returns a tuple with the SystemTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTeam

`func (o *Team) SetSystemTeam(v bool)`

SetSystemTeam sets SystemTeam field to given value.

### HasSystemTeam

`func (o *Team) HasSystemTeam() bool`

HasSystemTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


