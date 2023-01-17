package http

import (
	"context"
	"encoding/json"
	"fmt"
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
	client  rest.RESTClient
}

func (httpClient *HttpClient) Client(client rest.RESTClient) {
	httpClient.client = client
}

func (httpClient *HttpClient) BaseUrl(baseUrl string) {
	httpClient.baseUrl = baseUrl
}

func (httpClient *HttpClient) GetBaseUrl() string {
	return httpClient.baseUrl
}

func (httpClient HttpClient) Get(path string, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(http.MethodGet, path, nil, queryParams...)
}

func (httpClient HttpClient) Post(path string, body []byte, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(http.MethodPost, path, body, queryParams...)
}

func (httpClient HttpClient) Delete(path string, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(http.MethodDelete, path, nil, queryParams...)
}

func (httpClient HttpClient) Put(path string, body []byte, queryParams ...QueryParam) ([]byte, error) {
	return httpClient.sendRequest(http.MethodPut, path, body, queryParams...)
}

func (httpClient HttpClient) sendRequest(method string, path string, body []byte, queryParams ...QueryParam) ([]byte, error) {
	client := httpClient.client
	if client.Client == nil {
		return nil, fmt.Errorf("http client is uninitialized")
	}
	theUrl := httpClient.createUrl(path, queryParams...)
	req := client.Verb(method)
	req.Body(body)
	req.RequestURI(theUrl)

	result := req.Do(context.Background())
	if result.Error() != nil {
		return nil, result.Error()
	}

	data, err := result.Raw()
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	var statusCode int
	result.StatusCode(&statusCode)
	//TODO: handle 3xx status codes
	if statusCode >= 400 && statusCode <= 599 {
		return nil, fmt.Errorf("%v - %s", statusCode, string(data))
	}

	if statusCode >= 299 {
		var respError json.RawMessage
		unMarshalErr := json.Unmarshal(data, &respError)
		if unMarshalErr != nil {
			return nil, fmt.Errorf("%v - %s", statusCode, unMarshalErr)
		}
		return nil, fmt.Errorf("%v - %s", statusCode, respError)
	}

	return data, nil
}

func encodeQueryParams(params []QueryParam) string {
	values := make(url.Values)
	for _, param := range params {
		values.Add(param.key, param.value)
	}
	return values.Encode()
}

func (httpClient HttpClient) createUrl(api string, params ...QueryParam) string {
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
