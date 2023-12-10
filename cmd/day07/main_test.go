package main

import "testing"

const testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestPart1(t *testing.T) {
	const want = 6440
	got := part1(testInput)
	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	const want = 5905
	got := part2(testInput)
	if got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
