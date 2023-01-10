package test

import (
	"io"
	"k8s.io/client-go/rest"
	"net/http"
	"net/url"
	"os"
)

type MockBody struct {
	offset     int
	response   []byte
	filename   string
	statusCode int
}

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

func (b *MockBody) Close() error {
	b.offset = 0
	return nil
}

type MockResult struct {
	Method     string
	Path       string
	Filename   string
	Data       []byte
	StatusCode int
}

type MockHttpClient struct {
	mocks map[string]map[string]*MockBody
}

var the404Response = &http.Response{
	Body: &MockBody{
		response: []byte("Mock not found - failing with 404"),
	},
	StatusCode: 404,
}

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

func getPath(url *url.URL) string {
	if url.RawPath != "" {
		return url.RawPath
	} else {
		return url.Path
	}
}

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
