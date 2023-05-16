# Facet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to [**FacetScope**](FacetScope.md) |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**ConfigurationUri** | Pointer to **string** |  | [optional] 
**VariableMapping** | Pointer to **map[string]string** |  | [optional] 
**VariableUsages** | Pointer to [**[]UsagePoint**](UsagePoint.md) |  | [optional] 
**PropertiesWithVariables** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewFacet

`func NewFacet() *Facet`

NewFacet instantiates a new Facet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetWithDefaults

`func NewFacetWithDefaults() *Facet`

NewFacetWithDefaults instantiates a new Facet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Facet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Facet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Facet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Facet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Facet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Facet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Facet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Facet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScope

`func (o *Facet) GetScope() FacetScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Facet) GetScopeOk() (*FacetScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Facet) SetScope(v FacetScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Facet) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTargetId

`func (o *Facet) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Facet) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Facet) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Facet) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetConfigurationUri

`func (o *Facet) GetConfigurationUri() string`

GetConfigurationUri returns the ConfigurationUri field if non-nil, zero value otherwise.

### GetConfigurationUriOk

`func (o *Facet) GetConfigurationUriOk() (*string, bool)`

GetConfigurationUriOk returns a tuple with the ConfigurationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUri

`func (o *Facet) SetConfigurationUri(v string)`

SetConfigurationUri sets ConfigurationUri field to given value.

### HasConfigurationUri

`func (o *Facet) HasConfigurationUri() bool`

HasConfigurationUri returns a boolean if a field has been set.

### GetVariableMapping

`func (o *Facet) GetVariableMapping() map[string]string`

GetVariableMapping returns the VariableMapping field if non-nil, zero value otherwise.

### GetVariableMappingOk

`func (o *Facet) GetVariableMappingOk() (*map[string]string, bool)`

GetVariableMappingOk returns a tuple with the VariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMapping

`func (o *Facet) SetVariableMapping(v map[string]string)`

SetVariableMapping sets VariableMapping field to given value.

### HasVariableMapping

`func (o *Facet) HasVariableMapping() bool`

HasVariableMapping returns a boolean if a field has been set.

### GetVariableUsages

`func (o *Facet) GetVariableUsages() []UsagePoint`

GetVariableUsages returns the VariableUsages field if non-nil, zero value otherwise.

### GetVariableUsagesOk

`func (o *Facet) GetVariableUsagesOk() (*[]UsagePoint, bool)`

GetVariableUsagesOk returns a tuple with the VariableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUsages

`func (o *Facet) SetVariableUsages(v []UsagePoint)`

SetVariableUsages sets VariableUsages field to given value.

### HasVariableUsages

`func (o *Facet) HasVariableUsages() bool`

HasVariableUsages returns a boolean if a field has been set.

### GetPropertiesWithVariables

`func (o *Facet) GetPropertiesWithVariables() []interface{}`

GetPropertiesWithVariables returns the PropertiesWithVariables field if non-nil, zero value otherwise.

### GetPropertiesWithVariablesOk

`func (o *Facet) GetPropertiesWithVariablesOk() (*[]interface{}, bool)`

GetPropertiesWithVariablesOk returns a tuple with the PropertiesWithVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesWithVariables

`func (o *Facet) SetPropertiesWithVariables(v []interface{})`

SetPropertiesWithVariables sets PropertiesWithVariables field to given value.

### HasPropertiesWithVariables

`func (o *Facet) HasPropertiesWithVariables() bool`

HasPropertiesWithVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


