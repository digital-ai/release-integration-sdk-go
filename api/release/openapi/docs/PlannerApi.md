# \PlannerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveReleases**](PlannerApi.md#GetActiveReleases) | **Get** /api/v1/analytics/planner/active | 
[**GetCompletedReleases**](PlannerApi.md#GetCompletedReleases) | **Get** /api/v1/analytics/planner/completed | 
[**GetReleasesByIds**](PlannerApi.md#GetReleasesByIds) | **Post** /api/v1/analytics/planner/byIds | 



## GetActiveReleases

> []ProjectedRelease GetActiveReleases(ctx).Page(page).ResultsPerPage(resultsPerPage).Execute()



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
    page := int64(789) // int64 |  (optional) (default to 0)
    resultsPerPage := int64(789) // int64 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlannerApi.GetActiveReleases(context.Background()).Page(page).ResultsPerPage(resultsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.GetActiveReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveReleases`: []ProjectedRelease
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.GetActiveReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **resultsPerPage** | **int64** |  | [default to 100]

### Return type

[**[]ProjectedRelease**](ProjectedRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompletedReleases

> []ProjectedRelease GetCompletedReleases(ctx).Page(page).ResultsPerPage(resultsPerPage).Since(since).Execute()



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
    page := int64(789) // int64 |  (optional) (default to 0)
    resultsPerPage := int64(789) // int64 |  (optional) (default to 100)
    since := int64(789) // int64 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlannerApi.GetCompletedReleases(context.Background()).Page(page).ResultsPerPage(resultsPerPage).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.GetCompletedReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompletedReleases`: []ProjectedRelease
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.GetCompletedReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompletedReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **resultsPerPage** | **int64** |  | [default to 100]
 **since** | **int64** |  | [default to 0]

### Return type

[**[]ProjectedRelease**](ProjectedRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleasesByIds

> []ProjectedRelease GetReleasesByIds(ctx).RequestBody(requestBody).Execute()



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
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlannerApi.GetReleasesByIds(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.GetReleasesByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleasesByIds`: []ProjectedRelease
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.GetReleasesByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasesByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**[]ProjectedRelease**](ProjectedRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

