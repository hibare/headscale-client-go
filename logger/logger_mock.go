package logger

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockLogger is a mock implementation of the Logger interface for testing.
type MockLogger struct {
	mock.Mock
}

// Info logs an mock info message.
func (m *MockLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	m.Called(ctx, msg, keysAndValues)
}

// Error logs an mock error message.
func (m *MockLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	m.Called(ctx, msg, keysAndValues)
}

// Warn logs an mock warning message.
func (m *MockLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	m.Called(ctx, msg, keysAndValues)
}

// Debug logs an mock debug message.
func (m *MockLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	m.Called(ctx, msg, keysAndValues)
}
