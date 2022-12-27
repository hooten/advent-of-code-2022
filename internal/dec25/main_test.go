package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSnaffuToDec(t *testing.T) {
	testCases := []struct {
		s string
		d int64
	}{
		{s: "10000", d: 625},
		{s: "1====", d: 313},
		{s: "2222", d: 312},
		{s: "1===============================================================================2222222222222222222222222222222222222222222222222222222", d: 312},
	}
	for _, tc := range testCases {
		require.Equal(t, tc.d, SnafuToDecimal(tc.s))
	}
}
