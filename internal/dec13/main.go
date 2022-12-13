package main

import (
	"encoding/json"
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"github.com/kr/pretty"
	"log"
	"sort"
	"strconv"
	"strings"
)

type ListOrInt struct {
	List  []ListOrInt
	Int   int
	IsInt bool
}

func NewInt(i int) ListOrInt {
	return ListOrInt{
		Int:   i,
		IsInt: true,
	}
}

func NewListOrIntFromAnies(anies []any) ListOrInt {
	var list []ListOrInt
	for _, a := range anies {
		bytes, err := json.Marshal(a)
		if err != nil {
			log.Fatal(err)
		}
		item := ParseListOrInt(string(bytes))
		list = append(list, item)
	}
	return ListOrInt{List: list}
}

func NewList(list ListOrInt) ListOrInt {
	return ListOrInt{
		List: []ListOrInt{list},
	}
}

func ParseListOrInt(s string) ListOrInt {
	var i int
	if err := json.Unmarshal([]byte(s), &i); err == nil {
		return NewInt(i)
	}
	var l []any
	err := json.Unmarshal([]byte(s), &l)
	if err == nil {
		return NewListOrIntFromAnies(l)
	}
	log.Fatal(err)
	return ListOrInt{}
}

func (l ListOrInt) Items() []ListOrInt {
	if l.IsInt {
		log.Fatal("cannot take items of int")
	}
	return l.List
}

func (l ListOrInt) Tail() ListOrInt {
	if l.IsInt {
		log.Fatal("cannot take items of int")
	}
	return ListOrInt{List: l.List[1:]}
}

func (l ListOrInt) String() string {
	if l.IsInt {
		return strconv.Itoa(l.Int)
	}
	var s []string
	for _, item := range l.List {
		s = append(s, item.String())
	}
	return "[" + strings.Join(s, ",") + "]"
}

func (l ListOrInt) Less(r ListOrInt, indent int, debug bool) int {
	if debug {
		fmt.Printf("%s- Compare %s vs %s\n", strings.Repeat(" ", indent), l.String(), r.String())
	}
	if l.IsInt && r.IsInt {
		if l.Int < r.Int {
			if debug {
				fmt.Printf("%s- Left side is smaller, so inputs are in the right order\n", strings.Repeat(" ", indent+2))
			}
			return -1
		}
		if l.Int > r.Int {
			if debug {
				fmt.Printf("%s- Right side is smaller, so inputs are not in the right order\n", strings.Repeat(" ", indent+2))
			}
			return 1
		}
		return 0
	}

	if l.IsInt && !r.IsInt {
		if debug {
			fmt.Printf("%s- Mixed types; convert left to [%d] and retry comparison\n", strings.Repeat(" ", indent+2), l.Int)
		}
		return (NewList(l)).Less(r, indent+2, debug)
	}
	if !l.IsInt && r.IsInt {
		if debug {
			fmt.Printf("%s- Mixed types; convert right to [%d] and retry comparison\n", strings.Repeat(" ", indent+2), r.Int)
		}
		return l.Less(NewList(r), indent+2, debug)
	}

	// both lists
	if !l.IsInt && !r.IsInt {
		left := l.Items()
		right := r.Items()
		for i := range left {
			if i >= len(right) {
				if debug {
					fmt.Printf("%s- Right side ran out of items, so inputs are not in the right order\n", strings.Repeat(" ", indent+2))
				}
				return 1
			}
			cmp := (left[i]).Less(right[i], indent+2, debug)
			if cmp != 0 {
				return cmp
			}
		}
		if len(left) < len(right) {
			if debug {
				fmt.Printf("%s- Left side ran out of items, so inputs are in the right order\n", strings.Repeat(" ", indent+2))
			}
			return -1
		}
		return 0
	}
	log.Fatal("why are we here", l, r)
	return 0
}

func main() {
	filename := "./internal/dec13/input.txt"
	file := util.MustReadFile(filename)
	packetPairs := util.SplitByBlock(file)
	correctOrder := 0
	idxs := []string{}
	for i, packetPair := range packetPairs {
		packets := strings.Split(packetPair, "\n")
		if len(packets) < 2 {
			pretty.Print(packets)
			log.Fatal("not right packets", packets)
		}
		leftStr := packets[0]
		rightStr := packets[1]
		if i == 104 {
			fmt.Println(leftStr)
			fmt.Println(rightStr)
		}
		left := ParseListOrInt(leftStr)
		right := ParseListOrInt(rightStr)
		fmt.Printf("\n== Pair %d ==\n", i+1)
		cmp := 0
		for {
			cmp = left.Less(right, 0, true)
			if cmp != 0 {
				break
			}
			left = left.Tail()
			right = right.Tail()
		}

		if cmp == -1 {
			idxs = append(idxs, strconv.Itoa(i+1))
			correctOrder += i + 1
		}
	}
	fmt.Println("correct order score")
	fmt.Println(correctOrder)

	lines := util.SplitByLine(file)
	listOrInts := util.Map(func(s string) ListOrInt {
		return ParseListOrInt(s)
	}, lines)
	dividerA := "[[2]]"
	dividerB := "[[6]]"
	listOrInts = append(listOrInts, ParseListOrInt(dividerA), ParseListOrInt(dividerB))
	sort.Slice(listOrInts, func(i, j int) bool {
		less := listOrInts[i].Less(listOrInts[j], 0, false)
		return less == -1
	})
	decoderKey := 1
	for i, listOrInt := range listOrInts {
		s := listOrInt.String()
		fmt.Println(s)
		if s == dividerA {
			decoderKey *= i + 1
		}
		if s == dividerB {
			decoderKey *= i + 1
		}
	}
	fmt.Println("decoder keey")
	fmt.Println(decoderKey)

}

// part 1: 6395
// part 2: 24921
