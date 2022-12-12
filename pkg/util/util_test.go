package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestMap(t *testing.T) {
	t.Run("string to int", func(t *testing.T) {
		actual := Map(func(s string) int {
			return int(s[0])
		}, []string{"a", "b", "c"})
		assert.Equal(t, []int{97, 98, 99}, actual)
	})
	t.Run("int to rune", func(t *testing.T) {
		actual := Map(func(i int) rune {
			return rune(i)
		}, []int{97, 98, 99})
		assert.Equal(t, []rune{'a', 'b', 'c'}, actual)
	})
	t.Run("int to int", func(t *testing.T) {
		actual := Map(func(i int) int {
			return i + 100
		}, []int{97, 98, 99})
		assert.Equal(t, []int{197, 198, 199}, actual)
	})
}

func TestFilter(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		actual := Filter(func(t string) bool {
			return len(t) > 2
		}, []string{"yo", "hi", "hello", "bonjour"})
		assert.Equal(t, []string{"hello", "bonjour"}, actual)
	})
	t.Run("int", func(t *testing.T) {
		actual := Filter(func(t int) bool {
			return t > 2
		}, []int{0, 1, 2, 3, 4, 5, 0, -1, 20})
		assert.Equal(t, []int{3, 4, 5, 20}, actual)
	})
}

func TestToSet(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		actual := ToSet([]string{"yo", "hi", "hello", "yo"})
		assert.Equal(t, map[string]bool{
			"yo":    true,
			"hi":    true,
			"hello": true,
		}, actual)
	})
}

func TestKeys(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		expected := []string{"yo", "hi", "hello"}
		set := map[string]bool{
			"yo":    true,
			"hi":    true,
			"hello": true,
		}
		assert.ElementsMatch(t, expected, Keys(set))
	})
}

func TestReduce(t *testing.T) {
	actual := Reduce(func(t int, s string) int {
		return t + int(s[0])
	}, []string{"a"}, 0)
	require.Equal(t, 97, actual)
}

func TestSplitByChar(t *testing.T) {
	file := MustReadFile("./testdata/string.txt")
	chars := SplitByChar(file)
	assert.NotNil(t, chars)
}
