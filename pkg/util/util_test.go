package util

import "testing"
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
