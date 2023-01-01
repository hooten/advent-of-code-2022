package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"math"
	"os"
	"strings"
)

type Board [][]string

func (b Board) String() string {
	return util.Reduce(func(s string, row []string) string {
		return s + strings.Join(row, "") + "\n"
	}, b, "") + "\n"
}

func (b Board) Cell(x int, y int) (string, error) {
	height := len(b)
	width := len(b[0])

	if x < 0 || y < 0 || x >= width || y >= height {
		return "", fmt.Errorf("(%d, %d) out of bounds", x, y)
	}
	return b[y][x], nil
}

func main() {
	bytes, err := os.ReadFile("./internal/dec23/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := util.Filter(func(line string) bool {
		return line != ""
	}, strings.Split(string(bytes), "\n"))
	board := util.Map(func(line string) []string {
		return strings.Split(line, "")
	}, lines)
	maxRounds := 10
	end, _ := Play(board, 0, maxRounds)

	trimmed := util.Map(func(y int) []string {
		return util.Map(func(x int) string {
			cell, err := end.Cell(x, y)
			if err != nil {
				log.Fatal(err)
			}
			return cell
		}, util.NewRange(MinX(end), MaxX(end)))
	}, util.NewRange(MinY(end), MaxY(end)))

	emptyTiles := util.Reduce(func(tiles int, row []string) int {
		empties := util.Filter(func(cell string) bool {
			return cell == "."
		}, row)
		return tiles + len(empties)
	}, trimmed, 0)

	fmt.Println(Board(trimmed).String())
	fmt.Println("Part 1:", emptyTiles)

	_, endRound := Play(board, 0, math.MaxInt)
	fmt.Println("Part 2:", endRound+1)

}

func MinY(board Board) int {
	i, _ := util.First(func(row []string) bool {
		join := strings.Join(row, "")
		return strings.Contains(join, "#")
	}, board)
	return i
}

func MinX(board Board) int {
	return util.Reduce(func(min int, row []string) int {
		x, _ := util.First(func(cell string) bool {
			return cell == "#"
		}, row)
		if x == -1 {
			return min
		}
		if x < min {
			return x
		}
		return min
	}, board, math.MaxInt)
}

func MaxY(board Board) int {
	i, _ := util.Last(func(row []string) bool {
		join := strings.Join(row, "")
		return strings.Contains(join, "#")
	}, board)
	return i
}

func MaxX(board Board) int {
	return util.Reduce(func(max int, row []string) int {
		x, _ := util.Last(func(cell string) bool {
			return cell == "#"
		}, row)
		if x == -1 {
			return max
		}
		if x > max {
			return x
		}
		return max
	}, board, math.MinInt)
}

func Play(board Board, round int, maxRounds int) (Board, int) {
	height := len(board)
	width := len(board[0])
	fmt.Printf("Round %d\n", round)
	if round == maxRounds {
		return board, round
	}
	proposedActions := util.Filter(func(action Action) bool {
		return action != nil
	}, util.FlatMapWithIndex(func(row []string, y int) []Action {
		return util.MapWithIndex(func(cell string, x int) Action {
			s, err := board.Cell(x, y)
			if err != nil {
				log.Fatal(err)
			}
			if s != "#" {
				return nil
			}
			adj := AdjacentPositions(board, y, x)
			if len(adj) == 8 {
				return Stay{position: Position{x: x, y: y}}
			}
			return NewMove(board, y, x, round)
		}, row)
	}, board))

	moves := util.Filter(func(action Action) bool {
		return action.Start() != action.End()
	}, proposedActions)
	if len(moves) == 0 {
		return board, round
	}

	proposedActionsMap := util.ReduceWithIndex(func(m map[Position][]Position, action Action, y int) map[Position][]Position {
		start := action.Start()
		end := action.End()
		return util.Assoc(m, end, append(m[end], start))
	}, proposedActions, map[Position][]Position{})

	actionsMap := util.Reduce(func(m map[Position]bool, end Position) map[Position]bool {
		starts := proposedActionsMap[end]
		if len(starts) == 1 {
			return util.Assoc(m, end, true)
		}
		return util.Reduce(func(m map[Position]bool, start Position) map[Position]bool {
			return util.Assoc(m, start, true)
		}, starts, m)
	}, util.Keys(proposedActionsMap), map[Position]bool{})

	minX := util.Reduce(func(min int, position Position) int {
		if position.x < min {
			return position.x
		}
		return min
	}, util.Keys(actionsMap), 0)

	maxX := util.Reduce(func(max int, position Position) int {
		if position.x > max {
			return position.x
		}
		return max
	}, util.Keys(actionsMap), width-1)

	minY := util.Reduce(func(min int, position Position) int {
		if position.y < min {
			return position.y
		}
		return min
	}, util.Keys(actionsMap), 0)

	maxY := util.Reduce(func(max int, position Position) int {
		if position.y > max {
			return position.y
		}
		return max
	}, util.Keys(actionsMap), height-1)

	newBoard := util.Map(func(y int) []string {
		return util.Map(func(x int) string {
			position := Position{x: x, y: y}
			_, ok := actionsMap[position]
			if !ok {
				return "."
			}
			return "#"
		}, util.NewRange(minX, maxX))
	}, util.NewRange(minY, maxY))
	return Play(newBoard, round+1, maxRounds)
}

