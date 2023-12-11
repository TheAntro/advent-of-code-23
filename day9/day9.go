package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TheAntro/advent-of-code/aocutils"
)

func extrapolateValue(history []int) int {
	var differences []int
	allZeroes := true
	for i := 1; i < len(history); i++ {
		diff := history[i] - history[i-1]
		differences = append(differences, diff)
		if history[i] != 0 {
			allZeroes = false
		}
	}
	if allZeroes == true {
		return 0

	}
	// part 1
	// return history[len(history)-1] + extrapolateValue(differences)
	// part 2
	return history[0] - extrapolateValue(differences)
}

func main() {
	lines := aocutils.ReadLines("history.txt")
	var histories [][]int
	for i := range lines {
		var history []int
		for _, field := range strings.Fields(lines[i]) {
			historyNum, _ := strconv.Atoi(field)
			history = append(history, historyNum)
		}
		histories = append(histories, history)
	}
	sum := 0
	for j := range histories {
		sum += extrapolateValue(histories[j])
	}
	fmt.Println(sum)
}
