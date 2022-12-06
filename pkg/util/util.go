package util

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

func ToList[T comparable](set map[T]bool) []T {
	var xs []T
	for elem := range set {
		xs = append(xs, elem)
	}
	return xs
}

func Keys[K comparable, V any](m map[K]V) []K {
	var xs []K
	for key := range m {
		xs = append(xs, key)
	}
	return xs
}

func Values[K comparable, V any](m map[K]V) []V {
	var xs []V
	for _, val := range m {
		xs = append(xs, val)
	}
	return xs
}