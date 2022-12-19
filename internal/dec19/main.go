package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"math"
)

// AllRobots Order matters.
var AllRobots = []string{
	"geode",
	"obsidian",
	"clay",
	"ore",
	"none",
}

type Cost struct {
	Ore      int
	Clay     int
	Obsidian int
}

type Blueprint struct {
	Index             int
	OreRobotCost      Cost
	ClayRobotCost     Cost
	ObsidianRobotCost Cost
	GeodeRobotCost    Cost
	MaxCost           Cost
}

func NewBluePrint(index int, oreRobotCost Cost, clayRobotCost Cost, obsidianRobotCost Cost, geodeRobotCost Cost) Blueprint {
	return Blueprint{
		Index:             index,
		OreRobotCost:      oreRobotCost,
		ClayRobotCost:     clayRobotCost,
		ObsidianRobotCost: obsidianRobotCost,
		GeodeRobotCost:    geodeRobotCost,
		MaxCost: Cost{
			Ore:      util.Max(oreRobotCost.Ore, clayRobotCost.Ore, obsidianRobotCost.Ore),
			Clay:     obsidianRobotCost.Clay,
			Obsidian: geodeRobotCost.Obsidian,
		},
	}
}

func (b Blueprint) String() string {
	return fmt.Sprintf(
		"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		b.Index,
		b.OreRobotCost.Ore,
		b.ClayRobotCost.Ore,
		b.ObsidianRobotCost.Ore,
		b.ObsidianRobotCost.Clay,
		b.GeodeRobotCost.Ore,
		b.GeodeRobotCost.Obsidian,
	)
}

type Pack struct {
	blueprint Blueprint

	Ore      int
	Clay     int
	Obsidian int
	Geode    int

	OreRobot      int
	ClayRobot     int
	ObsidianRobot int
	GeodeRobot    int
}

func (p Pack) Affordable(kind string) bool {
	switch kind {
	case "none":
		return true
	case "ore":
		return p.Ore >= p.blueprint.OreRobotCost.Ore
	case "clay":
		return p.Ore >= p.blueprint.ClayRobotCost.Ore
	case "obsidian":
		cost0 := p.Ore >= p.blueprint.ObsidianRobotCost.Ore
		cost1 := p.Clay >= p.blueprint.ObsidianRobotCost.Clay
		return cost0 && cost1
	case "geode":
		cost0 := p.Ore >= p.blueprint.GeodeRobotCost.Ore
		cost1 := p.Obsidian >= p.blueprint.GeodeRobotCost.Obsidian
		return cost0 && cost1
	default:
		log.Fatal(kind)
	}
	return false
}

func (p Pack) OrderRobot(kind string) Pack {
	switch kind {
	case "none":
		return p
	case "ore":
		p.Ore -= p.blueprint.OreRobotCost.Ore
		return p
	case "clay":
		p.Ore -= p.blueprint.ClayRobotCost.Ore
		return p
	case "obsidian":
		p.Ore -= p.blueprint.ObsidianRobotCost.Ore
		p.Clay -= p.blueprint.ObsidianRobotCost.Clay
		return p
	case "geode":
		p.Ore -= p.blueprint.GeodeRobotCost.Ore
		p.Obsidian -= p.blueprint.GeodeRobotCost.Obsidian
		return p
	default:
		log.Fatal(kind)
	}
	return p
}

func (p Pack) CollectGems(debug bool) Pack {
	p.Ore += p.OreRobot
	p.Clay += p.ClayRobot
	p.Obsidian += p.ObsidianRobot
	p.Geode += p.GeodeRobot
	if debug {
		if p.OreRobot > 0 {
			fmt.Printf("%d ore-collecting robot collects %d ore; you now have %d ore.\n", p.OreRobot, p.OreRobot, p.Ore)
		}
		if p.ClayRobot > 0 {
			fmt.Printf("%d clay-collecting robot collects %d clay; you now have %d clay.\n", p.ClayRobot, p.ClayRobot, p.Clay)
		}
		if p.ObsidianRobot > 0 {
			fmt.Printf("%d obsidian-collecting robot collects %d obsidian; you now have %d obsidian.\n", p.ObsidianRobot, p.ObsidianRobot, p.Obsidian)
		}
		if p.GeodeRobot > 0 {
			fmt.Printf("%d geode-cracking robot collects %d geode; you now have %d geode.\n", p.GeodeRobot, p.GeodeRobot, p.Geode)
		}
	}
	return p
}

func (p Pack) AddRobot(kind string) Pack {
	switch kind {
	case "none":
		return p
	case "ore":
		p.OreRobot++
		return p
	case "clay":
		p.ClayRobot++
		return p
	case "obsidian":
		p.ObsidianRobot++
		return p
	case "geode":
		p.GeodeRobot++
		return p
	default:
		log.Fatal(kind)
	}
	return p
}

var packMax = 0

