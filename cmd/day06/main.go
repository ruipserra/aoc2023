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
		d := calculateDistance(t, time)
		if d > distance {
			ways++
		}
	}
	return ways
}

func calculateDistance(t int, time int) int {
	speed := t
	timeLeft := time - t
	return speed * timeLeft
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

func part2(input string) int {
	time := parseTime(input)
	distance := parseDistance(input)

	start := 0
	for t := 0; t < time; t++ {
		d := calculateDistance(t, time)
		if d > distance {
			start = t
			break
		}
	}

	if start == 0 {
		return 0
	}

	end := start
	for t := time; t > end; t-- {
		d := calculateDistance(t, time)
		if d > distance {
			end = t
			break
		}
	}

	return end - start + 1
}

func parseTime(input string) int {
	input, _, _ = strings.Cut(input, "\n")
	_, input, _ = strings.Cut(input, ":")
	return parseSingleInt(input)
}

func parseDistance(input string) int {
	_, input, _ = strings.Cut(input, "\n")
	_, input, _ = strings.Cut(input, ":")
	input, _, _ = strings.Cut(input, "\n")
	return parseSingleInt(input)
}

func parseSingleInt(input string) int {
	input = strings.ReplaceAll(input, " ", "")
	i, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
