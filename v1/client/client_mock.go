package client

import (
	"github.com/hibare/headscale-client-go/v1/apikeys"
	"github.com/hibare/headscale-client-go/v1/nodes"
	"github.com/hibare/headscale-client-go/v1/policy"
	"github.com/hibare/headscale-client-go/v1/preauthkeys"
	"github.com/hibare/headscale-client-go/v1/users"
	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of ClientInterface for testing.
type MockClient struct {
	mock.Mock
}

// APIKeys returns the mock APIKeyResource for managing API keys.
func (m *MockClient) APIKeys() apikeys.APIKeyResourceInterface {
	args := m.Called()
	return args.Get(0).(apikeys.APIKeyResourceInterface) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Nodes returns the mock NodeResource for managing nodes.
func (m *MockClient) Nodes() nodes.NodeResourceInterface {
	args := m.Called()
	return args.Get(0).(nodes.NodeResourceInterface) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Policy returns the mock PolicyResource for managing policies.
func (m *MockClient) Policy() policy.PolicyResourceInterface {
	args := m.Called()
	return args.Get(0).(policy.PolicyResourceInterface) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Users returns the mock UserResource for managing users.
func (m *MockClient) Users() users.UserResourceInterface {
	args := m.Called()
	return args.Get(0).(users.UserResourceInterface) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// PreAuthKeys returns the mock PreAuthKeyResource for managing pre-auth keys.
func (m *MockClient) PreAuthKeys() preauthkeys.PreAuthKeyResourceInterface {
	args := m.Called()
	return args.Get(0).(preauthkeys.PreAuthKeyResourceInterface) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}
