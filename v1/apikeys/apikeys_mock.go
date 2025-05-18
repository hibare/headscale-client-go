package apikeys

import (
	"context"
	"time"

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
func (m *MockAPIKeyResource) Create(ctx context.Context, expiration time.Time) (CreateAPIKeyResponse, error) {
	args := m.Called(ctx, expiration)
	return args.Get(0).(CreateAPIKeyResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Expire expires a mock API key from the Headscale.
func (m *MockAPIKeyResource) Expire(ctx context.Context, prefix string) error {
	args := m.Called(ctx, prefix)
	return args.Error(0)
}

// Delete deletes a mock API key from the Headscale.
func (m *MockAPIKeyResource) Delete(ctx context.Context, prefix string) error {
	args := m.Called(ctx, prefix)
	return args.Error(0)
}