func NewMove(board Board, y int, x int, round int) Action {
	sets := NewDirectionSets(round)
	validSets := util.Filter(func(fs []func(int, int) Position) bool {
		emptyCells := util.Filter(func(f func(int, int) Position) bool {
			newPosition := f(x, y)
			cell, err := board.Cell(newPosition.x, newPosition.y)
			if err != nil {
				// Board can expand.
				return true
			}
			return cell == "."
		}, fs)
		return len(emptyCells) == 3
	}, sets)
	if len(validSets) == 0 {
		return Stay{position: Position{x: x, y: y}}
	}
	f := validSets[0][0]
	return Move{
		position: Position{x: x, y: y},
		f:        f,
	}
}

var AllDirections = []string{"N", "S", "W", "E"}

var DirectionSets = map[string][]func(int, int) Position{
	"N": {N, NE, NW},
	"S": {S, SE, SW},
	"W": {W, NW, SW},
	"E": {E, NE, SE},
}

func NewDirectionSets(round int) [][]func(int, int) Position {
	iRange := []int{0, 1, 2, 3}
	directions := util.Map(func(i int) string {
		directionIdx := int(math.Abs(float64(round+i))) % 4
		return AllDirections[directionIdx]
	}, iRange)
	return util.Map(func(direction string) []func(int, int) Position {
		return DirectionSets[direction]
	}, directions)
}

type Position struct {
	x, y int
}

type Action interface {
	Start() Position
	End() Position
}

type Move struct {
	position Position
	f        func(int, int) Position
}

func (m Move) Start() Position {
	return m.position
}

func (m Move) End() Position {
	return m.f(m.position.x, m.position.y)
}

type Stay struct {
	position Position
}

func (s Stay) Start() Position {
	return s.position
}

func (s Stay) End() Position {
	return s.position
}

func AdjacentPositions(board Board, y int, x int) []Position {
	height := len(board)
	width := len(board[0])
	return util.Filter(func(position Position) bool {
		if position.x < 0 || position.y < 0 || position.x >= width || position.y >= height {
			return true
		}
		return board[position.y][position.x] == "."
	}, []Position{
		NW(x, y),
		N(x, y),
		NE(x, y),
		E(x, y),
		SE(x, y),
		S(x, y),
		SW(x, y),
		W(x, y),
	})
}

func W(x int, y int) Position {
	return Position{x - 1, y}
}

func SW(x int, y int) Position {
	return Position{x - 1, y + 1}
}

func S(x int, y int) Position {
	return Position{x, y + 1}
}

func SE(x int, y int) Position {
	return Position{x + 1, y + 1}
}

func E(x int, y int) Position {
	return Position{x + 1, y}
}

func NE(x int, y int) Position {
	return Position{x + 1, y - 1}
}

func N(x int, y int) Position {
	return Position{x, y - 1}
}

func NW(x int, y int) Position {
	return Position{x - 1, y - 1}
}
