package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/math"
	"github.com/hooten/advent-of-code-2022/pkg/slices"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("./internal/dec01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	groupStrs := strings.Split(string(bytes), "\n\n")
	groups := util.Map(func(group string) []int64 {
		countStrs := util.Filter(func(line string) bool {
			return line != ""
		}, strings.Split(group, "\n"))
		return util.Map(func(countStr string) int64 {
			count, err := strconv.ParseInt(countStr, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			return count
		}, countStrs)
	}, groupStrs)

	sums := util.Map(func(counts []int64) int64 {
		return math.Sum(counts...)
	}, groups)

	max := math.Max(sums...)
	fmt.Printf("Part 1: The largest calorie count is %d.\n", max)

	sorted := slices.Sort(sums)
	topThree := math.Sum(sorted[len(sorted)-3:]...)
	fmt.Printf("Part 2: The sum of the top three calorie counts is %d.\n", topThree)

}
