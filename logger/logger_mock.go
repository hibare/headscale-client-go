package logger

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockLogger is a mock implementation of the Logger interface for testing.
type MockLogger struct {
	mock.Mock
}

// Info logs a mock info message.
func (m *MockLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	m.Called(ctx, msg, keysAndValues)
}

// Error logs a mock error message.
func (m *MockLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	m.Called(ctx, msg, keysAndValues)
}

// Warn logs a mock warning message.
func (m *MockLogger) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	m.Called(ctx, msg, keysAndValues)
}

// Debug logs a mock debug message.
func (m *MockLogger) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	m.Called(ctx, msg, keysAndValues)
}
