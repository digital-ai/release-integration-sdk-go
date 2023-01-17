# \ReleaseApiApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVariable**](ReleaseApiApi.md#DeleteVariable) | **Delete** /releases/{variableId} | 
[**GetVariablePossibleValues**](ReleaseApiApi.md#GetVariablePossibleValues) | **Get** /releases/{variableId}/possibleValues | 
[**IsVariableUsed**](ReleaseApiApi.md#IsVariableUsed) | **Get** /releases/{variableId}/used | 
[**ReplaceVariable**](ReleaseApiApi.md#ReplaceVariable) | **Post** /releases/{variableId}/replace | 



## DeleteVariable

> DeleteVariable(ctx, variableId).Execute()



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
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApiApi.DeleteVariable(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApiApi.DeleteVariable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteVariableRequest struct via the builder pattern


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


## GetVariablePossibleValues

> []map[string]interface{} GetVariablePossibleValues(ctx, variableId).Execute()



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
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApiApi.GetVariablePossibleValues(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApiApi.GetVariablePossibleValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariablePossibleValues`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApiApi.GetVariablePossibleValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariablePossibleValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsVariableUsed

> bool IsVariableUsed(ctx, variableId).Execute()



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
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApiApi.IsVariableUsed(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApiApi.IsVariableUsed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsVariableUsed`: bool
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApiApi.IsVariableUsed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsVariableUsedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceVariable

> ReplaceVariable(ctx, variableId).VariableOrValue(variableOrValue).Execute()



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
    variableId := "variableId_example" // string | 
    variableOrValue := *openapiclient.NewVariableOrValue() // VariableOrValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApiApi.ReplaceVariable(context.Background(), variableId).VariableOrValue(variableOrValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApiApi.ReplaceVariable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReplaceVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableOrValue** | [**VariableOrValue**](VariableOrValue.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

