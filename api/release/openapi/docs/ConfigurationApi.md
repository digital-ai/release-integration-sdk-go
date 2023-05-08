# \ConfigurationApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConfiguration**](ConfigurationApi.md#AddConfiguration) | **Post** /api/v1/config | 
[**AddGlobalVariable**](ConfigurationApi.md#AddGlobalVariable) | **Post** /api/v1/config/Configuration/variables/global | 
[**CheckStatus**](ConfigurationApi.md#CheckStatus) | **Post** /api/v1/config/status | 
[**CheckStatus1**](ConfigurationApi.md#CheckStatus1) | **Post** /api/v1/config/{configurationId}/status | 
[**DeleteConfiguration**](ConfigurationApi.md#DeleteConfiguration) | **Delete** /api/v1/config/{configurationId} | 
[**DeleteGlobalVariable**](ConfigurationApi.md#DeleteGlobalVariable) | **Delete** /api/v1/config/{variableId} | 
[**GetConfiguration**](ConfigurationApi.md#GetConfiguration) | **Get** /api/v1/config/{configurationId} | 
[**GetGlobalVariable**](ConfigurationApi.md#GetGlobalVariable) | **Get** /api/v1/config/{variableId} | 
[**GetGlobalVariableValues**](ConfigurationApi.md#GetGlobalVariableValues) | **Get** /api/v1/config/Configuration/variableValues/global | 
[**GetGlobalVariables**](ConfigurationApi.md#GetGlobalVariables) | **Get** /api/v1/config/Configuration/variables/global | 
[**GetSystemMessage**](ConfigurationApi.md#GetSystemMessage) | **Get** /api/v1/config/system-message | 
[**SearchByTypeAndTitle**](ConfigurationApi.md#SearchByTypeAndTitle) | **Get** /api/v1/config/byTypeAndTitle | 
[**UpdateConfiguration**](ConfigurationApi.md#UpdateConfiguration) | **Put** /api/v1/config/{configurationId} | 
[**UpdateGlobalVariable**](ConfigurationApi.md#UpdateGlobalVariable) | **Put** /api/v1/config/{variableId} | 
[**UpdateSystemMessage**](ConfigurationApi.md#UpdateSystemMessage) | **Put** /api/v1/config/system-message | 



## AddConfiguration

> ReleaseConfiguration AddConfiguration(ctx).ReleaseConfiguration(releaseConfiguration).Execute()



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
    releaseConfiguration := *openapiclient.NewReleaseConfiguration() // ReleaseConfiguration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.AddConfiguration(context.Background()).ReleaseConfiguration(releaseConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.AddConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConfiguration`: ReleaseConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.AddConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseConfiguration** | [**ReleaseConfiguration**](ReleaseConfiguration.md) |  | 

### Return type

[**ReleaseConfiguration**](ReleaseConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGlobalVariable

> Variable AddGlobalVariable(ctx).Variable1(variable1).Execute()



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
    variable1 := *openapiclient.NewVariable1() // Variable1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.AddGlobalVariable(context.Background()).Variable1(variable1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.AddGlobalVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGlobalVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.AddGlobalVariable`: %v\n", resp)
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

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckStatus

> SharedConfigurationStatusResponse CheckStatus(ctx).ConfigurationView(configurationView).Execute()



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
    configurationView := *openapiclient.NewConfigurationView() // ConfigurationView |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.CheckStatus(context.Background()).ConfigurationView(configurationView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.CheckStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckStatus`: SharedConfigurationStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.CheckStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationView** | [**ConfigurationView**](ConfigurationView.md) |  | 

### Return type

[**SharedConfigurationStatusResponse**](SharedConfigurationStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckStatus1

> SharedConfigurationStatusResponse CheckStatus1(ctx, configurationId).Execute()



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
    configurationId := "configurationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.CheckStatus1(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.CheckStatus1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckStatus1`: SharedConfigurationStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.CheckStatus1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckStatus1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SharedConfigurationStatusResponse**](SharedConfigurationStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration

> DeleteConfiguration(ctx, configurationId).Execute()



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
    configurationId := "configurationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigurationApi.DeleteConfiguration(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.DeleteConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


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


## DeleteGlobalVariable

> DeleteGlobalVariable(ctx, variableId).Execute()



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
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigurationApi.DeleteGlobalVariable(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.DeleteGlobalVariable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGlobalVariableRequest struct via the builder pattern


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


## GetConfiguration

> ReleaseConfiguration GetConfiguration(ctx, configurationId).Execute()



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
    configurationId := "configurationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.GetConfiguration(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration`: ReleaseConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseConfiguration**](ReleaseConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalVariable

> Variable GetGlobalVariable(ctx, variableId).Execute()



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
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.GetGlobalVariable(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetGlobalVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetGlobalVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Variable**](Variable.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.GetGlobalVariableValues(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetGlobalVariableValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalVariableValues`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetGlobalVariableValues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalVariableValuesRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalVariables

> []Variable GetGlobalVariables(ctx).Execute()



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
    resp, r, err := apiClient.ConfigurationApi.GetGlobalVariables(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetGlobalVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalVariables`: []Variable
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetGlobalVariables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalVariablesRequest struct via the builder pattern


### Return type

[**[]Variable**](Variable.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMessage

> SystemMessageSettings GetSystemMessage(ctx).Execute()



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
    resp, r, err := apiClient.ConfigurationApi.GetSystemMessage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetSystemMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemMessage`: SystemMessageSettings
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetSystemMessage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMessageRequest struct via the builder pattern


### Return type

[**SystemMessageSettings**](SystemMessageSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchByTypeAndTitle

> []ReleaseConfiguration SearchByTypeAndTitle(ctx).ConfigurationType(configurationType).FolderId(folderId).FolderOnly(folderOnly).Title(title).Execute()



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
    configurationType := "configurationType_example" // string |  (optional)
    folderId := "folderId_example" // string |  (optional)
    folderOnly := true // bool |  (optional)
    title := "title_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.SearchByTypeAndTitle(context.Background()).ConfigurationType(configurationType).FolderId(folderId).FolderOnly(folderOnly).Title(title).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.SearchByTypeAndTitle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchByTypeAndTitle`: []ReleaseConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.SearchByTypeAndTitle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchByTypeAndTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationType** | **string** |  | 
 **folderId** | **string** |  | 
 **folderOnly** | **bool** |  | 
 **title** | **string** |  | 

### Return type

[**[]ReleaseConfiguration**](ReleaseConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> ReleaseConfiguration UpdateConfiguration(ctx, configurationId).ReleaseConfiguration(releaseConfiguration).Execute()



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
    configurationId := "configurationId_example" // string | 
    releaseConfiguration := *openapiclient.NewReleaseConfiguration() // ReleaseConfiguration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.UpdateConfiguration(context.Background(), configurationId).ReleaseConfiguration(releaseConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: ReleaseConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseConfiguration** | [**ReleaseConfiguration**](ReleaseConfiguration.md) |  | 

### Return type

[**ReleaseConfiguration**](ReleaseConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalVariable

> Variable UpdateGlobalVariable(ctx, variableId).Variable(variable).Execute()



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
    variableId := "variableId_example" // string | 
    variable := *openapiclient.NewVariable() // Variable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.UpdateGlobalVariable(context.Background(), variableId).Variable(variable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.UpdateGlobalVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGlobalVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.UpdateGlobalVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variable** | [**Variable**](Variable.md) |  | 

### Return type

[**Variable**](Variable.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSystemMessage

> SystemMessageSettings UpdateSystemMessage(ctx).SystemMessageSettings(systemMessageSettings).Execute()



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
    systemMessageSettings := *openapiclient.NewSystemMessageSettings() // SystemMessageSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.UpdateSystemMessage(context.Background()).SystemMessageSettings(systemMessageSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.UpdateSystemMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSystemMessage`: SystemMessageSettings
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.UpdateSystemMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemMessageSettings** | [**SystemMessageSettings**](SystemMessageSettings.md) |  | 

### Return type

[**SystemMessageSettings**](SystemMessageSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

