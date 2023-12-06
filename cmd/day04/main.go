package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"math"
	"slices"
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
	score := 0

	eachLine(input, func(line string) {
		card := parseCard(line)
		score += calculateScore(card)
	})

	return score
}

func part2(input string) int {
	cards := []card{}
	eachLine(input, func(line string) {
		cards = append(cards, parseCard(line))
	})

	// Let's say originals and copies are the same thing so we can count
	// them all together.
	copies := make([]int, len(cards))
	for i := range copies {
		copies[i] = 1
	}

	for i, card := range cards {
		matches := countMatches(card)
		for j := 0; j < matches; j++ {
			idx := i + j + 1
			if idx >= len(copies) {
				break
			}
			copies[idx] += copies[i]
		}
	}

	total := 0
	for _, n := range copies {
		total += n
	}
	return total
}

func eachLine(input string, fn func(line string)) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fn(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanning input: %s", err)
	}
}

type card struct {
	want []int
	have []int
}

func parseCard(line string) card {
	colonPos := strings.IndexRune(line, ':')
	pipePos := strings.IndexRune(line, '|')
	if colonPos == -1 || pipePos == -1 {
		log.Fatalf("invalid card: %v", line)
	}
	want := parseInts(line[colonPos+1 : pipePos])
	have := parseInts(line[pipePos+1:])
	return card{want, have}
}

func parseInts(input string) []int {
	ints := []int{}
	for _, s := range strings.Fields(input) {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("parsing int: %s", err)
		}
		ints = append(ints, n)
	}
	return ints
}

func calculateScore(card card) int {
	matches := countMatches(card)
	if matches == 0 {
		return 0
	}
	//nolint:gomnd // The score increases in powers of 2
	return int(math.Pow(2, float64(matches-1)))
}

func countMatches(card card) int {
	matches := 0
	for _, n := range card.want {
		if slices.Contains(card.have, n) {
			matches++
		}
	}
	return matches
}
