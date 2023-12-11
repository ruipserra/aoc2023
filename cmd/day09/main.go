package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		ints := parseInts(line)
		prediction := predictNext(ints, false)
		sum += prediction
	}
	return sum
}

func parseInts(line string) []int {
	parts := strings.Fields(line)
	ints := make([]int, len(parts))
	for i, s := range parts {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("parseInts: %s", err)
		}
		ints[i] = n
	}
	return ints
}

func predictNext(values []int, allZeros bool) int {
	if allZeros {
		return 0
	}

	allZeros = true
	deltas := make([]int, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		deltas[i] = values[i+1] - values[i]
		if deltas[i] != 0 {
			allZeros = false
		}
	}
	return values[len(values)-1] + predictNext(deltas, allZeros)
}
