package util

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
	"sort"
)

func NewStep[N constraints.Integer](m, n N) []N {
	var ns []N
	for i := m; i <= n; i++ {
		ns = append(ns, i)
	}
	return ns
}

func NewRange[N constraints.Integer](start, end N) []N {
	ints := make([]N, 0)
	if start <= end {
		for i := start; i <= end; i++ {
			ints = append(ints, i)
		}
		return ints
	}
	for i := start; i >= end; i-- {
		ints = append(ints, i)
	}
	return ints
}

type LazyRange struct {
	Start int
	End   int
}

func NewLazyRange(start, end int) LazyRange {
	if start > end {
		log.Fatal("ordering invariant error, start=", start, " end=", end)
	}
	return LazyRange{
		Start: start,
		End:   end,
	}
}

func (base LazyRange) Merge(addl LazyRange) (LazyRange, bool) {
	if addl.Start <= base.Start && addl.End >= base.End {
		return addl, true
	}
	if base.Start <= addl.Start && base.End >= addl.End {
		return base, true
	}
	if base.Start <= addl.Start && addl.Start <= base.End {
		return NewLazyRange(base.Start, addl.End), true
	}
	if addl.Start <= base.Start && base.Start <= addl.End {
		return NewLazyRange(addl.Start, base.End), true
	}
	fmt.Println("cannot merge ", base, " ", addl)
	return LazyRange{}, false
}

type LazyRanges []LazyRange

func (rs LazyRanges) Sort() {
	sort.Slice(rs, func(i, j int) bool {
		if rs[i].Start < rs[j].Start {
			return true
		}
		if rs[i].Start > rs[j].Start {
			return false
		}
		return rs[i].End < rs[j].End
	})
}

func (rs LazyRanges) Merge() LazyRanges {
	rs.Sort()
	var merged LazyRanges
	for i := 0; i < len(rs); i++ {
		if len(merged) == 0 {
			merged = append(merged, rs[i])
			continue
		}
		if newLr, ok := merged[len(merged)-1].Merge(rs[i]); ok {
			merged[len(merged)-1] = newLr
			continue
		}
		merged = append(merged, rs[i])
	}
	return merged
}
