package test

import (
	"io"
	"k8s.io/client-go/rest"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

// MockBody is a mock implementation of the http.Response.Body interface.
type MockBody struct {
	offset     int
	response   []byte
	filename   string
	statusCode int
}

// Read reads data from the mock response body.
func (b *MockBody) Read(p []byte) (int, error) {
	if b.response == nil {
		b.response, _ = os.ReadFile(b.filename)
	}
	if b.offset >= len(b.response) {
		return 0, io.EOF
	}
	n := copy(p, b.response[b.offset:])
	b.offset += n
	return n, nil
}

// Close resets the offset of the mock response body.
func (b *MockBody) Close() error {
	b.offset = 0
	return nil
}

// MockResult represents a mocked HTTP response.
type MockResult struct {
	Method      string
	Path        string
	QueryParams url.Values
	Filename    string
	Data        []byte
	StatusCode  int
}

// MockHttpClient is a mock implementation of the rest.HTTPClient interface.
type MockHttpClient struct {
	mocks       map[string]map[string]*MockBody
	queryParams map[string]map[string]url.Values
}

// the404Response represents HTTP mock response with 404 code
var the404Response = &http.Response{
	Body: &MockBody{
		response: []byte("Mock not found - failing with 404"),
	},
	StatusCode: 404,
}

// NewMockHttpClient creates a new instance of MockHttpClient based on the provided mock results.
func NewMockHttpClient(mocks []MockResult) rest.HTTPClient {
	mockBodiesMap := make(map[string]map[string]*MockBody)
	mockQueriesMap := make(map[string]map[string]url.Values)
	for _, mock := range mocks {
		mockBodiesForMethodMap := mockBodiesMap[mock.Method]
		if mockBodiesForMethodMap == nil {
			mockBodiesForMethodMap = make(map[string]*MockBody)
		}
		mockQueriesMapForMethodMap := mockQueriesMap[mock.Method]
		if mockQueriesMapForMethodMap == nil {
			mockQueriesMapForMethodMap = make(map[string]url.Values)
		}
		if mock.Filename != "" {
			mockBodiesForMethodMap[mock.Path] = &MockBody{
				filename:   mock.Filename,
				statusCode: mock.StatusCode,
			}
		} else {
			mockBodiesForMethodMap[mock.Path] = &MockBody{
				response:   mock.Data,
				statusCode: mock.StatusCode,
			}
		}
		if len(mock.QueryParams) > 0 {
			mockQueriesMapForMethodMap[mock.Path] = mock.QueryParams
		}
		mockBodiesMap[mock.Method] = mockBodiesForMethodMap
		mockQueriesMap[mock.Method] = mockQueriesMapForMethodMap
	}
	return &MockHttpClient{
		mocks:       mockBodiesMap,
		queryParams: mockQueriesMap,
	}
}

// getPath returns the path from the provided URL.
func getPath(url *url.URL) string {
	if url.RawPath != "" {
		return url.RawPath
	} else {
		return url.Path
	}
}

// Do processes the mock HTTP request.
func (c MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	path := getPath(req.URL)
	requestQueryParams, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return nil, err
	}
	queryParams, exists := c.queryParams[req.Method][path]
	matchedQueryParams := true
	if len(queryParams) > 0 {
		matchedQueryParams = reflect.DeepEqual(requestQueryParams, queryParams)
	}

	mockedResponseBody, exists := c.mocks[req.Method][path]
	if matchedQueryParams && exists {
		return &http.Response{
			Body:       mockedResponseBody,
			StatusCode: mockedResponseBody.statusCode,
		}, nil
	}
	return the404Response, nil
}
