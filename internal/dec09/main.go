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
	bytes, err := os.ReadFile("./internal/dec09/input.txt")
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

func key(x, y int) string {
	return fmt.Sprintf("%d-%d\n", x, y)
}

type Pair struct {
	x int
	y int
}

func newPair() *Pair {
	return &Pair{
		x: 0,
		y: 0,
	}
}

func main() {
	lines := getFileContentsBy(newline, []string{empty, newline})
	var xs []*Pair
	knots := 2
	for e := 0; e < knots; e++ {
		xs = append(xs, newPair())
	}
	tailVisited := map[string]bool{
		key(0, 0): true,
	}
	for i, line := range lines {
		fmt.Sprintln(i)
		//if i > 25 {
		//	break
		//}
		re := regexp.MustCompile("([UDLR]) (.*)")
		submatch := re.FindAllStringSubmatch(line, -1)
		direction := submatch[0][1]
		n, _ := strconv.Atoi(submatch[0][2])
		fmt.Println(direction, n)
		for i := 0; i < n; i++ {
			switch direction {
			case "U":
				xs[0].y += 1
			case "D":
				xs[0].y -= 1
			case "L":
				xs[0].x -= 1
			case "R":
				xs[0].x += 1
			default:
				log.Fatal("bad direction", direction)
			}

			for curr := 1; curr < knots; curr++ {
				prev := curr - 1
				if (math.Abs(float64(xs[prev].x-xs[curr].x)) > 1 ||
					math.Abs(float64(xs[prev].y-xs[curr].y)) > 1) &&
					(xs[prev].x != xs[curr].x && xs[prev].y != xs[curr].y) {
					// must diag
					if xs[prev].x > xs[curr].x && xs[prev].y > xs[curr].y {
						xs[curr].x++
						xs[curr].y++
					}
					if xs[prev].x < xs[curr].x && xs[prev].y < xs[curr].y {
						xs[curr].x--
						xs[curr].y--
					}
					if xs[prev].x > xs[curr].x && xs[prev].y < xs[curr].y {
						xs[curr].x++
						xs[curr].y--
					}
					if xs[prev].x < xs[curr].x && xs[prev].y > xs[curr].y {
						xs[curr].x--
						xs[curr].y++
					}
				}

				if math.Abs(float64(xs[prev].x-xs[curr].x)) > 1 && xs[prev].y == xs[curr].y {
					if xs[prev].x > xs[curr].x {
						xs[curr].x++
					} else {
						xs[curr].x--
					}
				}

				if math.Abs(float64(xs[prev].y-xs[curr].y)) > 1 && xs[prev].x == xs[curr].x {
					if xs[prev].y > xs[curr].y {
						xs[curr].y++
					} else {
						xs[curr].y--
					}
				}
				if curr == knots-1 {
					tailVisited[key(xs[curr].x, xs[curr].y)] = true
				}

				fmt.Println("position", "(", xs[prev].x, xs[prev].y, ")", "(", xs[curr].x, xs[curr].y, ")")
			}
		}
	}
	//pretty.Print(headVisited)
	//pretty.Print(tailVisited)
	fmt.Println(len(tailVisited))
}

// 10054 is too high
