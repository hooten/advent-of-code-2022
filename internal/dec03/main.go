package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hooten/advent-of-code-2022/pkg/math"
	"github.com/hooten/advent-of-code-2022/pkg/util"
)

func main() {
	fmt.Printf("Part 1: Rucksack priority sum equals %d\n", rucksackPrioritySum())
	fmt.Printf("Part 2: Elf Badge priority sum equals %d\n", elfBadgePrioritySum())
}

func elfBadgePrioritySum() int64 {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := util.Filter(func(line string) bool {
		return line != ""
	}, strings.Split(string(bytes), "\n"))
	elfGroups := util.Reduce(func(groups [][]string, line string) [][]string {
		last := groups[len(groups)-1]
		if len(last) == 3 {
			return append(groups, []string{line})
		}
		return append(groups[:len(groups)-1], append(last, line))
	}, lines, [][]string{[]string{}})
	elfBadges := util.Map(func(group []string) string {
		return Common(group)
	}, elfGroups)
	priorities := util.Map(Priority, elfBadges)
	return math.Sum(priorities...)
}

type Rucksack struct {
	First  string
	Second string
}

func ParseRucksack(s string) Rucksack {
	return Rucksack{
		First:  (s[:len(s)/2]),
		Second: (s[len(s)/2:]),
	}
}

func Common(xs []string) string {
	if len(xs) == 0 {
		return ""
	}
	if len(xs) == 1 {
		return xs[0]
	}
	firstSet := util.ToSet(strings.Split(xs[0], ""))
	secondSet := util.ToSet(strings.Split(xs[1], ""))
	intersection := util.Intersection(firstSet, secondSet)
	common := strings.Join(util.Keys(intersection), "")
	return Common(append([]string{common}, xs[2:]...))
}

func (r Rucksack) Common() string {
	return Common([]string{r.First, r.Second})
}

func Priority(s string) int64 {
	if len(s) != 1 {
		log.Fatalf("expected %q to have length of 1", s)
	}
	if strings.ToLower(s) == s {
		return int64(s[0]) - 96
	}
	return int64(s[0]) - 38
}

func rucksackPrioritySum() int64 {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := util.Filter(func(line string) bool {
		return line != ""
	}, strings.Split(string(bytes), "\n"))
	rucksacks := util.Map(ParseRucksack, lines)
	priorities := util.Map(func(r Rucksack) int64 {
		common := r.Common()
		return Priority(common)
	}, rucksacks)
	return math.Sum(priorities...)
}
