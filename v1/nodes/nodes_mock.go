package nodes

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockNodeResource is a mock implementation of NodeResourceInterface for testing.
type MockNodeResource struct {
	mock.Mock
}

// List returns a mock list of nodes from the Headscale.
func (m *MockNodeResource) List(ctx context.Context, filter NodeListFilter) (NodesResponse, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(NodesResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Get returns a mock node from the Headscale.
func (m *MockNodeResource) Get(ctx context.Context, id string) (NodeResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(NodeResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Register registers a mock node with the Headscale.
func (m *MockNodeResource) Register(ctx context.Context, user, key string) (NodeResponse, error) {
	args := m.Called(ctx, user, key)
	return args.Get(0).(NodeResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// Delete deletes a mock node from the Headscale.
func (m *MockNodeResource) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Expire expires a mock node from the Headscale.
func (m *MockNodeResource) Expire(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Rename renames a mock node from the Headscale.
func (m *MockNodeResource) Rename(ctx context.Context, id, name string) (NodeResponse, error) {
	args := m.Called(ctx, id, name)
	return args.Get(0).(NodeResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// AddTags adds tags to a mock node from the Headscale.
func (m *MockNodeResource) AddTags(ctx context.Context, id string, tags []string) (NodeResponse, error) {
	args := m.Called(ctx, id, tags)
	return args.Get(0).(NodeResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}

// BackFillIP backfills the IP of a mock node from the Headscale.
func (m *MockNodeResource) BackFillIP(ctx context.Context, id string) (BackfillIPsResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(BackfillIPsResponse), args.Error(1) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
}
