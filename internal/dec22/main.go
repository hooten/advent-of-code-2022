package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"github.com/kr/pretty"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var SquareEdge = 50

func main() {
	file := util.MustReadFile("./internal/dec22/input.txt")
	lines := util.SplitByLine(file)
	board := NewBoard(lines)
	path := NewPath(lines)
	position := NewPosition(board)
	mapEnd := Play(board, path, position, false)
	pretty.Println("Part 1:")
	pretty.Println(mapEnd)
	pretty.Printf("Elements: y=%d, x=%d, facing=%d\n", mapEnd.Y+1, mapEnd.X, facing(mapEnd.Direction))
	pretty.Println("Password: ", 1000*(mapEnd.Y+1)+4*(mapEnd.X+1)+facing(mapEnd.Direction))

	cubeEnd := Play(board, path, position, true)
	pretty.Println("Part 2:")
	pretty.Println(cubeEnd)
	pretty.Printf("Elements: y=%d, x=%d, facing=%d\n", cubeEnd.Y+1, cubeEnd.X, facing(cubeEnd.Direction))
	pretty.Println("Password: ", 1000*(cubeEnd.Y+1)+4*(cubeEnd.X+1)+facing(cubeEnd.Direction))
}

func Play(board [][]string, path []string, position Position, cubeRules bool) Position {
	if len(path) == 0 {
		return position
	}
	move := path[0]
	if directionRe.MatchString(move) {
		newPosition := Position{
			X:         position.X,
			Y:         position.Y,
			Direction: nextDirection(position.Direction, move),
		}
		return Play(board, path[1:], newPosition, cubeRules)
	}
	steps, err := strconv.Atoi(move)
	if err != nil {
		log.Fatal(err)
	}
	x, y, direction := nextPosition(board, position, steps, cubeRules)
	newPosition := Position{
		X:         x,
		Y:         y,
		Direction: direction,
	}
	return Play(board, path[1:], newPosition, cubeRules)
}

func nextPosition(board [][]string, position Position, steps int, cubeRules bool) (int, int, string) {
	x := position.X
	y := position.Y
	dir := position.Direction

	if steps == 0 {
		return x, y, dir
	}
	next, err := verifyFree(board, wrap(board, nextBoardUnawarePosition(position), cubeRules))
	if err != nil {
		return x, y, dir
	}

	return nextPosition(board, next, steps-1, cubeRules)
}

func verifyFree(board [][]string, position Position) (Position, error) {
	x := position.X
	y := position.Y

	if board[y][x] == "#" {
		return position, fmt.Errorf("position (%d, %d) is a wall", x, y)
	}
	return position, nil
}

func wrap(board [][]string, position Position, cubeRules bool) Position {
	width := len(board[0])
	height := len(board)

	x := position.X
	y := position.Y
	dir := position.Direction

	if cubeRules {
		return cubeWrap(board, position)
	} else {
		if x < 0 {
			return wrap(board, Position{
				X:         width - 1,
				Y:         y,
				Direction: dir,
			}, cubeRules)
		}
		if y < 0 {
			return wrap(board, Position{
				X:         x,
				Y:         height - 1,
				Direction: dir,
			}, cubeRules)
		}
		if x >= width {
			return wrap(board, Position{
				X:         0,
				Y:         y,
				Direction: dir,
			}, cubeRules)
		}
		if y >= height {
			return wrap(board, Position{
				X:         x,
				Y:         0,
				Direction: dir,
			}, cubeRules)
		}
	}
	if board[y][x] == " " {
		return wrap(board, nextBoardUnawarePosition(position), cubeRules)
	}
	return position
}

func In(x int, lowerBound int, upperBound int) bool {
	return lowerBound <= x && x <= upperBound
}

