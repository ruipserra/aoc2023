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
	printAnswer("Part 1", part1(input))
	printAnswer("Part 2", part2(input))
}

func printAnswer(prefix string, answer int) {
	fmt.Printf("%s: %v\n", prefix, answer)
}

func eachLine(input string, f func(string)) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		f(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("reading input: %s", err)
	}
}

func part1(input string) int {
	result := 0
	eachLine(input, func(line string) {
		firstDigit := findFirstDigit(line)
		lastDigit := findLastDigit(line)
		i, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			log.Fatalf("part1: atoi: %s", err)
		}
		result += i
	})
	return result
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func findFirst(s string, backwards bool, finder func(substr string) string) string {
	for i := range s {
		substr := s[i:]
		if backwards {
			substr = s[len(s)-1-i:]
		}

		if found := finder(substr); found != "" {
			return found
		}
	}
	return ""
}

func findDigit(s string, backwards bool) string {
	return findFirst(s, backwards, func(substr string) string {
		if isDigit(substr[0]) {
			return string(substr[0])
		}
		return ""
	})
}

func findFirstDigit(s string) string {
	return findDigit(s, false)
}

func findLastDigit(s string) string {
	return findDigit(s, true)
}

func part2(input string) int {
	result := 0
	eachLine(input, func(line string) {
		firstDigit := findFirstDigitWithWords(line)
		lastDigit := findLastDigitWithWords(line)
		i, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			log.Fatalf("part2: atoi: %s", err)
		}
		result += i
	})
	return result
}

var words = map[string]string{
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
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
	return findFirst(s, false, func(substr string) string {
		if digit := findDigitFromWords(substr); digit != "" {
			return digit
		}
		return ""
	})
}

func findLastDigitWithWords(s string) string {
	return findFirst(s, true, func(substr string) string {
		if digit := findDigitFromWords(substr); digit != "" {
			return digit
		}
		return ""
	})
}
