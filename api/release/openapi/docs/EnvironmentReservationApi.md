# \EnvironmentReservationApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplication**](EnvironmentReservationApi.md#AddApplication) | **Post** /api/v1/environments/reservations/{environmentReservationId} | 
[**CreateReservation**](EnvironmentReservationApi.md#CreateReservation) | **Post** /api/v1/environments/reservations | 
[**DeleteEnvironmentReservation**](EnvironmentReservationApi.md#DeleteEnvironmentReservation) | **Delete** /api/v1/environments/reservations/{environmentReservationId} | 
[**GetReservation**](EnvironmentReservationApi.md#GetReservation) | **Get** /api/v1/environments/reservations/{environmentReservationId} | 
[**SearchReservations**](EnvironmentReservationApi.md#SearchReservations) | **Post** /api/v1/environments/reservations/search | 
[**UpdateReservation**](EnvironmentReservationApi.md#UpdateReservation) | **Put** /api/v1/environments/reservations/{environmentReservationId} | 



## AddApplication

> AddApplication(ctx, environmentReservationId).ApplicationId(applicationId).Execute()



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
    environmentReservationId := "environmentReservationId_example" // string | 
    applicationId := "applicationId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnvironmentReservationApi.AddApplication(context.Background(), environmentReservationId).ApplicationId(applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReservationApi.AddApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentReservationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationId** | **string** |  | 

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


## CreateReservation

> EnvironmentReservationView CreateReservation(ctx).EnvironmentReservationForm(environmentReservationForm).Execute()



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
    environmentReservationForm := *openapiclient.NewEnvironmentReservationForm() // EnvironmentReservationForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentReservationApi.CreateReservation(context.Background()).EnvironmentReservationForm(environmentReservationForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReservationApi.CreateReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReservation`: EnvironmentReservationView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentReservationApi.CreateReservation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentReservationForm** | [**EnvironmentReservationForm**](EnvironmentReservationForm.md) |  | 

### Return type

[**EnvironmentReservationView**](EnvironmentReservationView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentReservation

> DeleteEnvironmentReservation(ctx, environmentReservationId).Execute()



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
    environmentReservationId := "environmentReservationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnvironmentReservationApi.DeleteEnvironmentReservation(context.Background(), environmentReservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReservationApi.DeleteEnvironmentReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentReservationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentReservationRequest struct via the builder pattern


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


## GetReservation

> EnvironmentReservationView GetReservation(ctx, environmentReservationId).Execute()



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
    environmentReservationId := "environmentReservationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentReservationApi.GetReservation(context.Background(), environmentReservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReservationApi.GetReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReservation`: EnvironmentReservationView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentReservationApi.GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentReservationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentReservationView**](EnvironmentReservationView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReservations

> []EnvironmentReservationSearchView SearchReservations(ctx).ReservationFilters(reservationFilters).Execute()



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
    reservationFilters := *openapiclient.NewReservationFilters() // ReservationFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentReservationApi.SearchReservations(context.Background()).ReservationFilters(reservationFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReservationApi.SearchReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchReservations`: []EnvironmentReservationSearchView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentReservationApi.SearchReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reservationFilters** | [**ReservationFilters**](ReservationFilters.md) |  | 

### Return type

[**[]EnvironmentReservationSearchView**](EnvironmentReservationSearchView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReservation

> EnvironmentReservationView UpdateReservation(ctx, environmentReservationId).EnvironmentReservationForm(environmentReservationForm).Execute()



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
    environmentReservationId := "environmentReservationId_example" // string | 
    environmentReservationForm := *openapiclient.NewEnvironmentReservationForm() // EnvironmentReservationForm |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentReservationApi.UpdateReservation(context.Background(), environmentReservationId).EnvironmentReservationForm(environmentReservationForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReservationApi.UpdateReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReservation`: EnvironmentReservationView
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentReservationApi.UpdateReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentReservationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentReservationForm** | [**EnvironmentReservationForm**](EnvironmentReservationForm.md) |  | 

### Return type

[**EnvironmentReservationView**](EnvironmentReservationView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

