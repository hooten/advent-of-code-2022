package util

func Expand[T any](ch chan T) []T {
	var ts []T
	for t := range ch {
		ts = append(ts, t)
	}
	return ts
}
