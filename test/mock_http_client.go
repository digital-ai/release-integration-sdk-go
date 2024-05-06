package test

import (
	"io"
	"k8s.io/client-go/rest"
	"net/http"
	"net/url"
	"os"
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
	Method     string
	Path       string
	Filename   string
	Data       []byte
	StatusCode int
}

// MockHttpClient is a mock implementation of the rest.HTTPClient interface.
type MockHttpClient struct {
	mocks map[string]map[string]*MockBody
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
	for _, mock := range mocks {
		ref := mockBodiesMap[mock.Method]
		if ref == nil {
			ref = make(map[string]*MockBody)
		}
		if mock.Filename != "" {
			ref[mock.Path] = &MockBody{
				filename:   mock.Filename,
				statusCode: mock.StatusCode,
			}
		} else {
			ref[mock.Path] = &MockBody{
				response:   mock.Data,
				statusCode: mock.StatusCode,
			}
		}
		mockBodiesMap[mock.Method] = ref
	}
	return &MockHttpClient{
		mocks: mockBodiesMap,
	}
}

// getPath returns the path from the provided URL.
func getPath(url *url.URL) string {
	var path string
	if url.RawPath != "" {
		path = url.RawPath
	} else {
		path = url.Path
	}
	if url.RawQuery != "" {
		return path + "?" + url.RawQuery
	}
	return path
}

// Do performs the mock HTTP request.
func (c MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	mock, exists := c.mocks[req.Method][getPath(req.URL)]
	if exists {
		return &http.Response{
			Body:       mock,
			StatusCode: mock.statusCode,
		}, nil
	}
	return the404Response, nil
}
