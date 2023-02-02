# \ApplicationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](ApplicationApi.md#CreateApplication) | **Post** /api/v1/applications | 
[**DeleteApplication**](ApplicationApi.md#DeleteApplication) | **Delete** /api/v1/applications/{applicationId} | 
[**GetApplication**](ApplicationApi.md#GetApplication) | **Get** /api/v1/applications/{applicationId} | 
[**SearchApplications**](ApplicationApi.md#SearchApplications) | **Post** /api/v1/applications/search | 
[**UpdateApplication**](ApplicationApi.md#UpdateApplication) | **Put** /api/v1/applications/{applicationId} | 



## CreateApplication

> ApplicationView CreateApplication(ctx).ApplicationForm(applicationForm).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationForm := *openapiclient.NewApplicationForm() // ApplicationForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationApi.CreateApplication(context.Background()).ApplicationForm(applicationForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `ApplicationApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationForm** | [**ApplicationForm**](ApplicationForm.md) |  | 

### Return type

[**ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, applicationId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := "applicationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationApi.DeleteApplication(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> ApplicationView GetApplication(ctx, applicationId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := "applicationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationApi.GetApplication(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `ApplicationApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApplications

> []ApplicationView SearchApplications(ctx).ApplicationFilters(applicationFilters).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationFilters := *openapiclient.NewApplicationFilters() // ApplicationFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationApi.SearchApplications(context.Background()).ApplicationFilters(applicationFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.SearchApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchApplications`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `ApplicationApi.SearchApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationFilters** | [**ApplicationFilters**](ApplicationFilters.md) |  | 

### Return type

[**[]ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> ApplicationView UpdateApplication(ctx, applicationId).ApplicationForm(applicationForm).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := "applicationId_example" // string | 
    applicationForm := *openapiclient.NewApplicationForm() // ApplicationForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationApi.UpdateApplication(context.Background(), applicationId).ApplicationForm(applicationForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `ApplicationApi.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationForm** | [**ApplicationForm**](ApplicationForm.md) |  | 

### Return type

[**ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

