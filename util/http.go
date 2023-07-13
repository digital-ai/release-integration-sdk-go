package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// GetContentFromUrl retrieves the content from the specified URL.
// It sends an HTTP GET request to the URL and returns the response body as a byte slice.
// The function handles the HTTP response status code and error conditions appropriately.
// If the request is successful (status code 200), the response body is returned without errors.
// Otherwise, an error is returned with relevant information.
func GetContentFromUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %v", err)
	}

	return data, nil
}
