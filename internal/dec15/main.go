package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"math"
	"sort"
)

func main() {
	test := false
	debug := false
	filename, part1Y, part2Max := getInputs(test)
	file := util.MustReadFile(filename)
	lines := util.SplitByLine(file)
	lineData, bounds := readLineData(lines, debug)

	partOne(lineData, part1Y, bounds, debug)
	partTwo(lineData, part2Max, debug)
}

func getInputs(test bool) (filename string, part1Y int, part2Max int) {
	if test {
		return "./internal/dec15/test.txt", 10, 20
	}
	return "./internal/dec15/input.txt", 2000000, 4000000
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

type LineDatum struct {
	sensor   *util.Pair
	distance int
}

func partOne(lineData []LineDatum, row int, bounds Bounds, debug bool) {
	fmt.Println("========== PART ONE ==========")
	lrs := generateLazyRanges(lineData, row, bounds.MinX, bounds.MaxX, debug)
	merge := lrs.Merge()
	if len(merge) == 0 || len(merge) > 2 {
		log.Fatal("bad merge ", merge)
	}
	if len(merge) == 1 {
		fmt.Println("lazy range for row =", row, " : ", merge[0].Start, "to", merge[0].End)
		fmt.Println("part one answer =", merge[0].End-merge[0].Start)
	}
	fmt.Println("==========   END   ==========")
	fmt.Println("")
}

func partTwo(lineData []LineDatum, max int, debug bool) {
	fmt.Println("========== PART TWO ==========")
	for row := 0; row <= max; row++ {
		lrs := generateLazyRanges(lineData, row, 0, max, debug)
		merge := lrs.Merge()
		if len(merge) == 0 || len(merge) > 2 {
			log.Fatal("bad merge ", merge)
		}
		if len(merge) == 1 {
			if debug {
				fmt.Println("lazy range for row=", row, " : ", merge[0].Start, "-", merge[0].End)
			}
		}
		if len(merge) == 2 {
			fmt.Println("found! ", merge)
			finalX := merge[0].End + 1
			finalY := row
			fmt.Println("x =", finalX, ", y =", finalY)
			fmt.Println("part two answer =", finalX*4000000+finalY)
			break
		}
	}
	fmt.Println("==========   END   ==========")
	fmt.Println("")
}

type Bounds struct {
	MinX int
	MaxX int
}

func readLineData(lines []string, debug bool) ([]LineDatum, Bounds) {
	var lineData []LineDatum
	var maxX int
	var maxDistance int
	for _, line := range lines {
		if debug {
			fmt.Println("reading ", line)
		}
		matches, ok := util.RegexpMatch("Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)", line)
		if !ok {
			log.Fatal("bad match ", line)
		}
		sensor := util.NewPair(util.MustAtoi(matches[1]), util.MustAtoi(matches[2]))
		if debug {
			fmt.Println("sensor ", sensor)
		}

		beacon := util.NewPair(util.MustAtoi(matches[3]), util.MustAtoi(matches[4]))
		if debug {
			fmt.Println("beacon ", beacon)
		}

		distance := taxicabDistance(sensor.X, beacon.X, sensor.Y, beacon.Y)
		if debug {
			fmt.Println("distance ", distance)
		}

		if sensor.X > maxX {
			maxX = sensor.X
		}
		if distance > maxDistance {
			maxDistance = distance
		}

		lineData = append(lineData, LineDatum{
			sensor:   sensor,
			distance: distance,
		})
	}
	return lineData, Bounds{
		MinX: -maxDistance,
		MaxX: maxX + maxDistance,
	}
}

func generateLazyRanges(lineData []LineDatum, row int, min int, max int, debug bool) LazyRanges {
	var lrs LazyRanges
	for i, lineDatum := range lineData {
		distance := lineDatum.distance
		sensor := lineDatum.sensor
		yDelta := int(math.Abs(float64(sensor.Y - row)))
		if debug {
			fmt.Printf("- Computing line %d. Sensor=%s  Distance=%d yDelta=%d\n. ", i, sensor.Key(), distance, yDelta)
		}

		offset := distance - yDelta
		start := sensor.X - offset
		if start < min {
			start = min
		}
		endX := sensor.X + offset
		if endX > max {
			endX = max
		}
		if start > endX {
			continue
		}
		lazyRange := NewLazyRange(start, endX)
		lrs = append(lrs, lazyRange)
	}
	lrs.Sort()
	return lrs
}

func taxicabDistance(x1 int, x2 int, y1 int, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}
