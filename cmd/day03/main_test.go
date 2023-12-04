package main

import "testing"

var testInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart1(t *testing.T) {
	want := 4361
	got := part1(testInput)

	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 467835
	got := part2(testInput)

	if got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
