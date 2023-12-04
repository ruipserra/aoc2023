package main

import "testing"

func TestMain(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	want := 4361
	got := part1(input)

	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}
