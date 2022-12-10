package util

import (
	"log"
	"strconv"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	file := MustReadFile("./testdata/matrix.txt")
	if file == "" {
		t.Fatal("empty")
	}
	matrix := NewMatrixFromFile(file, "", func(s string) int {
		atoi, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return atoi
	})
	if matrix == nil {
		t.Fatal("nil")
	}
}
