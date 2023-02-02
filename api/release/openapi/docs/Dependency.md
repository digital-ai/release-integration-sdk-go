# Dependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GateTask** | Pointer to [**GateTask**](GateTask.md) |  | [optional] 
**Target** | Pointer to [**PlanItem**](PlanItem.md) |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**ArchivedTargetTitle** | Pointer to **string** |  | [optional] 
**ArchivedTargetId** | Pointer to **string** |  | [optional] 
**ArchivedAsResolved** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Done** | Pointer to **bool** |  | [optional] 
**Aborted** | Pointer to **bool** |  | [optional] 
**TargetDisplayPath** | Pointer to **string** |  | [optional] 
**TargetTitle** | Pointer to **string** |  | [optional] 
**ValidAllowedPlanItemId** | Pointer to **bool** |  | [optional] 

## Methods

### NewDependency

`func NewDependency() *Dependency`

NewDependency instantiates a new Dependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependencyWithDefaults

`func NewDependencyWithDefaults() *Dependency`

NewDependencyWithDefaults instantiates a new Dependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateTask

`func (o *Dependency) GetGateTask() GateTask`

GetGateTask returns the GateTask field if non-nil, zero value otherwise.

### GetGateTaskOk

`func (o *Dependency) GetGateTaskOk() (*GateTask, bool)`

GetGateTaskOk returns a tuple with the GateTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateTask

`func (o *Dependency) SetGateTask(v GateTask)`

SetGateTask sets GateTask field to given value.

### HasGateTask

`func (o *Dependency) HasGateTask() bool`

HasGateTask returns a boolean if a field has been set.

### GetTarget

`func (o *Dependency) GetTarget() PlanItem`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Dependency) GetTargetOk() (*PlanItem, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Dependency) SetTarget(v PlanItem)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Dependency) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetId

`func (o *Dependency) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Dependency) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Dependency) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Dependency) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetArchivedTargetTitle

`func (o *Dependency) GetArchivedTargetTitle() string`

GetArchivedTargetTitle returns the ArchivedTargetTitle field if non-nil, zero value otherwise.

### GetArchivedTargetTitleOk

`func (o *Dependency) GetArchivedTargetTitleOk() (*string, bool)`

GetArchivedTargetTitleOk returns a tuple with the ArchivedTargetTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedTargetTitle

`func (o *Dependency) SetArchivedTargetTitle(v string)`

SetArchivedTargetTitle sets ArchivedTargetTitle field to given value.

### HasArchivedTargetTitle

`func (o *Dependency) HasArchivedTargetTitle() bool`

HasArchivedTargetTitle returns a boolean if a field has been set.

### GetArchivedTargetId

`func (o *Dependency) GetArchivedTargetId() string`

GetArchivedTargetId returns the ArchivedTargetId field if non-nil, zero value otherwise.

### GetArchivedTargetIdOk

`func (o *Dependency) GetArchivedTargetIdOk() (*string, bool)`

GetArchivedTargetIdOk returns a tuple with the ArchivedTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedTargetId

`func (o *Dependency) SetArchivedTargetId(v string)`

SetArchivedTargetId sets ArchivedTargetId field to given value.

### HasArchivedTargetId

`func (o *Dependency) HasArchivedTargetId() bool`

HasArchivedTargetId returns a boolean if a field has been set.

### GetArchivedAsResolved

`func (o *Dependency) GetArchivedAsResolved() bool`

GetArchivedAsResolved returns the ArchivedAsResolved field if non-nil, zero value otherwise.

### GetArchivedAsResolvedOk

`func (o *Dependency) GetArchivedAsResolvedOk() (*bool, bool)`

GetArchivedAsResolvedOk returns a tuple with the ArchivedAsResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAsResolved

`func (o *Dependency) SetArchivedAsResolved(v bool)`

SetArchivedAsResolved sets ArchivedAsResolved field to given value.

### HasArchivedAsResolved

`func (o *Dependency) HasArchivedAsResolved() bool`

HasArchivedAsResolved returns a boolean if a field has been set.

### GetMetadata

`func (o *Dependency) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Dependency) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Dependency) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Dependency) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetArchived

`func (o *Dependency) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Dependency) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Dependency) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Dependency) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDone

`func (o *Dependency) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *Dependency) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *Dependency) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *Dependency) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetAborted

`func (o *Dependency) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *Dependency) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *Dependency) SetAborted(v bool)`

SetAborted sets Aborted field to given value.

### HasAborted

`func (o *Dependency) HasAborted() bool`

HasAborted returns a boolean if a field has been set.

### GetTargetDisplayPath

`func (o *Dependency) GetTargetDisplayPath() string`

GetTargetDisplayPath returns the TargetDisplayPath field if non-nil, zero value otherwise.

### GetTargetDisplayPathOk

`func (o *Dependency) GetTargetDisplayPathOk() (*string, bool)`

GetTargetDisplayPathOk returns a tuple with the TargetDisplayPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDisplayPath

`func (o *Dependency) SetTargetDisplayPath(v string)`

SetTargetDisplayPath sets TargetDisplayPath field to given value.

### HasTargetDisplayPath

`func (o *Dependency) HasTargetDisplayPath() bool`

HasTargetDisplayPath returns a boolean if a field has been set.

### GetTargetTitle

`func (o *Dependency) GetTargetTitle() string`

GetTargetTitle returns the TargetTitle field if non-nil, zero value otherwise.

### GetTargetTitleOk

`func (o *Dependency) GetTargetTitleOk() (*string, bool)`

GetTargetTitleOk returns a tuple with the TargetTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTitle

`func (o *Dependency) SetTargetTitle(v string)`

SetTargetTitle sets TargetTitle field to given value.

### HasTargetTitle

`func (o *Dependency) HasTargetTitle() bool`

HasTargetTitle returns a boolean if a field has been set.

### GetValidAllowedPlanItemId

`func (o *Dependency) GetValidAllowedPlanItemId() bool`

GetValidAllowedPlanItemId returns the ValidAllowedPlanItemId field if non-nil, zero value otherwise.

### GetValidAllowedPlanItemIdOk

`func (o *Dependency) GetValidAllowedPlanItemIdOk() (*bool, bool)`

GetValidAllowedPlanItemIdOk returns a tuple with the ValidAllowedPlanItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidAllowedPlanItemId

`func (o *Dependency) SetValidAllowedPlanItemId(v bool)`

SetValidAllowedPlanItemId sets ValidAllowedPlanItemId field to given value.

### HasValidAllowedPlanItemId

`func (o *Dependency) HasValidAllowedPlanItemId() bool`

HasValidAllowedPlanItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


