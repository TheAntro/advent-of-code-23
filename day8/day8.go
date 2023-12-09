package main

import (
	"fmt"
	"strings"

	"github.com/TheAntro/advent-of-code/aocutils"
)

func parseNodes(nodeData []string) (map[string][]string, []string) {
	nodeMap := make(map[string][]string)
	var startNodes []string
	for i := range nodeData {
		s := strings.Split(nodeData[i], " = ")
		node := s[0]
		if string(node[2]) == "A" {
			startNodes = append(startNodes, node)
		}
		lr := strings.Split(strings.TrimLeft(strings.TrimRight(s[1], ")"), "("), ", ")
		nodeMap[node] = lr
	}
	return nodeMap, startNodes
}

func main() {
	lines := aocutils.ReadLines("nodes.txt")
	instructions := strings.Split(lines[0], "")
	nodes, startNodes := parseNodes(lines[2:])
	for _, node := range startNodes {
		var count int
		for i := 0; ; i++ {
			if i%1000000 == 0 {
			}
			if string(node[2]) == "Z" {
				count = i
				break
			}
			if instructions[i%len(instructions)] == "L" {
				node = nodes[node][0]
			} else {
				node = nodes[node][1]
			}
		}
		// Plugged these into an LCM calculator
		fmt.Println(count)
	}
}
