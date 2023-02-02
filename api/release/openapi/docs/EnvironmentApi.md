# \EnvironmentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](EnvironmentApi.md#CreateEnvironment) | **Post** /api/v1/environments | 
[**DeleteEnvironment**](EnvironmentApi.md#DeleteEnvironment) | **Delete** /api/v1/environments/{environmentId} | 
[**GetDeployableApplicationsForEnvironment**](EnvironmentApi.md#GetDeployableApplicationsForEnvironment) | **Get** /api/v1/environments/{environmentId}/applications | 
[**GetEnvironment**](EnvironmentApi.md#GetEnvironment) | **Get** /api/v1/environments/{environmentId} | 
[**GetReservationsForEnvironment**](EnvironmentApi.md#GetReservationsForEnvironment) | **Get** /api/v1/environments/{environmentId}/reservations | 
[**SearchEnvironments**](EnvironmentApi.md#SearchEnvironments) | **Post** /api/v1/environments/search | 
[**UpdateEnvironment**](EnvironmentApi.md#UpdateEnvironment) | **Put** /api/v1/environments/{environmentId} | 



## CreateEnvironment

> EnvironmentView CreateEnvironment(ctx).EnvironmentForm(environmentForm).Execute()



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
    environmentForm := *openapiclient.NewEnvironmentForm() // EnvironmentForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.CreateEnvironment(context.Background()).EnvironmentForm(environmentForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.CreateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironment`: EnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentForm** | [**EnvironmentForm**](EnvironmentForm.md) |  | 

### Return type

[**EnvironmentView**](EnvironmentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, environmentId).Execute()



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
    environmentId := "environmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.DeleteEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.DeleteEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


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


## GetDeployableApplicationsForEnvironment

> []BaseApplicationView GetDeployableApplicationsForEnvironment(ctx, environmentId).Execute()



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
    environmentId := "environmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.GetDeployableApplicationsForEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.GetDeployableApplicationsForEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeployableApplicationsForEnvironment`: []BaseApplicationView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.GetDeployableApplicationsForEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployableApplicationsForEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BaseApplicationView**](BaseApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironment

> EnvironmentView GetEnvironment(ctx, environmentId).Execute()



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
    environmentId := "environmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.GetEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.GetEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironment`: EnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentView**](EnvironmentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReservationsForEnvironment

> []EnvironmentReservationView GetReservationsForEnvironment(ctx, environmentId).Execute()



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
    environmentId := "environmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.GetReservationsForEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.GetReservationsForEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReservationsForEnvironment`: []EnvironmentReservationView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.GetReservationsForEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReservationsForEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EnvironmentReservationView**](EnvironmentReservationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEnvironments

> []EnvironmentView SearchEnvironments(ctx).EnvironmentFilters(environmentFilters).Execute()



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
    environmentFilters := *openapiclient.NewEnvironmentFilters() // EnvironmentFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.SearchEnvironments(context.Background()).EnvironmentFilters(environmentFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.SearchEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchEnvironments`: []EnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.SearchEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentFilters** | [**EnvironmentFilters**](EnvironmentFilters.md) |  | 

### Return type

[**[]EnvironmentView**](EnvironmentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironment

> EnvironmentView UpdateEnvironment(ctx, environmentId).EnvironmentForm(environmentForm).Execute()



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
    environmentId := "environmentId_example" // string | 
    environmentForm := *openapiclient.NewEnvironmentForm() // EnvironmentForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.UpdateEnvironment(context.Background(), environmentId).EnvironmentForm(environmentForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.UpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironment`: EnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.UpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentForm** | [**EnvironmentForm**](EnvironmentForm.md) |  | 

### Return type

[**EnvironmentView**](EnvironmentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

