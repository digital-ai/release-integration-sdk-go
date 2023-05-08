# \FolderApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFolder**](FolderApi.md#AddFolder) | **Post** /api/v1/folders/{folderId} | 
[**CreateFolderVariable**](FolderApi.md#CreateFolderVariable) | **Post** /api/v1/folders/{folderId}/variables | 
[**DeleteFolder**](FolderApi.md#DeleteFolder) | **Delete** /api/v1/folders/{folderId} | 
[**DeleteFolderVariable**](FolderApi.md#DeleteFolderVariable) | **Delete** /api/v1/folders/{folderId}/{variableId} | 
[**Find**](FolderApi.md#Find) | **Get** /api/v1/folders/find | 
[**GetFolder**](FolderApi.md#GetFolder) | **Get** /api/v1/folders/{folderId} | 
[**GetFolderPermissions**](FolderApi.md#GetFolderPermissions) | **Get** /api/v1/folders/permissions | 
[**GetFolderTeams**](FolderApi.md#GetFolderTeams) | **Get** /api/v1/folders/{folderId}/teams | 
[**GetFolderTemplates**](FolderApi.md#GetFolderTemplates) | **Get** /api/v1/folders/{folderId}/templates | 
[**GetFolderVariable**](FolderApi.md#GetFolderVariable) | **Get** /api/v1/folders/{folderId}/{variableId} | 
[**IsFolderOwner**](FolderApi.md#IsFolderOwner) | **Get** /api/v1/folders/{folderId}/folderOwner | 
[**List**](FolderApi.md#List) | **Get** /api/v1/folders/{folderId}/list | 
[**ListRoot**](FolderApi.md#ListRoot) | **Get** /api/v1/folders/list | 
[**ListVariableValues**](FolderApi.md#ListVariableValues) | **Get** /api/v1/folders/{folderId}/variableValues | 
[**ListVariables**](FolderApi.md#ListVariables) | **Get** /api/v1/folders/{folderId}/variables | 
[**Move**](FolderApi.md#Move) | **Post** /api/v1/folders/{folderId}/move | 
[**MoveTemplate**](FolderApi.md#MoveTemplate) | **Post** /api/v1/folders/{folderId}/templates/{templateId} | 
[**RenameFolder**](FolderApi.md#RenameFolder) | **Post** /api/v1/folders/{folderId}/rename | 
[**SearchReleasesFolder**](FolderApi.md#SearchReleasesFolder) | **Post** /api/v1/folders/{folderId}/releases | 
[**SetFolderTeams**](FolderApi.md#SetFolderTeams) | **Post** /api/v1/folders/{folderId}/teams | 
[**UpdateFolderVariable**](FolderApi.md#UpdateFolderVariable) | **Put** /api/v1/folders/{folderId}/{variableId} | 



## AddFolder

> Folder AddFolder(ctx, folderId).Folder(folder).Execute()



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
    folderId := "folderId_example" // string | 
    folder := *openapiclient.NewFolder() // Folder |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.AddFolder(context.Background(), folderId).Folder(folder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.AddFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.AddFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folder** | [**Folder**](Folder.md) |  | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFolderVariable

> Variable CreateFolderVariable(ctx, folderId).Variable1(variable1).Execute()



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
    folderId := "folderId_example" // string | 
    variable1 := *openapiclient.NewVariable1() // Variable1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.CreateFolderVariable(context.Background(), folderId).Variable1(variable1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.CreateFolderVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFolderVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.CreateFolderVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderVariableRequest struct via the builder pattern


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


## DeleteFolder

> DeleteFolder(ctx, folderId).Execute()



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
    folderId := "folderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FolderApi.DeleteFolder(context.Background(), folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.DeleteFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderRequest struct via the builder pattern


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


## DeleteFolderVariable

> DeleteFolderVariable(ctx, folderId, variableId).Execute()



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
    folderId := "folderId_example" // string | 
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FolderApi.DeleteFolderVariable(context.Background(), folderId, variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.DeleteFolderVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderVariableRequest struct via the builder pattern


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


## Find

> Folder Find(ctx).ByPath(byPath).Depth(depth).Execute()



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
    byPath := "byPath_example" // string |  (optional)
    depth := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.Find(context.Background()).ByPath(byPath).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.Find``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Find`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.Find`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **byPath** | **string** |  | 
 **depth** | **int32** |  | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolder

> Folder GetFolder(ctx, folderId).Depth(depth).Execute()



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
    folderId := "folderId_example" // string | 
    depth := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.GetFolder(context.Background(), folderId).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.GetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.GetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int32** |  | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderPermissions

> []string GetFolderPermissions(ctx).Execute()



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
    resp, r, err := apiClient.FolderApi.GetFolderPermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.GetFolderPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolderPermissions`: []string
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.GetFolderPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderPermissionsRequest struct via the builder pattern


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


## GetFolderTeams

> []TeamView GetFolderTeams(ctx, folderId).Execute()



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
    folderId := "folderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.GetFolderTeams(context.Background(), folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.GetFolderTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolderTeams`: []TeamView
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.GetFolderTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TeamView**](TeamView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderTemplates

> []Release GetFolderTemplates(ctx, folderId).Depth(depth).Page(page).ResultsPerPage(resultsPerPage).Execute()



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
    folderId := "folderId_example" // string | 
    depth := int32(56) // int32 |  (optional)
    page := int64(789) // int64 |  (optional)
    resultsPerPage := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.GetFolderTemplates(context.Background(), folderId).Depth(depth).Page(page).ResultsPerPage(resultsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.GetFolderTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolderTemplates`: []Release
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.GetFolderTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int32** |  | 
 **page** | **int64** |  | 
 **resultsPerPage** | **int64** |  | 

### Return type

[**[]Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderVariable

> Variable GetFolderVariable(ctx, folderId, variableId).Execute()



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
    folderId := "folderId_example" // string | 
    variableId := "variableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.GetFolderVariable(context.Background(), folderId, variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.GetFolderVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolderVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.GetFolderVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderVariableRequest struct via the builder pattern


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


## IsFolderOwner

> bool IsFolderOwner(ctx, folderId).Execute()



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
    folderId := "folderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.IsFolderOwner(context.Background(), folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.IsFolderOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsFolderOwner`: bool
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.IsFolderOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsFolderOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Folder List(ctx, folderId).Depth(depth).Page(page).Permissions(permissions).ResultsPerPage(resultsPerPage).Execute()



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
    folderId := "folderId_example" // string | 
    depth := int32(56) // int32 |  (optional)
    page := int64(789) // int64 |  (optional)
    permissions := true // bool |  (optional)
    resultsPerPage := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.List(context.Background(), folderId).Depth(depth).Page(page).Permissions(permissions).ResultsPerPage(resultsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int32** |  | 
 **page** | **int64** |  | 
 **permissions** | **bool** |  | 
 **resultsPerPage** | **int64** |  | 

### Return type

[**[]Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoot

> []Folder ListRoot(ctx).Depth(depth).Page(page).Permissions(permissions).ResultsPerPage(resultsPerPage).Execute()



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
    depth := int32(56) // int32 |  (optional)
    page := int64(789) // int64 |  (optional)
    permissions := true // bool |  (optional)
    resultsPerPage := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.ListRoot(context.Background()).Depth(depth).Page(page).Permissions(permissions).ResultsPerPage(resultsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.ListRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoot`: []Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.ListRoot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depth** | **int32** |  | 
 **page** | **int64** |  | 
 **permissions** | **bool** |  | 
 **resultsPerPage** | **int64** |  | 

### Return type

[**[]Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVariableValues

> map[string]string ListVariableValues(ctx, folderId).FolderOnly(folderOnly).Execute()



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
    folderId := "folderId_example" // string | 
    folderOnly := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.ListVariableValues(context.Background(), folderId).FolderOnly(folderOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.ListVariableValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVariableValues`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.ListVariableValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVariableValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderOnly** | **bool** |  | 

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


## ListVariables

> []Variable ListVariables(ctx, folderId).FolderOnly(folderOnly).Execute()



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
    folderId := "folderId_example" // string | 
    folderOnly := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.ListVariables(context.Background(), folderId).FolderOnly(folderOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.ListVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVariables`: []Variable
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.ListVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderOnly** | **bool** |  | 

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


## Move

> Move(ctx, folderId).NewParentId(newParentId).Execute()



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
    folderId := "folderId_example" // string | 
    newParentId := "newParentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FolderApi.Move(context.Background(), folderId).NewParentId(newParentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.Move``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newParentId** | **string** |  | 

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


## MoveTemplate

> MoveTemplate(ctx, folderId, templateId).MergePermissions(mergePermissions).Execute()



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
    folderId := "folderId_example" // string | 
    templateId := "templateId_example" // string | 
    mergePermissions := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FolderApi.MoveTemplate(context.Background(), folderId, templateId).MergePermissions(mergePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.MoveTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mergePermissions** | **bool** |  | 

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


## RenameFolder

> RenameFolder(ctx, folderId).NewName(newName).Execute()



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
    folderId := "folderId_example" // string | 
    newName := "newName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FolderApi.RenameFolder(context.Background(), folderId).NewName(newName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.RenameFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newName** | **string** |  | 

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


## SearchReleasesFolder

> []Release SearchReleasesFolder(ctx, folderId).Depth(depth).Numberbypage(numberbypage).Page(page).ReleasesFilters(releasesFilters).Execute()



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
    folderId := "folderId_example" // string | 
    depth := int32(56) // int32 |  (optional)
    numberbypage := int64(789) // int64 |  (optional)
    page := int64(789) // int64 |  (optional)
    releasesFilters := *openapiclient.NewReleasesFilters() // ReleasesFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.SearchReleasesFolder(context.Background(), folderId).Depth(depth).Numberbypage(numberbypage).Page(page).ReleasesFilters(releasesFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.SearchReleasesFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchReleasesFolder`: []Release
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.SearchReleasesFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchReleasesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int32** |  | 
 **numberbypage** | **int64** |  | 
 **page** | **int64** |  | 
 **releasesFilters** | [**ReleasesFilters**](ReleasesFilters.md) |  | 

### Return type

[**[]Release**](Release.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFolderTeams

> []TeamView SetFolderTeams(ctx, folderId).TeamView(teamView).Execute()



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
    folderId := "folderId_example" // string | 
    teamView := []openapiclient.TeamView{*openapiclient.NewTeamView()} // []TeamView |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.SetFolderTeams(context.Background(), folderId).TeamView(teamView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.SetFolderTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetFolderTeams`: []TeamView
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.SetFolderTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFolderTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamView** | [**[]TeamView**](TeamView.md) |  | 

### Return type

[**[]TeamView**](TeamView.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolderVariable

> Variable UpdateFolderVariable(ctx, folderId, variableId).Variable(variable).Execute()



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
    folderId := "folderId_example" // string | 
    variableId := "variableId_example" // string | 
    variable := *openapiclient.NewVariable() // Variable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.UpdateFolderVariable(context.Background(), folderId, variableId).Variable(variable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.UpdateFolderVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFolderVariable`: Variable
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.UpdateFolderVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 
**variableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderVariableRequest struct via the builder pattern


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

