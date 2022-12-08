package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile() string {
	bytes, err := os.ReadFile("./internal/dec08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func getFileContentsBy(sep string, omitLines []string) []string {
	fileBlob := readFile()
	rawLines := strings.Split(fileBlob, sep)
	return util.Filter(
		func(s string) bool {
			return !util.HasElem(omitLines, s)
		},
		rawLines,
	)
}

const newline = "\n"
const empty = ""

type Pair struct {
	x int
	y int
}

func key(x int, y int) string {
	return strconv.Itoa(x) + "-" + strconv.Itoa(y)
}

func coords(key string) (int, int) {
	split := strings.Split(key, "-")
	if len(split) != 2 {
		log.Fatal("bad key", key)
	}
	x, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal("bad coord", split[0])
	}
	y, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal("bad coord", split[1])
	}
	return x, y
}

func main() {
	rows := getFileContentsBy(newline, []string{empty, newline})
	var matrix [][]int
	for i, row := range rows {
		matrix = append(matrix, []int{})
		for _, col := range strings.Split(row, empty) {
			cell, errr := strconv.Atoi(col)
			if errr != nil {
				log.Fatal(errr)
			}
			matrix[i] = append(matrix[i], cell)
		}
	}

	visible := map[string]bool{}
	numCols := len(matrix[0])
	fmt.Println("numcols", numCols)
	numRows := len(rows)
	fmt.Println("numrows", numRows)
	// from top
	for j := 0; j < numCols; j++ {
		highest := -1
		for i := 0; i < numRows; i++ {
			cell := matrix[i][j]
			if cell > highest {
				visible[key(i, j)] = true
				highest = cell
			}
			if highest == 9 {
				break
			}
		}
	}
	//from bottom
	for j := 0; j < numCols; j++ {
		highest := -1
		for i := numRows - 1; i >= 0; i-- {
			cell := matrix[i][j]
			if cell > highest {
				visible[key(i, j)] = true
				highest = cell
			}
			if highest == 9 {
				break
			}
		}
	}
	//from right
	for i := 0; i < numRows; i++ {
		highest := -1
		for j := numCols - 1; j >= 0; j-- {
			cell := matrix[i][j]
			if cell > highest {
				visible[key(i, j)] = true
				highest = cell
			}
			if highest == 9 {
				break
			}
		}
	}
	// from left
	for i := 0; i < numRows; i++ {
		highest := -1
		for j := 0; j < numCols; j++ {
			cell := matrix[i][j]
			if cell > highest {
				visible[key(i, j)] = true
				highest = cell
			}
			if highest == 9 {
				break
			}
		}
	}
	fmt.Println(len(visible))

	scores := map[string]int{}
	for x := 0; x < numRows; x++ {
		for y := 0; y < numCols; y++ {

			var up, down, left, right int
			height := matrix[x][y]
			//right
			for j := y + 1; j < numCols; j++ {
				if matrix[x][j] < height {
					right++
				}
				if matrix[x][j] >= height {
					right++
					break
				}
			}
			// left
			for j := y - 1; j >= 0; j-- {
				if matrix[x][j] < height {
					left++
				}
				if matrix[x][j] >= height {
					left++
					break
				}
			}

			// down
			for i := x + 1; i < numRows; i++ {
				if matrix[i][y] < height {
					down++
				}
				if matrix[i][y] >= height {
					down++
					break
				}
			}

			// up
			for i := x - 1; i >= 0; i-- {
				if matrix[i][y] < height {
					up++
				}
				if matrix[i][y] >= height {
					up++
					break
				}
			}
			scores[key(x, y)] = up * down * left * right
		}
	}

	values := util.Values(scores)
	sort.Ints(values)
	fmt.Println(len(values))
	top := values[len(values)-1]
	fmt.Println(top)
	fmt.Println(top < 3285839138400 && top > 8464)
}

// 3285839138400 too high
// 8464 too low
