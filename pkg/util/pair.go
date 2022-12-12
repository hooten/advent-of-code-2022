package util

import (
	"fmt"
	"strconv"
)

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

func NewPairFromKey(key string) (*Pair, error) {
	match, ok := RegexpMatch("(-?\\d*), (-?\\d*)", key)
	if !ok {
		return nil, fmt.Errorf("no match for %s", key)
	}
	if len(match) != 3 {
		println(match[3])
		return nil, fmt.Errorf("expected 3 matches got %v with len %d", match, len(match))
	}
	x, err := strconv.Atoi(match[1])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(match[2])
	if err != nil {
		return nil, err
	}
	return &Pair{
		X: x,
		Y: y,
	}, nil
}

func (p *Pair) Key() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
