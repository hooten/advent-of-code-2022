package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("./internal/dec01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	caloriesInput := strings.Split(string(bytes), "\n")
	var calorieCounts []int
	total := 0
	for _, calorieStr := range caloriesInput {
		if calorieStr == "" {
			calorieCounts = append(calorieCounts, total)
			total = 0
			continue
		}
		calorieCount, err := strconv.Atoi(calorieStr)
		if err != nil {
			log.Fatal(err)
		}
		total += calorieCount
	}
	sort.Ints(calorieCounts)
	n := len(calorieCounts)
	fmt.Printf("The largest calorie count is %d\n", calorieCounts[n-1])
	fmt.Printf("The sum of the top three elves is %d\n",
		calorieCounts[n-1]+
			calorieCounts[n-2]+
			calorieCounts[n-3],
	)
}
