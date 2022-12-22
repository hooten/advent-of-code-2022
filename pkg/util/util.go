package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Concat[T any](lists ...[]T) []T {
	var concatenated []T
	for _, list := range lists {
		concatenated = append(concatenated, list...)
	}
	return concatenated
}

func Filter[T any](f func(T) bool, xs []T) []T {
	var ys []T
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

func Map[E any, T any](f func(E) T, xs []E) []T {
	var ys []T
	for _, x := range xs {
		elem := f(x)
		ys = append(ys, elem)
	}
	return ys
}

func MapWithIndex[E any, T any](f func(E, int) T, xs []E) []T {
	var ys []T
	for i, x := range xs {
		elem := f(x, i)
		ys = append(ys, elem)
	}
	return ys
}

func Take[T any](n int, xs []T) []T {
	var ys []T
	for i := 0; i < n; i++ {
		if i >= len(xs) {
			break
		}
		ys = append(ys, xs[i])
	}
	return ys
}

func FlatMap[E any, T any](f func(E) []T, xs []E) []T {
	var ys []T
	for _, x := range xs {
		elem := f(x)
		ys = append(ys, elem...)
	}
	return ys
}

func Reduce[E any, T any](f func(T, E) T, xs []E, init T) T {
	var acc = init
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

func ToSet[T comparable](xs []T) map[T]bool {
	m := map[T]bool{}
	for _, x := range xs {
		m[x] = true
	}
	return m
}

func Contains[T comparable](xs []T, elem T) bool {
	for _, x := range xs {
		if x == elem {
			return true
		}
	}
	return false
}

func HasKey[K comparable, V any](m map[K]V, key K) bool {
	for mapKey := range m {
		if mapKey == key {
			return true
		}
	}
	return false
}

func HasValue[K comparable, V comparable](m map[K]V, val V) bool {
	for _, mapVal := range m {
		if mapVal == val {
			return true
		}
	}
	return false
}

func Keys[K comparable, V any](m map[K]V) []K {
	var xs []K
	for key := range m {
		xs = append(xs, key)
	}
	return xs
}

func Assoc[K comparable, V any](m map[K]V, key K, val V) map[K]V {
	n := map[K]V{}
	for k, v := range m {
		n[k] = v
	}
	n[key] = val
	return n
}

func SelectKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	set := ToSet(keys)
	selected := map[K]V{}
	for key, val := range m {
		if _, contains := set[key]; contains {
			selected[key] = val
		}
	}
	return selected
}

func Values[K comparable, V any](m map[K]V) []V {
	var xs []V
	for _, val := range m {
		xs = append(xs, val)
	}
	return xs
}

func InitializeMap[K comparable, V any](keys []K, initial V) map[K]V {
	var m map[K]V
	for _, key := range keys {
		m[key] = initial
	}
	return m
}

func MustReadFile(name string) string {
	bytes, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func SplitByBlock(file string) []string {
	return splitBy(file, "\n\n")
}

func SplitByLine(file string) []string {
	return splitBy(file, "\n")
}

func SplitByChar(file string) []string {
	return splitBy(file, "")
}

func splitBy(file string, sep string) []string {
	raw := strings.Split(file, sep)
	return Filter(func(s string) bool {
		return s != "\n" && s != ""
	}, raw)
}

func MustAtoi64(s string) int64 {
	atoi, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		log.Fatal(s, err)
	}
	return atoi
}

func MustAtoi(s string) int {
	atoi, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		log.Fatal(s, err)
	}
	return int(atoi)
}

func Ternary[T any](ok bool, a T, b T) T {
	if ok {
		return a
	}
	return b
}
