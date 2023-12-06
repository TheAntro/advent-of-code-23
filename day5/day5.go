package main

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/TheAntro/advent-of-code/aocutils"
)

type GardeningMap struct {
	src  string
	dest string
	m    []MapRange
}

type MapRange struct {
	dest int
	src  int
	len  int
}

func findMap(src string, dest string, maps []GardeningMap) (GardeningMap, error) {
	for _, m := range maps {
		if m.src == src && m.dest == dest {
			return m, nil
		}
	}
	var notFound []MapRange
	return GardeningMap{"", "", notFound}, errors.New("not found")
}

func getDest(src int, m map[int]int) int {
	dest := m[src]
	if dest == 0 {
		dest = src
	}
	return dest
}

func lowestLocWithMatchingSeed(seedsLine []int, maps []GardeningMap) int {
	locMap, _ := findMap("humidity", "location", maps)
	locationSlice := locMap.m[:]
	slices.Reverse(maps)
	sort.Slice(locationSlice, func(a, b int) bool {
		return locationSlice[a].dest < locationSlice[b].dest
	})
	highestLoc := locationSlice[len(locationSlice)-1].dest + locationSlice[len(locationSlice)-1].len - 1
	for loc := 0; loc <= highestLoc; loc++ {
		if loc%1000000 == 0 {
			fmt.Printf("Trying location %d\n", loc)
		}
		num := loc
		for _, m := range maps {
			for _, n := range m.m {
				if num >= n.dest {
					if num-n.dest < n.len {
						num = n.src + (num - n.dest)
						break
					}
				}
			}
		}
		for i, seed := range seedsLine {
			if i%2 == 0 {
				if num >= seed && num < seed+seedsLine[i+1] {
					return loc
				}
			}
		}
	}
	return -1
}

func main() {
	lines := aocutils.ReadLines("gardening_map.txt")
	seeds := strings.Fields(strings.Split(lines[0], ":")[1])
	var maps []GardeningMap
	for i, line := range lines {
		if strings.Contains(line, "map") {
			var newMap []MapRange
			mapParts := strings.Split(strings.Split(line, " ")[0], "-")
			src := mapParts[0]
			dest := mapParts[2]
			for j := i + 1; j < len(lines); j++ {
				if lines[j] == "" {
					maps = append(maps, GardeningMap{src, dest, newMap})
					break
				}
				rangeParts := strings.Fields(lines[j])
				srcStart, _ := strconv.Atoi(rangeParts[1])
				destStart, _ := strconv.Atoi(rangeParts[0])
				rangeLen, _ := strconv.Atoi(rangeParts[2])
				newMap = append(newMap, MapRange{destStart, srcStart, rangeLen})
				if j+1 == len(lines) {
					maps = append(maps, GardeningMap{src, dest, newMap})
				}
			}
		}
	}
	lowestLoc := -1
	for _, seed := range seeds {
		num, _ := strconv.Atoi(seed)
		for _, m := range maps {
			for _, n := range m.m {
				if num > n.src {
					if num-n.src < n.len {
						num = n.dest + (num - n.src)
						break
					}
				}
			}
		}
		if lowestLoc == -1 {
			lowestLoc = num
		} else if num < lowestLoc {
			lowestLoc = num
		}
	}
	fmt.Printf("Part 1: %d\n", lowestLoc)
	var seedNums []int
	for _, seed := range seeds {
		seedNum, _ := strconv.Atoi(seed)
		seedNums = append(seedNums, seedNum)
	}
	lowestLocWithSeed := lowestLocWithMatchingSeed(seedNums, maps)
	fmt.Printf("Part 2: %d\n", lowestLocWithSeed)
}
