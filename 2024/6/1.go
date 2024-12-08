package main

import (
	"advent-of-code/2024/6/utils"
	"fmt"
)

func main() {
	guardMap, x, y, direction := utils.GetInput()
	visited := make(map[string]int)
	steps := 0

	visited[utils.GetMapKey(x, y)] = steps

	for y > 0 && x > 0 && x < len(guardMap[0]) && y < len(guardMap) {
		steps++
		visited[utils.GetMapKey(x, y)] = steps
		guardMap[y][x] = utils.Ground
		x, y, direction = utils.GetNext(x, y, guardMap, direction)
		if x >= 0 && y >= 0 {
			guardMap[y][x] = direction
		}
	}

	fmt.Println(len(visited))
}
