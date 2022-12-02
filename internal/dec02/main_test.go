package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGetMyHand(t *testing.T) {
	testCases := []struct {
		elfHand    string
		gamePoints int
		expected   string
	}{
		{
			elfHand:    "Paper",
			gamePoints: 0,
			expected:   "Rock",
		},
		{
			elfHand:    "Paper",
			gamePoints: 3,
			expected:   "Paper",
		},
		{
			elfHand:    "Paper",
			gamePoints: 6,
			expected:   "Scissors",
		},
		{
			elfHand:    "Scissors",
			gamePoints: 6,
			expected:   "Rock",
		},
		{
			elfHand:    "Scissors",
			gamePoints: 3,
			expected:   "Scissors",
		},
		{
			elfHand:    "Scissors",
			gamePoints: 0,
			expected:   "Paper",
		},
		{
			elfHand:    "Rock",
			gamePoints: 3,
			expected:   "Rock",
		},
		{
			elfHand:    "Rock",
			gamePoints: 6,
			expected:   "Paper",
		},
		{
			elfHand:    "Rock",
			gamePoints: 0,
			expected:   "Scissors",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.elfHand+strconv.Itoa(tc.gamePoints), func(t *testing.T) {
			actual := getMyHand(tc.elfHand, tc.gamePoints)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
