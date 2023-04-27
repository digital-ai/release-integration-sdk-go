package http

import (
	"bytes"
	"fmt"
	"io"
	"k8s.io/client-go/rest"
	"net/http"
	"net/url"
	"strings"
)

type QueryParam struct {
	key   string
	value string
}

func (q *QueryParam) Pair(key string, value string) {
	q.key = key
	q.value = value
}

type HttpClient struct {
	baseUrl string
	client  rest.HTTPClient
	headers map[string][]string
}

type RequestConfig struct {
	method      string
	Body        []byte
	Headers     map[string][]string
	Path        string
	QueryParams []QueryParam
}

func (httpClient *HttpClient) Client(client rest.HTTPClient) {
	httpClient.client = client
}

func (httpClient *HttpClient) BaseUrl(baseUrl string) {
	httpClient.baseUrl = baseUrl
}

func (httpClient *HttpClient) GetBaseUrl() string {
	return httpClient.baseUrl
}

func (httpClient *HttpClient) Get(path string, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(http.MethodGet, path, nil, queryParams...)
}

func (httpClient *HttpClient) GetWithConfig(config *RequestConfig) ([]byte, error) {
	config.method = http.MethodGet
	return httpClient.sendRequestWithCustomHeaders(config)
}

func (httpClient *HttpClient) Post(path string, body []byte, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(http.MethodPost, path, body, queryParams...)
}

func (httpClient *HttpClient) PostWithConfig(config *RequestConfig) ([]byte, error) {
	config.method = http.MethodPost
	return httpClient.sendRequestWithCustomHeaders(config)
}

func (httpClient *HttpClient) Delete(path string, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(http.MethodDelete, path, nil, queryParams...)
}

func (httpClient *HttpClient) DeleteWithConfig(config *RequestConfig) ([]byte, error) {
	config.method = http.MethodDelete
	return httpClient.sendRequestWithCustomHeaders(config)
}

func (httpClient *HttpClient) Put(path string, body []byte, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(http.MethodPut, path, body, queryParams...)
}

func (httpClient *HttpClient) PutWithConfig(config *RequestConfig) ([]byte, error) {
	config.method = http.MethodPut
	return httpClient.sendRequestWithCustomHeaders(config)
}

func (httpClient *HttpClient) sendRequest(method string, path string, body []byte, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequestWithCustomHeaders(&RequestConfig{
		method:      method,
		Body:        body,
		Path:        path,
		QueryParams: queryParams,
	})
}

func (httpClient *HttpClient) sendRequestWithCustomHeaders(config *RequestConfig) ([]byte, error) {
	client := httpClient.client
	if client == nil {
		return nil, fmt.Errorf("http client is uninitialized")
	}
	theUrl := httpClient.createUrl(config.Path, config.QueryParams...)
	req, err := http.NewRequest(config.method, theUrl, bytes.NewBuffer(config.Body))
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

func setHeaders(request *http.Request, headers map[string][]string) {
	if headers != nil {
		for header, value := range headers {
			request.Header[header] = value
		}
	}
}

func encodeQueryParams(params []QueryParam) string {
	values := make(url.Values)
	for _, param := range params {
		values.Add(param.key, param.value)
	}
	return values.Encode()
}

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
