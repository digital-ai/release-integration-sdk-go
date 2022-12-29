package http

import (
	"bytes"
	"encoding/json"
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
}

func (httpClient *HttpClient) Client(client rest.HTTPClient) {
	httpClient.client = client
}

func (httpClient *HttpClient) BaseUrl(baseUrl string) {
	httpClient.baseUrl = baseUrl
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
	if client == nil {
		return nil, fmt.Errorf("http client is uninitialized")
	}
	theUrl := httpClient.createUrl(path, queryParams...)
	req, err := http.NewRequest(method, theUrl, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s error: %v", method, err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
		return nil, fmt.Errorf("%v - %s", resp.StatusCode, string(data))
	}

	if resp.StatusCode != http.StatusOK {
		var respError json.RawMessage
		unMarshalErr := json.Unmarshal(data, &respError)
		if unMarshalErr != nil {
			return nil, fmt.Errorf("%v - %s", resp.StatusCode, unMarshalErr)
		}
		return nil, fmt.Errorf("%v - %s", resp.StatusCode, respError)
	}

	return data, nil
}

func encodeQueryParams(params []QueryParam) string {
	var values url.Values
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
