package main

import (
	_ "embed"
	"fmt"
	"log"
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
	seeds, maps := parseAlmanac(input)
	locations := followMappings(seeds, maps)
	return slices.Min(locations)
}

type mapping struct {
	source   int
	dest     int
	rangeLen int
}

func parseAlmanac(input string) ([]int, [][]mapping) {
	seeds, input := parseSeeds(input)
	maps := parseMaps(input)
	return seeds, maps
}

func parseSeeds(input string) ([]int, string) {
	before, after, found := strings.Cut(input, "\n\n")
	if !found {
		log.Fatal("invalid input")
	}

	str, found := strings.CutPrefix(before, "seeds: ")
	if !found {
		log.Fatal("invalid seeds")
	}

	seeds := parseInts(str)
	return seeds, after
}

func parseInts(input string) []int {
	strs := strings.Fields(input)
	ints := make([]int, len(strs))

	for i, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = n
	}

	return ints
}

func parseMaps(input string) [][]mapping {
	maps := [][]mapping{}
	for input != "" {
		var m []mapping
		m, input = parseMap(input)
		maps = append(maps, m)
	}
	return maps
}

func parseMap(input string) ([]mapping, string) {
	_, input, found := strings.Cut(input, "map:\n")
	if !found {
		log.Fatal("invalid input")
	}

	mappings := []mapping{}
	for {
		var line string
		line, input, _ = strings.Cut(input, "\n")
		if line == "" {
			break
		}
		ints := parseInts(line)
		mappings = append(mappings, mapping{
			dest:     ints[0],
			source:   ints[1],
			rangeLen: ints[2],
		})
	}

	return mappings, input
}

func followMappings(seeds []int, maps [][]mapping) []int {
	locations := make([]int, len(seeds))

	for i, seed := range seeds {
		curr := seed
		for _, mappings := range maps {
			curr = translateMappings(mappings, curr)
		}
		locations[i] = curr
	}

	return locations
}

func translateMappings(mappings []mapping, n int) int {
	for _, mapping := range mappings {
		srcStart := mapping.source
		srcEnd := mapping.source + mapping.rangeLen

		if n >= srcStart && n < srcEnd {
			delta := n - srcStart
			return mapping.dest + delta
		}
	}

	return n
}
