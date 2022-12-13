package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func main() {
	input, _ := os.ReadFile("input.txt")
	partOneCount := 0
	partTwoCount := 0
	assignmentsPairs := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, ap := range assignmentsPairs {
		pair := strings.Split(ap, ",")
		A := strings.Split(pair[0], "-")
		B := strings.Split(pair[1], "-")
		if pInt(A[0]) <= pInt(B[0]) && pInt(B[1]) <= pInt(A[1]) || pInt(B[0]) <= pInt(A[0]) && pInt(A[1]) <= pInt(B[1]) {
			partOneCount += 1
		}
		if (pInt(B[0]) <= pInt(A[0]) && pInt(A[0]) <= pInt(B[1]) || pInt(B[0]) <= pInt(A[1]) && pInt(A[1]) <= pInt(B[1])) ||
			(pInt(A[0]) <= pInt(B[0]) && pInt(B[0]) <= pInt(A[1]) || pInt(A[0]) <= pInt(B[1]) && pInt(B[1]) <= pInt(A[1])) {
			partTwoCount += 1
		}
	}
	fmt.Println("Part 1:", partOneCount)
	fmt.Println("Part 2:", partTwoCount)
}
