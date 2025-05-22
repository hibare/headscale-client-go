package policy

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockPolicyResource is a mock implementation of PolicyResourceInterface for testing.
type MockPolicyResource struct {
	mock.Mock
}

// Get returns a mock policy from the Headscale.
func (m *MockPolicyResource) Get(ctx context.Context) (Policy, error) {
	args := m.Called(ctx)
	return args.Get(0).(Policy), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Update updates a mock policy from the Headscale.
func (m *MockPolicyResource) Update(ctx context.Context, policyStr string) (UpdatePolicyResponse, error) {
	args := m.Called(ctx, policyStr)
	return args.Get(0).(UpdatePolicyResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}
