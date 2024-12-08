package main

import (
	"advent-of-code/2024/6/utils"
	"fmt"
	"strconv"
	"strings"
)

func IsLoop(startX int, startY int, startDirection string, guardMap [][]string) bool {
	x := startX
	y := startY
	direction := startDirection
	visited := make(map[string]int)

	for utils.PositionIsWithinBounds(x, y, guardMap) {
		guardMap[y][x] = utils.Ground
		x, y, direction = utils.GetNext(x, y, guardMap, direction)

		position := fmt.Sprintf("%d,%d|%s", x, y, direction)
		if visited[position] == 1 {
			return true
		}

		visited[position]++
	}

	return false
}

func CopyMap(guardMap [][]string) [][]string {
	copiedMap := make([][]string, len(guardMap))
	for i := range guardMap {
		copiedMap[i] = make([]string, len(guardMap[i]))
		copy(copiedMap[i], guardMap[i])
	}
	return copiedMap
}

func AddObstacleToOnlyVisitedPositions(startX int, startY int, startDirection string, guardMap [][]string, visited map[string]int) int {
	loops := 0
	for visitedPosition := range visited {
		positions := strings.Split(visitedPosition, ",")
		visitedX, _ := strconv.Atoi(positions[0])
		visitedY, _ := strconv.Atoi(positions[1])

		if visitedX == startX && visitedY == startY {
			continue
		}

		guardMapCopy := CopyMap(guardMap)
		guardMapCopy[visitedY][visitedX] = utils.Obstacle

		if IsLoop(startX, startY, startDirection, guardMapCopy) {
			loops++
		}
	}

	return loops
}

func main() {
	guardMap, startX, startY, startDirection := utils.GetInput()

	visited := utils.GetVisitedPath(startX, startY, startDirection, guardMap)

	// 2188 is  correct
	fmt.Println(AddObstacleToOnlyVisitedPositions(startX, startY, startDirection, guardMap, visited))
}
