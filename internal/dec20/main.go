package main

import (
	"container/list"
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"github.com/kr/pretty"
	"log"
	"strconv"
)

const DecryptionKey int64 = 811589153

func toIntSlice(l *list.List) []int64 {
	var slice []int64
	for e := l.Front(); e != nil; e = e.Next() {
		slice = append(slice, e.Value.(int64))
	}
	return slice
}

func toElemSlice(l *list.List) []*list.Element {
	var slice []*list.Element
	for e := l.Front(); e != nil; e = e.Next() {
		slice = append(slice, e)
	}
	return slice
}

func newLinkedList(numbers []int64) *list.List {
	return util.Reduce(func(l *list.List, n int64) *list.List {
		l.PushBack(n)
		return l
	}, numbers, list.New())
}

func main() {
	file := util.MustReadFile("./internal/dec20/input.txt")
	lines := util.SplitByLine(file)
	part2 := true
	numbers := util.Map(func(line string) int64 {
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if part2 {
			return value * DecryptionKey
		}
		return value
	}, lines)

	l := newLinkedList(numbers)
	orig := toElemSlice(l)

	iterations := util.Ternary(part2, 10, 1)
	for i := 0; i < iterations; i++ {
		Mix(l, orig)
		pretty.Println(toIntSlice(l))
	}
	fmt.Println("Grove Sum", NewGroveSum(l))
}

func Find(l *list.List, n int64) *list.Element {
	return findElem(l.Front(), n)
}

func findElem(e *list.Element, n int64) *list.Element {
	value := e.Value.(int64)
	if value == n {
		return e
	}
	return findElem(e.Next(), n)
}

func Mix(l *list.List, orig []*list.Element) {
	for _, e := range orig {
		if e.Value.(int64) == 0 {
			continue
		}
		n := e.Value.(int64) % int64(l.Len()-1)
		mark := e
		if n > 0 {
			for j := int64(0); j <= n; j++ {
				mark = mark.Next()
				if mark == nil {
					mark = l.Front()
				}
			}
			l.MoveBefore(e, mark)
		} else {
			for j := int64(0); j >= n; j-- {
				mark = mark.Prev()
				if mark == nil {
					mark = l.Back()
				}
			}
			l.MoveAfter(e, mark)
		}
	}
}

func NewGroveSum(l *list.List) int64 {
	var elem = Find(l, 0)

	groveSum := int64(0)
	for m := 1; m <= 3; m++ {
		curr := elem
		offset := (m * 1000) % l.Len()
		for i := 0; i < offset; i++ {
			curr = curr.Next()
			if curr == nil {
				curr = l.Front()
			}
		}
		val := curr.Value.(int64)
		groveSum += val
	}
	return groveSum
}
