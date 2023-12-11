package main

import "testing"

const testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func TestPart1(t *testing.T) {
	const want = 114
	got := part1(testInput)
	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}
