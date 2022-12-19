package main

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	var xs = []int{1}
	for b := 0; b < len(xs); b++ {
		fmt.Println(xs[b])
		if xs[b] == 1 {
			xs = append(xs, 2)
		}
	}
}
