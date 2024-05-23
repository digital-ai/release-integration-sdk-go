package test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"k8s.io/client-go/rest"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strings"
)

// Mock is a definition of mock consisting of mocked path, queries and mock response body which is returned on matched mock pattern
type Mock struct {
	method       string
	path         string
	queryParams  url.Values
	responseBody *MockBody
}

// flattenValues flat values map to be represented as string, with keeping in mind order of keys and paired values array
func flattenValues(values url.Values) string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var flattened []string
	for _, key := range keys {
		params := values[key]
		sort.Strings(params)
		flattened = append(flattened, fmt.Sprintf("%s:%s", key, strings.Join(params, ",")))
	}

	return strings.Join(flattened, ";")
}

// hash creates a unique hash key from method, path and query params map
func (m Mock) hash() string {
	return calculateHash(m.method, m.path, m.queryParams)
}

// calculateHash creates a unique hash key from method, path and query params map
func calculateHash(method string, path string, queryParams url.Values) string {
	combinedStr := fmt.Sprintf("%s::%s", method, path)
	flattenedValues := ""
	if len(queryParams) > 0 {
		flattenedValues = flattenValues(queryParams)
	}
	finalStr := fmt.Sprintf("%s::%s", combinedStr, flattenedValues)
	hash := sha256.Sum256([]byte(finalStr))
	hashKey := hex.EncodeToString(hash[:])
	return hashKey
}

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

func (mr MockResult) ToMock() *Mock {
	mockBody := &MockBody{statusCode: mr.StatusCode}
	if mr.Filename != "" {
		mockBody.filename = mr.Filename
	} else {
		mockBody.response = mr.Data
	}
	return &Mock{
		method:       mr.Method,
		path:         mr.Path,
		queryParams:  mr.QueryParams,
		responseBody: mockBody,
	}
}

// MockHttpClient is a mock implementation of the rest.HTTPClient interface.
type MockHttpClient struct {
	mocks map[string]*Mock
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
	mockMap := make(map[string]*Mock)
	for _, mockResults := range mocks {
		mock := mockResults.ToMock()
		mockMap[mock.hash()] = mock
	}
	return &MockHttpClient{
		mocks: mockMap,
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
	// TODO hash doesn't have the request params and the request has we need to fallback
	hash := calculateHash(req.Method, path, requestQueryParams)
	mockData, exists := c.mocks[hash]

	skipQueryCheck := len(mockData.queryParams) == 0
	// fallback and check if we ignore query params check
	if skipQueryCheck && !exists {
		hash = calculateHash(req.Method, path, nil)
		mockData, exists = c.mocks[hash]
		skipQueryCheck = true
	}
	if exists {
		if skipQueryCheck || reflect.DeepEqual(requestQueryParams, mockData.queryParams) {
			return &http.Response{
				Body:       mockData.responseBody,
				StatusCode: mockData.responseBody.statusCode,
			}, nil
		}
	}
	return the404Response, nil
}
