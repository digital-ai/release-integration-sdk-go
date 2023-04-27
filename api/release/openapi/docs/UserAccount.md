# UserAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**External** | Pointer to **bool** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PreviousPassword** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**LoginAllowed** | Pointer to **bool** |  | [optional] 
**DateFormat** | Pointer to **string** |  | [optional] 
**TimeFormat** | Pointer to **string** |  | [optional] 
**FirstDayOfWeek** | Pointer to **int32** |  | [optional] 
**LastActive** | Pointer to **string** |  | [optional] 
**AnalyticsEnabled** | Pointer to **bool** |  | [optional] 
**TaskDrawerEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserAccount

`func NewUserAccount() *UserAccount`

NewUserAccount instantiates a new UserAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountWithDefaults

`func NewUserAccountWithDefaults() *UserAccount`

NewUserAccountWithDefaults instantiates a new UserAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserAccount) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserAccount) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetExternal

`func (o *UserAccount) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *UserAccount) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *UserAccount) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *UserAccount) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetProfileId

`func (o *UserAccount) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *UserAccount) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *UserAccount) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *UserAccount) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetEmail

`func (o *UserAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *UserAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPreviousPassword

`func (o *UserAccount) GetPreviousPassword() string`

GetPreviousPassword returns the PreviousPassword field if non-nil, zero value otherwise.

### GetPreviousPasswordOk

`func (o *UserAccount) GetPreviousPasswordOk() (*string, bool)`

GetPreviousPasswordOk returns a tuple with the PreviousPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPassword

`func (o *UserAccount) SetPreviousPassword(v string)`

SetPreviousPassword sets PreviousPassword field to given value.

### HasPreviousPassword

`func (o *UserAccount) HasPreviousPassword() bool`

HasPreviousPassword returns a boolean if a field has been set.

### GetFullName

`func (o *UserAccount) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserAccount) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserAccount) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UserAccount) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetExternalId

`func (o *UserAccount) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UserAccount) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UserAccount) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UserAccount) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLoginAllowed

`func (o *UserAccount) GetLoginAllowed() bool`

GetLoginAllowed returns the LoginAllowed field if non-nil, zero value otherwise.

### GetLoginAllowedOk

`func (o *UserAccount) GetLoginAllowedOk() (*bool, bool)`

GetLoginAllowedOk returns a tuple with the LoginAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAllowed

`func (o *UserAccount) SetLoginAllowed(v bool)`

SetLoginAllowed sets LoginAllowed field to given value.

### HasLoginAllowed

`func (o *UserAccount) HasLoginAllowed() bool`

HasLoginAllowed returns a boolean if a field has been set.

### GetDateFormat

`func (o *UserAccount) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *UserAccount) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *UserAccount) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *UserAccount) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetTimeFormat

`func (o *UserAccount) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *UserAccount) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *UserAccount) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *UserAccount) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### GetFirstDayOfWeek

`func (o *UserAccount) GetFirstDayOfWeek() int32`

GetFirstDayOfWeek returns the FirstDayOfWeek field if non-nil, zero value otherwise.

### GetFirstDayOfWeekOk

`func (o *UserAccount) GetFirstDayOfWeekOk() (*int32, bool)`

GetFirstDayOfWeekOk returns a tuple with the FirstDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDayOfWeek

`func (o *UserAccount) SetFirstDayOfWeek(v int32)`

SetFirstDayOfWeek sets FirstDayOfWeek field to given value.

### HasFirstDayOfWeek

`func (o *UserAccount) HasFirstDayOfWeek() bool`

HasFirstDayOfWeek returns a boolean if a field has been set.

### GetLastActive

`func (o *UserAccount) GetLastActive() string`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *UserAccount) GetLastActiveOk() (*string, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *UserAccount) SetLastActive(v string)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *UserAccount) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetAnalyticsEnabled

`func (o *UserAccount) GetAnalyticsEnabled() bool`

GetAnalyticsEnabled returns the AnalyticsEnabled field if non-nil, zero value otherwise.

### GetAnalyticsEnabledOk

`func (o *UserAccount) GetAnalyticsEnabledOk() (*bool, bool)`

GetAnalyticsEnabledOk returns a tuple with the AnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsEnabled

`func (o *UserAccount) SetAnalyticsEnabled(v bool)`

SetAnalyticsEnabled sets AnalyticsEnabled field to given value.

### HasAnalyticsEnabled

`func (o *UserAccount) HasAnalyticsEnabled() bool`

HasAnalyticsEnabled returns a boolean if a field has been set.

### GetTaskDrawerEnabled

`func (o *UserAccount) GetTaskDrawerEnabled() bool`

GetTaskDrawerEnabled returns the TaskDrawerEnabled field if non-nil, zero value otherwise.

### GetTaskDrawerEnabledOk

`func (o *UserAccount) GetTaskDrawerEnabledOk() (*bool, bool)`

GetTaskDrawerEnabledOk returns a tuple with the TaskDrawerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDrawerEnabled

`func (o *UserAccount) SetTaskDrawerEnabled(v bool)`

SetTaskDrawerEnabled sets TaskDrawerEnabled field to given value.

### HasTaskDrawerEnabled

`func (o *UserAccount) HasTaskDrawerEnabled() bool`

HasTaskDrawerEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


