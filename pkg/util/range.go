package util

import "golang.org/x/exp/constraints"

func NewRange[N constraints.Integer](m, n N) []N {
	var ns []N
	for i := m; i <= n; i++ {
		ns = append(ns, i)
	}
	return ns
}
