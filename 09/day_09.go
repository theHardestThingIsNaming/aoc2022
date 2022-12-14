package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
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
	partTwo := map[Point]bool{head: true}
	for _, move := range lines {
		m := strings.Split(move, " ")
		steps, _ := strconv.Atoi(m[1])
		for i := 0; i < steps; i++ {
			switch m[0] {
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
			tail = updateTail(head, tail)
			for i := 1; i < len(rope); i++ {
				rope[i] = updateTail(rope[i-1], rope[i])
			}
			partOne[tail] = true
			partTwo[rope[9]] = true
		}
	}
	fmt.Println("partOne: ", len(partOne), "\npartTwo: ", len(partTwo))
}
