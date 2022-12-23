package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
)

type Operation struct {
	A        string
	Operator string
	B        string
}

func (o *Operation) String() string {
	return fmt.Sprintf("%s %s %s", o.A, o.Operator, o.B)
}

type Monkey struct {
	Name      string
	Number    int64
	Operation *Operation
}

func (m *Monkey) String() string {
	if m.Operation == nil {
		return fmt.Sprintf("%s: %d", m.Name, m.Number)
	}
	return fmt.Sprintf("%s: %s", m.Name, m.Operation.String())
}

func Validate(line string, monkey *Monkey) error {
	if monkey.String() != line {
		return fmt.Errorf("parse error: expected %s, got %s", line, monkey.String())
	}
	return nil
}

func main() {
	file := util.MustReadFile("./internal/dec21/input.txt")
	lines := util.SplitByLine(file)
	monkeys := util.Map(func(line string) *Monkey {
		numberMatch, ok := util.RegexpMatch("([a-z]{4}): (\\d+)", line)
		if ok {
			monkey := &Monkey{
				Name:      numberMatch[1],
				Number:    util.MustAtoi64(numberMatch[2]),
				Operation: nil,
			}
			if err := Validate(line, monkey); err != nil {
				log.Fatal(err)
			}
			return monkey
		}
		operationMatch, ok := util.RegexpMatch("([a-z]{4}): ([a-z]{4}) ([-+/*]) ([a-z]{4})", line)
		if ok {
			monkey := &Monkey{
				Name:   operationMatch[1],
				Number: 0,
				Operation: &Operation{
					A:        operationMatch[2],
					Operator: operationMatch[3],
					B:        operationMatch[4],
				},
			}
			if err := Validate(line, monkey); err != nil {
				log.Fatal(err)
			}
			return monkey
		}
		log.Fatal("Could not parse line ", line)
		return nil
	}, lines)

	monkeyMap := util.Reduce(
		func(m map[string]*Monkey, monkey *Monkey) map[string]*Monkey {
			return util.Assoc(m, monkey.Name, monkey)
		},
		monkeys,
		map[string]*Monkey{},
	)

	rootSolutiono := Eval(monkeyMap, monkeyMap["root"])
	fmt.Println(rootSolutiono)
}

var Operations = map[string]func(int64, int64) int64{
	"+": func(a int64, b int64) int64 {
		return a + b
	},
	"-": func(a int64, b int64) int64 {
		return a - b
	},
	"*": func(a int64, b int64) int64 {
		return a * b
	},
	"/": func(a int64, b int64) int64 {
		return a / b
	},
}

func Eval(monkeyMap map[string]*Monkey, monkey *Monkey) int64 {
	if monkey.Operation == nil {
		return monkey.Number
	}
	a := monkey.Operation.A
	b := monkey.Operation.B
	f := Operations[monkey.Operation.Operator]
	return f(Eval(monkeyMap, monkeyMap[a]), Eval(monkeyMap, monkeyMap[b]))
}
