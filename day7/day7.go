package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/TheAntro/advent-of-code/aocutils"
)

type Hand struct {
	hand   []int
	rank   int
	bid    int
	jokers int
}

func getRankAndJokers(hand []int) (int, int) {
	var values [13]int
	var jokers int
	for _, hand := range hand {
		if hand == 1 {
			jokers++
		} else {
			values[hand-2] = values[hand-2] + 1
		}
	}
	var highestCount int
	var secondHighestCount int
	for _, count := range values {
		var lastHighestCount int
		if count > highestCount {
			lastHighestCount = highestCount
			highestCount = count
		} else if count > secondHighestCount {
			secondHighestCount = count
		}
		if lastHighestCount > secondHighestCount {
			secondHighestCount = lastHighestCount
		}
	}
	highestCount += jokers
	var rank int
	switch highestCount {
	case 5:
		rank = 0
	case 4:
		rank = 1
	case 3:
		if secondHighestCount == 2 {
			rank = 2
		} else {
			rank = 3
		}
	case 2:
		if secondHighestCount == 2 {
			rank = 4
		} else {
			rank = 5
		}
	default:
		rank = 6
	}
	return rank, jokers
}

func insert(a []Hand, index int, value Hand) []Hand {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func parseHands(lines []string) []Hand {
	var hands []Hand
	for _, line := range lines {
		fields := strings.Fields(line)
		var hand []int
		for _, card := range strings.Split(fields[0], "") {
			var cardValue int
			switch card {
			case "A":
				cardValue = 14
			case "K":
				cardValue = 13
			case "Q":
				cardValue = 12
			case "J":
				cardValue = 1
			case "T":
				cardValue = 10
			default:
				cardValue, _ = strconv.Atoi(card)
			}
			hand = append(hand, cardValue)
		}
		bid, _ := strconv.Atoi(fields[1])
		rank, jokers := getRankAndJokers(hand)
		hands = append(hands, Hand{hand, rank, bid, jokers})
	}
	sort.Slice(hands, func(a, b int) bool {
		if hands[a].rank != hands[b].rank {
			return hands[a].rank > hands[b].rank
		} else if hands[a].hand[0] != hands[b].hand[0] {
			return hands[a].hand[0] < hands[b].hand[0]
		} else if hands[a].hand[1] != hands[b].hand[1] {
			return hands[a].hand[1] < hands[b].hand[1]
		} else if hands[a].hand[2] != hands[b].hand[2] {
			return hands[a].hand[2] < hands[b].hand[2]
		} else if hands[a].hand[3] != hands[b].hand[3] {
			return hands[a].hand[3] < hands[b].hand[3]
		} else if hands[a].hand[4] != hands[b].hand[4] {
			return hands[a].hand[4] < hands[b].hand[4]
		} else if hands[a].jokers != hands[b].jokers {
			return hands[a].jokers > hands[b].jokers
		}
		return false
	})
	return hands
}

func main() {
	lines := aocutils.ReadLines("hands.txt")
	hands := parseHands(lines)
	var total int
	for i := range hands {
		total += hands[i].bid * (i + 1)
		fmt.Println(hands[i])
	}
	fmt.Println(total)
}
