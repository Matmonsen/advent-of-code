package utils

import (
	"fmt"
	"os"
	"strings"
)

const Ground = "."
const AntiNode = "#"

type Location struct {
	X   int
	Y   int
	Key string
}

type Item struct {
	Location Location
	Symbol   string
}

type SignalFrequency struct {
	AntennaA  Item
	NodeA     Item
	AntennaB  Item
	NodeB     Item
	DistanceX int
	DistanceY int
}

func PositionIsNotOutOfBounds(x int, y int, maxX int, maxY int) bool {
	return y > -1 && x > -1 && y < maxY && x < maxX
}

func CreateLocation(x int, y int) Location {
	return Location{
		x,
		y,
		GetMapKey(x, y),
	}
}

func FindAntiNodes(antenna Item, items []Item) []SignalFrequency {
	signalFrequencies := make([]SignalFrequency, 0)

	for _, item := range items {
		if item.Location.Key == antenna.Location.Key {
			continue
		}

		distanceX := antenna.Location.X - item.Location.X
		distanceY := antenna.Location.Y - item.Location.Y

		signalFrequencies = append(signalFrequencies, SignalFrequency{
			antenna,
			Item{
				CreateLocation(item.Location.X-distanceX, item.Location.Y-distanceY),
				AntiNode,
			},
			item,
			Item{
				CreateLocation(antenna.Location.X+distanceX, antenna.Location.Y+distanceY),
				AntiNode,
			},
			distanceX,
			distanceY,
		})
	}

	return signalFrequencies
}

func GetMapKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func GetInput() ([][]string, map[string][]Item, int, int) {
	fileData, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(fileData), "\n")
	puzzleInput := make([][]string, len(rows))
	antennaMap := make(map[string][]Item)

	for y, line := range rows {
		puzzleInput[y] = strings.Split(line, "")
		for x, char := range puzzleInput[y] {
			if char != Ground {
				_, exists := antennaMap[char]
				if !exists {
					antennaMap[char] = make([]Item, 0)
				}

				antennaMap[char] = append(antennaMap[char], Item{
					Location{x, y, GetMapKey(x, y)},
					char,
				})
			}

		}
	}

	outerY := len(puzzleInput)
	outerX := len(puzzleInput[0])

	return puzzleInput, antennaMap, outerX, outerY
}
