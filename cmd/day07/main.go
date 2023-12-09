package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input string) int {
	handsAndBids := parseHandsAndBids(input)
	sortHands(handsAndBids)
	winnings := 0
	for i, handAndBid := range handsAndBids {
		rank := i + 1
		winnings += rank * handAndBid.bid
	}
	return winnings
}

type handAndBid struct {
	hand string
	bid  int
}

func parseHandsAndBids(input string) []handAndBid {
	handsAndBids := []handAndBid{}
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			handsAndBids = append(handsAndBids, parseHandAndBid(line))
		}
	}
	return handsAndBids
}

func parseHandAndBid(line string) handAndBid {
	hand, bidStr, _ := strings.Cut(line, " ")
	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		log.Fatalf("parseHandAndBid: %s", err)
	}
	return handAndBid{hand: hand, bid: bid}
}

func sortHands(handsAndBids []handAndBid) {
	slices.SortFunc(handsAndBids, func(a, b handAndBid) int {
		aType := handType(a.hand)
		bType := handType(b.hand)

		if aType < bType {
			return -1
		}

		if aType > bType {
			return 1
		}

		return compareCards(a.hand, b.hand)
	})
}

var handTypes = [][]byte{{1, 1, 1, 1, 1}, {1, 1, 1, 2}, {1, 2, 2}, {1, 1, 3}, {2, 3}, {1, 4}, {5}}

func handType(hand string) int {
	cards := map[byte]byte{}
	for _, r := range hand {
		cards[byte(r)]++
	}

	cardCounts := []byte{}
	for _, count := range cards {
		cardCounts = append(cardCounts, count)
	}

	slices.Sort(cardCounts)
	result := slices.IndexFunc(handTypes, func(other []byte) bool {
		return bytes.Equal(cardCounts, other)
	})
	if result == -1 {
		log.Fatalf("unknown hand type: %s", hand)
	}

	return result
}

var cardLabels = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func compareCards(a, b string) int {
	for i := 0; i < len(a); i++ {
		ia := slices.Index(cardLabels, a[i])
		ib := slices.Index(cardLabels, b[i])
		if ia != ib {
			return ia - ib
		}
	}
	return 0
}
