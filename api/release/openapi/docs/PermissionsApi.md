# \PermissionsApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalPermissions**](PermissionsApi.md#GetGlobalPermissions) | **Get** /api/v1/global-permissions | 



## GetGlobalPermissions

> []string GetGlobalPermissions(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.GetGlobalPermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.GetGlobalPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalPermissions`: []string
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.GetGlobalPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalPermissionsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

