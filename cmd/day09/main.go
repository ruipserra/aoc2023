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
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input string) int {
	sum := 0
	eachLine(input, func(ints []int) {
		prediction := predict(ints, dirRight)
		sum += prediction
	})
	return sum
}

func part2(input string) int {
	sum := 0
	eachLine(input, func(ints []int) {
		prediction := predict(ints, dirLeft)
		sum += prediction
	})
	return sum
}

func eachLine(input string, f func([]int)) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		ints := parseInts(line)
		f(ints)
	}
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

type direction byte

const (
	dirLeft direction = iota
	dirRight
)

func predict(values []int, dir direction) int {
	deltas, zeros := calculateDeltas(values)
	if zeros {
		switch dir {
		case dirLeft:
			return values[0]
		case dirRight:
			return values[len(values)-1]
		default:
			panic("invalid prediction direction")
		}
	}

	switch dir {
	case dirLeft:
		return values[0] - predict(deltas, dir)
	case dirRight:
		return values[len(values)-1] + predict(deltas, dir)
	default:
		panic("invalid prediction direction")
	}
}

func calculateDeltas(values []int) ([]int, bool) {
	zeros := true
	deltas := make([]int, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		deltas[i] = values[i+1] - values[i]
		if deltas[i] != 0 {
			zeros = false
		}
	}
	return deltas, zeros
}
