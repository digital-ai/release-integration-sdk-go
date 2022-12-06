package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var client *http.Client

func SetClient(c *http.Client) {
	client = c
}

func GetClient() *http.Client {
	return client
}

func Get(url string) ([]byte, error) {
	return SendRequest(http.MethodGet, url, nil)
}

func Post(url string, body []byte) ([]byte, error) {
	return SendRequest(http.MethodPost, url, body)
}

func Delete(url string) ([]byte, error) {
	return SendRequest(http.MethodDelete, url, nil)
}

func Put(url string, body []byte) ([]byte, error) {
	return SendRequest(http.MethodPut, url, body)
}

func SendRequest(method string, url string, body []byte) ([]byte, error) {
	if client == nil {
		return nil, fmt.Errorf("Http client is uninitialized")
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
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

func CreateUrl(host string, api string) string {
	if strings.HasSuffix(host, "/") {
		host += api
	} else {
		host += "/" + api
	}
	return host
}
