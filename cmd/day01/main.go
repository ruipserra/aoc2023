package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./cmd/day01/input.txt")
	if err != nil {
		log.Fatalf("opening input file: %s", err)
	}
	fmt.Printf("Part 1: %v\n", part1(f))
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatalf("rewinding input file: %s", err)
	}
	fmt.Printf("Part 2: %v\n", part2(f))
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

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func findFirstDigit(s string) string {
	for i := range s {
		if isDigit(s[i]) {
			return string(s[i])
		}
	}
	return ""
}

func findLastDigit(s string) string {
	for i := range s {
		c := s[len(s)-1-i]
		if isDigit(c) {
			return string(c)
		}
	}
	return ""
}

func part2(r io.Reader) int {
	result := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := findFirstDigitWithWords(line)
		lastDigit := findLastDigitWithWords(line)
		i, err := strconv.ParseInt(firstDigit+lastDigit, 10, 64)
		if err != nil {
			log.Fatalf("part2: ParseInt: %s", err)
		}
		result += int(i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("part2: reading input: %s", err)
	}
	return result
}

var words = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"zero":  "0",
}

func findDigitFromWords(s string) string {
	for word, digit := range words {
		if strings.HasPrefix(s, word) {
			return digit
		}
	}
	return ""
}

func findFirstDigitWithWords(s string) string {
	for i := range s {
		if isDigit(s[i]) {
			return string(s[i])
		}
		digit := findDigitFromWords(s[i:])
		if digit != "" {
			return digit
		}
	}
	return "0"
}

func findLastDigitWithWords(s string) string {
	for i := range s {
		c := s[len(s)-1-i]
		if isDigit(c) {
			return string(c)
		}
		digit := findDigitFromWords(s[len(s)-1-i:])
		if digit != "" {
			return digit
		}
	}
	return "0"
}
