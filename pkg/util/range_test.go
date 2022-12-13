package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRange(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		actual := NewRange(1, 5)
		require.Equal(t, []int{1, 2, 3, 4, 5}, actual)
	})
	t.Run("zero", func(t *testing.T) {
		actual := NewRange(0, 0)
		require.Equal(t, []int{0}, actual)
	})

	t.Run("bad", func(t *testing.T) {
		actual := NewRange(0, -10)
		require.Equal(t, []int(nil), actual)
	})

}
