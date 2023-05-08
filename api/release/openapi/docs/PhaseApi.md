# \PhaseApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPhase**](PhaseApi.md#AddPhase) | **Post** /api/v1/phases/{releaseId}/phase | 
[**AddTaskToContainer**](PhaseApi.md#AddTaskToContainer) | **Post** /api/v1/phases/{containerId}/tasks | 
[**CopyPhase**](PhaseApi.md#CopyPhase) | **Post** /api/v1/phases/{phaseId}/copy | 
[**DeletePhase**](PhaseApi.md#DeletePhase) | **Delete** /api/v1/phases/{phaseId} | 
[**GetPhase**](PhaseApi.md#GetPhase) | **Get** /api/v1/phases/{phaseId} | 
[**SearchPhases**](PhaseApi.md#SearchPhases) | **Get** /api/v1/phases/search | 
[**SearchPhasesByTitle**](PhaseApi.md#SearchPhasesByTitle) | **Get** /api/v1/phases/byTitle | 
[**UpdatePhase**](PhaseApi.md#UpdatePhase) | **Put** /api/v1/phases/{phaseId} | 



## AddPhase

> Phase AddPhase(ctx, releaseId).Position(position).Phase(phase).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    releaseId := "releaseId_example" // string | 
    position := int32(56) // int32 |  (optional)
    phase := *openapiclient.NewPhase() // Phase |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhaseApi.AddPhase(context.Background(), releaseId).Position(position).Phase(phase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhaseApi.AddPhase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPhase`: Phase
    fmt.Fprintf(os.Stdout, "Response from `PhaseApi.AddPhase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPhaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **position** | **int32** |  | 
 **phase** | [**Phase**](Phase.md) |  | 

### Return type

[**Phase**](Phase.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTaskToContainer

> Task AddTaskToContainer(ctx, containerId).Position(position).Task(task).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    containerId := "containerId_example" // string | 
    position := int32(56) // int32 |  (optional)
    task := *openapiclient.NewTask() // Task |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhaseApi.AddTaskToContainer(context.Background(), containerId).Position(position).Task(task).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhaseApi.AddTaskToContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTaskToContainer`: Task
    fmt.Fprintf(os.Stdout, "Response from `PhaseApi.AddTaskToContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTaskToContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **position** | **int32** |  | 
 **task** | [**Task**](Task.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyPhase

> Phase CopyPhase(ctx, phaseId).TargetPosition(targetPosition).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    phaseId := "phaseId_example" // string | 
    targetPosition := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhaseApi.CopyPhase(context.Background(), phaseId).TargetPosition(targetPosition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhaseApi.CopyPhase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyPhase`: Phase
    fmt.Fprintf(os.Stdout, "Response from `PhaseApi.CopyPhase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyPhaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetPosition** | **int32** |  | 

### Return type

[**Phase**](Phase.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePhase

> DeletePhase(ctx, phaseId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    phaseId := "phaseId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PhaseApi.DeletePhase(context.Background(), phaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhaseApi.DeletePhase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePhaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhase

> Phase GetPhase(ctx, phaseId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    phaseId := "phaseId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhaseApi.GetPhase(context.Background(), phaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhaseApi.GetPhase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhase`: Phase
    fmt.Fprintf(os.Stdout, "Response from `PhaseApi.GetPhase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Phase**](Phase.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPhases

> []Phase SearchPhases(ctx).PhaseTitle(phaseTitle).PhaseVersion(phaseVersion).ReleaseId(releaseId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    phaseTitle := "phaseTitle_example" // string |  (optional)
    phaseVersion := openapiclient.PhaseVersion("LATEST") // PhaseVersion |  (optional)
    releaseId := "releaseId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhaseApi.SearchPhases(context.Background()).PhaseTitle(phaseTitle).PhaseVersion(phaseVersion).ReleaseId(releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhaseApi.SearchPhases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPhases`: []Phase
    fmt.Fprintf(os.Stdout, "Response from `PhaseApi.SearchPhases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPhasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phaseTitle** | **string** |  | 
 **phaseVersion** | [**PhaseVersion**](PhaseVersion.md) |  | 
 **releaseId** | **string** |  | 

### Return type

[**[]Phase**](Phase.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPhasesByTitle

> []Phase SearchPhasesByTitle(ctx).PhaseTitle(phaseTitle).ReleaseId(releaseId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    phaseTitle := "phaseTitle_example" // string |  (optional)
    releaseId := "releaseId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhaseApi.SearchPhasesByTitle(context.Background()).PhaseTitle(phaseTitle).ReleaseId(releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhaseApi.SearchPhasesByTitle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPhasesByTitle`: []Phase
    fmt.Fprintf(os.Stdout, "Response from `PhaseApi.SearchPhasesByTitle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPhasesByTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phaseTitle** | **string** |  | 
 **releaseId** | **string** |  | 

### Return type

[**[]Phase**](Phase.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhase

> Phase UpdatePhase(ctx, phaseId).Phase(phase).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    phaseId := "phaseId_example" // string | 
    phase := *openapiclient.NewPhase() // Phase |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhaseApi.UpdatePhase(context.Background(), phaseId).Phase(phase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhaseApi.UpdatePhase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePhase`: Phase
    fmt.Fprintf(os.Stdout, "Response from `PhaseApi.UpdatePhase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phase** | [**Phase**](Phase.md) |  | 

### Return type

[**Phase**](Phase.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

