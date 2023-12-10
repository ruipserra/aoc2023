package main

import "testing"

const testInput1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

const testInput2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

func TestPart1(t *testing.T) {
	tests := []struct {
		desc  string
		input string
		want  int
	}{
		{"testInput1", testInput1, 2},
		{"testInput2", testInput2, 6},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := part1(test.input)
			if got != test.want {
				t.Errorf("part1() = %v, want %v", got, test.want)
			}
		})
	}
}
