package main

import (
	"fmt"
	"os"
)

func findMarker(bigString string, distinctCount int) {
	for i := range bigString {
		m := make(map[byte]bool)
		if i+distinctCount >= len(bigString) {
			break
		}
		for p := i; p < i+distinctCount; p++ {
			m[bigString[p]] = true
		}
		if len(m) == distinctCount {
			fmt.Println(bigString[i : i+distinctCount])
			fmt.Println("index: ", i+distinctCount)
			break
		}
	}
}

func main() {
	input, _ := os.ReadFile("input.txt")
	bigString := string(input)
	findMarker(bigString, 4)
	findMarker(bigString, 14)
}

