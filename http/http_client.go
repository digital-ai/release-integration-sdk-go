package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"k8s.io/client-go/rest"
	"net/http"
	"net/url"
	"strings"
)

// QueryParam represents a query parameter used in HTTP requests.
type QueryParam struct {
	key   string
	value string
}

// Pair sets the key-value pair for the QueryParam.
func (q *QueryParam) Pair(key string, value string) {
	q.key = key
	q.value = value
}

// HttpClient represents an HTTP client for making HTTP requests.
// It contains the base URL for the client, the underlying HTTP client implementation, and headers to be included in each request made by the client.
type HttpClient struct {
	baseUrl string
	client  rest.HTTPClient
	headers map[string][]string
}

// RequestConfig represents a configuration for an HTTP request.
// It contains the HTTP method, request body, headers, endpoint path, and query parameters to be included in the request.
type RequestConfig struct {
	method      string
	Body        []byte
	Headers     map[string][]string
	Path        string
	QueryParams []QueryParam
}

// Client sets the underlying HTTP client for the HttpClient.
func (httpClient *HttpClient) Client(client rest.HTTPClient) {
	httpClient.client = client
}

// BaseUrl sets the base URL for the HttpClient.
func (httpClient *HttpClient) BaseUrl(baseUrl string) {
	httpClient.baseUrl = baseUrl
}

// GetBaseUrl returns the base URL of the HttpClient.
func (httpClient *HttpClient) GetBaseUrl() string {
	return httpClient.baseUrl
}

func (httpClient *HttpClient) Get(ctx context.Context, path string, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(ctx, http.MethodGet, path, nil, queryParams...)
}

func (httpClient *HttpClient) GetWithConfig(ctx context.Context, config *RequestConfig) ([]byte, error) {
	config.method = http.MethodGet
	return httpClient.sendRequestWithCustomHeaders(ctx, config)
}

func (httpClient *HttpClient) Post(ctx context.Context, path string, body []byte, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(ctx, http.MethodPost, path, body, queryParams...)
}

func (httpClient *HttpClient) PostWithConfig(ctx context.Context, config *RequestConfig) ([]byte, error) {
	config.method = http.MethodPost
	return httpClient.sendRequestWithCustomHeaders(ctx, config)
}

func (httpClient *HttpClient) Delete(ctx context.Context, path string, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(ctx, http.MethodDelete, path, nil, queryParams...)
}

func (httpClient *HttpClient) DeleteWithConfig(ctx context.Context, config *RequestConfig) ([]byte, error) {
	config.method = http.MethodDelete
	return httpClient.sendRequestWithCustomHeaders(ctx, config)
}

func (httpClient *HttpClient) Put(ctx context.Context, path string, body []byte, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(ctx, http.MethodPut, path, body, queryParams...)
}

func (httpClient *HttpClient) PutWithConfig(ctx context.Context, config *RequestConfig) ([]byte, error) {
	config.method = http.MethodPut
	return httpClient.sendRequestWithCustomHeaders(ctx, config)
}

func (httpClient *HttpClient) sendRequest(ctx context.Context, method string, path string, body []byte, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequestWithCustomHeaders(ctx, &RequestConfig{
		method:      method,
		Body:        body,
		Path:        path,
		QueryParams: queryParams,
	})
}

func (httpClient *HttpClient) sendRequestWithCustomHeaders(ctx context.Context, config *RequestConfig) ([]byte, error) {
	client := httpClient.client
	if client == nil {
		return nil, fmt.Errorf("http client is uninitialized")
	}
	theUrl := httpClient.createUrl(config.Path, config.QueryParams...)
	req, err := http.NewRequestWithContext(ctx, config.method, theUrl, bytes.NewBuffer(config.Body))
	if err != nil {
		return nil, err
	}
	setHeaders(req, httpClient.headers)
	setHeaders(req, config.Headers)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s error: %v", config.method, err)
	}

	defer func(Body io.ReadCloser) {
		if deferredErr := Body.Close(); deferredErr != nil {
			err = deferredErr
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %v", err)
	}

	//TODO: handle 3xx statuses
	if resp.StatusCode >= 299 {
		return data, fmt.Errorf("%v - %s", resp.StatusCode, string(data[:]))
	}

	return data, err
}

// setHeaders sets the headers of an HTTP request based on the provided map of headers.
func setHeaders(request *http.Request, headers map[string][]string) {
	if headers != nil {
		for header, value := range headers {
			request.Header[header] = value
		}
	}
}

// encodeQueryParams encodes a list of QueryParam objects into a URL-encoded query string.
func encodeQueryParams(params []QueryParam) string {
	values := make(url.Values)
	for _, param := range params {
		values.Add(param.key, param.value)
	}
	return values.Encode()
}

// createUrl creates a complete URL for an API endpoint based on the HttpClient's base URL and provided API path and query parameters.
func (httpClient *HttpClient) createUrl(api string, params ...QueryParam) string {
	host := httpClient.baseUrl

	if strings.HasSuffix(host, "/") {
		host += api
	} else {
		host += "/" + api
	}
	if len(params) > 0 {
		return host + "?" + encodeQueryParams(params)
	}
	return host
}
