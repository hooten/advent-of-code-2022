package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//partOne()
	partTwo()
}

func readFile() string {
	bytes, err := os.ReadFile("./internal/dec04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func partTwo() {
	input := readFile()
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		fmt.Println(line)
		assignments := strings.Split(line, ",")
		if len(assignments) < 2 {
			log.Fatal("bad length", len(assignments))
		}
		m := make(map[int]int)
		for _, assignment := range assignments {
			rangeStr := strings.Split(assignment, "-")
			start, _ := strconv.Atoi(rangeStr[0])
			end, _ := strconv.Atoi(rangeStr[1])
			for i := start; i <= end; i++ {
				m[i]++
			}
		}
		for i, n := range m {
			if n > 1 {
				total++
				fmt.Println("overlap at", i)
				break
			}
		}

	}
	fmt.Println("total", total)
}

func partOne() {
	input := readFile()
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		assignments := strings.Split(line, ",")
		if len(assignments) < 2 {
			fmt.Printf("%q", line)
			log.Fatal("bad length")
		}
		oneAssign := strings.Split(assignments[0], "-")
		twoAssign := strings.Split(assignments[1], "-")
		one0, _ := strconv.Atoi(oneAssign[0])
		two0, _ := strconv.Atoi(twoAssign[0])
		one1, _ := strconv.Atoi(oneAssign[1])
		two1, _ := strconv.Atoi(twoAssign[1])
		if one0 <= two0 && one1 >= two1 {
			fmt.Printf("one: %q is contained in %q\n", twoAssign, oneAssign)
			total++
			continue
		}
		if two0 <= one0 && two1 >= one1 {
			fmt.Printf("two: %q is contained in %q\n", oneAssign, twoAssign)
			total++
			continue
		}
		fmt.Printf("%q is NOT contained in %q or vice versa\n", oneAssign, twoAssign)
	}
	fmt.Println("total", total)
}

func toSet(s string) map[string]bool {
	as := strings.Split(s, "")
	m := map[string]bool{}
	for _, a := range as {
		m[a] = true
	}
	return m
}
