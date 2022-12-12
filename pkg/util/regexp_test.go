package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegexpMatch(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		match, ok := RegexpMatch("a (.) c", "a b c")
		require.True(t, ok)
		require.Equal(t, []string{"a b c", "b"}, match)
		require.Len(t, match, 2)
	})
	t.Run("no match", func(t *testing.T) {
		match, ok := RegexpMatch("a (.) c", "xyz")
		require.False(t, ok)
		require.Nil(t, match)
		require.Len(t, match, 0)
	})
}

func TestRegexpMatchMultiLine(t *testing.T) {
	matches := RegexpMatchMultiLine(`Monkey (.*):
  Starting items: (.*)
  Operation: new = old * (.*)
  Test: divisible by (.*)
    If true: throw to monkey 6
    If false: throw to monkey 4
`, `Monkey 0:
  Starting items: 65, 58, 93, 57, 66
  Operation: new = old * 7
  Test: divisible by 19
    If true: throw to monkey 6
    If false: throw to monkey 4
`)
	require.NotNil(t, matches)
	require.Len(t, matches, 6)
	require.Equal(t, "0", matches[0][1])
	require.Equal(t, "65, 58, 93, 57, 66", matches[1][1])
	require.Equal(t, "* 7", matches[2][1])
	require.Equal(t, "19", matches[3][1])
}
