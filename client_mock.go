package headscale

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of the HeadscaleClientInterface.
type MockClient struct {
	mock.Mock
}

// NewMockClient creates a new instance of MockClient.
func NewMockClient() ClientInterface {
	return &MockClient{}
}

func (m *MockClient) buildURL(pathParts ...any) *url.URL {
	args := m.Called(pathParts...)
	if u, ok := args.Get(0).(*url.URL); ok {
		return u
	}
	panic("buildURL: unexpected return type")
}

func (m *MockClient) buildRequest(ctx context.Context, method string, uri *url.URL, opt requestOptions) (*http.Request, error) {
	args := m.Called(ctx, method, uri, opt)
	req, ok := args.Get(0).(*http.Request)
	if !ok {
		return nil, errors.New("buildRequest: unexpected return type")
	}
	return req, args.Error(1)
}

func (m *MockClient) do(_ context.Context, req *http.Request, v interface{}) error {
	args := m.Called(req, v)
	return args.Error(0)
}

// APIKeys returns the mock APIKeyResource for managing API keys.
func (m *MockClient) APIKeys() *APIKeyResource {
	args := m.Called()
	if resource, ok := args.Get(0).(*APIKeyResource); ok {
		return resource
	}
	panic("APIKeys: unexpected return type")
}

// Nodes returns the mock NodeResource for managing nodes.
func (m *MockClient) Nodes() *NodeResource {
	args := m.Called()
	if resource, ok := args.Get(0).(*NodeResource); ok {
		return resource
	}
	panic("Nodes: unexpected return type")
}

// Policy returns the mock PolicyResource for managing policies.
func (m *MockClient) Policy() *PolicyResource {
	args := m.Called()
	if resource, ok := args.Get(0).(*PolicyResource); ok {
		return resource
	}
	panic("Policy: unexpected return type")
}

// Users returns the mock UserResource for managing users.
func (m *MockClient) Users() *UserResource {
	args := m.Called()
	if resource, ok := args.Get(0).(*UserResource); ok {
		return resource
	}
	panic("Users: unexpected return type")
}

// Routes returns the mock RoutesResource for managing routes.
func (m *MockClient) Routes() *RoutesResource {
	args := m.Called()
	if resource, ok := args.Get(0).(*RoutesResource); ok {
		return resource
	}
	panic("Routes: unexpected return type")
}

// PreAuthKeys returns the mock PreAuthKeyResource for managing pre-auth keys.
func (m *MockClient) PreAuthKeys() *PreAuthKeyResource {
	args := m.Called()
	if resource, ok := args.Get(0).(*PreAuthKeyResource); ok {
		return resource
	}
	panic("PreAuthKeys: unexpected return type")
}
