package requests

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stretchr/testify/mock"
)

// MockRequest is a reusable mock for RequestInterface for use in tests.
type MockRequest struct {
	mock.Mock
}

// BuildURL is a mock for the BuildURL method.
func (m *MockRequest) BuildURL(pathParts ...any) *url.URL {
	args := m.Called(pathParts...)
	return args.Get(0).(*url.URL) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// BuildRequest is a mock for the BuildRequest method.
func (m *MockRequest) BuildRequest(ctx context.Context, method string, uri *url.URL, opt RequestOptions) (*http.Request, error) {
	args := m.Called(ctx, method, uri, opt)
	return args.Get(0).(*http.Request), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Do is a mock for the Do method.
func (m *MockRequest) Do(ctx context.Context, req *http.Request, v interface{}) error {
	args := m.Called(ctx, req, v)
	return args.Error(0)
}
