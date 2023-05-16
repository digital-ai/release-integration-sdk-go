# RiskAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 
**RiskAssessorId** | Pointer to **string** |  | [optional] 
**Risk** | Pointer to [**Risk**](Risk.md) |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Headline** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 

## Methods

### NewRiskAssessment

`func NewRiskAssessment() *RiskAssessment`

NewRiskAssessment instantiates a new RiskAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAssessmentWithDefaults

`func NewRiskAssessmentWithDefaults() *RiskAssessment`

NewRiskAssessmentWithDefaults instantiates a new RiskAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskAssessment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskAssessment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskAssessment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskAssessment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RiskAssessment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskAssessment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskAssessment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RiskAssessment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariableUsages

`func (o *RiskAssessment) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *RiskAssessment) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *RiskAssessment) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *RiskAssessment) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.

### GetRiskAssessorId

`func (o *RiskAssessment) GetRiskAssessorId() string`

GetRiskAssessorId returns the RiskAssessorId field if non-nil, zero value otherwise.

### GetRiskAssessorIdOk

`func (o *RiskAssessment) GetRiskAssessorIdOk() (*string, bool)`

GetRiskAssessorIdOk returns a tuple with the RiskAssessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAssessorId

`func (o *RiskAssessment) SetRiskAssessorId(v string)`

SetRiskAssessorId sets RiskAssessorId field to given value.

### HasRiskAssessorId

`func (o *RiskAssessment) HasRiskAssessorId() bool`

HasRiskAssessorId returns a boolean if a field has been set.

### GetRisk

`func (o *RiskAssessment) GetRisk() Risk`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *RiskAssessment) GetRiskOk() (*Risk, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *RiskAssessment) SetRisk(v Risk)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *RiskAssessment) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetScore

`func (o *RiskAssessment) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RiskAssessment) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RiskAssessment) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *RiskAssessment) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetHeadline

`func (o *RiskAssessment) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *RiskAssessment) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *RiskAssessment) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *RiskAssessment) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetMessages

`func (o *RiskAssessment) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *RiskAssessment) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *RiskAssessment) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *RiskAssessment) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetIcon

`func (o *RiskAssessment) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RiskAssessment) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RiskAssessment) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RiskAssessment) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


