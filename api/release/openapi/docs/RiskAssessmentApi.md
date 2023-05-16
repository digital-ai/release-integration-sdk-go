# \RiskAssessmentApi

All URIs are relative to *http://localhost:5516*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssessment**](RiskAssessmentApi.md#GetAssessment) | **Get** /api/v1/risks/assessments/{riskAssessmentId} | 



## GetAssessment

> RiskAssessment GetAssessment(ctx, riskAssessmentId).Execute()



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
    riskAssessmentId := "riskAssessmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskAssessmentApi.GetAssessment(context.Background(), riskAssessmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskAssessmentApi.GetAssessment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssessment`: RiskAssessment
    fmt.Fprintf(os.Stdout, "Response from `RiskAssessmentApi.GetAssessment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskAssessmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssessmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskAssessment**](RiskAssessment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [patAuth](../README.md#patAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

