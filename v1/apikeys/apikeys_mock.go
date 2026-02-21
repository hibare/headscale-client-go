package apikeys

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockAPIKeyResource is a mock implementation of APIKeyResourceInterface for testing.
type MockAPIKeyResource struct {
	mock.Mock
}

// List returns a mock list of API keys from the Headscale.
func (m *MockAPIKeyResource) List(ctx context.Context) (APIKeysResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).(APIKeysResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Create creates a mock API key from the Headscale.
func (m *MockAPIKeyResource) Create(ctx context.Context, createAPIKeyRequest CreateAPIKeyRequest) (CreateAPIKeyResponse, error) {
	args := m.Called(ctx, createAPIKeyRequest)
	return args.Get(0).(CreateAPIKeyResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Expire expires a mock API key from the Headscale.
func (m *MockAPIKeyResource) Expire(ctx context.Context, prefix string) error {
	args := m.Called(ctx, prefix)
	return args.Error(0)
}

// ExpireByID expires a mock API key by ID from the Headscale.
func (m *MockAPIKeyResource) ExpireByID(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Delete deletes a mock API key from the Headscale.
func (m *MockAPIKeyResource) Delete(ctx context.Context, prefix string) error {
	args := m.Called(ctx, prefix)
	return args.Error(0)
}
