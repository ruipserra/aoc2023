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
