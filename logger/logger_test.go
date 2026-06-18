package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w //nolint:reassign // reason: capturing stdout for testing

	f()

	_ = w.Close()
	os.Stdout = old //nolint:reassign // reason: restoring stdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestNewDefaultLogger(t *testing.T) {
	tests := []struct {
		name          string
		level         LogLevel
		logFunc       func(logger *DefaultLogger, ctx context.Context)
		expectedLevel string
		expectedMsg   string
		expectedKey   string
		expectedValue string
		shouldLog     bool
	}{
		{
			name:  "Debug level - Debug log",
			level: LevelDebug,
			logFunc: func(logger *DefaultLogger, ctx context.Context) {
				logger.Debug(ctx, "debug message", "key1", "val1")
			},
			expectedLevel: "DEBUG",
			expectedMsg:   "debug message",
			expectedKey:   "key1",
			expectedValue: "val1",
			shouldLog:     true,
		},
		{
			name:  "Info level - Debug log (filtered)",
			level: LevelInfo,
			logFunc: func(logger *DefaultLogger, ctx context.Context) {
				logger.Debug(ctx, "debug message", "key1", "val1")
			},
			shouldLog: false,
		},
		{
			name:  "Info level - Info log",
			level: LevelInfo,
			logFunc: func(logger *DefaultLogger, ctx context.Context) {
				logger.Info(ctx, "info message")
			},
			expectedLevel: "INFO",
			expectedMsg:   "info message",
			shouldLog:     true,
		},
		{
			name:  "Warn level - Warn log",
			level: LevelWarn,
			logFunc: func(logger *DefaultLogger, ctx context.Context) {
				logger.Warn(ctx, "warn message")
			},
			expectedLevel: "WARN",
			expectedMsg:   "warn message",
			shouldLog:     true,
		},
		{
			name:  "Error level - Error log",
			level: LevelError,
			logFunc: func(logger *DefaultLogger, ctx context.Context) {
				logger.Error(ctx, "error message")
			},
			expectedLevel: "ERROR",
			expectedMsg:   "error message",
			shouldLog:     true,
		},
		{
			name:  "Unknown level - defaults to Info",
			level: LogLevel(99),
			logFunc: func(logger *DefaultLogger, ctx context.Context) {
				logger.Info(ctx, "unknown level log")
			},
			expectedLevel: "INFO",
			expectedMsg:   "unknown level log",
			shouldLog:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var logger *DefaultLogger
			ctx := t.Context()

			output := captureStdout(func() {
				logger = NewDefaultLogger(tt.level)
				tt.logFunc(logger, ctx)
			})

			if !tt.shouldLog {
				assert.Empty(t, output)
				return
			}

			require.NotEmpty(t, output)
			var parsed map[string]any
			err := json.Unmarshal([]byte(output), &parsed)
			require.NoError(t, err)

			assert.Equal(t, tt.expectedLevel, parsed["level"])
			assert.Equal(t, tt.expectedMsg, parsed["msg"])
			if tt.expectedKey != "" {
				assert.Equal(t, tt.expectedValue, parsed[tt.expectedKey])
			}
		})
	}
}

func TestMockLogger(t *testing.T) {
	mockLogger := new(MockLogger)
	ctx := t.Context()

	mockLogger.On("Info", ctx, "info message", mock.Anything).Return()
	mockLogger.On("Warn", ctx, "warn message", mock.Anything).Return()
	mockLogger.On("Error", ctx, "error message", mock.Anything).Return()
	mockLogger.On("Debug", ctx, "debug message", mock.Anything).Return()

	mockLogger.Info(ctx, "info message")
	mockLogger.Warn(ctx, "warn message")
	mockLogger.Error(ctx, "error message")
	mockLogger.Debug(ctx, "debug message")

	mockLogger.AssertExpectations(t)
}

func TestMockLogger_DirectExecution(t *testing.T) {
	// Simple sanity test for mock implementations
	ml := &MockLogger{}
	ctx := t.Context()

	ml.On("Info", ctx, "hello", "a", "b").Return()
	ml.Info(ctx, "hello", "a", "b")
	ml.AssertExpectations(t)
}
