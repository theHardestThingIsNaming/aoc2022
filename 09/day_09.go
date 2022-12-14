package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func addSign(x int) int {
	if x == 0 {
		return 0
	} else if x < 0 {
		return -1
	}
	return 1
}

func updateTail(head Point, tail Point) Point {
	dx := head.x - tail.x
	dy := head.y - tail.y
	if abs(dx) > 1 || abs(dy) > 1 {
		return Point{tail.x + addSign(dx), tail.y + addSign(dy)}
	}
	return tail
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	head, tail := Point{x: 0, y: 0}, Point{x: 0, y: 0}
	rope := make([]Point, 10)
	partOne := map[Point]bool{head: true}
	partTwo := map[Point]bool{{x: 0, y: 0}: true}
	for _, move := range lines {
		i := strings.Split(move, " ")
		direction := i[0]
		steps, _ := strconv.Atoi(i[1])
		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				head.x++
				rope[0].x++
			case "L":
				head.x--
				rope[0].x--
			case "D":
				head.y--
				rope[0].y--
			case "U":
				head.y++
				rope[0].y++
			}
			//update part one short rope
			tail = updateTail(head, tail)
			//update part two long rope
			for i := 1; i < len(rope); i++ {
				rope[i] = updateTail(rope[i-1], rope[i])
			}
			partOne[tail] = true
			partTwo[rope[9]] = true // tail
		}
	}
	fmt.Println("partOne: ", len(partOne))
	fmt.Println("partTwo: ", len(partTwo))
}
