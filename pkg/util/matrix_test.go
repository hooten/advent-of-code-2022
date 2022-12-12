package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	file := MustReadFile("./testdata/matrix.txt")
	if file == "" {
		t.Fatal("empty")
	}
	matrix := NewMatrixFromFile(file, "", MustAtoi)
	if matrix == nil {
		t.Fatal("nil")
	}
	require.Len(t, matrix, 99)
	require.Len(t, matrix[0], 99)
}
