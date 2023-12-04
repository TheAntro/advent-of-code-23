package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(fileName string) []string {
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

func readMatrix(fileName string) [][]string {
	var matrix [][]string
	lines := readLines(fileName)
	for i := range lines {
		matrix = append(matrix, strings.Split(lines[i], ""))
	}
	return matrix
}

func isSymbol(char string) bool {
	if char == "." {
		return false
	} else if isInt(char) {
		return false
	} else {
		return true
	}
}

func isInt(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
    return true
	}
	return false
}

type Coord struct {
	x, y int
}

func getAdjacentPositions(matrix [][]string, symbolX int, symbolY int) []Coord {
	var adjacentPositions []Coord
	if (symbolX > 0) {
		adjacentPositions = append(adjacentPositions, Coord{symbolX - 1, symbolY})
	}
	if (symbolY > 0) {
		adjacentPositions = append(adjacentPositions, Coord{symbolX, symbolY - 1})
	}
	if (symbolX > 0 && symbolY > 0) {
		adjacentPositions = append(adjacentPositions, Coord{symbolX - 1, symbolY - 1})
	}
	if (symbolX < len(matrix[0]) && symbolY > 0) {
		adjacentPositions = append(adjacentPositions, Coord{symbolX + 1, symbolY - 1})
	}
	if (symbolY < len(matrix)) {
		adjacentPositions = append(adjacentPositions, Coord{symbolX, symbolY + 1})
	}
	if (symbolX < len(matrix[0])) {
		adjacentPositions = append(adjacentPositions, Coord{symbolX + 1, symbolY})
	}
	if (symbolY < len(matrix) && symbolX < len(matrix[0])) {
		adjacentPositions = append(adjacentPositions, Coord{symbolX + 1, symbolY + 1})
	}
	if (symbolX > 0 && symbolY < len(matrix)) {
		adjacentPositions = append(adjacentPositions, Coord{symbolX - 1, symbolY + 1})
	}
	return adjacentPositions
}

func resolveNumber(matrix [][]string, numberX int, numberY int) (int, int) {
	var startX int
	if (numberX != 0) {
		for x := numberX; isInt(matrix[numberY][x]); {
			startX = x
			if (x > 0) {
				x--
			} else {
				break
			}
		}
	}
	var endX int
	if (numberX == len(matrix[0])) {
		endX = numberX
	} else {
		for x := numberX; isInt(matrix[numberY][x]); {
			endX = x
			if x < len(matrix[0]) - 1 {
				x++
			} else {
				break
			}
		}
	}
	var numStr string 
	for x := startX; x <= endX; x++ {
		numStr += matrix[numberY][x]
	}
	number, _ := strconv.Atoi(numStr)
	return number, startX
}

func getAdjacentParts(matrix [][]string, symbolX int, symbolY int) map[string]int {
	adjacentPositions := getAdjacentPositions(matrix, symbolX, symbolY)
	adjacentParts :=  make(map[string]int)
	for _, pos := range adjacentPositions {
		if isInt(matrix[pos.y][pos.x]) {
			// resolveNumber(matrix, pos.x, pos.y)
			number, startX := resolveNumber(matrix, pos.x, pos.y)
			adjacentParts["x" + strconv.Itoa(startX) + "y" + strconv.Itoa(pos.y)] = number
		}
	}
	return adjacentParts
}

func parseSchematicMatrix(matrix [][]string) (map[string]int, int) {
	partMap := make(map[string]int)
	var gearRatios int
	for y := range matrix {
		for x := range matrix[y] {
			if isSymbol(matrix[y][x]) {
				adjacentParts := getAdjacentParts(matrix, x, y)
				for pos, num := range adjacentParts {
					partMap[pos] = num
				}
				if (matrix[y][x] == "*" && len(adjacentParts) == 2) {
					gearRatio := 1
					for _, num := range adjacentParts {
						gearRatio *= num
					}
					gearRatios += gearRatio
				}
			}
		} 
	}
	return partMap, gearRatios
}

func main() {
	schematicMatrix := readMatrix("schematic.txt")
	partsMap, gearRatio := parseSchematicMatrix(schematicMatrix)
	var sum int
	for _, partNum := range partsMap {
		sum += partNum
	}
	fmt.Println(sum)
	fmt.Printf("Gear ratios: %d\n", gearRatio)
}