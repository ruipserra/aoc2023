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
	times := parseTimes(input)
	distances := parseDistances(input)
	if len(times) != len(distances) {
		log.Fatal("len(times) != len(distances)")
	}

	result := 1
	for i := range times {
		result *= howManyWaysToBeatRecord(times[i], distances[i])
	}
	return result
}

func howManyWaysToBeatRecord(time int, distance int) int {
	ways := 0
	for t := 0; t < time; t++ {
		speed := t
		timeLeft := time - t
		d := speed * timeLeft
		if d > distance {
			ways++
		}
	}
	return ways
}

func parseTimes(input string) []int {
	input, _ = strings.CutPrefix(input, "Time:")
	input, _, _ = strings.Cut(input, "\n")
	return parseInts(input)
}

func parseDistances(input string) []int {
	_, input, _ = strings.Cut(input, "\n")
	input, _ = strings.CutPrefix(input, "Distance:")
	return parseInts(input)
}

func parseInts(input string) []int {
	ints := []int{}
	for _, str := range strings.Fields(input) {
		i, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}
	return ints
}
