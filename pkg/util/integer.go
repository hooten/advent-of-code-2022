package util

func Max(ints ...int) int {
	return Reduce(func(max int, i int) int {
		if max > i {
			return max
		}
		return i
	}, append([]int{}, ints...), 0)
}

func Min(ints ...int) int {
	return Reduce(func(min int, i int) int {
		if min < i {
			return min
		}
		return i
	}, append([]int{}, ints...), 0)
}