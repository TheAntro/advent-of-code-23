package aocutils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(fileName string) []string {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
			fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Contains[T comparable](slice []T, value T) bool {
	for _, elem := range slice {
		if (elem == value) {
			return true
		}
	}
	return false
}