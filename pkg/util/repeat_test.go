package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	testCases := []struct {
		x        int
		n        int
		expected []int
	}{
		{x: 8, n: 6, expected: []int{8, 8, 8, 8, 8, 8}},
	}
	for _, tc := range testCases {
		assert.Equalf(t, tc.expected, Repeat(tc.x, tc.n), "Repeat(%v, %v)", tc.x, tc.n)
	}
}

func TestRepeatString(t *testing.T) {
	repeat := strings.Repeat("a", 5)
	require.Equal(t, "aaaaa", repeat)
}
