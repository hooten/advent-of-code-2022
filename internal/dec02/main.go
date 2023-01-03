package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/math"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: Score is %d.\n", strategyGuidePoints(true))
	fmt.Printf("Part 2: Score is %d.\n", strategyGuidePoints(false))
}

func strategyGuidePoints(xyzAsHand bool) int64 {
	const filename = "./internal/dec02/input.txt"
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := util.Filter(func(line string) bool {
		return line != ""
	}, strings.Split(string(bytes), "\n"))

	games := util.Map(func(line string) Game {
		return ParseGame(line, xyzAsHand)
	}, lines)

	pointsPerGame := util.Map(func(game Game) int64 {
		return game.Play()
	}, games)

	return math.Sum(pointsPerGame...)
}

type Hand int64

const (
	Rock Hand = iota
	Paper
	Scissors
)

var HandPointsLegend = map[Hand]int64{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var HandLegend = map[string]Hand{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func ParseHand(s string) Hand {
	hand, ok := HandLegend[s]
	if !ok {
		log.Fatalf("could not parse hand %q", s)
	}
	return hand
}

type Game struct {
	player1 Hand
	player2 Hand
}

func ParseGame(s string, xyzAsHand bool) Game {
	hands := strings.Split(s, " ")
	if len(hands) != 2 {
		log.Fatalf("expected 2 hands, got %q", s)
	}
	player1Hand := ParseHand(hands[0])
	if xyzAsHand {
		return Game{
			player1: player1Hand,
			player2: ParseHand(hands[1]),
		}
	}
	outcome := ParseOutcome(hands[1])
	hand := DetermineHand(player1Hand, outcome)
	return Game{
		player1: player1Hand,
		player2: hand,
	}
}

func DetermineHand(hand Hand, outcome Outcome) Hand {
	switch outcome {
	case Win:
		return (hand + 1) % 3
	case Draw:
		return hand
	case Loss:
		return (hand + 2) % 3
	}
	log.Fatalf("could not determine hand from hand %d and outcome %d", hand, outcome)
	return -1
}

var OutcomeLegend = map[string]Outcome{
	"X": Loss,
	"Y": Draw,
	"Z": Win,
}

func ParseOutcome(s string) Outcome {
	outcome, ok := OutcomeLegend[s]
	if !ok {
		log.Fatalf("could not parse outcome %q", s)
	}
	return outcome
}

func (g Game) Play() int64 {
	outcome := g.play()
	handPoints, ok := HandPointsLegend[g.player2]
	if !ok {
		log.Fatalf("could not determine hand points for %q", g.player2)
	}
	return int64(outcome) + handPoints
}

type Outcome int64

const (
	Win  Outcome = 6
	Draw         = 3
	Loss         = 0
)

var plays = map[int64]Outcome{
	-2: Loss,
	-1: Win,
	0:  Draw,
	1:  Loss,
	2:  Win,
}

func (g Game) play() Outcome {
	play := g.player1 - g.player2
	outcome, ok := plays[int64(play)]
	if !ok {
		log.Fatalf("could not determine outcome from play %d", play)
	}
	return outcome
}
