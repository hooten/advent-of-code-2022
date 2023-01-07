package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPriority(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{
			input:    "a",
			expected: 1,
		},
		{
			input:    "z",
			expected: 26,
		},
		{
			input:    "A",
			expected: 27,
		},
		{
			input:    "Z",
			expected: 52,
		},
	}
	for _, tc := range testCases {
		actual := Priority(tc.input)
		require.Equal(t, tc.expected, actual)
	}
}
