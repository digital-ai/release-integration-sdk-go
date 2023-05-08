# Risk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**TotalScore** | Pointer to **int32** |  | [optional] 
**RiskAssessments** | Pointer to [**[]RiskAssessment**](RiskAssessment.md) |  | [optional] 

## Methods

### NewRisk

`func NewRisk() *Risk`

NewRisk instantiates a new Risk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskWithDefaults

`func NewRiskWithDefaults() *Risk`

NewRiskWithDefaults instantiates a new Risk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Risk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Risk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Risk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Risk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Risk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Risk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Risk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Risk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariableUsages

`func (o *Risk) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *Risk) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *Risk) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *Risk) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.

### GetScore

`func (o *Risk) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Risk) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Risk) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Risk) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTotalScore

`func (o *Risk) GetTotalScore() int32`

GetTotalScore returns the TotalScore field if non-nil, zero value otherwise.

### GetTotalScoreOk

`func (o *Risk) GetTotalScoreOk() (*int32, bool)`

GetTotalScoreOk returns a tuple with the TotalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalScore

`func (o *Risk) SetTotalScore(v int32)`

SetTotalScore sets TotalScore field to given value.

### HasTotalScore

`func (o *Risk) HasTotalScore() bool`

HasTotalScore returns a boolean if a field has been set.

### GetRiskAssessments

`func (o *Risk) GetRiskAssessments() []RiskAssessment`

GetRiskAssessments returns the RiskAssessments field if non-nil, zero value otherwise.

### GetRiskAssessmentsOk

`func (o *Risk) GetRiskAssessmentsOk() (*[]RiskAssessment, bool)`

GetRiskAssessmentsOk returns a tuple with the RiskAssessments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAssessments

`func (o *Risk) SetRiskAssessments(v []RiskAssessment)`

SetRiskAssessments sets RiskAssessments field to given value.

### HasRiskAssessments

`func (o *Risk) HasRiskAssessments() bool`

HasRiskAssessments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


