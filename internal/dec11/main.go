package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"github.com/kr/pretty"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Monkey struct {
	Items   []int64
	Op      func(int64) int64
	Test    func(int64) bool
	IFTrue  int
	IFFalse int
}

func main() {
	file := util.MustReadFile("./internal/dec11/input.txt")

	rawMonkeyText := strings.Split(file, "\n\n")
	monkeysText := util.Filter(func(line string) bool {
		return line != "\n" && line != ""
	}, rawMonkeyText)

	var common = int64(1)
	var monkeys []*Monkey
	for _, text := range monkeysText {
		monkey := &Monkey{}
		for _, line := range strings.Split(text, "\n") {
			startingItems, _ := util.RegexpMatch("Starting items: (.*)", line)
			if len(startingItems) > 0 {
				xs := strings.Split(startingItems[1], ", ")
				monkey.Items = util.Map(util.MustAtoi64, xs)
			}
			ops, _ := util.RegexpMatch("Operator: new = old (.) (.*)", line)
			if len(ops) > 0 {
				if len(ops) < 2 {
					log.Fatal(line)
				}
				s := ops[2]
				operand, err := strconv.ParseInt(s, 10, 0)
				if err != nil {
					operand = -100
				}
				monkey.Op = func(x int64) int64 {
					start := time.Now()
					defer func() {
						dur := time.Since(start)
						if dur.Seconds() > 1 {
							fmt.Println("op time for", ops[1], ops[2], dur)
						}
					}()
					switch ops[1] {
					case "+":
						if operand == -100 {
							return x + x
						}
						return x + operand
					case "*":
						if operand == -100 {
							return x * x
						}
						return x * operand
					}
					log.Fatal(ops)
					return x
				}
			}
			test, _ := util.RegexpMatch("Test: divisible by (.*)", line)
			if len(test) > 0 {
				s := test[1]
				atoi, err := strconv.ParseInt(s, 10, 0)
				if err != nil {
					log.Fatal(err)
				}
				common *= atoi
				monkey.Test = func(x int64) bool {
					start := time.Now()
					defer func() {
						dur := time.Since(start)
						if dur.Seconds() > 1 {
							fmt.Println("test time", dur)
						}
					}()
					return x%atoi == 0
				}
			}
			iftrue, _ := util.RegexpMatch("If true: throw to monkey (.*)", line)
			if len(iftrue) > 0 {
				monkey.IFTrue = int(util.MustAtoi64(iftrue[1]))
			}
			iffalse, _ := util.RegexpMatch("If false: throw to monkey (.*)", line)
			if len(iffalse) > 0 {
				monkey.IFFalse = int(util.MustAtoi64(iffalse[1]))
			}
		}
		monkeys = append(monkeys, monkey)
	}
	pretty.Print(monkeys)

	inspections := make([]int, len(monkeys))
	rounds := 10000
	for r := 0; r < rounds; r++ {
		fmt.Println("round", r)
		for i := range monkeys {
			for len(monkeys[i].Items) > 0 {
				item := monkeys[i].Items[0]
				start := time.Now()
				inspections[i]++

				worry := monkeys[i].Op(item)

				worry %= common

				var throwTo int
				test := monkeys[i].Test(worry)
				if test {
					throwTo = int(monkeys[i].IFTrue)
				} else {
					throwTo = int(monkeys[i].IFFalse)
				}
				monkeys[throwTo].Items = append(monkeys[throwTo].Items, worry)
				monkeys[i].Items = monkeys[i].Items[1:]

				dur := time.Since(start)
				if dur.Seconds() > 1 {
					fmt.Println(dur)

				}
			}
		}
	}

	sort.Ints(inspections)
	bus := inspections[len(inspections)-2] * inspections[len(inspections)-1]
	pretty.Print(inspections)
	pretty.Print(bus)
}
