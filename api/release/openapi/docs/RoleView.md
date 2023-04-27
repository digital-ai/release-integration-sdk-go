# RoleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Principals** | Pointer to [**[]PrincipalView**](PrincipalView.md) |  | [optional] 

## Methods

### NewRoleView

`func NewRoleView() *RoleView`

NewRoleView instantiates a new RoleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleViewWithDefaults

`func NewRoleViewWithDefaults() *RoleView`

NewRoleViewWithDefaults instantiates a new RoleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *RoleView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermissions

`func (o *RoleView) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleView) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleView) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RoleView) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPrincipals

`func (o *RoleView) GetPrincipals() []PrincipalView`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *RoleView) GetPrincipalsOk() (*[]PrincipalView, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *RoleView) SetPrincipals(v []PrincipalView)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *RoleView) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


