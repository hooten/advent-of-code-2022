package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAfford(t *testing.T) {
	pack := NewPack(Blueprint{})
	pack.Ore = 11
	canBuy := pack.Affordable("ore")
	require.Equal(t, 5, canBuy)

}

func TestOptimizedBestCaseGeodes(t *testing.T) {
	pack := Pack{Geode: 0}
	require.Equal(t, 300, pack.OptimizedBestCaseGeodes(24))
	require.Equal(t, 1, pack.OptimizedBestCaseGeodes(1))
}
