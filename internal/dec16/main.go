package main

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"github.com/ernestosuarez/itertools"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"strings"
)

const Test = false

func main() {
	filename := getFilename(Test)
	file := util.MustReadFile(filename)
	lines := util.SplitByLine(file)

	nodes := parseNodes(lines)
	allNodes := util.Reduce(func(m map[string]Node, node Node) map[string]Node {
		return util.Assoc(m, node.Valve, node)
	}, nodes, map[string]Node{})

	shortestPaths := ShortestPaths(allNodes)

	positiveFlowNodes := util.Filter(func(node Node) bool {
		return node.Flow > 0
	}, util.Values(allNodes))

	//partOne(allNodes, positiveFlowNodes, shortestPaths)
	partTwo(allNodes, shortestPaths, positiveFlowNodes)
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
	Flow      int
	Neighbors []string
}

func NewNode(Valve string, FlowRate int, Neighbors []string) Node {
	return Node{
		Valve:     Valve,
		Flow:      FlowRate,
		Neighbors: Neighbors,
	}
}

func (n Node) String() string {
	return fmt.Sprintf("Valve %s has flow rate=%d; tunnel lead to valve %s", n.Valve, n.Flow, strings.Join(n.Neighbors, ", "))
}

func parseNodes(lines []string) []Node {
	return util.Map(func(line string) Node {
		match, ok := util.RegexpMatch("Valve ([A-Z]{2}) has flow rate=(\\d+); tunnels? leads? to valves? (.*)", line)
		if !ok {
			log.Fatal("match ", line)
		}
		valve := match[1]
		flowRate := util.MustAtoi(match[2])
		nodesStrs := strings.Split(match[3], ", ")
		node := NewNode(valve, flowRate, nodesStrs)
		normalizedLine := strings.Replace(strings.Replace(strings.Replace(line, "tunnels", "tunnel", -1), "leads", "lead", -1), "valves", "valve", -1)
		if node.String() != normalizedLine {
			log.Fatal("bad parsing. expected \"", normalizedLine, "\", got \"", node.String(), "\"")
		}
		return node
	}, lines)
}

func partOne(allNodes map[string]Node, positiveFlowNodes []Node, shortestPaths map[string]map[string]int) {
	fmt.Println("========== PART ONE ==========")
	max := MaxPressure(allNodes, positiveFlowNodes, shortestPaths, "AA", 30, 0, 0)
	fmt.Println("max :", max)
	fmt.Println("")
}

func MaxPressure(nodes map[string]Node, nodesRemaining []Node, paths map[string]map[string]int, currentPosition string, minutesRemaining int, currentFlow int, currentPressure int) int {
	if minutesRemaining == 0 {
		return currentPressure
	}

	availablePaths := util.Filter(func(nextNode Node) bool {
		enoughMinutesRemain := minutesRemaining-paths[currentPosition][nextNode.Valve] >= 0
		return enoughMinutesRemain
	}, nodesRemaining)

	possibleMaxes := util.Map(func(nextNode Node) int {
		nextPosition := nextNode.Valve
		nextNodesRemaining := util.Filter(func(node Node) bool {
			return node.Valve != nextPosition
		}, nodesRemaining)
		moveCost := paths[currentPosition][nextPosition]
		openAndMoveCost := 1 + moveCost
		nextPressure := currentPressure + currentFlow*openAndMoveCost
		nextFlow := currentFlow + nodes[nextPosition].Flow
		nextTime := minutesRemaining - openAndMoveCost
		return MaxPressure(nodes, nextNodesRemaining, paths, nextPosition, nextTime, nextFlow, nextPressure)
	}, availablePaths)
	return util.Max(util.Max(possibleMaxes...), currentPressure+currentFlow*minutesRemaining)
}

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

func partTwo(allNodes map[string]Node, shortestPaths map[string]map[string]int, positiveFlowNodes []Node) {
	fmt.Println("========== PART TWO ==========")
	allTunnels := util.Map(func(node Node) string {
		return node.Valve
	}, positiveFlowNodes)

	toNodes := func(xs []string) []Node {
		return util.Map(func(s string) Node {
			return allNodes[s]
		}, xs)
	}

	possibleMaxes := util.Map(func(myShare int) int {
		possibleMaxesForShare := util.Map(func(myTunnels []string) int {
			elephantTunnels := util.Filter(func(s string) bool {
				return !util.Contains(myTunnels, s)
			}, allTunnels)
			myMax := MaxPressure(allNodes, toNodes(myTunnels), shortestPaths, "AA", 26, 0, 0)
			elephantMax := MaxPressure(allNodes, toNodes(elephantTunnels), shortestPaths, "AA", 26, 0, 0)
			return myMax + elephantMax
		}, util.Expand(itertools.CombinationsStr(allTunnels, myShare)))
		return util.Max(possibleMaxesForShare...)
	}, util.NewRange(1, len(positiveFlowNodes)/2))

	fmt.Println("max :", util.Max(possibleMaxes...))
	fmt.Println("")
}
