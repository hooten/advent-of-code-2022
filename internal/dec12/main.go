package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"sort"
)

func main() {
	filename := "./internal/dec12/input.txt"
	file := util.MustReadFile(filename)
	lines := util.SplitByLine(file)
	matrix := util.NewMatrix(lines, "")
	heightMap := util.NewMatrixWithMap(lines, "", func(s string) int {
		if s == "S" {
			return 0
		}
		if s == "E" {
			return 'z' - 'a'
		}
		return int(s[0]) - 'a'
	})

	var possibleStarts []*util.Pair
	var start *util.Pair
	var end *util.Pair
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == "S" {
				start = util.NewPair(i, j)
				possibleStarts = append(possibleStarts, start)
			}
			if matrix[i][j] == "E" {
				end = util.NewPair(i, j)
			}
			if matrix[i][j] == "a" {
				possibleStarts = append(possibleStarts, util.NewPair(i, j))
			}
		}
	}

	distances := map[string]int{}
	setDistance(heightMap, distances, nil, start, 0)
	fmt.Println("Part 1:", distances[end.Key()])

	var distancesToEnd []int
	for _, possibleStart := range possibleStarts {
		setDistance(heightMap, distances, nil, possibleStart, 0)
		distancesToEnd = append(distancesToEnd, distances[end.Key()])
	}
	sort.Ints(distancesToEnd)
	fmt.Println("Part 2:", distancesToEnd[0])
}

func setDistance(heightMap [][]int, distances map[string]int, previous *util.Pair, current *util.Pair, altDistance int) {
	if current.X < 0 || current.X >= len(heightMap) || current.Y < 0 || current.Y >= len(heightMap[0]) {
		return
	}

	if distance, ok := distances[current.Key()]; ok {
		if distance <= altDistance {
			return
		}
	}

	if previous != nil {
		if heightMap[previous.X][previous.Y]+1 < heightMap[current.X][current.Y] {
			return
		}
	}

	distances[current.Key()] = altDistance

	for _, neighbor := range []*util.Pair{
		{X: current.X + 1, Y: current.Y},
		{X: current.X - 1, Y: current.Y},
		{X: current.X, Y: current.Y - 1},
		{X: current.X, Y: current.Y + 1},
	} {
		setDistance(heightMap, distances, current, neighbor, altDistance+1)
	}
}