func cubeWrap(board [][]string, position Position) Position {
	if validQuadrant(board, position) {
		return position
	}

	x := position.X
	y := position.Y

	//Q1 => Q4
	if x == SquareEdge-1 && In(y, 0, SquareEdge-1) {
		newX := 0
		newY := 3*SquareEdge - y - 1
		newDir := ">"
		return Position{X: newX, Y: newY, Direction: newDir}
	}
	//Q4 => Q1
	if x == -1 && In(y, 2*SquareEdge, 3*SquareEdge-1) {
		newX := SquareEdge
		newY := SquareEdge - (y - 2*SquareEdge) - 1
		newDir := ">"
		return Position{X: newX, Y: newY, Direction: newDir}
	}

	//Q3 => Q4
	if x == SquareEdge-1 && In(y, SquareEdge, 2*SquareEdge-1) {
		newX := y % SquareEdge
		newY := 2 * SquareEdge
		newDir := "v"
		return Position{X: newX, Y: newY, Direction: newDir}
	}
	//Q4 => Q3
	if In(x, 0, SquareEdge-1) && y == 2*SquareEdge-1 {
		newX := SquareEdge
		newY := SquareEdge + x
		newDir := ">"
		return Position{X: newX, Y: newY, Direction: newDir}
	}

	//Q1 => Q6
	if In(x, SquareEdge, 2*SquareEdge-1) && y == -1 {
		newX := 0
		newY := 3*SquareEdge + x%SquareEdge
		newDir := ">"
		return Position{X: newX, Y: newY, Direction: newDir}
	}
	//Q6 => Q1
	if x == -1 && In(y, 3*SquareEdge, 4*SquareEdge-1) {
		newX := y%SquareEdge + SquareEdge
		newY := 0
		newDir := "v"
		return Position{X: newX, Y: newY, Direction: newDir}
	}

	//Q2 => Q6
	if In(x, 2*SquareEdge, 3*SquareEdge-1) && y == -1 {
		newX := x % SquareEdge
		newY := 4*SquareEdge - 1
		newDir := "^"
		return Position{X: newX, Y: newY, Direction: newDir}
	}
	//Q6 => Q2
	if In(x, 0, SquareEdge-1) && y == 4*SquareEdge {
		newX := 2*SquareEdge + x
		newY := 0
		newDir := "v"
		return Position{X: newX, Y: newY, Direction: newDir}
	}

	//Q2 => Q5
	if x == 3*SquareEdge && In(y, 0, SquareEdge-1) {
		newX := 2*SquareEdge - 1
		newY := 3*SquareEdge - y - 1
		newDir := "<"
		return Position{X: newX, Y: newY, Direction: newDir}
	}
	//Q5 => Q2
	if x == 2*SquareEdge && In(y, 2*SquareEdge, 3*SquareEdge-1) {
		newX := 3*SquareEdge - 1
		newY := SquareEdge - y%SquareEdge - 1
		newDir := "<"
		return Position{X: newX, Y: newY, Direction: newDir}
	}

	//Q2 => Q3
	if In(x, 2*SquareEdge, 3*SquareEdge-1) && y == SquareEdge {
		newX := 2*SquareEdge - 1
		newY := x%SquareEdge + SquareEdge
		newDir := "<"
		return Position{X: newX, Y: newY, Direction: newDir}
	}
	//Q3 => Q2
	if x == 2*SquareEdge && In(y, SquareEdge, 2*SquareEdge-1) {
		newX := y%SquareEdge + 2*SquareEdge
		newY := SquareEdge - 1
		newDir := "^"
		return Position{X: newX, Y: newY, Direction: newDir}
	}

	//Q5 => Q6
	if In(x, SquareEdge, 2*SquareEdge-1) && y == 3*SquareEdge {
		newX := SquareEdge - 1
		newY := x%SquareEdge + 3*SquareEdge
		newDir := "<"
		return Position{X: newX, Y: newY, Direction: newDir}
	}
	//Q6 => Q5
	if x == SquareEdge && In(y, 3*SquareEdge, 4*SquareEdge-1) {
		newX := y%SquareEdge + SquareEdge
		newY := 3*SquareEdge - 1
		newDir := "^"
		return Position{X: newX, Y: newY, Direction: newDir}
	}
	log.Fatalf("case not caught: %v", position)
	return position
}

