package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var elfHandLegend = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var handPointsLegend = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var myHandPointsLegend = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

var gamePointsLegend = map[string]int{
	"A X": 3, // Rock Rock = Draw
	"A Y": 6, // Rock Paper = Win
	"A Z": 0, // Rock Scissors = Loss
	"B X": 0, // Paper Rock = Loss
	"B Y": 3, // Paper Paper = Draw
	"B Z": 6, // Paper Scissors = Win
	"C X": 6, // Scissors Rock = Win
	"C Y": 0, // Scissors Paper = Loss
	"C Z": 3, // Scissors Scissors = Draw
}

var resultLegend = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

func main() {
	// partOne()
	partTwo()
}

func partTwo() {
	file, err := os.Open("./internal/dec02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var myPoints int
	for scanner.Scan() {
		game := scanner.Text()
		plays := strings.Split(game, " ")
		if len(plays) != 2 {
			log.Fatalf("unexpected length %d", len(plays))
		}
		elfHand := elfHandLegend[plays[0]]
		myPlay := plays[1]
		gamePoints := resultLegend[myPlay]
		myHand := getMyHand(elfHand, gamePoints)
		handPoints := myHandPointsLegend[myHand]
		myPoints += handPoints + gamePoints
		fmt.Printf("{elf: %s, me: %s, game: %d + %d, total: %d}\n", elfHand, myHand, handPoints, gamePoints, myPoints)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("my score is %d", myPoints)
}

func getMyHand(elfHand string, gamePoints int) string {
	switch elfHand {
	case "Rock":
		switch gamePoints {
		case 0:
			return "Scissors"
		case 3:
			return "Rock"
		case 6:
			return "Paper"
		default:
			log.Fatalf("bad elf hand %q", elfHand)
		}
	case "Paper":
		switch gamePoints {
		case 0:
			return "Rock"
		case 3:
			return "Paper"
		case 6:
			return "Scissors"
		default:
			log.Fatalf("bad elf hand %q", elfHand)
		}
	case "Scissors":
		switch gamePoints {
		case 0:
			return "Paper"
		case 3:
			return "Scissors"
		case 6:
			return "Rock"
		default:
			log.Fatalf("bad elf hand %q", elfHand)
		}
	}
	log.Fatalf("bad elf hand %s", elfHand)
	return ""
}

func partOne() {
	file, err := os.Open("./internal/dec02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var myPoints int
	for scanner.Scan() {
		game := scanner.Text()
		plays := strings.Split(game, " ")
		if len(plays) != 2 {
			log.Fatalf("unexpected length %d", len(plays))
		}
		myPlay := plays[1]
		handPoints := handPointsLegend[myPlay]
		gamePoints := gamePointsLegend[game]
		myPoints += gamePoints + handPoints
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("my score is %d", myPoints)
}
