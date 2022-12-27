package main

import (
	"container/list"
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"strconv"
	"strings"
)

func main() {
	file := util.MustReadFile("./internal/dec24/input.txt")
	lines := util.SplitByLine(file)
	valley := util.Map(func(line string) []string {
		return strings.Split(line, "")
	}, lines)
	height := len(valley)
	width := len(valley[0])

	start := &Pair{X: 1, Y: 0, Valley: valley}
	end := FewestMinutes(start, &Pair{X: width - 2, Y: height - 1})
	fmt.Println("Part 1")
	minutesOut := len(end.Path()) - 1
	fmt.Println("Fewest minutes out:", minutesOut)

	fmt.Println("Part 2")
	start2 := &Pair{X: end.X, Y: end.Y, Valley: end.Valley}
	backToStart := FewestMinutes(start2, &Pair{X: 1, Y: 0})
	minutesBackToStart := len(backToStart.Path()) - 1
	fmt.Println("Fewest minutes back:", minutesBackToStart)

	start3 := &Pair{X: backToStart.X, Y: backToStart.Y, Valley: backToStart.Valley}
	backToEnd := FewestMinutes(start3, &Pair{X: width - 2, Y: height - 1})
	minutesBackOut := len(backToEnd.Path()) - 1
	fmt.Println("Fewest minutes back out:", minutesBackOut)

	fmt.Println("Total minutes:", minutesOut+minutesBackToStart+minutesBackOut)

}

func Neighbors(valley [][]string, pair *Pair) []*Pair {
	x := pair.X
	y := pair.Y
	pairs := []*Pair{
		{Valley: valley, Prev: pair, X: x + 1, Y: y},
		{Valley: valley, Prev: pair, X: x, Y: y + 1},
		{Valley: valley, Prev: pair, X: x, Y: y}, // Allowed to remain.
		{Valley: valley, Prev: pair, X: x - 1, Y: y},
		{Valley: valley, Prev: pair, X: x, Y: y - 1},
	}
	valid := util.Filter(func(pair *Pair) bool {
		if pair.X < 0 || pair.Y < 0 {
			return false
		}
		if pair.X >= len(valley[0]) || pair.Y >= len(valley) {
			return false
		}
		return valley[pair.Y][pair.X] == "."
	}, pairs)
	return valid
}

type Pair struct {
	X      int
	Y      int
	Valley [][]string
	Prev   *Pair
}

func (pair *Pair) Path() []Pair {
	var path []Pair
	curr := pair
	for curr != nil {
		c := *curr
		path = append([]Pair{{X: c.X, Y: c.Y, Valley: c.Valley}}, path...)
		curr = curr.Prev
	}
	return path
}

func (pair *Pair) String() string {
	s, err := String(pair.Valley, *pair)
	if err != nil {
		log.Fatal(err.Error())
	}
	return s
}

func FewestMinutes(start *Pair, end *Pair) *Pair {

	visited := map[string]bool{}

	queue := list.New()
	queue.PushBack(start)

	currValley := start.Valley

	for queue.Len() > 0 {
		e := queue.Front()
		v := queue.Remove(e).(*Pair)
		state := Pair{X: v.X, Y: v.Y, Valley: v.Valley} // don't include prev
		key := state.String()
		if visited[key] {
			continue
		}
		visited[key] = true
		currValley = v.Valley
		if v.X == end.X && v.Y == end.Y {
			return v
		}
		currValley = NextValley(currValley)
		neighbors := Neighbors(currValley, v)
		for _, neighbor := range neighbors {
			queue.PushBack(neighbor)
		}
	}
	return nil
}

func NewEmptyValley(height, width int) [][]string {
	var nextValley [][]string
	for y := 0; y < height; y++ {
		row := make([]string, width)
		for x := 0; x < width; x++ {
			if y == 0 && x != 1 {
				row[x] = "#"
				continue
			}
			if y == height-1 && x != width-2 {
				row[x] = "#"
				continue
			}
			if x == 0 || x == width-1 {
				row[x] = "#"
				continue
			}
			row[x] = "."
		}
		nextValley = append(nextValley, row)
	}
	return nextValley
}

func String(valley [][]string, pair Pair) (string, error) {
	var sb strings.Builder
	var err error
	for y := 0; y < len(valley); y++ {
		for x := 0; x < len(valley[0]); x++ {
			cell := valley[y][x]
			if pair.X == x && pair.Y == y {
				if cell != "." {
					err = fmt.Errorf("position in non-empty space (%d, %d)", x, y)
				} else {
					sb.WriteString("E")
					continue
				}
			}
			if len(cell) > 1 {
				sb.WriteString(strconv.Itoa(len(cell)))
				continue
			}
			sb.WriteString(cell)
		}
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	return sb.String(), err
}

func Draw(valley [][]string, pair Pair) {
	s, err := String(valley, pair)
	fmt.Println(s)
	if err != nil {
		log.Fatal(err)
	}
}

func NextValley(valley [][]string) [][]string {
	height := len(valley)
	width := len(valley[0])
	nextValley := NewEmptyValley(height, width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currs := strings.Split(valley[y][x], "")
			for _, curr := range currs {
				newX, newY := NextBlizzard(curr, x, y, width, height)
				if curr == "." {
					continue
				}
				if nextValley[newY][newX] == "#" && curr == "#" {
					continue
				}
				if nextValley[newY][newX] == "#" && curr != "#" {
					log.Fatalf("something went wrong, unexpected valley wall at (%d, %d) for blizzard %s", newX, newY, currs)
				}
				if nextValley[newY][newX] == "." {
					nextValley[newY][newX] = curr
					continue
				}
				nextValley[newY][newX] = nextValley[newY][newX] + curr
			}
		}
	}
	return nextValley
}

func NextBlizzard(curr string, x, y, width, height int) (int, int) {
	switch curr {
	// cases "#", "." already considered in new empty valley
	case ".":
		return x, y
	case "#":
		return x, y
	case "<":
		if x-1 < 1 {
			return width - 2, y
		}
		return x - 1, y
	case ">":
		if x+1 > width-2 {
			return 1, y
		}
		return x + 1, y
	case "^":
		if y-1 < 1 {
			return x, height - 2
		}
		return x, y - 1
	case "v":
		if y+1 > height-2 {
			return x, 1
		}
		return x, y + 1
	}
	log.Fatalf("unexpected blizzard %s at coordinates (%d, %d)", curr, x, y)
	return 0, 0
}
