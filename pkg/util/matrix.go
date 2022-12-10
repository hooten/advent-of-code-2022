package util

import "strings"

func NewMatrix[T comparable](
	lines []string,
	sep string,
	f func(s string) T,
) [][]T {
	var matrix [][]T
	for i, row := range lines {
		matrix = append(matrix, []T{})
		for _, cell := range strings.Split(row, sep) {
			s := cell
			matrix[i] = append(matrix[i], f(s))
		}
	}
	return matrix
}

func NewMatrixFromFile[T comparable](
	file string,
	sep string,
	f func(s string) T,
) [][]T {
	rawLines := strings.Split(file, "\n")
	lines := Filter(
		func(s string) bool {
			return !HasElem([]string{"\n", ""}, s)
		},
		rawLines,
	)
	return NewMatrix(lines, sep, f)
}
