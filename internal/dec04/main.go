package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hooten/advent-of-code-2022/pkg/util"
)

type Range struct {
	Start, End int64
}

func ParseRange(s string) (Range, error) {
	ends := strings.Split(s, "-")
	if len(ends) != 2 {
		return Range{}, fmt.Errorf("cannot parse range %q", ends)
	}
	start, err := strconv.ParseInt(ends[0], 10, 64)
	if err != nil {
		return Range{}, fmt.Errorf("cannot parse range start %q", ends[0])
	}
	end, err := strconv.ParseInt(ends[1], 10, 64)
	if err != nil {
		return Range{}, fmt.Errorf("cannot parse range end %q", ends[1])
	}
	return Range{
		Start: start,
		End:   end,
	}, nil
}

type Pair struct {
	First, Second Range
}

func ParsePair(s string) (Pair, error) {
	ranges := strings.Split(s, ",")
	if len(ranges) != 2 {
		return Pair{}, fmt.Errorf("cannot parse pair %q", ranges)
	}
	first, err := ParseRange(ranges[0])
	if err != nil {
		return Pair{}, err
	}
	second, err := ParseRange(ranges[1])
	if err != nil {
		return Pair{}, err
	}
	return Pair{
		First:  first,
		Second: second,
	}, nil
}

func (p Pair) Overlapping(strict bool) bool {
	return overlapping(p.First, p.Second, strict) || overlapping(p.Second, p.First, strict)
}

func overlapping(a, b Range, strict bool) bool {
	if strict {
		return a.Start <= b.Start && a.End >= b.End
	}
	return b.Start <= a.End && a.End <= b.End
}

func main() {
	fmt.Printf("Part 1: There are %d fully overlapping pairs.\n", overlappingPairs(true))
	fmt.Printf("Part 2: There are %d partly overlapping pairs.\n", overlappingPairs(false))
}

func overlappingPairs(strict bool) int64 {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	unfilteredLines := strings.Split(string(bytes), "\n")
	lines := util.Filter(func(line string) bool {
		return line != ""
	}, unfilteredLines)
	pairs := util.Map(func(line string) Pair {
		pair, err := ParsePair(line)
		if err != nil {
			log.Fatal(err)
		}
		return pair
	}, lines)
	overlapping := util.Filter(func(pair Pair) bool {
		return pair.Overlapping(strict)
	}, pairs)

	return int64(len(overlapping))
}
