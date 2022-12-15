package util

import "golang.org/x/exp/constraints"

func NewStep[N constraints.Integer](m, n N) []N {
	var ns []N
	for i := m; i <= n; i++ {
		ns = append(ns, i)
	}
	return ns
}

func NewRange[N constraints.Integer](start, end N) []N {
	ints := make([]N, 0)
	if start <= end {
		for i := start; i <= end; i++ {
			ints = append(ints, i)
		}
		return ints
	}
	for i := start; i >= end; i-- {
		ints = append(ints, i)
	}
	return ints
}
