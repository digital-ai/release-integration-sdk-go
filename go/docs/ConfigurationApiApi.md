# \ConfigurationApiApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGlobalVariable**](ConfigurationApiApi.md#AddGlobalVariable) | **Post** /config/Configuration/variables/global | 
[**GetGlobalVariableValues**](ConfigurationApiApi.md#GetGlobalVariableValues) | **Get** /config/Configuration/variableValues/global | 



## AddGlobalVariable

> Variable AddGlobalVariable(ctx).Variable1(variable1).Execute()



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
    variable1 := *openapiclient.NewVariable1() // Variable1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApiApi.AddGlobalVariable(context.Background()).Variable1(variable1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApiApi.AddGlobalVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGlobalVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApiApi.AddGlobalVariable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGlobalVariableRequest struct via the builder pattern


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


## GetGlobalVariableValues

> map[string]string GetGlobalVariableValues(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApiApi.GetGlobalVariableValues(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApiApi.GetGlobalVariableValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalVariableValues`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApiApi.GetGlobalVariableValues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalVariableValuesRequest struct via the builder pattern


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

