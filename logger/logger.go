// Package logger provides a logger for the Headscale client.
package logger

import (
	"context"
	"log/slog"
	"os"
)

// LogLevel represents the level of logging.
type LogLevel int

const (
	// LevelDebug is the debug log level.
	LevelDebug LogLevel = iota

	// LevelInfo is the info log level.
	LevelInfo

	// LevelWarn is the warn log level.
	LevelWarn

	// LevelError is the error log level.
	LevelError

	// LevelCritical is the critical log level.
	LevelCritical
)

// Logger is an interface for logging messages.
type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}

// DefaultLogger is a simple implementation of the Logger interface.
type DefaultLogger struct {
	logger *slog.Logger
}

// NewDefaultLogger creates a new DefaultLogger.
func NewDefaultLogger(level LogLevel) *DefaultLogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: mapLogLevel(level)})
	return &DefaultLogger{
		logger: slog.New(handler),
	}
}

// mapLogLevel maps CustomLogLevel to slog.Level.
func mapLogLevel(level LogLevel) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	case LevelCritical:
		return slog.LevelError // Assuming Critical maps to Error for simplicity
	default:
		return slog.LevelInfo
	}
}

// Info logs an informational message.
func (l *DefaultLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logger.InfoContext(ctx, msg, keysAndValues...)
}

// Error logs an error message.
func (l *DefaultLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logger.ErrorContext(ctx, msg, keysAndValues...)
}

// Warn logs a warning message.
func (l *DefaultLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logger.WarnContext(ctx, msg, keysAndValues...)
}

// Debug logs a debug message.
func (l *DefaultLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.logger.DebugContext(ctx, msg, keysAndValues...)
}
