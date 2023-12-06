package main

import (
	"fmt"
	"strings"

	"github.com/TheAntro/advent-of-code/aocutils"
)

type CardData struct {
	wnums []string
	mynums []string
}

func generateNumbers(numstr string) []string {
	var nums []string
		for _, num := range strings.Split(numstr, " ") {
			if (num != "" && num != " ") {
				nums = append(nums, num)
			}
		}
	return nums
}

func parseCards(lines []string) []CardData {
	var cards []CardData
	for _, line := range lines {
		cardParts := strings.Split(line, ":")
		numParts := strings.Split(cardParts[1], "|")
		wnums := generateNumbers(numParts[0])
		mynums := generateNumbers(numParts[1])
		cards = append(cards, CardData{wnums, mynums})
	}
	return cards
}

func resolveCardValue(count int) int {
	value := 0
	for i := 0; i < count; i++ {
		if (value == 0) {
			value = 1
		} else {
			value *= 2
		}
	}
	return value
}

func main() {
	lines := aocutils.ReadLines("scratchcards.txt")
	cards := parseCards((lines))
	var wonCards []int
	for range lines {
		wonCards = append(wonCards, 0)
	} 
	for i := 0; i < len(cards); i++ {
		count := 0
		for _, mynum := range cards[i].mynums {
			if aocutils.Contains(cards[i].wnums, mynum) {
				count++
			}
		}
		fmt.Printf("Current index: %d\n", i)
		fmt.Printf("Current count: %d\n", count)
		fmt.Printf("Won cards: %d\n", wonCards[i])
		for j := 0; j < count; j++ {
			addingTo := i + j + 1
			newCards := 1 + 1 * wonCards[i]
			fmt.Printf("Adding %d new cards to %d\n", newCards, addingTo)
			
			wonCards[i + j + 1] = wonCards[i + j + 1] + 1 + 1 * wonCards[i]
		}
	}
	var sum int
	for _, cards := range wonCards {
		sum += cards + 1
	}
	fmt.Println(sum)
}