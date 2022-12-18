package main

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"sort"
	"strings"
)

const Test = false
const Debug = false

func main() {
	filename := getFilename(Test)
	file := util.MustReadFile(filename)
	lines := util.SplitByLine(file)

	nodes := parse(lines)
	keys := util.Keys(nodes)
	sort.Slice(keys, func(i, j int) bool {
		return nodes[keys[i]].FlowRate > nodes[keys[j]].FlowRate
	})

	shortestPaths := ShortestPaths(nodes)
	a := util.SelectKeys(nodes, keys[:8])

	partOne(a, shortestPaths, Debug)
	partTwo(Debug)
}

func getFilename(test bool) string {
	base := "./internal/dec16"
	if test {
		return base + "/test.txt"
	}
	return base + "/input.txt"
}

type Node struct {
	Valve     string
	FlowRate  int
	Neighbors []string
}

func NewNode(Valve string, FlowRate int, Neighbors []string) Node {
	return Node{
		Valve:     Valve,
		FlowRate:  FlowRate,
		Neighbors: Neighbors,
	}
}

func (n Node) String() string {
	return fmt.Sprintf("Valve %s has flow rate=%d; tunnel lead to valve %s", n.Valve, n.FlowRate, strings.Join(n.Neighbors, ", "))
}

func parse(lines []string) map[string]Node {
	var m = map[string]Node{}
	for _, line := range lines {
		normalizedLine := strings.Replace(strings.Replace(strings.Replace(line, "tunnels", "tunnel", -1), "leads", "lead", -1), "valves", "valve", -1)
		match, ok := util.RegexpMatch("Valve ([A-Z]{2}) has flow rate=(\\d+); tunnels? leads? to valves? (.*)", line)
		if !ok {
			log.Fatal("match ", line)
		}
		valve := match[1]
		flowRate := util.MustAtoi(match[2])
		nodesStrs := strings.Split(match[3], ", ")
		node := NewNode(valve, flowRate, nodesStrs)
		if node.String() != normalizedLine {
			log.Fatal("bad parsing. expected \"", normalizedLine, "\", got \"", node.String(), "\"")
		}
		m[valve] = node
	}
	return m
}

func partOne(nodes map[string]Node, shortestPaths map[string]map[string]int, debug bool) {
	fmt.Println("========== PART ONE ==========")
	fmt.Println("========== PART END ==========")

}

// possible : 2582 1720

func ShortestPaths(nodes map[string]Node) map[string]map[string]int {
	graph := dijkstra.NewGraph()

	for id := range nodes {
		graph.AddMappedVertex(id)
	}

	for source, node := range nodes {
		for _, neighbor := range node.Neighbors {
			if err := graph.AddMappedArc(source, neighbor, 1); err != nil {
				log.Fatalf("bad mapped arc from source %q to dest %q", source, neighbor)
			}
		}
	}

	shortestPaths := map[string]map[string]int{}
	for source := range nodes {
		shortestPaths[source] = map[string]int{}
		i, err := graph.GetMapping(source)
		if err != nil {
			log.Fatalf("bad mapping %q", source)
		}
		for dest := range nodes {
			j, err := graph.GetMapping(dest)
			if err != nil {
				log.Fatalf("bad mapping %q", dest)
			}
			if i == j {
				continue
			}
			best, err := graph.Shortest(i, j)
			if err != nil {
				log.Fatalf("bad calculation of shortest %q to %q", source, dest)
			}
			shortestPaths[source][dest] = int(best.Distance)
		}
	}

	return shortestPaths
}

func partTwo(debug bool) {
	fmt.Println("========== PART TWO ==========")

	fmt.Println("==========   END   ==========")
	fmt.Println("")
}

/// 	iterator := itertools.PermutationsStr(toVisit, len(toVisit))
//
//	var maxPressure int // (answer is max)
//	for permutation := range iterator {
//		source := "AA"
//		flow := 0
//		pressure := 0
//		minute := 0
//		for _, destination := range permutation {
//			moves := shortestPaths[source][destination]
//			open := 1
//			untilNextFlow := moves + open
//			if debug {
//				fmt.Print("(minutes = ", untilNextFlow, ", flow = ", flow, ") ")
//			}
//			pressureAdded := untilNextFlow * flow
//			pressure += pressureAdded
//
//			flow += nodes[destination].FlowRate
//			source = destination
//			minute += untilNextFlow
//			if minute >= 30 {
//				break
//			}
//		}
//		if minute < 30 {
//			untilEnd := 30 - minute
//			if debug {
//				fmt.Print("(minutes = ", untilEnd, ", flow = ", flow, ") ")
//			}
//			pressure = pressure + untilEnd*flow
//		}
//		fmt.Println(permutation, pressure)
//		if pressure > maxPressure {
//			maxPressure = pressure
//		}
//	}
//
//	fmt.Println("max pressure :", maxPressure)
