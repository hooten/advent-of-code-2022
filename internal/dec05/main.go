package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"os"
	"strings"
)

func main() {
	partOne()
	//partTwo()
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
	lines := readFile()
	for _, line := range lines {
		fmt.Println(line)
	}
}

func partTwo() {
	lines := readFile()
	for _, line := range lines {
		fmt.Println(line)
	}
}
