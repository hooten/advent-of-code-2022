package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"github.com/kr/pretty"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//partOne()
	partTwo()
}

func readFile() []string {
	bytes, err := os.ReadFile("./internal/dec05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rawLines := strings.Split(string(bytes), "\n")
	return util.Filter(
		func(s string) bool {
			return s != ""
		},
		rawLines,
	)
}

func partOne() {
	matrix := [][]string{
		{"R", "G", "J", "B", "T", "V", "Z"},
		{"J", "R", "V", "L"},
		{"S", "Q", "F"},
		{"Z", "H", "N", "L", "F", "V", "Q", "G"},
		{"R", "Q", "T", "J", "C", "S", "M", "W"},
		{"S", "W", "T", "C", "H", "F"},
		{"D", "Z", "C", "V", "F", "N", "J"},
		{"L", "G", "Z", "D", "W", "R", "F", "Q"},
		{"J", "B", "W", "V", "P"},
	}
	lines := util.Filter(
		func(line string) bool {
			return strings.HasPrefix(line, "move")
		},
		readFile(),
	)
	pretty.Print(matrix)
	for _, line := range lines {
		total := 0
		fmt.Println(line)
		re := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
		submatch := re.FindAllStringSubmatch(line, -1)
		if len(submatch[0]) != 4 {
			log.Fatal("bad match", submatch[0])
		}
		qty, _ := strconv.Atoi(submatch[0][1])
		srcRaw, _ := strconv.Atoi(submatch[0][2])
		src := srcRaw - 1
		dstRaw, _ := strconv.Atoi(submatch[0][3])
		dst := dstRaw - 1
		fmt.Println(qty, src, dst)
		for i := 0; i < qty; i++ {
			index := len(matrix[src]) - 1
			elem := matrix[src][index:][0]
			matrix[src] = matrix[src][:index]
			matrix[dst] = append(matrix[dst], elem)
		}
		pretty.Print(matrix)
		for _, row := range matrix {

			for i := 0; i < len(row); i++ {
				total++
			}
		}
		fmt.Println("total", total)
	}
	s := ""
	for _, row := range matrix {
		char := row[len(row)-1]
		s = s + char
	}
	fmt.Println(s)
}

func partTwo() {
	matrix := [][]string{
		{"R", "G", "J", "B", "T", "V", "Z"},
		{"J", "R", "V", "L"},
		{"S", "Q", "F"},
		{"Z", "H", "N", "L", "F", "V", "Q", "G"},
		{"R", "Q", "T", "J", "C", "S", "M", "W"},
		{"S", "W", "T", "C", "H", "F"},
		{"D", "Z", "C", "V", "F", "N", "J"},
		{"L", "G", "Z", "D", "W", "R", "F", "Q"},
		{"J", "B", "W", "V", "P"},
	}
	lines := util.Filter(
		func(line string) bool {
			return strings.HasPrefix(line, "move")
		},
		readFile(),
	)
	pretty.Print(matrix)
	for _, line := range lines {
		total := 0
		fmt.Println(line)
		re := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
		submatch := re.FindAllStringSubmatch(line, -1)
		if len(submatch[0]) != 4 {
			log.Fatal("bad match", submatch[0])
		}
		qty, _ := strconv.Atoi(submatch[0][1])
		srcRaw, _ := strconv.Atoi(submatch[0][2])
		src := srcRaw - 1
		dstRaw, _ := strconv.Atoi(submatch[0][3])
		dst := dstRaw - 1
		fmt.Println(qty, src, dst)

		index := len(matrix[src]) - qty
		elems := matrix[src][index:]
		matrix[src] = matrix[src][:index]
		matrix[dst] = append(matrix[dst], elems...)
		pretty.Print(matrix)
		for _, row := range matrix {

			for i := 0; i < len(row); i++ {
				total++
			}
		}
		fmt.Println("total", total)
	}
	s := ""
	for _, row := range matrix {
		char := row[len(row)-1]
		s = s + char
	}
	fmt.Println(s)
}
