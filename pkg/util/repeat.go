package util

import "golang.org/x/exp/constraints"

func Repeat[T any, N constraints.Integer](x T, n N) []T {
	xs := make([]T, n)
	for i := N(0); i < n; i++ {
		xs[i] = x
	}
	return xs
}
