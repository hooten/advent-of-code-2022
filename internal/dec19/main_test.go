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
