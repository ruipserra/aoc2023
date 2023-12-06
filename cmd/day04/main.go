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
}

func part1(input string) int {
	score := 0

	eachLine(input, func(line string) {
		card := parseCard(line)
		score += calculateScore(card)
	})

	return score
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
	found := 0
	for _, n := range card.want {
		if slices.Contains(card.have, n) {
			found++
		}
	}
	if found == 0 {
		return 0
	}
	//nolint:gomnd // The score increases in powers of 2
	return int(math.Pow(2, float64(found-1)))
}
