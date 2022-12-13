package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func look(direction []int, treeHeight int, invert bool) (bool, int) {
	d := make([]int, len(direction))
	copy(d, direction) // can't fuck with the original
	if invert {
		for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
			d[i], d[j] = d[j], d[i]
		}
	}
	tallerThan := 1
	for _, v := range d {
		if treeHeight <= v {
			return false, tallerThan
		}
		tallerThan++
	}
	return true, tallerThan - 1
}

func treeInfo(i int, j int, grid [][]int) (bool, int) {
	h := grid[i][j]
	col := make([]int, len(grid[0]))
	for c := range grid {
		col[c] = grid[c][j]
	}
	l, ls := look(grid[i][:j], h, true)    // left
	r, rs := look(grid[i][j+1:], h, false) // right
	t, ts := look(col[:i], h, true)        // top
	b, bs := look(col[i+1:], h, false)     // bottom
	return l || r || t || b, ls * rs * ts * bs
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, treeHeight := range line {
			grid[i][j], _ = strconv.Atoi(string(treeHeight))
		}
	}
	partOne := (len(grid)*2 + len(grid[0])*2) - 4
	partTwo := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			visible, scenicScore := treeInfo(i, j, grid)
			if visible {
				partOne++
			}
			if scenicScore > partTwo {
				partTwo = scenicScore
			}
		}
	}
	fmt.Println("partOne: ", partOne)
	fmt.Println("partTwo: ", partTwo)
}
