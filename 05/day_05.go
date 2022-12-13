package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func partOne(moves []string, stacks map[int][]rune, stackKeys []int) {
	for _, move := range moves {
		var amount int
		var from int
		var to int
		fmt.Sscanf(move, "move %d from %d to %d", &amount, &from, &to)
		source := stacks[from]
		destination := stacks[to]
		for i := 0; i < amount; i++ {
			n := len(source) - 1 // Top element
			destination = append(destination, source[n])
			source = source[:n] // Pop
		}
		stacks[from] = source
		stacks[to] = destination
	}
	partOne := ""
	for _, k := range stackKeys {
		partOne += string(stacks[k][len(stacks[k])-1])
	}
	fmt.Println(partOne)
}

func partTwo(moves []string, stacks map[int][]rune, stackKeys []int) {
	for _, move := range moves {
		var amount int
		var from int
		var to int
		fmt.Sscanf(move, "move %d from %d to %d", &amount, &from, &to)
		source := stacks[from]
		destination := stacks[to]
		n := len(source) - amount
		for _, r := range source[n:] {
			destination = append(destination, r)
		}
		source = source[:n]
		stacks[from] = source
		stacks[to] = destination
	}
	partTwo := ""
	for _, k := range stackKeys {
		partTwo += string(stacks[k][len(stacks[k])-1])
	}
	fmt.Println(partTwo)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(string(input), "\n\n")
	layout := strings.Split(split[0], "\n")
	stacks := make(map[int][]rune)
	stackKeys := []int{}
	for i := len(layout) - 1; i >= 0; i-- {
		if i == len(layout)-1 {
			for _, k := range strings.Split(strings.TrimSpace(layout[i]), "   ") {
				key, _ := strconv.Atoi(k)
				stacks[key] = []rune{}
				stackKeys = append(stackKeys, key)
			}
		} else {
			for i, c := range layout[i] {
				if !strings.ContainsAny(string(c), " []") {
					key := int(math.Ceil(float64(i) / 4))
					stacks[key] = append(stacks[key], c)
				}
			}
		}

	}
	moves := strings.Split(split[1], "\n")
	// partOne(moves, stacks, stackKeys)
	partTwo(moves, stacks, stackKeys)
}
