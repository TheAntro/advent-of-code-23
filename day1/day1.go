package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readLines() []string {
	file, err := os.Open("calibration.txt")
	if err != nil {
			fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	return lines
}


func main() {
	var digits []int
	numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var calibrationData = readLines()
	var calibrationSum = 0
	for _, line := range calibrationData {
		var firstDigit int
		var firstDigitIndexSet = -1
		var lastDigit int
		var firstNumber int
		lastDigitIndex := -1
	  firstNumberIndex := -1
		var lastNumber int
		lastNumberIndex := -1
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) == true {
				if (firstDigitIndexSet == -1) {
					firstDigit, _ = strconv.Atoi(string(line[i]))
					firstDigitIndexSet = i
				}
				lastDigit, _ = strconv.Atoi(string(line[i]))
				lastDigitIndex = i
			}
		}
		for i, number := range numbers {
			numberIndex := strings.Index(line, number)
			if (numberIndex != -1) {
				if (firstNumberIndex == -1) {
					firstNumberIndex = numberIndex
					firstNumber = i + 1
				} else if (numberIndex < firstNumberIndex) {
					firstNumberIndex = numberIndex
					firstNumber = i + 1
				}
			}
			lastIndex := strings.LastIndex(line, number)
			if (lastIndex != -1) {
				if (lastNumberIndex == -1) {
					lastNumberIndex = lastIndex
					lastNumber = i + 1
				} else if (lastNumberIndex < lastIndex) {
					lastNumberIndex = lastIndex
					lastNumber = i + 1
				}
			}
		}
		if (firstNumberIndex != -1) {
			if (firstNumberIndex < firstDigitIndexSet) {
				firstDigit = firstNumber
			}
		}
		if (lastNumberIndex != -1) {
			if (lastNumberIndex > lastDigitIndex) {
				lastDigit = lastNumber
			}
		}
		calibrationValue := 10 * firstDigit + lastDigit
		digits = append(digits, calibrationValue)
		calibrationSum += calibrationValue
	}
	fmt.Println(calibrationData)
	fmt.Println(digits)
	fmt.Println(calibrationSum)
}