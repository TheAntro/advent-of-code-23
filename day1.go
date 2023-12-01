package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"strconv"
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
	var calibrationData = readLines()
	var calibrationSum = 0
	for _, line := range calibrationData {
		var firstDigit string
		var firstDigitIndexSet = -1
		var lastDigit string
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) == true {
				if (firstDigitIndexSet == -1) {
					firstDigit = string(line[i])
					firstDigitIndexSet = i
				}
				lastDigit = string(line[i])
			}
		}
		var calibrationValue = firstDigit + lastDigit
		calibrationNumber, _ := strconv.Atoi(calibrationValue)
		calibrationSum += calibrationNumber
	}
	fmt.Println(calibrationData)
	fmt.Println(calibrationSum)
}