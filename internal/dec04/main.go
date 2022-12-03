package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	partOne()
	// partTwo()
}

func readFile() string {
	bytes, err := os.ReadFile("./internal/dec04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func partOne() {
	input := readFile()
}

func partTwo() {
	input := readFile()

}

func toSet(s string) map[string]bool {
	as := strings.Split(s, "")
	m := map[string]bool{}
	for _, a := range as {
		m[a] = true
	}
	return m
}
