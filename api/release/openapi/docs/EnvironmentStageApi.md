# \EnvironmentStageApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStage3**](EnvironmentStageApi.md#CreateStage3) | **Post** /api/v1/environments/stages | 
[**DeleteEnvironmentStage**](EnvironmentStageApi.md#DeleteEnvironmentStage) | **Delete** /api/v1/environments/stages/{environmentStageId} | 
[**GetStageById**](EnvironmentStageApi.md#GetStageById) | **Get** /api/v1/environments/stages/{environmentStageId} | 
[**SearchStages**](EnvironmentStageApi.md#SearchStages) | **Post** /api/v1/environments/stages/search | 
[**UpdateStageInEnvironment**](EnvironmentStageApi.md#UpdateStageInEnvironment) | **Put** /api/v1/environments/stages/{environmentStageId} | 



## CreateStage3

> EnvironmentStageView CreateStage3(ctx).EnvironmentStageForm(environmentStageForm).Execute()



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
    environmentStageForm := *openapiclient.NewEnvironmentStageForm() // EnvironmentStageForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentStageApi.CreateStage3(context.Background()).EnvironmentStageForm(environmentStageForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentStageApi.CreateStage3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStage3`: EnvironmentStageView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentStageApi.CreateStage3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStage3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentStageForm** | [**EnvironmentStageForm**](EnvironmentStageForm.md) |  | 

### Return type

[**EnvironmentStageView**](EnvironmentStageView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentStage

> DeleteEnvironmentStage(ctx, environmentStageId).Execute()



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
    environmentStageId := "environmentStageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnvironmentStageApi.DeleteEnvironmentStage(context.Background(), environmentStageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentStageApi.DeleteEnvironmentStage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentStageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentStageRequest struct via the builder pattern


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


## GetStageById

> EnvironmentStageView GetStageById(ctx, environmentStageId).Execute()



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
    environmentStageId := "environmentStageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentStageApi.GetStageById(context.Background(), environmentStageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentStageApi.GetStageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStageById`: EnvironmentStageView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentStageApi.GetStageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentStageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentStageView**](EnvironmentStageView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStages

> []EnvironmentStageView SearchStages(ctx).EnvironmentStageFilters(environmentStageFilters).Execute()



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
    environmentStageFilters := *openapiclient.NewEnvironmentStageFilters() // EnvironmentStageFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentStageApi.SearchStages(context.Background()).EnvironmentStageFilters(environmentStageFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentStageApi.SearchStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchStages`: []EnvironmentStageView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentStageApi.SearchStages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentStageFilters** | [**EnvironmentStageFilters**](EnvironmentStageFilters.md) |  | 

### Return type

[**[]EnvironmentStageView**](EnvironmentStageView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStageInEnvironment

> EnvironmentStageView UpdateStageInEnvironment(ctx, environmentStageId).EnvironmentStageForm(environmentStageForm).Execute()



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
    environmentStageId := "environmentStageId_example" // string | 
    environmentStageForm := *openapiclient.NewEnvironmentStageForm() // EnvironmentStageForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentStageApi.UpdateStageInEnvironment(context.Background(), environmentStageId).EnvironmentStageForm(environmentStageForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentStageApi.UpdateStageInEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStageInEnvironment`: EnvironmentStageView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentStageApi.UpdateStageInEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentStageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStageInEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentStageForm** | [**EnvironmentStageForm**](EnvironmentStageForm.md) |  | 

### Return type

[**EnvironmentStageView**](EnvironmentStageView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

