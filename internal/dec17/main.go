package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"strings"
)

func getFilename(test bool) string {
	base := "./internal/dec17"
	if test {
		return base + "/test.txt"
	}
	return base + "/input.txt"
}

type Chamber struct {
	tower [][]string
	rocks int
}

func NewChamber() *Chamber {
	tower := make([][]string, 0)
	return &Chamber{
		tower: tower,
	}
}

func (c *Chamber) Draw() {
	var sb strings.Builder
	for y := len(c.tower) - 1; y >= 0; y-- {
		sb.WriteString("|")
		for x := 0; x < len(c.tower[0]); x++ {
			row := c.tower[y]
			s := row[x]
			sb.WriteString(s)
		}
		sb.WriteString("|\n")
	}
	sb.WriteString("+-------+")
	fmt.Println(sb.String())
}

var EmptyRow = strings.Split(".......", "")

func (c *Chamber) AddRock(rock [][]string) {
	for y := 0; y < 3; y++ {
		c.tower = append(c.tower, EmptyRow)
	}
	c.tower = append(c.tower, rock...)
	c.rocks++
}

func RowEmpty(row []string) bool {
	for _, x := range row {
		if x != "." {
			return false
		}
	}
	return true
}

func (c *Chamber) TopRow() string {
	return strings.Join(c.tower[len(c.tower)-1], "")
}

func (c *Chamber) HasEmpty() (bool, int) {
	for y := len(c.tower) - 1; y >= 0; y-- {
		if RowEmpty(c.tower[y]) {
			return true, y
		}
	}
	return false, -1
}

func (c *Chamber) ApplyJet(jet Jet) bool {
	falling, ok := c.FindFalling()
	if !ok {
		log.Fatal("should have found falling (?)")
	}

	top := falling[0]
	bottom := falling[len(falling)-1]

	newSegment := make([][]string, len(falling))
	existingSegment := c.tower[bottom : top+1]
	for y := len(falling) - 1; y >= 0; y-- {
		existingRow := existingSegment[y]
		newRow := util.Map(func(s string) string {
			if s == "@" {
				return "."
			}
			return s
		}, existingRow)

		if jet.isRight() {
			if existingRow[6] == "@" {
				return false
			}
			for x := 0; x < 6; x++ {
				if existingRow[x] == "@" {
					if existingRow[x+1] == "#" {
						return false
					} else {
						newRow[x+1] = "@"
					}
				}
			}
		}
		if jet.isLeft() {
			if existingRow[0] == "@" {
				return false
			}
			for x := 1; x < 7; x++ {
				if existingRow[x] == "@" {
					if existingRow[x-1] == "#" {
						return false
					} else {
						newRow[x-1] = "@"
					}
				}
			}
		}
		newSegment[y] = newRow
	}

	for y := 0; y < len(newSegment); y++ {
		c.tower[bottom+y] = newSegment[y]
	}
	return true
}

func (c *Chamber) Fall() bool {
	if ok, y := c.HasEmpty(); ok {
		base := c.tower[:y]
		top := c.tower[y+1:]
		c.tower = append(base, top...)
		return true
	}
	falling, ok := c.FindFalling()
	if !ok {
		log.Fatal("should have found falling (?)")
	}

	top := falling[0]
	bottom := falling[len(falling)-1]
	if bottom == 0 {
		return false
	}
	existingSegment := c.tower[bottom-1 : top+1]
	newSegment, err := FallSegment(existingSegment)
	if err != nil {
		return false
	}
	for y := 0; y < len(newSegment); y++ {
		c.tower[bottom-1+y] = newSegment[y]
	}
	if RowEmpty(c.tower[len(c.tower)-1]) {
		c.tower = c.tower[:len(c.tower)-1]
	}
	return true
}

func FallSegment(existingSegmentRaw [][]string) ([][]string, error) {
	existingSegmentNormal := make([][]string, len(existingSegmentRaw))
	for y, row := range existingSegmentRaw {
		existingSegmentNormal[y] = util.Map(func(s string) string {
			if s == "@" {
				return "."
			}
			return s
		}, row)
	}

	newSegment := make([][]string, len(existingSegmentNormal))
	lastIndex := len(existingSegmentNormal) - 1
	for y := 0; y < lastIndex; y++ {
		newRow := make([]string, 7)
		for x := 0; x < 7; x++ {
			if existingSegmentRaw[y+1][x] == "@" {
				if existingSegmentNormal[y][x] == "#" {
					return nil, fmt.Errorf("blocked")
				}
				newRow[x] = "@"
			} else {
				newRow[x] = existingSegmentNormal[y][x]
			}
		}
		newSegment[y] = newRow
	}
	lastRow := existingSegmentNormal[lastIndex]
	newSegment[lastIndex] = lastRow
	return newSegment, nil
}

