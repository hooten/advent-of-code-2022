package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"sort"
	"strings"
)

func main() {
	filename := "./internal/dec14/input.txt"
	part1 := false
	file := util.MustReadFile(filename)
	lines := util.SplitByLine(file)

	var rocks = make([][]util.Pair, len(lines))
	for i, line := range lines {
		coordinates := util.Map(func(s string) util.Pair {
			pair, err := util.NewPairFromKey(s)
			if err != nil {
				log.Fatal("bad pair from key ", s)
			}
			return *pair
		}, strings.Split(line, " -> "))
		rocks[i] = coordinates
	}
	fmt.Println(rocks)

	cave := map[string]string{
		util.NewPair(500, 0).Key(): "+",
	}
	for _, rock := range rocks {
		for i := 0; i < len(rock)-1; i++ {
			curr := rock[i]
			next := rock[i+1]
			if curr.X == next.X {
				x := curr.X
				ys := newRange(curr.Y, next.Y)
				for _, y := range ys {
					pair := util.NewPair(x, y)
					cave[pair.Key()] = "#"
				}
			}
			if curr.Y == next.Y {
				y := curr.Y
				xs := newRange(curr.X, next.X)
				for _, x := range xs {
					pair := util.NewPair(x, y)
					cave[pair.Key()] = "#"
				}
			}
		}
	}

	keys := util.Keys(cave)
	xs := util.Map(func(key string) int {
		pair, err := util.NewPairFromKey(key)
		if err != nil {
			log.Fatal("bad pair x ", key, pair, err)
		}
		return pair.X
	}, keys)

	sort.Ints(xs)
	minX := xs[0]
	maxX := xs[len(xs)-1]

	ys := util.Map(func(key string) int {
		pair, err := util.NewPairFromKey(key)
		if err != nil {
			log.Fatal("bad pair y", pair, err)
		}
		return pair.Y
	}, keys)

	sort.Ints(ys)
	minY := ys[0]
	maxY := ys[len(ys)-1]

	for pair := range cave {
		fmt.Println(pair)
	}

	grains := 0
	done := false
	// multi grain loop
	for !done {
		grains++
		// single grain loop
		grain := util.NewPair(500, 0)
		for {
			moves := possibleMoves(cave, grain)
			if len(moves) == 0 {
				cave[grain.Key()] = "o"
				break
			}
			grain = moves[0]
			if grain.Y == maxY && (grain.X < minX || grain.X > maxX) {
				done = true
				break
			}
		}
	}

	if part1 {
		for y := minY; y <= maxY; y++ {
			fmt.Println(" ")
			for x := minX; x <= maxX; x++ {
				if s, ok := cave[util.NewPair(x, y).Key()]; ok {
					fmt.Print(s)
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println("")
		fmt.Println("grains:")
		fmt.Println(grains - 1)
		return
	}

	floor := maxY + 2
	for x := minX - 10000; x <= maxX+10000; x++ {
		cave[util.NewPair(x, floor).Key()] = "#"
	}

	// part 2
	sourceBlocked := false
	// multi grain loop
	for !sourceBlocked {
		// single grain loop
		grain := util.NewPair(500, 0)
		for {
			moves := possibleMoves(cave, grain)
			if len(moves) == 0 {
				cave[grain.Key()] = "o"
				if grain.X == 500 && grain.Y == 0 {
					sourceBlocked = true
				}
				break
			}
			grain = moves[0]
		}
	}

	for y := minY; y <= maxY+2; y++ {
		fmt.Println(" ")
		for x := minX - 20; x <= maxX+20; x++ {
			if s, ok := cave[util.NewPair(x, y).Key()]; ok {
				fmt.Print(s)
			} else {
				fmt.Print(".")
			}
		}
	}

	fmt.Println("")
	fmt.Println("units:")
	sands := util.Filter(func(s string) bool {
		return s == "o"
	}, util.Values(cave))
	fmt.Println(len(sands))

}

func possibleMoves(cave map[string]string, grain *util.Pair) []*util.Pair {
	filtered := util.Filter(func(pair *util.Pair) bool {
		_, contained := cave[pair.Key()]
		return !contained
	}, []*util.Pair{
		util.NewPair(grain.X, grain.Y+1),
		util.NewPair(grain.X-1, grain.Y+1),
		util.NewPair(grain.X+1, grain.Y+1),
	})
	return filtered
}

func newRange(start, end int) []int {
	ints := make([]int, 0)
	if start <= end {
		for i := start; i <= end; i++ {
			ints = append(ints, i)
		}
		return ints
	}
	for i := start; i >= end; i-- {
		ints = append(ints, i)
	}
	return ints
}
