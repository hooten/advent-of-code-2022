package util

import "fmt"

type Pair struct {
	X int
	Y int
}

func NewPair(x, y int) *Pair {
	return &Pair{
		X: x,
		Y: y,
	}
}

func (p *Pair) Key() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}