func (c *Chamber) Solidify() {
	for y := 0; y < len(c.tower); y++ {
		for x := 0; x < 7; x++ {
			s := c.tower[y][x]
			if s == "@" {
				c.tower[y][x] = "#"
			}
		}
	}
}

func (c *Chamber) FindFalling() ([]int, bool) {
	var rows []int
	for y := len(c.tower) - 1; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if c.tower[y][x] == "@" {
				rows = append(rows, y)
				break
			}
		}
		if len(rows) >= 4 {
			break
		}
	}
	return rows, len(rows) > 0
}

type Jet string

func (j Jet) Delta() int {
	if j.isRight() {
		return 1
	}
	return -1

}

func (j Jet) isLeft() bool {
	return j == "<"
}

func (j Jet) isRight() bool {
	return j == ">"
}

var RockMinus = []string{
	"..@@@@.",
}

var RockPlus = []string{
	"...@...",
	"..@@@..",
	"...@...",
}

var RockL = []string{
	"..@@@..",
	"....@..",
	"....@..",
}

var RockPipe = []string{
	"..@....",
	"..@....",
	"..@....",
	"..@....",
}

var RockSquare = []string{
	"..@@...",
	"..@@...",
}

func NormalizeRock(rock []string) [][]string {
	normal := make([][]string, len(rock))
	for i, straition := range rock {
		split := strings.Split(straition, "")
		normal[i] = split
	}
	return normal
}

type RockPattern struct {
	rocks [][][]string
	pos   int
}

func NewRockPattern() *RockPattern {
	return &RockPattern{
		rocks: [][][]string{
			NormalizeRock(RockMinus),
			NormalizeRock(RockPlus),
			NormalizeRock(RockL),
			NormalizeRock(RockPipe),
			NormalizeRock(RockSquare),
		},
		pos: -1,
	}
}

func (rp *RockPattern) Next() [][]string {
	rp.pos++
	if rp.pos >= len(rp.rocks) {
		rp.pos = 0
	}
	rock := rp.rocks[rp.pos]
	return rock
}

type JetPattern struct {
	raw []string
	pos int
}

func NewJetPattern(raw []string) *JetPattern {
	return &JetPattern{
		raw: raw,
		pos: -1,
	}
}

func (jp *JetPattern) Next() Jet {
	jp.pos++
	if jp.pos >= len(jp.raw) {
		jp.pos = 0
	}
	return Jet(jp.raw[jp.pos])
}

const Test = false
const Debug = false
const Limit = 2022

func main() {
	filename := getFilename(Test)
	file := util.MustReadFile(filename)
	chars := util.SplitByChar(file)
	solution(NewJetPattern(chars), NewRockPattern(), Debug, false)
	solution(NewJetPattern(chars), NewRockPattern(), Debug, true)
}

func solution(jetPattern *JetPattern, rockPattern *RockPattern, debug bool, part2 bool) {
	fmt.Printf("========== PART %s ==========\n", util.Ternary(part2, "TWO", "ONE"))
	chamber := NewChamber()
	if debug {
		fmt.Printf("\n--- Start state ---\n")
		chamber.Draw()
	}
	jet := jetPattern.Next()
	for i := 0; i < Limit; i++ {
		if debug {
			fmt.Printf("\n--- Iteration %d: Add Rock ---\n", i)
		}
		rock := rockPattern.Next()
		chamber.AddRock(rock)
		if debug {
			chamber.Draw()
		}

		var jetApplied = true
		var fell = true
		for fell {
			jetApplied = chamber.ApplyJet(jet)
			if debug {
				if jetApplied {
					fmt.Printf("\n--- Iteration %d: Jet Applied (%s) ---\n", i, jet)
				} else {
					fmt.Printf("\n--- Iteration %d: Jet Failed (%s) ---\n", i, jet)
				}
				chamber.Draw()
			}

			fell = chamber.Fall()
			if debug {
				if fell {
					fmt.Printf("\n--- Iteration %d: Fall Applied ---\n", i)
				} else {
					fmt.Printf("\n--- Iteration %d: Fall Failed ---\n", i)
				}
				chamber.Draw()
			}
			jet = jetPattern.Next()
		}
		chamber.Solidify()
		if debug {
			fmt.Printf("\n--- Iteration %d: Solidify ---\n", i)
			chamber.Draw()
		}
	}

	if debug {
		fmt.Printf("\n--- End State ---\n")
		chamber.Draw()
	}
	fmt.Println("height :", len(chamber.tower))
	fmt.Println("==========   END   ==========")
	fmt.Println("")
}
