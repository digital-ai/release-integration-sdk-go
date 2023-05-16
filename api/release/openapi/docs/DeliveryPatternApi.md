# \DeliveryPatternApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckTitle**](DeliveryPatternApi.md#CheckTitle) | **Post** /api/v1/delivery-patterns/checkTitle | 
[**CreateDeliveryFromPattern**](DeliveryPatternApi.md#CreateDeliveryFromPattern) | **Post** /api/v1/delivery-patterns/{patternId}/create | 
[**CreatePattern**](DeliveryPatternApi.md#CreatePattern) | **Post** /api/v1/delivery-patterns | 
[**CreateStage**](DeliveryPatternApi.md#CreateStage) | **Post** /api/v1/delivery-patterns/{patternId}/createStage | 
[**CreateStage1**](DeliveryPatternApi.md#CreateStage1) | **Post** /api/v1/delivery-patterns/{patternId}/stages | 
[**CreateStage2**](DeliveryPatternApi.md#CreateStage2) | **Post** /api/v1/delivery-patterns/{patternId}/stages/{position} | 
[**CreateTrackedItemInPattern**](DeliveryPatternApi.md#CreateTrackedItemInPattern) | **Post** /api/v1/delivery-patterns/{patternId}/tracked-items | 
[**CreateTransition**](DeliveryPatternApi.md#CreateTransition) | **Post** /api/v1/delivery-patterns/{stageId}/transitions | 
[**DeletePattern**](DeliveryPatternApi.md#DeletePattern) | **Delete** /api/v1/delivery-patterns/{patternId} | 
[**DeleteStage**](DeliveryPatternApi.md#DeleteStage) | **Delete** /api/v1/delivery-patterns/{stageId} | 
[**DeleteTrackedItemDeliveryPattern**](DeliveryPatternApi.md#DeleteTrackedItemDeliveryPattern) | **Delete** /api/v1/delivery-patterns/{itemId} | 
[**DeleteTransition**](DeliveryPatternApi.md#DeleteTransition) | **Delete** /api/v1/delivery-patterns/{transitionId} | 
[**DuplicatePattern**](DeliveryPatternApi.md#DuplicatePattern) | **Post** /api/v1/delivery-patterns/{patternId}/duplicate | 
[**GetPattern**](DeliveryPatternApi.md#GetPattern) | **Get** /api/v1/delivery-patterns/{patternId} | 
[**GetPatternByIdOrTitle**](DeliveryPatternApi.md#GetPatternByIdOrTitle) | **Get** /api/v1/delivery-patterns/{patternIdOrTitle} | 
[**GetStagesInPattern**](DeliveryPatternApi.md#GetStagesInPattern) | **Get** /api/v1/delivery-patterns/{patternId}/stages | 
[**GetTrackedItemsInPattern**](DeliveryPatternApi.md#GetTrackedItemsInPattern) | **Get** /api/v1/delivery-patterns/{patternId}/tracked-items | 
[**SearchPatterns**](DeliveryPatternApi.md#SearchPatterns) | **Post** /api/v1/delivery-patterns/search | 
[**UpdatePattern**](DeliveryPatternApi.md#UpdatePattern) | **Put** /api/v1/delivery-patterns/{patternId} | 
[**UpdateStageFromBatch**](DeliveryPatternApi.md#UpdateStageFromBatch) | **Put** /api/v1/delivery-patterns/{stageId}/batched | 
[**UpdateStageInPattern**](DeliveryPatternApi.md#UpdateStageInPattern) | **Put** /api/v1/delivery-patterns/{stageId} | 
[**UpdateTrackedItemInPattern**](DeliveryPatternApi.md#UpdateTrackedItemInPattern) | **Put** /api/v1/delivery-patterns/{itemId} | 
[**UpdateTransitionInPattern**](DeliveryPatternApi.md#UpdateTransitionInPattern) | **Put** /api/v1/delivery-patterns/{transitionId} | 



## CheckTitle

> bool CheckTitle(ctx).ValidatePattern(validatePattern).Execute()



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
    validatePattern := *openapiclient.NewValidatePattern() // ValidatePattern |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.CheckTitle(context.Background()).ValidatePattern(validatePattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.CheckTitle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckTitle`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.CheckTitle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validatePattern** | [**ValidatePattern**](ValidatePattern.md) |  | 

### Return type

**bool**

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeliveryFromPattern

> Delivery CreateDeliveryFromPattern(ctx, patternId).CreateDelivery(createDelivery).Execute()



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
    patternId := "patternId_example" // string | 
    createDelivery := *openapiclient.NewCreateDelivery() // CreateDelivery |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.CreateDeliveryFromPattern(context.Background(), patternId).CreateDelivery(createDelivery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.CreateDeliveryFromPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeliveryFromPattern`: Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.CreateDeliveryFromPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeliveryFromPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDelivery** | [**CreateDelivery**](CreateDelivery.md) |  | 

### Return type

[**Delivery**](Delivery.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePattern

> Delivery CreatePattern(ctx).Delivery(delivery).Execute()



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
    delivery := *openapiclient.NewDelivery() // Delivery |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.CreatePattern(context.Background()).Delivery(delivery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.CreatePattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePattern`: Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.CreatePattern`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delivery** | [**Delivery**](Delivery.md) |  | 

### Return type

[**Delivery**](Delivery.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStage

> Stage CreateStage(ctx, patternId).CreateDeliveryStage(createDeliveryStage).Execute()



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
    patternId := "patternId_example" // string | 
    createDeliveryStage := *openapiclient.NewCreateDeliveryStage() // CreateDeliveryStage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.CreateStage(context.Background(), patternId).CreateDeliveryStage(createDeliveryStage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.CreateStage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStage`: Stage
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.CreateStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeliveryStage** | [**CreateDeliveryStage**](CreateDeliveryStage.md) |  | 

### Return type

[**Stage**](Stage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStage1

> Stage CreateStage1(ctx, patternId).Stage(stage).Execute()



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
    patternId := "patternId_example" // string | 
    stage := *openapiclient.NewStage() // Stage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.CreateStage1(context.Background(), patternId).Stage(stage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.CreateStage1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStage1`: Stage
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.CreateStage1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStage1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stage** | [**Stage**](Stage.md) |  | 

### Return type

[**Stage**](Stage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStage2

> Stage CreateStage2(ctx, patternId, position).Stage(stage).Execute()



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
    patternId := "patternId_example" // string | 
    position := int32(56) // int32 | 
    stage := *openapiclient.NewStage() // Stage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.CreateStage2(context.Background(), patternId, position).Stage(stage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.CreateStage2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStage2`: Stage
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.CreateStage2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 
**position** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStage2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stage** | [**Stage**](Stage.md) |  | 

### Return type

[**Stage**](Stage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrackedItemInPattern

> TrackedItem CreateTrackedItemInPattern(ctx, patternId).TrackedItem(trackedItem).Execute()



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
    patternId := "patternId_example" // string | 
    trackedItem := *openapiclient.NewTrackedItem() // TrackedItem |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.CreateTrackedItemInPattern(context.Background(), patternId).TrackedItem(trackedItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.CreateTrackedItemInPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrackedItemInPattern`: TrackedItem
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.CreateTrackedItemInPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackedItemInPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trackedItem** | [**TrackedItem**](TrackedItem.md) |  | 

### Return type

[**TrackedItem**](TrackedItem.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransition

> Transition CreateTransition(ctx, stageId).Transition(transition).Execute()



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
    stageId := "stageId_example" // string | 
    transition := *openapiclient.NewTransition() // Transition |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.CreateTransition(context.Background(), stageId).Transition(transition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.CreateTransition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransition`: Transition
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.CreateTransition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transition** | [**Transition**](Transition.md) |  | 

### Return type

[**Transition**](Transition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePattern

> DeletePattern(ctx, patternId).Execute()



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
    patternId := "patternId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryPatternApi.DeletePattern(context.Background(), patternId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.DeletePattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePatternRequest struct via the builder pattern


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


## DeleteStage

> DeleteStage(ctx, stageId).Execute()



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
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryPatternApi.DeleteStage(context.Background(), stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.DeleteStage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStageRequest struct via the builder pattern


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


## DeleteTrackedItemDeliveryPattern

> DeleteTrackedItemDeliveryPattern(ctx, itemId).Execute()



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
    itemId := "itemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryPatternApi.DeleteTrackedItemDeliveryPattern(context.Background(), itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.DeleteTrackedItemDeliveryPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrackedItemDeliveryPatternRequest struct via the builder pattern


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


## DeleteTransition

> DeleteTransition(ctx, transitionId).Execute()



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
    transitionId := "transitionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryPatternApi.DeleteTransition(context.Background(), transitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.DeleteTransition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransitionRequest struct via the builder pattern


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


## DuplicatePattern

> Delivery DuplicatePattern(ctx, patternId).DuplicateDeliveryPattern(duplicateDeliveryPattern).Execute()



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
    patternId := "patternId_example" // string | 
    duplicateDeliveryPattern := *openapiclient.NewDuplicateDeliveryPattern() // DuplicateDeliveryPattern |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.DuplicatePattern(context.Background(), patternId).DuplicateDeliveryPattern(duplicateDeliveryPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.DuplicatePattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuplicatePattern`: Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.DuplicatePattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicatePatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duplicateDeliveryPattern** | [**DuplicateDeliveryPattern**](DuplicateDeliveryPattern.md) |  | 

### Return type

[**Delivery**](Delivery.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPattern

> Delivery GetPattern(ctx, patternId).Execute()



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
    patternId := "patternId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.GetPattern(context.Background(), patternId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.GetPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPattern`: Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.GetPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Delivery**](Delivery.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatternByIdOrTitle

> Delivery GetPatternByIdOrTitle(ctx, patternIdOrTitle).Execute()



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
    patternIdOrTitle := "patternIdOrTitle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.GetPatternByIdOrTitle(context.Background(), patternIdOrTitle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.GetPatternByIdOrTitle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPatternByIdOrTitle`: Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.GetPatternByIdOrTitle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternIdOrTitle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatternByIdOrTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Delivery**](Delivery.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStagesInPattern

> []Stage GetStagesInPattern(ctx, patternId).Execute()



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
    patternId := "patternId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.GetStagesInPattern(context.Background(), patternId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.GetStagesInPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStagesInPattern`: []Stage
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.GetStagesInPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStagesInPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Stage**](Stage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackedItemsInPattern

> []TrackedItem GetTrackedItemsInPattern(ctx, patternId).Execute()



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
    patternId := "patternId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.GetTrackedItemsInPattern(context.Background(), patternId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.GetTrackedItemsInPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrackedItemsInPattern`: []TrackedItem
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.GetTrackedItemsInPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackedItemsInPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TrackedItem**](TrackedItem.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPatterns

> []Delivery SearchPatterns(ctx).Page(page).ResultsPerPage(resultsPerPage).DeliveryPatternFilters(deliveryPatternFilters).Execute()



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
    page := int64(789) // int64 |  (optional) (default to 0)
    resultsPerPage := int64(789) // int64 |  (optional) (default to 100)
    deliveryPatternFilters := *openapiclient.NewDeliveryPatternFilters() // DeliveryPatternFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.SearchPatterns(context.Background()).Page(page).ResultsPerPage(resultsPerPage).DeliveryPatternFilters(deliveryPatternFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.SearchPatterns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPatterns`: []Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.SearchPatterns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPatternsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **resultsPerPage** | **int64** |  | [default to 100]
 **deliveryPatternFilters** | [**DeliveryPatternFilters**](DeliveryPatternFilters.md) |  | 

### Return type

[**[]Delivery**](Delivery.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePattern

> Delivery UpdatePattern(ctx, patternId).Delivery(delivery).Execute()



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
    patternId := "patternId_example" // string | 
    delivery := *openapiclient.NewDelivery() // Delivery |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.UpdatePattern(context.Background(), patternId).Delivery(delivery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.UpdatePattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePattern`: Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.UpdatePattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delivery** | [**Delivery**](Delivery.md) |  | 

### Return type

[**Delivery**](Delivery.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStageFromBatch

> Stage UpdateStageFromBatch(ctx, stageId).Stage(stage).Execute()



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
    stageId := "stageId_example" // string | 
    stage := *openapiclient.NewStage() // Stage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.UpdateStageFromBatch(context.Background(), stageId).Stage(stage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.UpdateStageFromBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStageFromBatch`: Stage
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.UpdateStageFromBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStageFromBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stage** | [**Stage**](Stage.md) |  | 

### Return type

[**Stage**](Stage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStageInPattern

> Stage UpdateStageInPattern(ctx, stageId).Stage(stage).Execute()



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
    stageId := "stageId_example" // string | 
    stage := *openapiclient.NewStage() // Stage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.UpdateStageInPattern(context.Background(), stageId).Stage(stage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.UpdateStageInPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStageInPattern`: Stage
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.UpdateStageInPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStageInPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stage** | [**Stage**](Stage.md) |  | 

### Return type

[**Stage**](Stage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrackedItemInPattern

> TrackedItem UpdateTrackedItemInPattern(ctx, itemId).TrackedItem(trackedItem).Execute()



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
    itemId := "itemId_example" // string | 
    trackedItem := *openapiclient.NewTrackedItem() // TrackedItem |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.UpdateTrackedItemInPattern(context.Background(), itemId).TrackedItem(trackedItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.UpdateTrackedItemInPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrackedItemInPattern`: TrackedItem
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.UpdateTrackedItemInPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrackedItemInPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trackedItem** | [**TrackedItem**](TrackedItem.md) |  | 

### Return type

[**TrackedItem**](TrackedItem.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransitionInPattern

> Transition UpdateTransitionInPattern(ctx, transitionId).Transition(transition).Execute()



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
    transitionId := "transitionId_example" // string | 
    transition := *openapiclient.NewTransition() // Transition |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryPatternApi.UpdateTransitionInPattern(context.Background(), transitionId).Transition(transition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryPatternApi.UpdateTransitionInPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransitionInPattern`: Transition
    fmt.Fprintf(os.Stdout, "Response from `DeliveryPatternApi.UpdateTransitionInPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransitionInPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transition** | [**Transition**](Transition.md) |  | 

### Return type

[**Transition**](Transition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

