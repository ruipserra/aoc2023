package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input string) int {
	const start = "AAA"
	const end = "ZZZ"

	directions, nodes := parseMap(input)
	steps := 0
	loc := start
	for loc != end {
		dir := directions[steps%len(directions)]
		loc = nextLocation(nodes, loc, dir)
		steps++
	}
	return steps
}

type nodeMap = map[string][2]string

func parseMap(input string) (string, nodeMap) {
	lines := strings.Split(input, "\n")
	directions := lines[0]
	nodes := make(nodeMap)
	for _, line := range lines[1:] {
		if line != "" {
			node, nextNodes := parseMapNode(line)
			nodes[node] = nextNodes
		}
	}
	return directions, nodes
}

func parseMapNode(input string) (string, [2]string) {
	node, input, _ := strings.Cut(input, " = ")
	input, _ = strings.CutPrefix(input, "(")
	input, _ = strings.CutSuffix(input, ")")
	nextNodes := strings.Split(input, ", ")
	//nolint:gomnd // Each map node has 2 directions
	if len(nextNodes) != 2 {
		log.Fatalf("failed to parse next nodes: %v", input)
	}
	return node, [2]string{nextNodes[0], nextNodes[1]}
}

func nextLocation(nodes nodeMap, curr string, dir byte) string {
	node := nodes[curr]
	switch dir {
	case byte('L'):
		return node[0]
	case byte('R'):
		return node[1]
	default:
		panic(fmt.Sprintf("invalid direction: %v", dir))
	}
}
