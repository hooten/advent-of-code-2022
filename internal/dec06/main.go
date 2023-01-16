package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hooten/advent-of-code-2022/pkg/util"
)

func main() {
	buffer := readFile()
	packetIdx := FindStartOfIdx(buffer, 0, 4)
	fmt.Printf("Part 1: The start of the first packet index is %d.\n", packetIdx)
	messageIdx := FindStartOfIdx(buffer, 0, 14)
	fmt.Printf("Part 2: The start of the first message index is %d.\n", messageIdx)
}

func readFile() string {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(string(bytes), "\n")
}

func FindStartOfIdx(buffer string, i int, packetLen int) int {
	if len(buffer) < packetLen {
		return -1
	}
	if len(util.ToSet(strings.Split(buffer[:packetLen], ""))) == packetLen {
		return i + packetLen
	}
	return FindStartOfIdx(buffer[1:], i+1, packetLen)
}
