package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./cmd/day01/input.txt")
	if err != nil {
		log.Fatalf("opening input file: %s", err)
	}

	fmt.Printf("Part 1: %v\n", part1(f))
}

func part1(r io.Reader) int {
	result := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := findFirstDigit(line)
		lastDigit := findLastDigit(line)
		i, err := strconv.ParseInt(firstDigit+lastDigit, 10, 64)
		if err != nil {
			log.Fatalf("part1: ParseInt: %s", err)
		}
		result += int(i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("part1: reading input: %s", err)
	}
	return result
}

func findFirstDigit(s string) string {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return string(c)
		}
	}
	return ""
}

func findLastDigit(s string) string {
	for i := range s {
		c := rune(s[len(s)-1-i])
		if c >= '0' && c <= '9' {
			return string(c)
		}
	}
	return ""
}
