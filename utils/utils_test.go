package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToPtr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		val := 42
		ptr := ToPtr(val)
		require.NotNil(t, ptr)
		require.Equal(t, val, *ptr)
	})

	t.Run("string", func(t *testing.T) {
		val := "hello"
		ptr := ToPtr(val)
		require.NotNil(t, ptr)
		require.Equal(t, val, *ptr)
	})

	t.Run("struct", func(t *testing.T) {
		type sample struct {
			A int
			B string
		}
		val := sample{A: 1, B: "test"}
		ptr := ToPtr(val)
		require.NotNil(t, ptr)
		require.Equal(t, val, *ptr)
	})
}