func (p Pack) MaxGeodes(minutesRemaining int) int {
	if minutesRemaining == 0 {
		return 0
	}
	if packMax >= p.Geode+p.OptimizedBestCaseGeodes(p.GeodeRobot, p.GeodeRobot+minutesRemaining-1) {
		return 0
	}
	// If there are enough ore and obsidian robots to make enough resources to afford geode robots every turn, do that.
	if p.OreRobot >= p.blueprint.GeodeRobotCost.Ore && p.ObsidianRobot >= p.blueprint.GeodeRobotCost.Obsidian {
		return p.OptimizedBestCaseGeodes(p.GeodeRobot, p.GeodeRobot+minutesRemaining-1)
	}

	affordableRobots := util.Filter(func(kind string) bool {
		return p.Affordable(kind) && p.Optimal(kind)
	}, AllRobots)
	possibleMaxes := util.Map(func(kind string) int {
		newP := p.OrderRobot(kind).CollectGems(false).AddRobot(kind)
		max := p.GeodeRobot + newP.MaxGeodes(minutesRemaining-1)
		return max
	}, affordableRobots)
	max := util.Reduce(func(max int, elem int) int {
		if elem > max {
			return elem
		}
		return max
	}, possibleMaxes, 0)
	packMax = int(math.Max(float64(max), float64(packMax)))
	return max
}

func (p Pack) OptimizedBestCaseGeodes(currGeodeRobots, endGeodeRobots int) int {
	return endGeodeRobots*(endGeodeRobots+1)/2 - ((currGeodeRobots - 1) * currGeodeRobots / 2)
}

func (p Pack) Optimal(kind string) bool {
	switch kind {
	case "none":
		return p.OreRobot < p.blueprint.MaxCost.Ore
	case "ore":
		return p.OreRobot < p.blueprint.MaxCost.Ore
	case "clay":
		return p.ClayRobot < p.blueprint.ObsidianRobotCost.Clay
	case "obsidian":
		return p.ObsidianRobot < p.blueprint.GeodeRobotCost.Obsidian
	case "geode":
		return true
	default:
		log.Fatal(kind)
	}
	return true

}

func NewPack(blueprint Blueprint) Pack {
	return Pack{
		blueprint: blueprint,
		OreRobot:  1,
	}
}

func main() {
	part1 := false
	file := util.MustReadFile("./internal/dec19/input.txt")
	lines := util.SplitByLine(file)
	blueprints := util.MapWithIndex(func(line string, i int) Blueprint {
		re := "Blueprint (\\d+): Each ore robot costs (\\d+) ore. Each clay robot costs (\\d+) ore. Each obsidian robot costs (\\d+) ore and (\\d+) clay. Each geode robot costs (\\d+) ore and (\\d+) obsidian."
		match, ok := util.RegexpMatch(re, line)
		if !ok {
			log.Fatalf("error match line %d: %q\n", i, line)
		}
		if len(match) != 8 {
			log.Fatalf("bad match line %d: %q\n", i, line)
		}
		parsed := util.Map(func(s string) int {
			return util.MustAtoi(s)
		}, match[1:])
		index := parsed[1-1]
		oreRobotCost := Cost{
			Ore: parsed[2-1],
		}
		clayRobotCost := Cost{
			Ore: parsed[3-1],
		}
		obsidianRobotCost := Cost{
			Ore:  parsed[4-1],
			Clay: parsed[5-1],
		}
		geodeRobotCost := Cost{
			Ore:      parsed[6-1],
			Obsidian: parsed[7-1],
		}

		b := NewBluePrint(index, oreRobotCost, clayRobotCost, obsidianRobotCost, geodeRobotCost)

		if b.String() != line {
			log.Fatal("bad parsing")
		}
		return b
	}, lines)

	fmt.Println("-- blueprints --")
	for _, b := range blueprints {
		fmt.Println(b)
	}
	fmt.Println(" ")

	// Part 1
	if part1 {
		fmt.Println("==== PART ONE ====")
		qualityLevelSum := 0
		for _, b := range blueprints {
			pack := NewPack(b)
			maxGeodes := pack.MaxGeodes(24)
			packMax = 0
			fmt.Printf("Geodes with blueprint %d: %d\n", b.Index, maxGeodes)
			qualityLevel := b.Index * maxGeodes
			fmt.Printf("Quality ID: %d\n", qualityLevel)
			qualityLevelSum += qualityLevel
		}
		fmt.Println("Quality Level Sum:", qualityLevelSum)
	}

	// Part 2
	fmt.Println("==== PART TWO ====")
	maxGeodes := util.Map(func(blueprint Blueprint) int {
		packMax = 0
		maxGeodes := NewPack(blueprint).MaxGeodes(32)
		fmt.Println("max:", maxGeodes)
		return maxGeodes
	}, blueprints[:3])
	result := util.Reduce(func(result int, max int) int {
		return result * max
	}, maxGeodes, 1)

	fmt.Println("Largest Blueprints Multiplied:", result)

}
