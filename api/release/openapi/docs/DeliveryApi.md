# \DeliveryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteStage**](DeliveryApi.md#CompleteStage) | **Post** /api/v1/deliveries/{stageId}/complete | 
[**CompleteTrackedItem**](DeliveryApi.md#CompleteTrackedItem) | **Put** /api/v1/deliveries/{stageId}/{itemId}/complete | 
[**CompleteTransition**](DeliveryApi.md#CompleteTransition) | **Post** /api/v1/deliveries/{transitionId}/complete | 
[**CreateTrackedItemInDelivery**](DeliveryApi.md#CreateTrackedItemInDelivery) | **Post** /api/v1/deliveries/{deliveryId}/tracked-items | 
[**DeleteDelivery**](DeliveryApi.md#DeleteDelivery) | **Delete** /api/v1/deliveries/{deliveryId} | 
[**DeleteTrackedItemDelivery**](DeliveryApi.md#DeleteTrackedItemDelivery) | **Delete** /api/v1/deliveries/{itemId} | 
[**DescopeTrackedItem**](DeliveryApi.md#DescopeTrackedItem) | **Put** /api/v1/deliveries/{itemId}/descope | 
[**GetDelivery**](DeliveryApi.md#GetDelivery) | **Get** /api/v1/deliveries/{deliveryId} | 
[**GetDeliveryTimeline**](DeliveryApi.md#GetDeliveryTimeline) | **Get** /api/v1/deliveries/{deliveryId}/timeline | 
[**GetReleasesForDelivery**](DeliveryApi.md#GetReleasesForDelivery) | **Get** /api/v1/deliveries/{deliveryId}/releases | 
[**GetStagesInDelivery**](DeliveryApi.md#GetStagesInDelivery) | **Get** /api/v1/deliveries/{deliveryId}/stages | 
[**GetTrackedItemsinDelivery**](DeliveryApi.md#GetTrackedItemsinDelivery) | **Get** /api/v1/deliveries/{deliveryId}/tracked-items | 
[**ReopenStage**](DeliveryApi.md#ReopenStage) | **Post** /api/v1/deliveries/{stageId}/reopen | 
[**RescopeTrackedItem**](DeliveryApi.md#RescopeTrackedItem) | **Put** /api/v1/deliveries/{itemId}/rescope | 
[**ResetTrackedItem**](DeliveryApi.md#ResetTrackedItem) | **Put** /api/v1/deliveries/{stageId}/{itemId}/reset | 
[**SearchDeliveries**](DeliveryApi.md#SearchDeliveries) | **Post** /api/v1/deliveries/search | 
[**SkipTrackedItem**](DeliveryApi.md#SkipTrackedItem) | **Put** /api/v1/deliveries/{stageId}/{itemId}/skip | 
[**UpdateDelivery**](DeliveryApi.md#UpdateDelivery) | **Put** /api/v1/deliveries/{deliveryId} | 
[**UpdateStageInDelivery**](DeliveryApi.md#UpdateStageInDelivery) | **Put** /api/v1/deliveries/{stageId} | 
[**UpdateTrackedItemInDelivery**](DeliveryApi.md#UpdateTrackedItemInDelivery) | **Put** /api/v1/deliveries/{itemId} | 
[**UpdateTransitionInDelivery**](DeliveryApi.md#UpdateTransitionInDelivery) | **Put** /api/v1/deliveries/{transitionId} | 



## CompleteStage

> CompleteStage(ctx, stageId).Execute()



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
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.CompleteStage(context.Background(), stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.CompleteStage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCompleteStageRequest struct via the builder pattern


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


## CompleteTrackedItem

> CompleteTrackedItem(ctx, itemId, stageId).Execute()



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
    itemId := "itemId_example" // string | 
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.CompleteTrackedItem(context.Background(), itemId, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.CompleteTrackedItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTrackedItemRequest struct via the builder pattern


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


## CompleteTransition

> CompleteTransition(ctx, transitionId).CompleteTransition(completeTransition).Execute()



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
    transitionId := "transitionId_example" // string | 
    completeTransition := *openapiclient.NewCompleteTransition() // CompleteTransition |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.CompleteTransition(context.Background(), transitionId).CompleteTransition(completeTransition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.CompleteTransition``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCompleteTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completeTransition** | [**CompleteTransition**](CompleteTransition.md) |  | 

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


## CreateTrackedItemInDelivery

> TrackedItem CreateTrackedItemInDelivery(ctx, deliveryId).TrackedItem(trackedItem).Execute()



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
    deliveryId := "deliveryId_example" // string | 
    trackedItem := *openapiclient.NewTrackedItem() // TrackedItem |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.CreateTrackedItemInDelivery(context.Background(), deliveryId).TrackedItem(trackedItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.CreateTrackedItemInDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrackedItemInDelivery`: TrackedItem
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.CreateTrackedItemInDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackedItemInDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trackedItem** | [**TrackedItem**](TrackedItem.md) |  | 

### Return type

[**TrackedItem**](TrackedItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelivery

> DeleteDelivery(ctx, deliveryId).Execute()



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
    deliveryId := "deliveryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.DeleteDelivery(context.Background(), deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.DeleteDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeliveryRequest struct via the builder pattern


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


## DeleteTrackedItemDelivery

> DeleteTrackedItemDelivery(ctx, itemId).Execute()



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
    itemId := "itemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.DeleteTrackedItemDelivery(context.Background(), itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.DeleteTrackedItemDelivery``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTrackedItemDeliveryRequest struct via the builder pattern


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


## DescopeTrackedItem

> DescopeTrackedItem(ctx, itemId).Execute()



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
    itemId := "itemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.DescopeTrackedItem(context.Background(), itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.DescopeTrackedItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDescopeTrackedItemRequest struct via the builder pattern


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


## GetDelivery

> Delivery GetDelivery(ctx, deliveryId).Execute()



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
    deliveryId := "deliveryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.GetDelivery(context.Background(), deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.GetDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDelivery`: Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.GetDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Delivery**](Delivery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeliveryTimeline

> DeliveryTimeline GetDeliveryTimeline(ctx, deliveryId).Execute()



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
    deliveryId := "deliveryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.GetDeliveryTimeline(context.Background(), deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.GetDeliveryTimeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeliveryTimeline`: DeliveryTimeline
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.GetDeliveryTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveryTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeliveryTimeline**](DeliveryTimeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleasesForDelivery

> []DeliveryFlowReleaseInfo GetReleasesForDelivery(ctx, deliveryId).Execute()



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
    deliveryId := "deliveryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.GetReleasesForDelivery(context.Background(), deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.GetReleasesForDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleasesForDelivery`: []DeliveryFlowReleaseInfo
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.GetReleasesForDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasesForDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeliveryFlowReleaseInfo**](DeliveryFlowReleaseInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStagesInDelivery

> []Stage GetStagesInDelivery(ctx, deliveryId).Execute()



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
    deliveryId := "deliveryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.GetStagesInDelivery(context.Background(), deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.GetStagesInDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStagesInDelivery`: []Stage
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.GetStagesInDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStagesInDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Stage**](Stage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackedItemsinDelivery

> []TrackedItem GetTrackedItemsinDelivery(ctx, deliveryId).Execute()



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
    deliveryId := "deliveryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.GetTrackedItemsinDelivery(context.Background(), deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.GetTrackedItemsinDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrackedItemsinDelivery`: []TrackedItem
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.GetTrackedItemsinDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackedItemsinDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TrackedItem**](TrackedItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReopenStage

> ReopenStage(ctx, stageId).Execute()



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
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.ReopenStage(context.Background(), stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.ReopenStage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReopenStageRequest struct via the builder pattern


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


## RescopeTrackedItem

> RescopeTrackedItem(ctx, itemId).Execute()



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
    itemId := "itemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.RescopeTrackedItem(context.Background(), itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.RescopeTrackedItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRescopeTrackedItemRequest struct via the builder pattern


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


## ResetTrackedItem

> ResetTrackedItem(ctx, itemId, stageId).Execute()



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
    itemId := "itemId_example" // string | 
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.ResetTrackedItem(context.Background(), itemId, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.ResetTrackedItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetTrackedItemRequest struct via the builder pattern


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


## SearchDeliveries

> []Delivery SearchDeliveries(ctx).OrderBy(orderBy).Page(page).ResultsPerPage(resultsPerPage).DeliveryFilters(deliveryFilters).Execute()



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
    orderBy := openapiclient.DeliveryOrderMode("START_DATE_DELIVERY") // DeliveryOrderMode |  (optional)
    page := int64(789) // int64 |  (optional) (default to 0)
    resultsPerPage := int64(789) // int64 |  (optional) (default to 100)
    deliveryFilters := *openapiclient.NewDeliveryFilters() // DeliveryFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.SearchDeliveries(context.Background()).OrderBy(orderBy).Page(page).ResultsPerPage(resultsPerPage).DeliveryFilters(deliveryFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.SearchDeliveries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDeliveries`: []Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.SearchDeliveries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | [**DeliveryOrderMode**](DeliveryOrderMode.md) |  | 
 **page** | **int64** |  | [default to 0]
 **resultsPerPage** | **int64** |  | [default to 100]
 **deliveryFilters** | [**DeliveryFilters**](DeliveryFilters.md) |  | 

### Return type

[**[]Delivery**](Delivery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipTrackedItem

> SkipTrackedItem(ctx, itemId, stageId).Execute()



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
    itemId := "itemId_example" // string | 
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.SkipTrackedItem(context.Background(), itemId, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.SkipTrackedItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkipTrackedItemRequest struct via the builder pattern


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


## UpdateDelivery

> Delivery UpdateDelivery(ctx, deliveryId).Delivery(delivery).Execute()



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
    deliveryId := "deliveryId_example" // string | 
    delivery := *openapiclient.NewDelivery() // Delivery |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.UpdateDelivery(context.Background(), deliveryId).Delivery(delivery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.UpdateDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDelivery`: Delivery
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.UpdateDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delivery** | [**Delivery**](Delivery.md) |  | 

### Return type

[**Delivery**](Delivery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStageInDelivery

> Stage UpdateStageInDelivery(ctx, stageId).Stage(stage).Execute()



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
    stageId := "stageId_example" // string | 
    stage := *openapiclient.NewStage() // Stage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.UpdateStageInDelivery(context.Background(), stageId).Stage(stage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.UpdateStageInDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStageInDelivery`: Stage
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.UpdateStageInDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStageInDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stage** | [**Stage**](Stage.md) |  | 

### Return type

[**Stage**](Stage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrackedItemInDelivery

> TrackedItem UpdateTrackedItemInDelivery(ctx, itemId).TrackedItem(trackedItem).Execute()



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
    itemId := "itemId_example" // string | 
    trackedItem := *openapiclient.NewTrackedItem() // TrackedItem |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.UpdateTrackedItemInDelivery(context.Background(), itemId).TrackedItem(trackedItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.UpdateTrackedItemInDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrackedItemInDelivery`: TrackedItem
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.UpdateTrackedItemInDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrackedItemInDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trackedItem** | [**TrackedItem**](TrackedItem.md) |  | 

### Return type

[**TrackedItem**](TrackedItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransitionInDelivery

> Transition UpdateTransitionInDelivery(ctx, transitionId).Transition(transition).Execute()



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
    transitionId := "transitionId_example" // string | 
    transition := *openapiclient.NewTransition() // Transition |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryApi.UpdateTransitionInDelivery(context.Background(), transitionId).Transition(transition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryApi.UpdateTransitionInDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransitionInDelivery`: Transition
    fmt.Fprintf(os.Stdout, "Response from `DeliveryApi.UpdateTransitionInDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransitionInDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transition** | [**Transition**](Transition.md) |  | 

### Return type

[**Transition**](Transition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

