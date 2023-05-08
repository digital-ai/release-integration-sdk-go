# \ReleaseApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Abort**](ReleaseApi.md#Abort) | **Post** /api/v1/releases/{releaseId}/abort | 
[**CountReleases**](ReleaseApi.md#CountReleases) | **Post** /api/v1/releases/count | 
[**CreateReleaseVariable**](ReleaseApi.md#CreateReleaseVariable) | **Post** /api/v1/releases/{releaseId}/variables | 
[**DeleteRelease**](ReleaseApi.md#DeleteRelease) | **Delete** /api/v1/releases/{releaseId} | 
[**DeleteReleaseVariable**](ReleaseApi.md#DeleteReleaseVariable) | **Delete** /api/v1/releases/{variableId} | 
[**DownloadAttachment**](ReleaseApi.md#DownloadAttachment) | **Get** /api/v1/releases/attachments/{attachmentId} | 
[**FullSearchReleases**](ReleaseApi.md#FullSearchReleases) | **Post** /api/v1/releases/fullSearch | 
[**GetActiveTasks**](ReleaseApi.md#GetActiveTasks) | **Get** /api/v1/releases/{releaseId}/active-tasks | 
[**GetArchivedRelease**](ReleaseApi.md#GetArchivedRelease) | **Get** /api/v1/releases/archived/{releaseId} | 
[**GetPossibleReleaseVariableValues**](ReleaseApi.md#GetPossibleReleaseVariableValues) | **Get** /api/v1/releases/{variableId}/possibleValues | 
[**GetRelease**](ReleaseApi.md#GetRelease) | **Get** /api/v1/releases/{releaseId} | 
[**GetReleasePermissions**](ReleaseApi.md#GetReleasePermissions) | **Get** /api/v1/releases/permissions | 
[**GetReleaseTeams**](ReleaseApi.md#GetReleaseTeams) | **Get** /api/v1/releases/{releaseId}/teams | 
[**GetReleaseVariable**](ReleaseApi.md#GetReleaseVariable) | **Get** /api/v1/releases/{variableId} | 
[**GetReleaseVariables**](ReleaseApi.md#GetReleaseVariables) | **Get** /api/v1/releases/{releaseId}/variables | 
[**GetReleases**](ReleaseApi.md#GetReleases) | **Get** /api/v1/releases | 
[**GetVariableValues**](ReleaseApi.md#GetVariableValues) | **Get** /api/v1/releases/{releaseId}/variableValues | 
[**IsVariableUsedRelease**](ReleaseApi.md#IsVariableUsedRelease) | **Get** /api/v1/releases/{variableId}/used | 
[**ReplaceReleaseVariables**](ReleaseApi.md#ReplaceReleaseVariables) | **Post** /api/v1/releases/{variableId}/replace | 
[**RestartPhases**](ReleaseApi.md#RestartPhases) | **Post** /api/v1/releases/{releaseId}/restart | 
[**Resume**](ReleaseApi.md#Resume) | **Post** /api/v1/releases/{releaseId}/resume | 
[**SearchReleasesByTitle**](ReleaseApi.md#SearchReleasesByTitle) | **Get** /api/v1/releases/byTitle | 
[**SearchReleasesRelease**](ReleaseApi.md#SearchReleasesRelease) | **Post** /api/v1/releases/search | 
[**SetReleaseTeams**](ReleaseApi.md#SetReleaseTeams) | **Post** /api/v1/releases/{releaseId}/teams | 
[**StartRelease**](ReleaseApi.md#StartRelease) | **Post** /api/v1/releases/{releaseId}/start | 
[**UpdateRelease**](ReleaseApi.md#UpdateRelease) | **Put** /api/v1/releases/{releaseId} | 
[**UpdateReleaseVariable**](ReleaseApi.md#UpdateReleaseVariable) | **Put** /api/v1/releases/{variableId} | 
[**UpdateReleaseVariables**](ReleaseApi.md#UpdateReleaseVariables) | **Put** /api/v1/releases/{releaseId}/variables | 



## Abort

> Release Abort(ctx, releaseId).AbortRelease(abortRelease).Execute()



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
    abortRelease := *openapiclient.NewAbortRelease() // AbortRelease |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.Abort(context.Background(), releaseId).AbortRelease(abortRelease).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.Abort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Abort`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.Abort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **abortRelease** | [**AbortRelease**](AbortRelease.md) |  | 

### Return type

[**Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountReleases

> ReleaseCountResults CountReleases(ctx).ReleasesFilters(releasesFilters).Execute()



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
    releasesFilters := *openapiclient.NewReleasesFilters() // ReleasesFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.CountReleases(context.Background()).ReleasesFilters(releasesFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.CountReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountReleases`: ReleaseCountResults
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.CountReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releasesFilters** | [**ReleasesFilters**](ReleasesFilters.md) |  | 

### Return type

[**ReleaseCountResults**](ReleaseCountResults.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReleaseVariable

> Variable CreateReleaseVariable(ctx, releaseId).Variable1(variable1).Execute()



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
    variable1 := *openapiclient.NewVariable1() // Variable1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.CreateReleaseVariable(context.Background(), releaseId).Variable1(variable1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.CreateReleaseVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReleaseVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.CreateReleaseVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variable1** | [**Variable1**](Variable1.md) |  | 

### Return type

[**Variable**](Variable.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRelease

> DeleteRelease(ctx, releaseId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReleaseApi.DeleteRelease(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.DeleteRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseRequest struct via the builder pattern


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


## DeleteReleaseVariable

> DeleteReleaseVariable(ctx, variableId).Execute()



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
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReleaseApi.DeleteReleaseVariable(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.DeleteReleaseVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseVariableRequest struct via the builder pattern


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


## DownloadAttachment

> DownloadAttachment(ctx, attachmentId).Execute()



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
    attachmentId := "attachmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReleaseApi.DownloadAttachment(context.Background(), attachmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.DownloadAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentRequest struct via the builder pattern


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


## FullSearchReleases

> ReleaseFullSearchResult FullSearchReleases(ctx).ArchivePage(archivePage).ArchiveResultsPerPage(archiveResultsPerPage).Page(page).ResultsPerPage(resultsPerPage).ReleasesFilters(releasesFilters).Execute()



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
    archivePage := int64(789) // int64 |  (optional)
    archiveResultsPerPage := int64(789) // int64 |  (optional)
    page := int64(789) // int64 |  (optional)
    resultsPerPage := int64(789) // int64 |  (optional)
    releasesFilters := *openapiclient.NewReleasesFilters() // ReleasesFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.FullSearchReleases(context.Background()).ArchivePage(archivePage).ArchiveResultsPerPage(archiveResultsPerPage).Page(page).ResultsPerPage(resultsPerPage).ReleasesFilters(releasesFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.FullSearchReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullSearchReleases`: ReleaseFullSearchResult
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.FullSearchReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFullSearchReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archivePage** | **int64** |  | 
 **archiveResultsPerPage** | **int64** |  | 
 **page** | **int64** |  | 
 **resultsPerPage** | **int64** |  | 
 **releasesFilters** | [**ReleasesFilters**](ReleasesFilters.md) |  | 

### Return type

[**ReleaseFullSearchResult**](ReleaseFullSearchResult.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveTasks

> []Task GetActiveTasks(ctx, releaseId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetActiveTasks(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetActiveTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveTasks`: []Task
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetActiveTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Task**](Task.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchivedRelease

> Release GetArchivedRelease(ctx, releaseId).RoleIds(roleIds).Execute()



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
    roleIds := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetArchivedRelease(context.Background(), releaseId).RoleIds(roleIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetArchivedRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchivedRelease`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetArchivedRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivedReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleIds** | **bool** |  | [default to false]

### Return type

[**Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPossibleReleaseVariableValues

> []map[string]interface{} GetPossibleReleaseVariableValues(ctx, variableId).Execute()



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
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetPossibleReleaseVariableValues(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetPossibleReleaseVariableValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPossibleReleaseVariableValues`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetPossibleReleaseVariableValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPossibleReleaseVariableValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelease

> Release GetRelease(ctx, releaseId).RoleIds(roleIds).Execute()



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
    roleIds := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetRelease(context.Background(), releaseId).RoleIds(roleIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelease`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleIds** | **bool** |  | [default to false]

### Return type

[**Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleasePermissions

> []string GetReleasePermissions(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetReleasePermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetReleasePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleasePermissions`: []string
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetReleasePermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasePermissionsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseTeams

> []TeamView GetReleaseTeams(ctx, releaseId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetReleaseTeams(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetReleaseTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleaseTeams`: []TeamView
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetReleaseTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TeamView**](TeamView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseVariable

> Variable GetReleaseVariable(ctx, variableId).Execute()



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
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetReleaseVariable(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetReleaseVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleaseVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetReleaseVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Variable**](Variable.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseVariables

> []Variable GetReleaseVariables(ctx, releaseId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetReleaseVariables(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetReleaseVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleaseVariables`: []Variable
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetReleaseVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Variable**](Variable.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleases

> []Release GetReleases(ctx).Depth(depth).Page(page).ResultsPerPage(resultsPerPage).Execute()



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
    depth := int32(56) // int32 |  (optional) (default to 1)
    page := int64(789) // int64 |  (optional) (default to 0)
    resultsPerPage := int64(789) // int64 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetReleases(context.Background()).Depth(depth).Page(page).ResultsPerPage(resultsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleases`: []Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depth** | **int32** |  | [default to 1]
 **page** | **int64** |  | [default to 0]
 **resultsPerPage** | **int64** |  | [default to 100]

### Return type

[**[]Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariableValues

> map[string]string GetVariableValues(ctx, releaseId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.GetVariableValues(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.GetVariableValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariableValues`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.GetVariableValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsVariableUsedRelease

> bool IsVariableUsedRelease(ctx, variableId).Execute()



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
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.IsVariableUsedRelease(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.IsVariableUsedRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsVariableUsedRelease`: bool
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.IsVariableUsedRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsVariableUsedReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceReleaseVariables

> ReplaceReleaseVariables(ctx, variableId).VariableOrValue(variableOrValue).Execute()



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
    variableId := "variableId_example" // string | 
    variableOrValue := *openapiclient.NewVariableOrValue() // VariableOrValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReleaseApi.ReplaceReleaseVariables(context.Background(), variableId).VariableOrValue(variableOrValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.ReplaceReleaseVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceReleaseVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableOrValue** | [**VariableOrValue**](VariableOrValue.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartPhases

> Release RestartPhases(ctx, releaseId).FromPhaseId(fromPhaseId).FromTaskId(fromTaskId).PhaseVersion(phaseVersion).Resume(resume).Execute()



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
    fromPhaseId := "fromPhaseId_example" // string |  (optional)
    fromTaskId := "fromTaskId_example" // string |  (optional)
    phaseVersion := openapiclient.PhaseVersion("LATEST") // PhaseVersion |  (optional)
    resume := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.RestartPhases(context.Background(), releaseId).FromPhaseId(fromPhaseId).FromTaskId(fromTaskId).PhaseVersion(phaseVersion).Resume(resume).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.RestartPhases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartPhases`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.RestartPhases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartPhasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromPhaseId** | **string** |  | 
 **fromTaskId** | **string** |  | 
 **phaseVersion** | [**PhaseVersion**](PhaseVersion.md) |  | 
 **resume** | **bool** |  | 

### Return type

[**Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Resume

> Release Resume(ctx, releaseId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.Resume(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.Resume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Resume`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.Resume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReleasesByTitle

> []Release SearchReleasesByTitle(ctx).ReleaseTitle(releaseTitle).Execute()



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
    releaseTitle := "releaseTitle_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.SearchReleasesByTitle(context.Background()).ReleaseTitle(releaseTitle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.SearchReleasesByTitle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchReleasesByTitle`: []Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.SearchReleasesByTitle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchReleasesByTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseTitle** | **string** |  | 

### Return type

[**[]Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReleasesRelease

> []Release SearchReleasesRelease(ctx).Page(page).PageIsOffset(pageIsOffset).ResultsPerPage(resultsPerPage).ReleasesFilters(releasesFilters).Execute()



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
    page := int64(789) // int64 |  (optional) (default to 0)
    pageIsOffset := true // bool |  (optional) (default to false)
    resultsPerPage := int64(789) // int64 |  (optional) (default to 100)
    releasesFilters := *openapiclient.NewReleasesFilters() // ReleasesFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.SearchReleasesRelease(context.Background()).Page(page).PageIsOffset(pageIsOffset).ResultsPerPage(resultsPerPage).ReleasesFilters(releasesFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.SearchReleasesRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchReleasesRelease`: []Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.SearchReleasesRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchReleasesReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageIsOffset** | **bool** |  | [default to false]
 **resultsPerPage** | **int64** |  | [default to 100]
 **releasesFilters** | [**ReleasesFilters**](ReleasesFilters.md) |  | 

### Return type

[**[]Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReleaseTeams

> []TeamView SetReleaseTeams(ctx, releaseId).TeamView(teamView).Execute()



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
    teamView := []openapiclient.TeamView{*openapiclient.NewTeamView()} // []TeamView |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.SetReleaseTeams(context.Background(), releaseId).TeamView(teamView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.SetReleaseTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetReleaseTeams`: []TeamView
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.SetReleaseTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetReleaseTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamView** | [**[]TeamView**](TeamView.md) |  | 

### Return type

[**[]TeamView**](TeamView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartRelease

> Release StartRelease(ctx, releaseId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.StartRelease(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.StartRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartRelease`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.StartRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRelease

> Release UpdateRelease(ctx, releaseId).Release(release).Execute()



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
    release := *openapiclient.NewRelease() // Release |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.UpdateRelease(context.Background(), releaseId).Release(release).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.UpdateRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRelease`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.UpdateRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **release** | [**Release**](Release.md) |  | 

### Return type

[**Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleaseVariable

> Variable UpdateReleaseVariable(ctx, variableId).Variable(variable).Execute()



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
    variableId := "variableId_example" // string | 
    variable := *openapiclient.NewVariable() // Variable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.UpdateReleaseVariable(context.Background(), variableId).Variable(variable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.UpdateReleaseVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReleaseVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.UpdateReleaseVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variable** | [**Variable**](Variable.md) |  | 

### Return type

[**Variable**](Variable.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleaseVariables

> []Variable UpdateReleaseVariables(ctx, releaseId).Variable(variable).Execute()



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
    variable := []openapiclient.Variable{*openapiclient.NewVariable()} // []Variable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.UpdateReleaseVariables(context.Background(), releaseId).Variable(variable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.UpdateReleaseVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReleaseVariables`: []Variable
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.UpdateReleaseVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variable** | [**[]Variable**](Variable.md) |  | 

### Return type

[**[]Variable**](Variable.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

