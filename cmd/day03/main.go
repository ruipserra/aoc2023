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

	lines := strings.Split(input, "\n")
	for row, line := range lines {
		width := len(line)

		col := 0
		for col < width {
			if isDigit(line[col]) {
				numberWidth := countDigits(line[col:])
				n, err := strconv.Atoi(line[col : col+numberWidth])

				if err != nil {
					log.Fatalf("parsing number: %s", err)
				}

				if isSymbolAdjacent(lines, row, col, numberWidth) {
					sum += n
				}

				col += numberWidth
			} else {
				col++
			}
		}
	}

	return sum
}

func part2(input string) int {
	const gearCount = 2

	sum := 0

	lines := strings.Split(input, "\n")
	for row, line := range lines {
		for col := 0; col < len(line); col++ {
			if line[col] != '*' {
				continue
			}

			adjacentNumbers := findAdjacentNumbers(lines, row, col)
			if len(adjacentNumbers) == gearCount {
				sum += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}

	return sum
}

func isSymbol(b byte) bool {
	return b != '.' && !isDigit(b)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func countDigits(s string) int {
	width := 0
	for i := 0; i < len(s) && isDigit(s[i]); i++ {
		width++
	}
	return width
}

func isSymbolAdjacent(lines []string, row int, col int, digitCount int) bool {
	for rowOffset := -1; rowOffset <= 1; rowOffset++ {
		rowIdx := row + rowOffset

		if rowIdx < 0 || rowIdx >= len(lines) || lines[rowIdx] == "" {
			continue
		}

		start := max(col-1, 0)
		end := min(col+digitCount+1, len(lines[rowIdx]))

		adjacent := strings.ContainsFunc(lines[rowIdx][start:end], func(r rune) bool {
			return isSymbol(byte(r))
		})

		if adjacent {
			return true
		}
	}

	return false
}

func findAdjacentNumbers(lines []string, row int, col int) []int {
	adjacent := []int{}

	for rowOffset := -1; rowOffset <= 1; rowOffset++ {
		rowIdx := row + rowOffset

		if rowIdx < 0 || rowIdx >= len(lines) || lines[rowIdx] == "" {
			continue
		}

		colIdx := col - 1
		for colIdx <= col+1 {
			if colIdx < 0 || colIdx >= len(lines[rowIdx]) {
				colIdx++
				continue
			}
			if isDigit(lines[rowIdx][colIdx]) {
				n, end := grabNumber(lines[rowIdx], colIdx)
				adjacent = append(adjacent, n)
				colIdx = end
			}

			colIdx++
		}
	}

	return adjacent
}

func grabNumber(s string, i int) (int, int) {
	start := i
	for j := i - 1; j >= 0 && isDigit(s[j]); j-- {
		start = j
	}

	width := countDigits(s[start:])
	n, err := strconv.Atoi(s[start : start+width])
	if err != nil {
		log.Fatalf("grabbing number: %s", err)
	}
	return n, start + width
}
