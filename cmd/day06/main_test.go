package main

import "testing"

const testInput = `Time:      7  15   30
Distance:  9  40  200`

func TestPart1(t *testing.T) {
	const want = 288
	got := part1(testInput)
	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}
