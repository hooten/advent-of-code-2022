package math

import "github.com/hooten/advent-of-code-2022/pkg/util"

func Sum(xs ...int64) int64 {
	if len(xs) == 0 {
		return 0
	}
	return xs[0] + Sum(xs[1:]...)
}

func Max(xs ...int64) int64 {
	return util.Reduce(func(max int64, sum int64) int64 {
		if sum > max {
			return sum
		}
		return max
	}, xs, 0)
}
