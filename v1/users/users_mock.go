package users

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockUserResource is a mock implementation of UserResourceInterface for testing.
type MockUserResource struct {
	mock.Mock
}

// List returns a mock list of users from the Headscale.
func (m *MockUserResource) List(ctx context.Context, filter UserListFilter) (UsersResponse, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(UsersResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Create creates a mock user from the Headscale.
func (m *MockUserResource) Create(ctx context.Context, request CreateUserRequest) (UserResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(UserResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Delete deletes a mock user from the Headscale.
func (m *MockUserResource) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Rename renames a mock user from the Headscale.
func (m *MockUserResource) Rename(ctx context.Context, id, newName string) (UserResponse, error) {
	args := m.Called(ctx, id, newName)
	return args.Get(0).(UserResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}
