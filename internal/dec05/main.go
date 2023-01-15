package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/hooten/advent-of-code-2022/pkg/util"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	stacks, moves, err := parseFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rearrangedStack := Rearrange(stacks, moves, false)
	fmt.Printf("Part 1: The CrateMover 9000 makes the top crates: %s.\n", StackTops(rearrangedStack))
}

func partTwo() {
	stacks, moves, err := parseFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rearrangedQueue := Rearrange(stacks, moves, true)
	fmt.Printf("Part 2: The CrateMover 9001 makes the top crates %s.\n", StackTops(rearrangedQueue))
}

func StackTops(stacks []Stack) string {
	return util.Reduce(func(s string, stack Stack) string {
		return s + stack.Peek()
	}, stacks, "")
}

func Rearrange(stacks []Stack, moves []Move, queue bool) []Stack {
	if len(moves) == 0 {
		return stacks
	}
	move := moves[0]
	newMoves := moves[1:]

	if queue {
		from := move.CrateSrc
		to := move.CrateDst
		newFrom, elems := stacks[from-1].Pop(move.N)
		newTo := stacks[to-1].Append(elems...)
		newStacks := make([]Stack, len(stacks))
		copy(newStacks, stacks)
		newStacks[from-1] = newFrom
		newStacks[to-1] = newTo
		return Rearrange(newStacks, newMoves, queue)
	}
	newStacks := util.Reduce(func(stacks []Stack, i int) []Stack {
		from := move.CrateSrc
		to := move.CrateDst
		newFrom, elems := stacks[from-1].Pop(1)
		newTo := stacks[to-1].Append(elems[0])
		newStacks := make([]Stack, len(stacks))
		copy(newStacks, stacks)
		newStacks[from-1] = newFrom
		newStacks[to-1] = newTo
		return newStacks
	}, util.NewRange(1, move.N), stacks)
	return Rearrange(newStacks, newMoves, queue)
}

type Stack struct {
	Crates []string
}

func (s Stack) Prepend(crate string) Stack {
	return Stack{
		Crates: append([]string{crate}, s.Crates...),
	}
}

func (s Stack) Append(crates ...string) Stack {
	return Stack{
		Crates: append(s.Crates, crates...),
	}
}

func (s Stack) Pop(n int) (Stack, []string) {
	return Stack{
		Crates: s.Crates[:len(s.Crates)-n],
	}, s.Crates[len(s.Crates)-n:]
}

func (s Stack) Peek() string {
	return s.Crates[len(s.Crates)-1]
}

func ParseStacks(lines []string) ([]Stack, error) {
	width := len(lines[0])
	n := (width + 1) / 4
	return util.Reduce(func(stacks []Stack, line string) []Stack {
		return util.MapWithIndex(func(stack Stack, i int) Stack {
			x := i*4 + 1
			crate := line[x : x+1]
			if regexp.MustCompile("[A-Z]").MatchString(crate) {
				return stack.Prepend(crate)
			}
			return stack
		}, stacks)
	}, lines, make([]Stack, n)), nil

}

type Move struct {
	N        int
	CrateSrc int
	CrateDst int
}

func ParseMoves(lines []string) []Move {
	re := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	return util.Map(func(line string) Move {
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			log.Fatalf("expected 4 matches, got %v for line %q", matches, line)
		}
		return Move{
			N:        util.MustAtoi(matches[1]),
			CrateSrc: util.MustAtoi(matches[2]),
			CrateDst: util.MustAtoi(matches[3]),
		}
	}, lines)
}

func parseFile(filename string) ([]Stack, []Move, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	parts := strings.Split(string(bytes), "\n\n")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("expected file to have 2 parts, got %d", len(parts))
	}
	firstPart := strings.Split(parts[0], "\n")
	stacks, err := ParseStacks(firstPart)
	if err != nil {
		return nil, nil, err
	}
	secondPart := util.Filter(func(line string) bool {
		return line != ""
	}, strings.Split(parts[1], "\n"))
	moves := ParseMoves(secondPart)
	return stacks, moves, nil
}
