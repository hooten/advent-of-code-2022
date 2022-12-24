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

	root := monkeyMap["root"]
	dissoc := util.Dissoc(monkeyMap, "root")
	reducedMonkeyMap := util.Assoc(ReduceMonkeyMap(dissoc, util.Keys(dissoc)), "root", root)

	part1Solution := Eval(reducedMonkeyMap, root)
	fmt.Println("Part 1:", part1Solution)

	FindHumn(reducedMonkeyMap, root)

}

func FindHumn(monkeyMap map[string]*Monkey, root *Monkey) (int, error) {
	lhs := monkeyMap[root.Operation.A]
	rhs := monkeyMap[root.Operation.B]

	a, strA := EvalString(monkeyMap, lhs)
	b, strB := EvalString(monkeyMap, rhs)
	if strA == "" && strB == "" {
		return 0, fmt.Errorf("something went wrong")
	}
	if strA != "" && strB == "" {
		fmt.Println(strA, "==", b)
	}
	if strA == "" && strB != "" {
		fmt.Println(a, "==", strB)
	}

	fmt.Println(strA, "==", strB)
	return 0, nil

}

func ReduceMonkeyMap(m map[string]*Monkey, keys []string) map[string]*Monkey {
	if len(keys) == 0 {
		return m
	}
	key := keys[len(keys)-1]
	newKeys := keys[:len(keys)-1]
	value, err := EvalIgnoringHumn(util.Dissoc(m, key), m[key])
	if err != nil {
		return ReduceMonkeyMap(m, newKeys)
	}
	newM := util.Assoc(m, key, &Monkey{Name: key, Number: value})
	return ReduceMonkeyMap(newM, newKeys)
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

func EvalIgnoringHumn(monkeyMap map[string]*Monkey, monkey *Monkey) (int64, error) {
	if monkey.Name == "humn" {
		return 0, fmt.Errorf("cannot evaluate humn")
	}
	if monkey.Operation == nil {
		return monkey.Number, nil
	}
	a := monkey.Operation.A
	b := monkey.Operation.B
	f := Operations[monkey.Operation.Operator]
	evalA, err := EvalIgnoringHumn(monkeyMap, monkeyMap[a])
	if err != nil {
		return 0, err
	}
	evalB, err := EvalIgnoringHumn(monkeyMap, monkeyMap[b])
	if err != nil {
		return 0, err
	}
	return f(evalA, evalB), nil
}

func EvalString(monkeyMap map[string]*Monkey, monkey *Monkey) (int64, string) {
	if monkey.Name == "humn" {
		return 0, "X"
	}
	if monkey.Operation == nil {
		return monkey.Number, ""
	}
	a := monkey.Operation.A
	b := monkey.Operation.B
	f := Operations[monkey.Operation.Operator]
	evalA, strA := EvalString(monkeyMap, monkeyMap[a])
	evalB, strB := EvalString(monkeyMap, monkeyMap[b])
	if strA == "" && strB == "" {
		return f(evalA, evalB), ""
	}
	if strA == "" && strB != "" {
		return 0, fmt.Sprintf("(%d %s %s)", evalA, monkey.Operation.Operator, strB)
	}
	if strA != "" && strB == "" {
		return 0, fmt.Sprintf("(%s %s %d)", strA, monkey.Operation.Operator, evalB)
	}
	return 0, fmt.Sprintf("(%s %s %s)", strA, monkey.Operation.Operator, strB)
}
