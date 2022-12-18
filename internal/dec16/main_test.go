package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseListOrInt(t *testing.T) {
	testCaes := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "",
		},
	}
	for _, tc := range testCaes {
		actual := tc.input
		require.Equal(t, tc.expected, actual)
	}

}
