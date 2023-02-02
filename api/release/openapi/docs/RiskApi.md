# \RiskApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyRiskProfile**](RiskApi.md#CopyRiskProfile) | **Post** /api/v1/risks/profiles/{riskProfileId}/copy | 
[**CreateRiskProfile**](RiskApi.md#CreateRiskProfile) | **Post** /api/v1/risks/profiles | 
[**DeleteRiskProfile**](RiskApi.md#DeleteRiskProfile) | **Delete** /api/v1/risks/profiles/{riskProfileId} | 
[**GetAllRiskAssessors**](RiskApi.md#GetAllRiskAssessors) | **Get** /api/v1/risks/assessors | 
[**GetRisk**](RiskApi.md#GetRisk) | **Get** /api/v1/risks/{riskId} | 
[**GetRiskGlobalThresholds**](RiskApi.md#GetRiskGlobalThresholds) | **Get** /api/v1/risks/config | 
[**GetRiskProfile**](RiskApi.md#GetRiskProfile) | **Get** /api/v1/risks/profiles/{riskProfileId} | 
[**GetRiskProfiles**](RiskApi.md#GetRiskProfiles) | **Get** /api/v1/risks/profiles | 
[**UpdateRiskGlobalThresholds**](RiskApi.md#UpdateRiskGlobalThresholds) | **Put** /api/v1/risks/config | 
[**UpdateRiskProfile**](RiskApi.md#UpdateRiskProfile) | **Put** /api/v1/risks/profiles/{riskProfileId} | 



## CopyRiskProfile

> RiskProfile CopyRiskProfile(ctx, riskProfileId).Execute()



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
    riskProfileId := "riskProfileId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskApi.CopyRiskProfile(context.Background(), riskProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.CopyRiskProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyRiskProfile`: RiskProfile
    fmt.Fprintf(os.Stdout, "Response from `RiskApi.CopyRiskProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyRiskProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskProfile**](RiskProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRiskProfile

> RiskProfile CreateRiskProfile(ctx).RiskProfile(riskProfile).Execute()



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
    riskProfile := *openapiclient.NewRiskProfile() // RiskProfile |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskApi.CreateRiskProfile(context.Background()).RiskProfile(riskProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.CreateRiskProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRiskProfile`: RiskProfile
    fmt.Fprintf(os.Stdout, "Response from `RiskApi.CreateRiskProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskProfile** | [**RiskProfile**](RiskProfile.md) |  | 

### Return type

[**RiskProfile**](RiskProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskProfile

> DeleteRiskProfile(ctx, riskProfileId).Execute()



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
    riskProfileId := "riskProfileId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskApi.DeleteRiskProfile(context.Background(), riskProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.DeleteRiskProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskProfileRequest struct via the builder pattern


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


## GetAllRiskAssessors

> []RiskAssessor GetAllRiskAssessors(ctx).Execute()



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
    resp, r, err := apiClient.RiskApi.GetAllRiskAssessors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.GetAllRiskAssessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRiskAssessors`: []RiskAssessor
    fmt.Fprintf(os.Stdout, "Response from `RiskApi.GetAllRiskAssessors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRiskAssessorsRequest struct via the builder pattern


### Return type

[**[]RiskAssessor**](RiskAssessor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRisk

> Risk GetRisk(ctx, riskId).Execute()



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
    riskId := "riskId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskApi.GetRisk(context.Background(), riskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.GetRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRisk`: Risk
    fmt.Fprintf(os.Stdout, "Response from `RiskApi.GetRisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Risk**](Risk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskGlobalThresholds

> RiskGlobalThresholds GetRiskGlobalThresholds(ctx).Execute()



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
    resp, r, err := apiClient.RiskApi.GetRiskGlobalThresholds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.GetRiskGlobalThresholds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskGlobalThresholds`: RiskGlobalThresholds
    fmt.Fprintf(os.Stdout, "Response from `RiskApi.GetRiskGlobalThresholds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskGlobalThresholdsRequest struct via the builder pattern


### Return type

[**RiskGlobalThresholds**](RiskGlobalThresholds.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskProfile

> RiskProfile GetRiskProfile(ctx, riskProfileId).Execute()



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
    riskProfileId := "riskProfileId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskApi.GetRiskProfile(context.Background(), riskProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.GetRiskProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskProfile`: RiskProfile
    fmt.Fprintf(os.Stdout, "Response from `RiskApi.GetRiskProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskProfile**](RiskProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskProfiles

> []RiskProfile GetRiskProfiles(ctx).Execute()



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
    resp, r, err := apiClient.RiskApi.GetRiskProfiles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.GetRiskProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskProfiles`: []RiskProfile
    fmt.Fprintf(os.Stdout, "Response from `RiskApi.GetRiskProfiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskProfilesRequest struct via the builder pattern


### Return type

[**[]RiskProfile**](RiskProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskGlobalThresholds

> RiskGlobalThresholds UpdateRiskGlobalThresholds(ctx).RiskGlobalThresholds(riskGlobalThresholds).Execute()



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
    riskGlobalThresholds := *openapiclient.NewRiskGlobalThresholds() // RiskGlobalThresholds |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskApi.UpdateRiskGlobalThresholds(context.Background()).RiskGlobalThresholds(riskGlobalThresholds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.UpdateRiskGlobalThresholds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskGlobalThresholds`: RiskGlobalThresholds
    fmt.Fprintf(os.Stdout, "Response from `RiskApi.UpdateRiskGlobalThresholds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskGlobalThresholdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskGlobalThresholds** | [**RiskGlobalThresholds**](RiskGlobalThresholds.md) |  | 

### Return type

[**RiskGlobalThresholds**](RiskGlobalThresholds.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskProfile

> RiskProfile UpdateRiskProfile(ctx, riskProfileId).RiskProfile(riskProfile).Execute()



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
    riskProfileId := "riskProfileId_example" // string | 
    riskProfile := *openapiclient.NewRiskProfile() // RiskProfile |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskApi.UpdateRiskProfile(context.Background(), riskProfileId).RiskProfile(riskProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskApi.UpdateRiskProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskProfile`: RiskProfile
    fmt.Fprintf(os.Stdout, "Response from `RiskApi.UpdateRiskProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskProfile** | [**RiskProfile**](RiskProfile.md) |  | 

### Return type

[**RiskProfile**](RiskProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

