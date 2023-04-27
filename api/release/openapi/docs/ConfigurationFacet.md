# ConfigurationFacet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**FacetScope**](FacetScope.md) |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**ConfigurationUri** | Pointer to **string** |  | [optional] 
**VariableMapping** | Pointer to **map[string]string** |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 
**PropertiesWithVariables** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewConfigurationFacet

`func NewConfigurationFacet() *ConfigurationFacet`

NewConfigurationFacet instantiates a new ConfigurationFacet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationFacetWithDefaults

`func NewConfigurationFacetWithDefaults() *ConfigurationFacet`

NewConfigurationFacetWithDefaults instantiates a new ConfigurationFacet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ConfigurationFacet) GetScope() FacetScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ConfigurationFacet) GetScopeOk() (*FacetScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ConfigurationFacet) SetScope(v FacetScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ConfigurationFacet) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTargetId

`func (o *ConfigurationFacet) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ConfigurationFacet) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ConfigurationFacet) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ConfigurationFacet) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetConfigurationUri

`func (o *ConfigurationFacet) GetConfigurationUri() string`

GetConfigurationUri returns the ConfigurationUri field if non-nil, zero value otherwise.

### GetConfigurationUriOk

`func (o *ConfigurationFacet) GetConfigurationUriOk() (*string, bool)`

GetConfigurationUriOk returns a tuple with the ConfigurationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUri

`func (o *ConfigurationFacet) SetConfigurationUri(v string)`

SetConfigurationUri sets ConfigurationUri field to given value.

### HasConfigurationUri

`func (o *ConfigurationFacet) HasConfigurationUri() bool`

HasConfigurationUri returns a boolean if a field has been set.

### GetVariableMapping

`func (o *ConfigurationFacet) GetVariableMapping() map[string]string`

GetVariableMapping returns the VariableMapping field if non-nil, zero value otherwise.

### GetVariableMappingOk

`func (o *ConfigurationFacet) GetVariableMappingOk() (*map[string]string, bool)`

GetVariableMappingOk returns a tuple with the VariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMapping

`func (o *ConfigurationFacet) SetVariableMapping(v map[string]string)`

SetVariableMapping sets VariableMapping field to given value.

### HasVariableMapping

`func (o *ConfigurationFacet) HasVariableMapping() bool`

HasVariableMapping returns a boolean if a field has been set.

### GetVariableUsages

`func (o *ConfigurationFacet) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *ConfigurationFacet) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *ConfigurationFacet) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *ConfigurationFacet) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.

### GetPropertiesWithVariables

`func (o *ConfigurationFacet) GetPropertiesWithVariables() []interface{}`

GetPropertiesWithVariables returns the PropertiesWithVariables field if non-nil, zero value otherwise.

### GetPropertiesWithVariablesOk

`func (o *ConfigurationFacet) GetPropertiesWithVariablesOk() (*[]interface{}, bool)`

GetPropertiesWithVariablesOk returns a tuple with the PropertiesWithVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesWithVariables

`func (o *ConfigurationFacet) SetPropertiesWithVariables(v []interface{})`

SetPropertiesWithVariables sets PropertiesWithVariables field to given value.

### HasPropertiesWithVariables

`func (o *ConfigurationFacet) HasPropertiesWithVariables() bool`

HasPropertiesWithVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


