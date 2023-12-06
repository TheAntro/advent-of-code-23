package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/TheAntro/advent-of-code/aocutils"
)

type Race struct {
	time, distance int
}

func generateRaces(data []string) []Race {
	var races []Race
	times := strings.Fields(strings.Split(data[0], ":")[1])
	distances := strings.Fields(strings.Split(data[1], ":")[1])
	for i := range times {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		races = append(races, Race{time, distance})
	}
	return races
}

func generateRaceP2(data []string) Race {
	time, _ := strconv.Atoi(strings.Join(strings.Fields(strings.Split(data[0], ":")[1]), ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(strings.Split(data[1], ":")[1]), ""))
	return Race{time, distance}
}

func calcPossibles(race Race) int {
	b := race.time
	c := race.distance
	higher := (float64(-b) - math.Sqrt(float64(b*b-4*c))) / -2
	lower := (float64(-b) + math.Sqrt(float64(b*b-4*c))) / -2
	ties := 0
	if higher == float64(int64(higher)) {
		ties++
	}
	possibles := int(higher) - int(lower) - ties
	return possibles
}

func main() {
	lines := aocutils.ReadLines("races.txt")
	races := generateRaces(lines)
	multiplied := 1
	for _, race := range races {
		possibles := calcPossibles(race)
		multiplied *= possibles
	}
	fmt.Println(multiplied)
	longRace := generateRaceP2(lines)
	possibles := calcPossibles(longRace)
	fmt.Println(possibles)
}
