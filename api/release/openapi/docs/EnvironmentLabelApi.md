# \EnvironmentLabelApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLabel**](EnvironmentLabelApi.md#CreateLabel) | **Post** /api/v1/environments/labels | 
[**DeleteEnvironmentLabel**](EnvironmentLabelApi.md#DeleteEnvironmentLabel) | **Delete** /api/v1/environments/labels/{environmentLabelId} | 
[**GetLabelById**](EnvironmentLabelApi.md#GetLabelById) | **Get** /api/v1/environments/labels/{environmentLabelId} | 
[**SearchLabels**](EnvironmentLabelApi.md#SearchLabels) | **Post** /api/v1/environments/labels/search | 
[**UpdateLabel**](EnvironmentLabelApi.md#UpdateLabel) | **Put** /api/v1/environments/labels/{environmentLabelId} | 



## CreateLabel

> EnvironmentLabelView CreateLabel(ctx).EnvironmentLabelForm(environmentLabelForm).Execute()



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
    environmentLabelForm := *openapiclient.NewEnvironmentLabelForm() // EnvironmentLabelForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentLabelApi.CreateLabel(context.Background()).EnvironmentLabelForm(environmentLabelForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentLabelApi.CreateLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLabel`: EnvironmentLabelView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentLabelApi.CreateLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentLabelForm** | [**EnvironmentLabelForm**](EnvironmentLabelForm.md) |  | 

### Return type

[**EnvironmentLabelView**](EnvironmentLabelView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentLabel

> DeleteEnvironmentLabel(ctx, environmentLabelId).Execute()



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
    environmentLabelId := "environmentLabelId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnvironmentLabelApi.DeleteEnvironmentLabel(context.Background(), environmentLabelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentLabelApi.DeleteEnvironmentLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentLabelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentLabelRequest struct via the builder pattern


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


## GetLabelById

> EnvironmentLabelView GetLabelById(ctx, environmentLabelId).Execute()



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
    environmentLabelId := "environmentLabelId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentLabelApi.GetLabelById(context.Background(), environmentLabelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentLabelApi.GetLabelById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabelById`: EnvironmentLabelView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentLabelApi.GetLabelById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentLabelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentLabelView**](EnvironmentLabelView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLabels

> []EnvironmentLabelView SearchLabels(ctx).EnvironmentLabelFilters(environmentLabelFilters).Execute()



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
    environmentLabelFilters := *openapiclient.NewEnvironmentLabelFilters() // EnvironmentLabelFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentLabelApi.SearchLabels(context.Background()).EnvironmentLabelFilters(environmentLabelFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentLabelApi.SearchLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchLabels`: []EnvironmentLabelView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentLabelApi.SearchLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentLabelFilters** | [**EnvironmentLabelFilters**](EnvironmentLabelFilters.md) |  | 

### Return type

[**[]EnvironmentLabelView**](EnvironmentLabelView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLabel

> EnvironmentLabelView UpdateLabel(ctx, environmentLabelId).EnvironmentLabelForm(environmentLabelForm).Execute()



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
    environmentLabelId := "environmentLabelId_example" // string | 
    environmentLabelForm := *openapiclient.NewEnvironmentLabelForm() // EnvironmentLabelForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentLabelApi.UpdateLabel(context.Background(), environmentLabelId).EnvironmentLabelForm(environmentLabelForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentLabelApi.UpdateLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLabel`: EnvironmentLabelView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentLabelApi.UpdateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentLabelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentLabelForm** | [**EnvironmentLabelForm**](EnvironmentLabelForm.md) |  | 

### Return type

[**EnvironmentLabelView**](EnvironmentLabelView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

