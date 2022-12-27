package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"math"
	"strings"
)

func main() {
	file := util.MustReadFile("./internal/dec25/input.txt")
	lines := util.SplitByLine(file)
	snafuNumbers := util.Map(func(line string) int64 {
		return SnafuToDecimal(line)
	}, lines)
	sum := util.Reduce(func(total int64, e int64) int64 {
		return total + e
	}, snafuNumbers, 0)
	fmt.Println("need", sum)
	fmt.Println(DecimalToSnafu(sum))

}

var SnafuToDecimalMap = map[string]int64{
	"2": 2,
	"1": 1,
	"0": 0,
	"-": -1,
	"=": -2,
}

var DecimalToSnafuMap = map[int64]string{
	2:  "2",
	1:  "1",
	0:  "0",
	-1: "-",
	-2: "=",
}

func SnafuToDecimal(s string) int64 {
	if len(s) == 0 {
		return 0
	}
	snafuDigits := strings.Split(s, "")
	multiplier := SnafuToDecimalMap[snafuDigits[0]]
	exponent := len(s) - 1
	digit := int64(math.Pow(5, float64(exponent)))
	return multiplier*digit + SnafuToDecimal(strings.Join(snafuDigits[1:], ""))
}

func DecimalToSnafu(m int64) string {
	_, lowerExp := surroundingPowersOfFive(m, 0)
	estimateHigh := "2-0=11=-0-2122222220"
	return decrementUntil(estimateHigh, m)
	estimate := "1" + strings.Repeat("0", int(lowerExp))
	return incrementUntil(estimate, m)
}

func decrementUntil(snafu string, target int64) string {
	decimal := SnafuToDecimal(snafu)
	fmt.Println(snafu, "     ", decimal)
	if decimal == target {
		return snafu
	}
	snafus := strings.Split(snafu, "")
	newSnafu := decrement(snafus)
	return decrementUntil(newSnafu, target)
}

func incrementUntil(snafu string, target int64) string {
	decimal := SnafuToDecimal(snafu)
	//fmt.Println(snafu, "      ", decimal)
	if decimal == target {
		return snafu
	}
	snafus := strings.Split(snafu, "")
	newSnafu := increment(snafus)
	return incrementUntil(newSnafu, target)
}

func decrement(snafus []string) string {
	secondaryDigitSet := util.ToSet(snafus[1:])
	if snafus[0] == "1" && len(secondaryDigitSet) == 1 && secondaryDigitSet["="] {
		return strings.Repeat("2", len(secondaryDigitSet))
	}
	ending := snafus[len(snafus)-1]
	first := snafus[:len(snafus)-1]
	if ending == "=" {
		return decrement(first) + "2"
	}
	return strings.Join(first, "") + prev(ending)
}

func increment(snafus []string) string {
	digitSet := util.ToSet(snafus)
	if len(digitSet) == 1 && digitSet["2"] {
		return "1" + strings.Repeat("=", len(snafus))
	}
	ending := snafus[len(snafus)-1]
	first := snafus[:len(snafus)-1]
	if ending == "2" {
		return increment(first) + "="
	}
	return strings.Join(first, "") + next(ending)
}

func next(ending string) string {
	switch ending {
	case "=":
		return "-"
	case "-":
		return "0"
	case "0":
		return "1"
	case "1":
		return "2"
	}
	log.Fatal("bad input", ending)
	return "F"
}
func prev(ending string) string {
	switch ending {
	case "2":
		return "1"
	case "1":
		return "0"
	case "0":
		return "-"
	case "-":
		return "="
	}
	log.Fatal("bad input", ending)
	return "F"
}
func surroundingPowersOfFive(i int64, exp int64) (int64, int64) {
	pow := powerOfFive(exp)
	if i < pow {
		return exp, exp - 1
	}
	return surroundingPowersOfFive(i, exp+1)
}

func powerOfFive(exp int64) int64 {
	return int64(math.Pow(5, float64(exp)))
}
