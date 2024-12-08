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

func GetNext(x int, y int, guardMap [][]string, direction string) (int, int, string) {

	if direction == Up {
		if y-1 < 0 {
			return x, -1, direction
		} else {
			if guardMap[y-1][x] == Ground {
				return x, y - 1, direction
			} else {
				if x+1 > len(guardMap[0]) {
					return -1, y, direction
				} else {
					return x + 1, y, Right
				}
			}
		}
	} else if direction == Down {
		if y+1 >= len(guardMap) {
			return x, -1, direction
		} else {
			if guardMap[y+1][x] == Ground {
				return x, y + 1, direction
			} else {
				if x-1 < 0 {
					return -1, y, direction
				} else {
					return x - 1, y, Left
				}
			}
		}
	} else if direction == Left {
		if x-1 < 0 {
			return -1, y, direction
		} else {
			if guardMap[y][x-1] == Ground {
				return x - 1, y, direction
			} else {
				if y-1 < 0 {
					return x, -1, direction
				} else {
					return x, y - 1, Up
				}

			}
		}
	} else if direction == Right {
		if x+1 >= len(guardMap[0]) {
			return -1, y, direction
		} else {
			if guardMap[y][x+1] == Ground {
				return x + 1, y, direction
			} else {

				if y >= len(guardMap) {
					return -1, y, direction
				} else {
					return x, y + 1, Down
				}

			}
		}
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
			if char == Up {
				direction = Up
				x = _x
				y = _y
			} else if char == Down {
				direction = Down
				x = _x
				y = _y
			} else if char == Right {
				direction = Right
				x = _x
				y = _y
			} else if char == Left {
				direction = Left
				x = _x
				y = _y
			}
		}
	}

	return guardMap, x, y, direction
}
