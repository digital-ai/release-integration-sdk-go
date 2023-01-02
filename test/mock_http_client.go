package test

import (
	"io"
	"k8s.io/client-go/rest"
	"net/http"
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
	return nil
}

type MockResult struct {
	Method     string
	Path       string
	Filename   string
	StatusCode int
}

type MockHttpClient struct {
	mocks map[string]map[string]*MockBody
}

func NewMockHttpClient(mocks []MockResult) rest.HTTPClient {
	mockBodiesMap := make(map[string]map[string]*MockBody)
	for _, mock := range mocks {
		ref := mockBodiesMap[mock.Method]
		if ref == nil {
			ref = make(map[string]*MockBody)
		}
		ref[mock.Path] = &MockBody{
			filename:   mock.Filename,
			statusCode: mock.StatusCode,
		}
		mockBodiesMap[mock.Method] = ref
	}
	return &MockHttpClient{
		mocks: mockBodiesMap,
	}
}

func (c MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	mock := c.mocks[req.Method][req.URL.Path]
	return &http.Response{
		Body:       mock,
		StatusCode: mock.statusCode,
	}, nil
}
