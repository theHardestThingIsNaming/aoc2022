package main

import (
	"fmt"
	"os"
	"strings"
)

// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors
// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 lost, 3 draw, and 6 won
//X lose, Y draw, and Z win

func main() {
	pointsForResult := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 7, "C Y": 2, "C Z": 6,
	}
	pointsForComplexResult := map[string]int{
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7,
	}
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")
	total := 0
	partTwoTotal := 0
	for _, s := range split {
		total += pointsForResult[s]
		partTwoTotal += pointsForComplexResult[s]
	}
	fmt.Println(total)
	fmt.Println(partTwoTotal)
}
