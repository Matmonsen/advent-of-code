package main

import (
	"advent-of-code/2024/6/utils"
	"fmt"
	"strconv"
	"strings"
)

func VisualizePath(startX int, startY int, guardMap [][]string, visited map[string]int, oX int, oY int) {
	for key := range visited {
		a := strings.Split(key, "|")
		b := strings.Split(a[0], ",")
		n1, _ := strconv.Atoi(b[0])
		n2, _ := strconv.Atoi(b[1])

		guardMap[n2][n1] = "X"
	}
	guardMap[oY][oX] = "O"
	guardMap[startY][startX] = "^"
	for _, row := range guardMap {
		fmt.Println(row)
	}
	fmt.Println()
}

func IsLoop(startX int, startY int, startDirection string, guardMap [][]string, oX int, oY int) bool {
	x := startX
	y := startY
	direction := startDirection
	visited := make(map[string]int)

	for y > 0 && x > 0 && x < len(guardMap[0]) && y < len(guardMap) {
		guardMap[y][x] = utils.Ground
		x, y, direction = utils.GetNext(x, y, guardMap, direction)
		position := fmt.Sprintf("%d,%d|%s", x, y, direction)
		if visited[position] == 4 {

			//VisualizePath(startX, startY, guardMap, visited, oX, oY)

			return true
		}
		if x >= 0 && y >= 0 {
			guardMap[y][x] = direction
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

		if IsLoop(startX, startY, startDirection, guardMapCopy, visitedX, visitedY) {
			loops++
		}
	}

	return loops
}

func AddObstacleToAllPositions(startX int, startY int, startDirection string, guardMap [][]string) int {
	loops := 0
	for i, row := range guardMap {
		fmt.Println("i: ", i)
		for j, _ := range row {
			guardMapCopy := CopyMap(guardMap)
			guardMapCopy[i][j] = utils.Obstacle

			if IsLoop(startX, startY, startDirection, guardMapCopy, j, i) {
				loops++
			}
		}

	}

	return loops
}

func main() {
	guardMap, startX, startY, startDirection := utils.GetInput()
	x := startX
	y := startY
	direction := startDirection
	steps := 0
	visited := make(map[string]int)

	for y > 0 && x > 0 && x < len(guardMap[0]) && y < len(guardMap) {
		steps++
		visited[utils.GetMapKey(x, y)] = steps
		guardMap[y][x] = utils.Ground
		x, y, direction = utils.GetNext(x, y, guardMap, direction)
		if x >= 0 && y >= 0 {
			guardMap[y][x] = direction
		}
	}

	// 2132 is too low....
	fmt.Println(AddObstacleToOnlyVisitedPositions(startX, startY, startDirection, guardMap, visited))
}
