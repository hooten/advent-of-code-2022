package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegexpMatch(t *testing.T) {
	match := RegexpMatch("a (.) c", "a b c")
	require.Equal(t, []string{"a b c", "b"}, match)
	ExpectMatchesLen(match, 2)
}
