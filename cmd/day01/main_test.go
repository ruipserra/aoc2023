package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	got := part1(strings.NewReader(input))
	want := 142

	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	got := part2(strings.NewReader(input))
	want := 281

	if got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