func validQuadrant(board [][]string, position Position) bool {
	width := len(board[0])
	height := len(board)

	x := position.X
	y := position.Y

	if x < 0 || y < 0 {
		return false
	}
	if x >= width || y >= height {
		return false
	}
	if x < SquareEdge && y < 2*SquareEdge {
		return false
	}
	if x >= 2*SquareEdge && y >= SquareEdge {
		return false
	}
	if x >= SquareEdge && y >= 3*SquareEdge {
		return false
	}
	_ = board[y][x]
	return true
}

func nextBoardUnawarePosition(position Position) Position {
	switch position.Direction {
	case ">":
		return Position{
			X:         position.X + 1,
			Y:         position.Y,
			Direction: position.Direction,
		}
	case "v":
		return Position{
			X:         position.X,
			Y:         position.Y + 1,
			Direction: position.Direction,
		}
	case "<":
		return Position{
			X:         position.X - 1,
			Y:         position.Y,
			Direction: position.Direction,
		}
	case "^":
		return Position{
			X:         position.X,
			Y:         position.Y - 1,
			Direction: position.Direction,
		}
	}
	log.Fatalf("could not get next position from %v", position)
	return Position{}
}

func facing(direction string) int {
	switch direction {
	case ">":
		return 0
	case "v":
		return 1
	case "<":
		return 2
	case "^":
		return 3
	}
	log.Fatalf("no code for direection %q", direction)
	return -1
}

func nextDirection(direction string, move string) string {
	if move == "L" {
		switch direction {
		case ">":
			return "^"
		case "v":
			return ">"
		case "<":
			return "v"
		case "^":
			return "<"
		}
	}
	if move == "R" {
		switch direction {
		case ">":
			return "v"
		case "v":
			return "<"
		case "<":
			return "^"
		case "^":
			return ">"
		}
	}
	log.Fatalf("cannot parse move %q and Direction %q", move, direction)
	return "fatal"
}

func NewPosition(board [][]string) Position {
	positionIdx, _ := util.First(func(char string) bool {
		return char == "."
	}, board[0])
	position := Position{
		X:         positionIdx,
		Y:         0,
		Direction: ">",
	}
	return position
}

type Position struct {
	X         int
	Y         int
	Direction string
}

func NewPath(lines []string) []string {
	pathStr := util.Filter(func(line string) bool {
		_, ok := util.RegexpMatch("([LR0-9]+)", line)
		return ok
	}, lines)
	if len(pathStr) != 1 {
		log.Fatal("bad path string: ", pathStr)
	}
	path := ParsePath(pathStr[0])
	return path
}

func NewBoard(lines []string) [][]string {
	boardRaw := util.Filter(func(line string) bool {
		return strings.HasPrefix(line, " ") ||
			strings.HasPrefix(line, ".") ||
			strings.HasPrefix(line, "#")
	}, lines)
	maxWidth := util.Reduce(func(max int, line string) int {
		return int(math.Max(float64(max), float64(len(line))))
	}, boardRaw, 0)
	boardStr := util.Map(func(line string) []string {
		if len(line) < maxWidth {
			extraChars := strings.Repeat(" ", maxWidth-len(line))
			return strings.Split(line+extraChars, "")
		}
		return strings.Split(line, "")
	}, boardRaw)
	return boardStr
}

var directionRe = regexp.MustCompile("[RL]")

func ParsePath(s string) []string {
	if s == "" {
		return []string{}
	}
	if strings.HasPrefix(s, "L") ||
		strings.HasPrefix(s, "R") {
		return append([]string{s[:1]}, ParsePath(s[1:])...)
	}
	split := directionRe.Split(s, -1)
	newS := strings.Replace(s, split[0], "", 1)
	return append([]string{split[0]}, ParsePath(newS)...)
}
