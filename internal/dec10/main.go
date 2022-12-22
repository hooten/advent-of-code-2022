package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile() string {
	bytes, err := os.ReadFile("./internal/dec10/input.txt")
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
			return !util.Contains(omitLines, s)
		},
		rawLines,
	)
}

const newline = "\n"
const empty = ""

func main() {
	lines := getFileContentsBy(newline, []string{empty, newline})
	XS := []*util.Pair{util.NewPair(1, 1)}
	for _, line := range lines {
		if line == "noop" {
			XS = append(XS, util.NewPair(XS[len(XS)-1].Y, XS[len(XS)-1].Y))
		} else {
			re := regexp.MustCompile("addx (.*)")
			submatch := re.FindAllStringSubmatch(line, -1)
			value, err := strconv.Atoi(submatch[0][1])
			if err != nil {
				log.Fatal(err)
			}

			XS = append(XS, util.NewPair(XS[len(XS)-1].Y, XS[len(XS)-1].Y))
			XS = append(XS, util.NewPair(XS[len(XS)-1].Y, XS[len(XS)-1].Y+value))
		}
	}
	fmt.Println(
		XS[20].X*20 +
			XS[60].X*60 +
			XS[100].X*100 +
			XS[140].X*140 +
			XS[180].X*180 +
			XS[220].X*220,
	)
	i := 1
	for row := 0; row < 6; row++ {
		fmt.Println("")
		for col := 0; col < 40; col++ {
			if col%5 == 0 {
				fmt.Print("     ")
			}
			if math.Abs(float64(XS[i].X-col)) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			i++
		}
	}
}

//not 14360
