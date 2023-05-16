# \FacetApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFacet**](FacetApi.md#CreateFacet) | **Post** /api/v1/facets | 
[**DeleteFacet**](FacetApi.md#DeleteFacet) | **Delete** /api/v1/facets/{facetId} | 
[**GetFacet**](FacetApi.md#GetFacet) | **Get** /api/v1/facets/{facetId} | 
[**GetFacetTypes**](FacetApi.md#GetFacetTypes) | **Get** /api/v1/facets/types | 
[**SearchFacets**](FacetApi.md#SearchFacets) | **Post** /api/v1/facets/search | 
[**UpdateFacet**](FacetApi.md#UpdateFacet) | **Put** /api/v1/facets/{facetId} | 



## CreateFacet

> Facet CreateFacet(ctx).ConfigurationFacet(configurationFacet).Execute()



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
    configurationFacet := *openapiclient.NewConfigurationFacet() // ConfigurationFacet |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FacetApi.CreateFacet(context.Background()).ConfigurationFacet(configurationFacet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FacetApi.CreateFacet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFacet`: Facet
    fmt.Fprintf(os.Stdout, "Response from `FacetApi.CreateFacet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFacetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationFacet** | [**ConfigurationFacet**](ConfigurationFacet.md) |  | 

### Return type

[**Facet**](Facet.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFacet

> DeleteFacet(ctx, facetId).Execute()



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
    facetId := "facetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FacetApi.DeleteFacet(context.Background(), facetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FacetApi.DeleteFacet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**facetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFacetRequest struct via the builder pattern


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


## GetFacet

> Facet GetFacet(ctx, facetId).Execute()



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
    facetId := "facetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FacetApi.GetFacet(context.Background(), facetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FacetApi.GetFacet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFacet`: Facet
    fmt.Fprintf(os.Stdout, "Response from `FacetApi.GetFacet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**facetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFacetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Facet**](Facet.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFacetTypes

> []interface{} GetFacetTypes(ctx).BaseType(baseType).Execute()



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
    baseType := "baseType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FacetApi.GetFacetTypes(context.Background()).BaseType(baseType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FacetApi.GetFacetTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFacetTypes`: []interface{}
    fmt.Fprintf(os.Stdout, "Response from `FacetApi.GetFacetTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFacetTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseType** | **string** |  | 

### Return type

**[]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFacets

> []Facet SearchFacets(ctx).FacetFilters(facetFilters).Execute()



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
    facetFilters := *openapiclient.NewFacetFilters() // FacetFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FacetApi.SearchFacets(context.Background()).FacetFilters(facetFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FacetApi.SearchFacets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchFacets`: []Facet
    fmt.Fprintf(os.Stdout, "Response from `FacetApi.SearchFacets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFacetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facetFilters** | [**FacetFilters**](FacetFilters.md) |  | 

### Return type

[**[]Facet**](Facet.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFacet

> Facet UpdateFacet(ctx, facetId).ConfigurationFacet(configurationFacet).Execute()



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
    facetId := "facetId_example" // string | 
    configurationFacet := *openapiclient.NewConfigurationFacet() // ConfigurationFacet |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FacetApi.UpdateFacet(context.Background(), facetId).ConfigurationFacet(configurationFacet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FacetApi.UpdateFacet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFacet`: Facet
    fmt.Fprintf(os.Stdout, "Response from `FacetApi.UpdateFacet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**facetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFacetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationFacet** | [**ConfigurationFacet**](ConfigurationFacet.md) |  | 

### Return type

[**Facet**](Facet.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

