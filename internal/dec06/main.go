package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readFile()
	compute(lines[0], 14)
}

func readFile() []string {
	bytes, err := os.ReadFile("./internal/dec06/input.txt")
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

func compute(line string, n int) {
	chars := strings.Split(line, "")
	for i := n - 1; i < len(chars); i++ {
		xs := chars[i-(n-1) : i+1]
		set := util.ToSet(xs)
		keys := util.Keys(set)
		if len(keys) == n {
			fmt.Println(i + 1)
			break
		}
	}
}

