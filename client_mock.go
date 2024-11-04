package headscale

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of the HeadscaleClientInterface.
type MockClient struct {
	mock.Mock
}

func NewMockClient() HeadscaleClientInterface {
	return &MockClient{}
}

func (m *MockClient) buildURL(pathParts ...any) *url.URL {
	args := m.Called(pathParts...)
	return args.Get(0).(*url.URL)
}

func (m *MockClient) buildRequest(ctx context.Context, method string, uri *url.URL, opt requestOptions) (*http.Request, error) {
	args := m.Called(ctx, method, uri, opt)
	return args.Get(0).(*http.Request), args.Error(1)
}

func (m *MockClient) do(ctx context.Context, req *http.Request, v interface{}) error {
	args := m.Called(req, v)
	return args.Error(0)
}

// Implementing the HeadscaleClientInterface methods
func (m *MockClient) APIKeys() *APIKeyResource {
	args := m.Called()
	return args.Get(0).(*APIKeyResource)
}

func (m *MockClient) Nodes() *NodeResource {
	args := m.Called()
	return args.Get(0).(*NodeResource)
}

func (m *MockClient) Policy() *PolicyResource {
	args := m.Called()
	return args.Get(0).(*PolicyResource)
}

func (m *MockClient) Users() *UserResource {
	args := m.Called()
	return args.Get(0).(*UserResource)
}

func (m *MockClient) Routes() *RoutesResource {
	args := m.Called()
	return args.Get(0).(*RoutesResource)
}
