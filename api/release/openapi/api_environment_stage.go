/*
Digital.ai Release API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// EnvironmentStageApiService EnvironmentStageApi service
type EnvironmentStageApiService service

type ApiCreateStage3Request struct {
	ctx context.Context
	ApiService *EnvironmentStageApiService
	environmentStageForm *EnvironmentStageForm
}

func (r ApiCreateStage3Request) EnvironmentStageForm(environmentStageForm EnvironmentStageForm) ApiCreateStage3Request {
	r.environmentStageForm = &environmentStageForm
	return r
}

func (r ApiCreateStage3Request) Execute() (*EnvironmentStageView, *http.Response, error) {
	return r.ApiService.CreateStage3Execute(r)
}

/*
CreateStage3 Method for CreateStage3

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateStage3Request
*/
func (a *EnvironmentStageApiService) CreateStage3(ctx context.Context) ApiCreateStage3Request {
	return ApiCreateStage3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EnvironmentStageView
func (a *EnvironmentStageApiService) CreateStage3Execute(r ApiCreateStage3Request) (*EnvironmentStageView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentStageView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentStageApiService.CreateStage3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/environments/stages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentStageForm
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteEnvironmentStageRequest struct {
	ctx context.Context
	ApiService *EnvironmentStageApiService
	environmentStageId string
}

func (r ApiDeleteEnvironmentStageRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEnvironmentStageExecute(r)
}

/*
DeleteEnvironmentStage Method for DeleteEnvironmentStage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentStageId
 @return ApiDeleteEnvironmentStageRequest
*/
func (a *EnvironmentStageApiService) DeleteEnvironmentStage(ctx context.Context, environmentStageId string) ApiDeleteEnvironmentStageRequest {
	return ApiDeleteEnvironmentStageRequest{
		ApiService: a,
		ctx: ctx,
		environmentStageId: environmentStageId,
	}
}

// Execute executes the request
func (a *EnvironmentStageApiService) DeleteEnvironmentStageExecute(r ApiDeleteEnvironmentStageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentStageApiService.DeleteEnvironmentStage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/environments/stages/{environmentStageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentStageId"+"}", url.PathEscape(parameterValueToString(r.environmentStageId, "environmentStageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetStageByIdRequest struct {
	ctx context.Context
	ApiService *EnvironmentStageApiService
	environmentStageId string
}

func (r ApiGetStageByIdRequest) Execute() (*EnvironmentStageView, *http.Response, error) {
	return r.ApiService.GetStageByIdExecute(r)
}

/*
GetStageById Method for GetStageById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentStageId
 @return ApiGetStageByIdRequest
*/
func (a *EnvironmentStageApiService) GetStageById(ctx context.Context, environmentStageId string) ApiGetStageByIdRequest {
	return ApiGetStageByIdRequest{
		ApiService: a,
		ctx: ctx,
		environmentStageId: environmentStageId,
	}
}

// Execute executes the request
//  @return EnvironmentStageView
func (a *EnvironmentStageApiService) GetStageByIdExecute(r ApiGetStageByIdRequest) (*EnvironmentStageView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentStageView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentStageApiService.GetStageById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/environments/stages/{environmentStageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentStageId"+"}", url.PathEscape(parameterValueToString(r.environmentStageId, "environmentStageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchStagesRequest struct {
	ctx context.Context
	ApiService *EnvironmentStageApiService
	environmentStageFilters *EnvironmentStageFilters
}

func (r ApiSearchStagesRequest) EnvironmentStageFilters(environmentStageFilters EnvironmentStageFilters) ApiSearchStagesRequest {
	r.environmentStageFilters = &environmentStageFilters
	return r
}

func (r ApiSearchStagesRequest) Execute() ([]EnvironmentStageView, *http.Response, error) {
	return r.ApiService.SearchStagesExecute(r)
}

/*
SearchStages Method for SearchStages

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchStagesRequest
*/
func (a *EnvironmentStageApiService) SearchStages(ctx context.Context) ApiSearchStagesRequest {
	return ApiSearchStagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EnvironmentStageView
func (a *EnvironmentStageApiService) SearchStagesExecute(r ApiSearchStagesRequest) ([]EnvironmentStageView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EnvironmentStageView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentStageApiService.SearchStages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/environments/stages/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentStageFilters
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateStageInEnvironmentRequest struct {
	ctx context.Context
	ApiService *EnvironmentStageApiService
	environmentStageId string
	environmentStageForm *EnvironmentStageForm
}

func (r ApiUpdateStageInEnvironmentRequest) EnvironmentStageForm(environmentStageForm EnvironmentStageForm) ApiUpdateStageInEnvironmentRequest {
	r.environmentStageForm = &environmentStageForm
	return r
}

func (r ApiUpdateStageInEnvironmentRequest) Execute() (*EnvironmentStageView, *http.Response, error) {
	return r.ApiService.UpdateStageInEnvironmentExecute(r)
}

/*
UpdateStageInEnvironment Method for UpdateStageInEnvironment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentStageId
 @return ApiUpdateStageInEnvironmentRequest
*/
func (a *EnvironmentStageApiService) UpdateStageInEnvironment(ctx context.Context, environmentStageId string) ApiUpdateStageInEnvironmentRequest {
	return ApiUpdateStageInEnvironmentRequest{
		ApiService: a,
		ctx: ctx,
		environmentStageId: environmentStageId,
	}
}

// Execute executes the request
//  @return EnvironmentStageView
func (a *EnvironmentStageApiService) UpdateStageInEnvironmentExecute(r ApiUpdateStageInEnvironmentRequest) (*EnvironmentStageView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentStageView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentStageApiService.UpdateStageInEnvironment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/environments/stages/{environmentStageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentStageId"+"}", url.PathEscape(parameterValueToString(r.environmentStageId, "environmentStageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentStageForm
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}