package util

import "strings"

func NewMatrix(
	lines []string,
	sep string,
) [][]string {
	var matrix [][]string
	for i, row := range lines {
		matrix = append(matrix, []string{})
		for _, cell := range strings.Split(row, sep) {
			matrix[i] = append(matrix[i], cell)
		}
	}
	return matrix
}

func NewMatrixWithMap[T comparable](
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
	lines := SplitByLine(file)
	return NewMatrixWithMap(lines, sep, f)
}
