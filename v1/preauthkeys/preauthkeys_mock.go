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
func (m *MockPreAuthKeyResource) List(ctx context.Context) (PreAuthKeysResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).(PreAuthKeysResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Create creates a mock pre-auth key from the Headscale.
func (m *MockPreAuthKeyResource) Create(ctx context.Context, createPreAuthKeyRequest CreatePreAuthKeyRequest) (PreAuthKeyResponse, error) {
	args := m.Called(ctx, createPreAuthKeyRequest)
	return args.Get(0).(PreAuthKeyResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Expire expires a mock pre-auth key from the Headscale.
func (m *MockPreAuthKeyResource) Expire(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Delete deletes a mock pre-auth key from the Headscale.
func (m *MockPreAuthKeyResource) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
