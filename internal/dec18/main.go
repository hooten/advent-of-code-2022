package main

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"math"
)

type Triple struct {
	X int
	Y int
	Z int
}

func NewTriple(x, y, z int) Triple {
	return Triple{
		X: x,
		Y: y,
		Z: z,
	}
}

func (t Triple) String() string {
	return fmt.Sprintf("%d,%d,%d", t.X, t.Y, t.Z)
}

func (t Triple) GetAdjacent() []Triple {
	x := t.X
	y := t.Y
	z := t.Z
	return []Triple{
		NewTriple(x+1, y, z),
		NewTriple(x-1, y, z),
		NewTriple(x, y+1, z),
		NewTriple(x, y-1, z),
		NewTriple(x, y, z+1),
		NewTriple(x, y, z-1),
	}
}

type TripleIndex struct {
	m       map[int]map[int]map[int]bool
	triples []Triple
	min     Triple // Min of all three coordinates, not a point.
	max     Triple // Max of all three coordinates, not a point.
}

func NewTripleIndex(triples []Triple) *TripleIndex {
	ti := &TripleIndex{
		m:       map[int]map[int]map[int]bool{},
		triples: triples,
		min:     Triple{X: math.MaxInt, Y: math.MaxInt, Z: math.MaxInt},
		max:     Triple{X: math.MinInt, Y: math.MinInt, Z: math.MinInt},
	}
	for _, triple := range triples {
		ti.addToMap(triple.X, triple.Y, triple.Z)
		ti.addMinsAndMaxes(triple)
	}
	return ti
}

func (ti *TripleIndex) AddAll(ti2 *TripleIndex) {
	for _, triple := range ti2.triples {
		ti.Add(triple)
	}
}
func (ti *TripleIndex) Add(triple Triple) {
	ti.addToMap(triple.X, triple.Y, triple.Z)
	ti.addMinsAndMaxes(triple)
	ti.triples = append(ti.triples, triple)
}

func (ti *TripleIndex) Draw() {
	for i := ti.min.X - 1; i <= ti.max.X; i++ {
		fmt.Println("x =", i)
		for j := ti.min.Y - 1; j <= ti.max.Y; j++ {
			for k := ti.min.Z - 1; k <= ti.max.Z; k++ {
				if ti.Get(i, j, k) {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println("")
		}
		fmt.Print("\n\n")
	}
}

func (ti *TripleIndex) GetExteriorSidesShowing() int {

	graph := dijkstra.NewGraph()
	potentialBubbles := NewTripleIndex([]Triple{})
	for i := ti.min.X - 1; i <= ti.max.X+1; i++ {
		for j := ti.min.Y - 1; j <= ti.max.Y+1; j++ {
			for k := ti.min.Z - 1; k <= ti.max.Z+1; k++ {
				if !ti.Get(i, j, k) {
					triple := NewTriple(i, j, k)
					graph.AddMappedVertex(triple.String())
					potentialBubbles.Add(triple)
				}
			}
		}
	}

	for _, triple := range potentialBubbles.triples {
		for _, adj := range triple.GetAdjacent() {
			if err := graph.AddMappedArc(triple.String(), adj.String(), 1); err != nil {
				log.Fatal("Line 111", err)
			}
		}
	}

	exterior := potentialBubbles.min
	if potentialBubbles.Get(exterior.X-1, exterior.Y-1, exterior.Z-1) {
		log.Fatal("bad assumption about exterior point")
	}
	exteriorVertex, err := graph.GetMapping(exterior.String())
	if err != nil {
		log.Fatal("Get Mapping for ", exterior, err)
	}

	bubbles := NewTripleIndex([]Triple{})
	for _, triple := range potentialBubbles.triples {
		vertex, err := graph.GetMapping(triple.String())
		if err != nil {
			log.Fatal(err)
		}
		if vertex != exteriorVertex {

			if _, err := graph.Shortest(vertex, exteriorVertex); err != nil {
				if err == dijkstra.ErrNoPath {
					bubbles.Add(triple)
				} else {
					log.Fatal(err)
				}
			}
		}
	}
	bubbles.AddAll(ti)
	return bubbles.GetAllSidesShowing()
}

func (ti *TripleIndex) GetAllSidesShowing() int {
	sidesShowing := 0
	for _, triple := range ti.triples {
		for _, adj := range triple.GetAdjacent() {
			if sideCovered := ti.Get(adj.X, adj.Y, adj.Z); !sideCovered {
				sidesShowing++
			}
		}
	}
	return sidesShowing
}

func (ti *TripleIndex) addToMap(x, y, z int) {
	if _, ok := ti.m[x]; !ok {
		ti.m[x] = map[int]map[int]bool{}
	}
	if _, ok := ti.m[x][y]; !ok {
		ti.m[x][y] = map[int]bool{}
	}
	ti.m[x][y][z] = true
}

func (ti *TripleIndex) Get(x, y, z int) bool {
	if _, ok := ti.m[x]; !ok {
		return false
	}
	if _, ok := ti.m[x][y]; !ok {
		return false
	}
	return ti.m[x][y][z]
}

func (ti *TripleIndex) addMinsAndMaxes(triple Triple) {
	if triple.X < ti.min.X {
		ti.min.X = triple.X
	}
	if triple.Y < ti.min.Y {
		ti.min.Y = triple.Y
	}
	if triple.Z < ti.min.Z {
		ti.min.Z = triple.Z
	}

	if triple.X > ti.max.X {
		ti.max.X = triple.X
	}
	if triple.Y > ti.max.Y {
		ti.max.Y = triple.Y
	}
	if triple.Z > ti.max.Z {
		ti.max.Z = triple.Z
	}
}

func main() {
	file := util.MustReadFile("./internal/dec18/input.txt")
	lines := util.SplitByLine(file)
	triples := util.Map(func(line string) Triple {
		match, ok := util.RegexpMatch("(\\d+),(\\d+),(\\d+)", line)
		if !ok || len(match) != 4 {
			log.Fatal("bad parsing ", match)
		}
		triple := Triple{
			X: util.MustAtoi(match[1]),
			Y: util.MustAtoi(match[2]),
			Z: util.MustAtoi(match[3]),
		}
		if line != triple.String() {
			log.Fatal("bad parsing ", line, " ", triple.String())
		}
		return triple
	}, lines)

	ti := NewTripleIndex(triples)

	// Part 1
	fmt.Println("==== PART ONE ====")
	fmt.Printf("There are %d sizes showing.\n", ti.GetAllSidesShowing())

	// Part 2

	fmt.Println("==== PART TWO ====")
	showing := ti.GetExteriorSidesShowing()
	if showing <= 2004 {
		fmt.Printf("The answer of %d is TOO LOW.\n", showing)
	} else {
		fmt.Printf("There are %d sizes showing.\n", showing)

	}

}
