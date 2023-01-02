package slices

import (
	"golang.org/x/exp/constraints"
	"sort"
)

func Sort[X constraints.Ordered](xs []X) []X {
	sorted := make([]X, len(xs))
	copy(sorted, xs)
	sort.Slice(
		sorted,
		func(i, j int) bool {
			return sorted[i] < sorted[j]
		},
	)
	return sorted
}
