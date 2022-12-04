package util

func Filter(f func(string) bool, xs []string) []string {
	var ys []string
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
