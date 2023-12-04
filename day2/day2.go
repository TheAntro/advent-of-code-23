package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	number int
	red int
	green int
	blue int
}

type BagContent struct {
	red int
	green int
	blue int
}

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

func Split(r rune) bool {
	return r == ';' || r == ','
}

func parseGameData(data []string) []Game {
	var games []Game
	for i := 0; i < len(data); i++ {
		splitData := strings.Split(data[i], ":")
		game, _ := strconv.Atoi(strings.Split(splitData[0], " ")[1])
		cubes := strings.FieldsFunc(splitData[1], Split)
		var red int
		var green int
		var blue int
		for j := 0; j < len(cubes); j++ {
			cubeData := strings.Split(cubes[j], " ")
			count, _ := strconv.Atoi(cubeData[1])
			color := cubeData[2]
			switch color {
			case "red":
				if count > red {
					red = count
				}
			case "green":
				if count > green {
					green = count
				}
			case "blue":
				if count > blue {
					blue = count
				}
			}
		}
		games = append(games, Game{game, red, green, blue})
	}
	return games
}

func getPossibleGames(games []Game, bag BagContent) []int {
	var possibleGames []int
	for i := 0; i < len(games); i++ {
		if bag.red >= games[i].red && bag.green >= games[i].green && bag.blue >= games[i].blue {
			possibleGames = append(possibleGames, games[i].number)
		}
	}
	return possibleGames
}

func getPowerOfCubes(games []Game) int {
	var sum int;
	for _, game := range games {
		sum += game.red * game.green * game.blue
	}
	return sum
}

func main() {
	lines := readLines("cube-conundrum")
	games := parseGameData(lines)
	bag := BagContent{12, 13, 14}
	possibleGames := getPossibleGames(games, bag)
	var sum int
	for i := 0; i < len(possibleGames); i++ {
		sum += possibleGames[i]
	}
	fmt.Println(sum)
	powerOfCubes := getPowerOfCubes((games))
	fmt.Printf("power: %d\n", powerOfCubes)
}