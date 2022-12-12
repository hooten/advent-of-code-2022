package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPairFromKey(t *testing.T) {
	testCases := []struct {
		key      string
		expected Pair
	}{
		{
			key: "(0, 0)",
			expected: Pair{
				X: 0,
				Y: 0,
			},
		},
		{
			key: "(19, -137)",
			expected: Pair{
				X: 19,
				Y: -137,
			},
		},
		{
			key: "(-1, -1)",
			expected: Pair{
				X: -1,
				Y: -1,
			},
		},
	}
	for _, tc := range testCases {
		actual, err := NewPairFromKey(tc.key)
		require.Nil(t, err)
		require.Equal(t, tc.expected, *actual)
	}

}
