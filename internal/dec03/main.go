package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	partTwo()
}

func partTwo() {
	bytes, err := os.ReadFile("./internal/dec03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	priorities := 0
	lines := strings.Split(string(bytes), "\n")
	for len(lines) > 1 { // newline at the end of the file
		elfGroup := lines[:3]
		fmt.Println("elf group", elfGroup)
		lines = lines[3:]
		freqs := make(map[string]int)
		for _, elf := range elfGroup {
			set := toSet(elf)
			for s := range set {
				freqs[s]++
			}
		}
		for s, n := range freqs {
			if n == 3 {
				if strings.ToLower(s) == s {
					code := int(s[0]) - int('a') + 1
					fmt.Println("common", s, code)
					priorities += code
				} else {
					code := int(s[0]) - int('A') + 1 + 26
					fmt.Println("common", s, code)
					priorities += code
				}
			}
		}
	}
	fmt.Println("priority", priorities)
}

// 1335 is too low
// 22782 is too high

func partOne() {
	bytes, err := os.ReadFile("./internal/dec03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	priorities := 0
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		prefix := toSet(line[:len(line)/2])
		suffix := toSet(line[len(line)/2:])
		m := map[string]int{}
		for s := range prefix {
			m[s]++
		}
		for s := range suffix {
			m[s]++
		}
		for s, n := range m {
			if n == 2 {
				if strings.ToLower(s) == s {
					code := int(s[0]) - int('a') + 1
					fmt.Println("common", s, code)
					priorities += code
				} else {
					code := int(s[0]) - int('A') + 1 + 26
					fmt.Println("common", s, code)
					priorities += code
				}
			}
		}
	}
	fmt.Println("priority", priorities)
}

func toSet(s string) map[string]bool {
	as := strings.Split(s, "")
	m := map[string]bool{}
	for _, a := range as {
		m[a] = true
	}
	return m
}
