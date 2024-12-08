package main

import (
	"advent-of-code/2024/6/utils"
	"fmt"
)

func main() {
	guardMap, startX, startY, startDirection := utils.GetInput()

	visited := utils.GetVisitedPath(startX, startY, startDirection, guardMap)

	// Answer is 5453
	fmt.Println(len(visited))
}
