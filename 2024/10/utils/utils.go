package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	X     int
	Y     int
	Value int
}

func (c Coords) GetKey() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func GetInput() ([][]int, []Coords) {
	fileData, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(fileData), "\n")

	data := make([][]int, len(rows))
	trailHeads := make([]Coords, 0)

	for i, line := range rows {
		row := strings.Split(strings.TrimSpace(line), "")
		data[i] = make([]int, len(row))
		for j, n := range row {
			if n == "." {
				continue
			}
			data[i][j], _ = strconv.Atoi(n)
			if data[i][j] == 0 {
				trailHeads = append(trailHeads, Coords{j, i, 0})
			}
		}
	}

	return data, trailHeads
}

func Copy2dArray(data [][]int) [][]int {
	copiedMap := make([][]int, len(data))
	for i := range data {
		copiedMap[i] = make([]int, len(data[i]))
		copy(copiedMap[i], data[i])
	}
	return copiedMap
}

func GetNeighbours(cord Coords, data [][]int) []Coords {
	neighbors := make([]Coords, 0)
	current := data[cord.Y][cord.X]
	nextHeight := current + 1

	left := cord.X - 1
	right := cord.X + 1
	up := cord.Y - 1
	down := cord.Y + 1

	leftIsWithinBounds := left > -1
	rightIsWithinBounds := right < len(data[0])
	downIsWithinBounds := down < len(data)
	upIsWithinBounds := up > -1

	// Up
	if upIsWithinBounds && data[up][cord.X] == nextHeight {
		neighbors = append(neighbors, Coords{X: cord.X, Y: up, Value: data[up][cord.X]})
	}
	// Down
	if downIsWithinBounds && data[down][cord.X] == nextHeight {
		neighbors = append(neighbors, Coords{X: cord.X, Y: down, Value: data[down][cord.X]})
	}
	// Left
	if leftIsWithinBounds && data[cord.Y][left] == nextHeight {
		neighbors = append(neighbors, Coords{X: left, Y: cord.Y, Value: data[cord.Y][left]})
	}
	// Right
	if rightIsWithinBounds && data[cord.Y][right] == nextHeight {
		neighbors = append(neighbors, Coords{X: right, Y: cord.Y, Value: data[cord.Y][right]})
	}

	return neighbors
}
