# \TriggersApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrigger**](TriggersApi.md#AddTrigger) | **Post** /api/v1/triggers | 
[**DisableAllTriggers**](TriggersApi.md#DisableAllTriggers) | **Post** /api/v1/triggers/disable/all | 
[**DisableTrigger**](TriggersApi.md#DisableTrigger) | **Put** /api/v1/triggers/{triggerId}/disable | 
[**DisableTriggers**](TriggersApi.md#DisableTriggers) | **Post** /api/v1/triggers/disable | 
[**EnableAllTriggers**](TriggersApi.md#EnableAllTriggers) | **Post** /api/v1/triggers/enable/all | 
[**EnableTrigger**](TriggersApi.md#EnableTrigger) | **Put** /api/v1/triggers/{triggerId}/enable | 
[**EnableTriggers**](TriggersApi.md#EnableTriggers) | **Post** /api/v1/triggers/enable | 
[**GetTrigger**](TriggersApi.md#GetTrigger) | **Get** /api/v1/triggers/{triggerId} | 
[**GetTypes**](TriggersApi.md#GetTypes) | **Get** /api/v1/triggers/types | 
[**RemoveTrigger**](TriggersApi.md#RemoveTrigger) | **Delete** /api/v1/triggers/{triggerId} | 
[**RunTrigger**](TriggersApi.md#RunTrigger) | **Post** /api/v1/triggers/{triggerId}/run | 
[**SearchTriggers**](TriggersApi.md#SearchTriggers) | **Get** /api/v1/triggers | 
[**UpdateTrigger**](TriggersApi.md#UpdateTrigger) | **Put** /api/v1/triggers/{triggerId} | 



## AddTrigger

> Trigger AddTrigger(ctx).Trigger(trigger).Execute()



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
    trigger := *openapiclient.NewTrigger() // Trigger |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersApi.AddTrigger(context.Background()).Trigger(trigger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.AddTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrigger`: Trigger
    fmt.Fprintf(os.Stdout, "Response from `TriggersApi.AddTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trigger** | [**Trigger**](Trigger.md) |  | 

### Return type

[**Trigger**](Trigger.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableAllTriggers

> BulkActionResultView DisableAllTriggers(ctx).Execute()



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
    resp, r, err := apiClient.TriggersApi.DisableAllTriggers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.DisableAllTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableAllTriggers`: BulkActionResultView
    fmt.Fprintf(os.Stdout, "Response from `TriggersApi.DisableAllTriggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableAllTriggersRequest struct via the builder pattern


### Return type

[**BulkActionResultView**](BulkActionResultView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableTrigger

> DisableTrigger(ctx, triggerId).Execute()



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
    triggerId := "triggerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TriggersApi.DisableTrigger(context.Background(), triggerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.DisableTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableTriggerRequest struct via the builder pattern


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


## DisableTriggers

> BulkActionResultView DisableTriggers(ctx).RequestBody(requestBody).Execute()



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
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersApi.DisableTriggers(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.DisableTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableTriggers`: BulkActionResultView
    fmt.Fprintf(os.Stdout, "Response from `TriggersApi.DisableTriggers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**BulkActionResultView**](BulkActionResultView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableAllTriggers

> BulkActionResultView EnableAllTriggers(ctx).Execute()



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
    resp, r, err := apiClient.TriggersApi.EnableAllTriggers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.EnableAllTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableAllTriggers`: BulkActionResultView
    fmt.Fprintf(os.Stdout, "Response from `TriggersApi.EnableAllTriggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAllTriggersRequest struct via the builder pattern


### Return type

[**BulkActionResultView**](BulkActionResultView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableTrigger

> EnableTrigger(ctx, triggerId).Execute()



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
    triggerId := "triggerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TriggersApi.EnableTrigger(context.Background(), triggerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.EnableTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableTriggerRequest struct via the builder pattern


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


## EnableTriggers

> BulkActionResultView EnableTriggers(ctx).RequestBody(requestBody).Execute()



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
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersApi.EnableTriggers(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.EnableTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableTriggers`: BulkActionResultView
    fmt.Fprintf(os.Stdout, "Response from `TriggersApi.EnableTriggers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**BulkActionResultView**](BulkActionResultView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrigger

> Trigger GetTrigger(ctx, triggerId).Execute()



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
    triggerId := "triggerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersApi.GetTrigger(context.Background(), triggerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.GetTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrigger`: Trigger
    fmt.Fprintf(os.Stdout, "Response from `TriggersApi.GetTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Trigger**](Trigger.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTypes

> []interface{} GetTypes(ctx).Execute()



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
    resp, r, err := apiClient.TriggersApi.GetTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.GetTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTypes`: []interface{}
    fmt.Fprintf(os.Stdout, "Response from `TriggersApi.GetTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTypesRequest struct via the builder pattern


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


## RemoveTrigger

> RemoveTrigger(ctx, triggerId).Execute()



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
    triggerId := "triggerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TriggersApi.RemoveTrigger(context.Background(), triggerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.RemoveTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTriggerRequest struct via the builder pattern


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


## RunTrigger

> RunTrigger(ctx, triggerId).Execute()



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
    triggerId := "triggerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TriggersApi.RunTrigger(context.Background(), triggerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.RunTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunTriggerRequest struct via the builder pattern


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


## SearchTriggers

> map[string]interface{} SearchTriggers(ctx).FolderId(folderId).FolderTitle(folderTitle).Page(page).ResultsPerPage(resultsPerPage).TemplateId(templateId).TemplateTitle(templateTitle).TriggerTitle(triggerTitle).TriggerType(triggerType).Execute()



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
    folderId := "folderId_example" // string |  (optional)
    folderTitle := "folderTitle_example" // string |  (optional)
    page := int32(56) // int32 |  (optional) (default to 0)
    resultsPerPage := int32(56) // int32 |  (optional) (default to 100)
    templateId := "templateId_example" // string |  (optional)
    templateTitle := "templateTitle_example" // string |  (optional)
    triggerTitle := "triggerTitle_example" // string |  (optional)
    triggerType := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersApi.SearchTriggers(context.Background()).FolderId(folderId).FolderTitle(folderTitle).Page(page).ResultsPerPage(resultsPerPage).TemplateId(templateId).TemplateTitle(templateTitle).TriggerTitle(triggerTitle).TriggerType(triggerType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.SearchTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTriggers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TriggersApi.SearchTriggers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string** |  | 
 **folderTitle** | **string** |  | 
 **page** | **int32** |  | [default to 0]
 **resultsPerPage** | **int32** |  | [default to 100]
 **templateId** | **string** |  | 
 **templateTitle** | **string** |  | 
 **triggerTitle** | **string** |  | 
 **triggerType** | **[]string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrigger

> Trigger UpdateTrigger(ctx, triggerId).Trigger(trigger).Execute()



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
    triggerId := "triggerId_example" // string | 
    trigger := *openapiclient.NewTrigger() // Trigger |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggersApi.UpdateTrigger(context.Background(), triggerId).Trigger(trigger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersApi.UpdateTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrigger`: Trigger
    fmt.Fprintf(os.Stdout, "Response from `TriggersApi.UpdateTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trigger** | [**Trigger**](Trigger.md) |  | 

### Return type

[**Trigger**](Trigger.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

