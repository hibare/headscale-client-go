package preauthkeys

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockPreAuthKeyResource is a mock implementation of PreAuthKeyResourceInterface for testing.
type MockPreAuthKeyResource struct {
	mock.Mock
}

// List returns a mock list of pre-auth keys from the Headscale.
func (m *MockPreAuthKeyResource) List(ctx context.Context, filter PreAuthKeyListFilter) (PreAuthKeysResponse, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(PreAuthKeysResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Create creates a mock pre-auth key from the Headscale.
func (m *MockPreAuthKeyResource) Create(ctx context.Context, createPreAuthKeyRequest CreatePreAuthKeyRequest) (PreAuthKeyResponse, error) {
	args := m.Called(ctx, createPreAuthKeyRequest)
	return args.Get(0).(PreAuthKeyResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Expire expires a mock pre-auth key from the Headscale.
func (m *MockPreAuthKeyResource) Expire(ctx context.Context, user string, key string) error {
	args := m.Called(ctx, user, key)
	return args.Error(0)
}
