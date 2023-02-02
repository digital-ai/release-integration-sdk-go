# RiskProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**DefaultProfile** | Pointer to **bool** |  | [optional] 
**RiskProfileAssessors** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewRiskProfile

`func NewRiskProfile() *RiskProfile`

NewRiskProfile instantiates a new RiskProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskProfileWithDefaults

`func NewRiskProfileWithDefaults() *RiskProfile`

NewRiskProfileWithDefaults instantiates a new RiskProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *RiskProfile) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RiskProfile) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RiskProfile) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RiskProfile) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTitle

`func (o *RiskProfile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RiskProfile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RiskProfile) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RiskProfile) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDefaultProfile

`func (o *RiskProfile) GetDefaultProfile() bool`

GetDefaultProfile returns the DefaultProfile field if non-nil, zero value otherwise.

### GetDefaultProfileOk

`func (o *RiskProfile) GetDefaultProfileOk() (*bool, bool)`

GetDefaultProfileOk returns a tuple with the DefaultProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProfile

`func (o *RiskProfile) SetDefaultProfile(v bool)`

SetDefaultProfile sets DefaultProfile field to given value.

### HasDefaultProfile

`func (o *RiskProfile) HasDefaultProfile() bool`

HasDefaultProfile returns a boolean if a field has been set.

### GetRiskProfileAssessors

`func (o *RiskProfile) GetRiskProfileAssessors() map[string]string`

GetRiskProfileAssessors returns the RiskProfileAssessors field if non-nil, zero value otherwise.

### GetRiskProfileAssessorsOk

`func (o *RiskProfile) GetRiskProfileAssessorsOk() (*map[string]string, bool)`

GetRiskProfileAssessorsOk returns a tuple with the RiskProfileAssessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskProfileAssessors

`func (o *RiskProfile) SetRiskProfileAssessors(v map[string]string)`

SetRiskProfileAssessors sets RiskProfileAssessors field to given value.

### HasRiskProfileAssessors

`func (o *RiskProfile) HasRiskProfileAssessors() bool`

HasRiskProfileAssessors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


