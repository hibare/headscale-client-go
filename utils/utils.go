// Package utils provides utility functions for the Headscale client.
//
//nolint:revive // Package name is acceptable for a small utility package
package utils

// ToPtr returns a pointer to the provided value v.
//
// This generic utility is useful for constructing pointer values
// for literals or values in a concise and type-safe manner.
// Example usage:
//
//	ptr := ToPtr(42) // ptr is of type *int
func ToPtr[T any](v T) *T {
	return &v
}
