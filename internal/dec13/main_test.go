package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseListOrInt(t *testing.T) {
	testCaes := []struct {
		input string
	}{
		{
			input: "[[[[4,7,3],[6,0],1,6],[6,1,4,5]],[]]",
		},
		{
			input: "[[[]]]",
		},
		{
			input: "[[8]]",
		},
		{
			input: "[10]",
		},
		{
			input: "[[3,10],[[4,7,[1,8,6,10,8],5]]]",
		},
	}
	for _, tc := range testCaes {
		actual := ParseListOrInt(tc.input)
		require.Equal(t, tc.input, actual.String())
	}

}
