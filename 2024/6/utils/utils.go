package utils

import (
	"fmt"
	"os"
	"strings"
)

const Up = "^"
const Down = "v"
const Left = "<"
const Right = ">"
const Obstacle = "#"
const Ground = "."

func GetNextWhenGoingUp(x int, y int, direction string, guardMap [][]string) (int, int, string) {

	if y-1 < 0 {
		return x, -1, direction
	} else {
		if guardMap[y-1][x] == Obstacle {
			if guardMap[y][x+1] == Obstacle {
				return x, y + 1, Down
			} else {
				return x + 1, y, Right
			}
		} else {
			return x, y - 1, Up
		}
	}
}

func GetNextWhenGoingDown(x int, y int, direction string, guardMap [][]string) (int, int, string) {

	maxY := len(guardMap)
	if y+1 >= maxY {
		return x, -1, direction
	} else {
		if guardMap[y+1][x] == Obstacle {
			if guardMap[y][x-1] == Obstacle {
				return x, y - 1, Up
			} else {
				return x - 1, y, Left
			}
		} else {
			return x, y + 1, Down
		}
	}
}

func GetNextWhenGoingLeft(x int, y int, direction string, guardMap [][]string) (int, int, string) {
	if x-1 < 0 {
		return -1, y, direction
	} else {
		if guardMap[y][x-1] == Obstacle {
			if guardMap[y-1][x] == Obstacle {
				return x + 1, y, Right
			} else {
				return x, y - 1, Up
			}
		} else {
			return x - 1, y, Left
		}
	}
}

func GetNextWhenGoingRight(x int, y int, direction string, guardMap [][]string) (int, int, string) {
	maxX := len(guardMap[0])

	if x+1 >= maxX {
		return -1, y, direction
	} else {
		if guardMap[y][x+1] == Obstacle {
			if guardMap[y+1][x] == Obstacle {
				return x - 1, y, Left
			} else {
				return x, y + 1, Down
			}
		} else {
			return x + 1, y, Right
		}
	}
}

func GetNext(x int, y int, guardMap [][]string, direction string) (int, int, string) {
	if direction == Up {
		return GetNextWhenGoingUp(x, y, direction, guardMap)
	} else if direction == Down {
		return GetNextWhenGoingDown(x, y, direction, guardMap)
	} else if direction == Left {
		return GetNextWhenGoingLeft(x, y, direction, guardMap)
	} else if direction == Right {
		return GetNextWhenGoingRight(x, y, direction, guardMap)
	} else {
		return -1, -1, direction
	}
}

func GetMapKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func GetInput() ([][]string, int, int, string) {
	fileData, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(fileData), "\n")
	guardMap := make([][]string, len(rows))
	x := 0
	y := 0
	direction := Up

	for _y, line := range rows {
		guardMap[_y] = strings.Split(line, "")

		for _x, char := range guardMap[_y] {
			if char != Ground && char != Obstacle {
				direction = char
				x = _x
				y = _y
			}
		}
	}

	return guardMap, x, y, direction
}

func PositionIsWithinBounds(x int, y int, guardMap [][]string) bool {
	return y > 0 && x > 0 && x < len(guardMap[0]) && y < len(guardMap)
}

func GetVisitedPath(startX int, startY int, startDirection string, guardMap [][]string) map[string]int {
	x := startX
	y := startY
	direction := startDirection
	visited := make(map[string]int)
	for PositionIsWithinBounds(x, y, guardMap) {
		visited[GetMapKey(x, y)] = 1
		guardMap[y][x] = Ground
		x, y, direction = GetNext(x, y, guardMap, direction)
	}
	return visited
}
