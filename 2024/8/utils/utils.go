package utils

import (
	"fmt"
	"os"
	"strings"
)

const Ground = "."
const Node = "#"

type Location struct {
	X      int
	Y      int
	Key    string
	Symbol string
}

type AntiNode struct {
	AntennaA Location
	NodeA    Location
	AntennaB Location
	NodeB    Location
	Distance Location
}

func FindAntiNodes(antenna Location, locations []Location) []AntiNode {
	antiNodes := make([]AntiNode, 0)

	for _, location := range locations {
		if location.Key == antenna.Key {
			continue
		}

		x := antenna.X - location.X
		y := antenna.Y - location.Y

		nodeAx := location.X - x
		nodeAy := location.Y - y

		nodeBx := antenna.X + x
		nodeBy := antenna.Y + y

		antiNodes = append(antiNodes, AntiNode{
			antenna,
			Location{
				nodeAx,
				nodeAy,
				GetMapKey(nodeAx, nodeAy),
				Node,
			},
			location,
			Location{
				nodeBx,
				nodeBy,
				GetMapKey(nodeBx, nodeBy),
				Node,
			},
			Location{
				x,
				y,
				GetMapKey(x, y),
				"",
			},
		})
	}

	return antiNodes
}

func GetMapKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func GetInput() ([][]string, map[string][]Location) {
	fileData, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(fileData), "\n")
	puzzleInput := make([][]string, len(rows))
	antennaMap := make(map[string][]Location)

	for y, line := range rows {
		puzzleInput[y] = strings.Split(line, "")
		for x, char := range puzzleInput[y] {
			if char != Ground {
				_, exists := antennaMap[char]
				if !exists {
					antennaMap[char] = make([]Location, 0)
				}

				antennaMap[char] = append(antennaMap[char], Location{x, y, GetMapKey(x, y), char})
			}

		}
	}
	return puzzleInput, antennaMap
}
