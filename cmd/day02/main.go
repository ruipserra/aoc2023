package main

import (
	"bufio"
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

func eachLine(input string, fn func(line string)) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fn(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanning lines: %s", err)
	}
}

type game struct {
	subsets []subset
	id      int
}

func (g game) isValid() bool {
	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	for _, subset := range g.subsets {
		if subset.red > maxRed ||
			subset.blue > maxBlue ||
			subset.green > maxGreen {
			return false
		}
	}
	return true
}

func (g game) power() int {
	var maxRed, maxGreen, maxBlue int

	for _, subset := range g.subsets {
		maxRed = max(maxRed, subset.red)
		maxBlue = max(maxBlue, subset.blue)
		maxGreen = max(maxGreen, subset.green)
	}

	return maxRed * maxGreen * maxBlue
}

type subset struct {
	red   int
	green int
	blue  int
}

func parseGame(s string) game {
	parser := parser{
		input: s,
		game:  &game{},
	}
	return parser.parse()
}

type parser struct {
	game  *game
	input string
}

func (p *parser) parse() game {
	p.parseGameID()
	p.parseSubsets()
	return *p.game
}

func (p *parser) parseGameID() {
	before, after, _ := strings.Cut(p.input, ":")
	idStr, found := strings.CutPrefix(before, "Game ")
	if !found {
		log.Fatalf("failed to parse game id from '%s'", idStr)
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Fatalf("parsing game id: %v", err)
	}

	p.game.id = int(id)
	p.input = after
}

func (p *parser) parseSubsets() {
	for p.input != "" {
		p.parseSubset()
	}
}

func (p *parser) parseSubset() {
	//nolint:govet // yes I know this is shadowing the type
	subset := subset{}

	before, after, _ := strings.Cut(p.input, ";")
	p.input = after

	parts := strings.Split(before, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		number, color, _ := strings.Cut(part, " ")
		count, err := strconv.ParseInt(number, 10, 32)
		if err != nil {
			log.Fatalf("parsing cube count: %s", err)
		}

		switch color {
		case "red":
			subset.red = int(count)
		case "green":
			subset.green = int(count)
		case "blue":
			subset.blue = int(count)
		}
	}

	p.game.subsets = append(p.game.subsets, subset)
}

func part1(input string) int {
	sum := 0

	eachLine(input, func(line string) {
		game := parseGame(line)
		if game.isValid() {
			sum += game.id
		}
	})

	return sum
}

func part2(input string) int {
	power := 0

	eachLine(input, func(line string) {
		game := parseGame(line)
		power += game.power()
	})

	return power
}
