# \ReleasesApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReleasesReleaseIdVariableValuesGet**](ReleasesApi.md#ReleasesReleaseIdVariableValuesGet) | **Get** /releases/{releaseId}/variableValues | 
[**ReleasesReleaseIdVariablesGet**](ReleasesApi.md#ReleasesReleaseIdVariablesGet) | **Get** /releases/{releaseId}/variables | 
[**ReleasesReleaseIdVariablesPost**](ReleasesApi.md#ReleasesReleaseIdVariablesPost) | **Post** /releases/{releaseId}/variables | 
[**ReleasesReleaseIdVariablesPut**](ReleasesApi.md#ReleasesReleaseIdVariablesPut) | **Put** /releases/{releaseId}/variables | 



## ReleasesReleaseIdVariableValuesGet

> map[string]string ReleasesReleaseIdVariableValuesGet(ctx, releaseId).Execute()



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
    releaseId := "releaseId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesApi.ReleasesReleaseIdVariableValuesGet(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.ReleasesReleaseIdVariableValuesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleasesReleaseIdVariableValuesGet`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `ReleasesApi.ReleasesReleaseIdVariableValuesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleasesReleaseIdVariableValuesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleasesReleaseIdVariablesGet

> []Variable ReleasesReleaseIdVariablesGet(ctx, releaseId).Execute()



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
    releaseId := "releaseId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesApi.ReleasesReleaseIdVariablesGet(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.ReleasesReleaseIdVariablesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleasesReleaseIdVariablesGet`: []Variable
    fmt.Fprintf(os.Stdout, "Response from `ReleasesApi.ReleasesReleaseIdVariablesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleasesReleaseIdVariablesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleasesReleaseIdVariablesPost

> Variable ReleasesReleaseIdVariablesPost(ctx, releaseId).Variable1(variable1).Execute()



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
    releaseId := "releaseId_example" // string | 
    variable1 := *openapiclient.NewVariable1() // Variable1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesApi.ReleasesReleaseIdVariablesPost(context.Background(), releaseId).Variable1(variable1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.ReleasesReleaseIdVariablesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleasesReleaseIdVariablesPost`: Variable
    fmt.Fprintf(os.Stdout, "Response from `ReleasesApi.ReleasesReleaseIdVariablesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleasesReleaseIdVariablesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variable1** | [**Variable1**](Variable1.md) |  | 

### Return type

[**Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleasesReleaseIdVariablesPut

> []Variable ReleasesReleaseIdVariablesPut(ctx, releaseId).Variable(variable).Execute()



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
    releaseId := "releaseId_example" // string | 
    variable := []openapiclient.Variable{*openapiclient.NewVariable()} // []Variable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesApi.ReleasesReleaseIdVariablesPut(context.Background(), releaseId).Variable(variable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.ReleasesReleaseIdVariablesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleasesReleaseIdVariablesPut`: []Variable
    fmt.Fprintf(os.Stdout, "Response from `ReleasesApi.ReleasesReleaseIdVariablesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleasesReleaseIdVariablesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variable** | [**[]Variable**](Variable.md) |  | 

### Return type

[**[]Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
