package main

import (
	"advent-of-code/2024/10/utils"
	"fmt"
)

func FindAllTrails(start utils.Coords, data [][]int) int {
	finishedTrails := make(map[string]utils.Coords)
	neighbours := make([]utils.Coords, 0)
	neighbours = append(neighbours, start)
	i := 0

	for i < len(neighbours) {
		current := neighbours[i]
		if current.Value == 9 {
			finishedTrails[current.GetKey()] = current
		} else {
			for _, newNeighbor := range utils.GetNeighbours(current, data) {
				neighbours = append(neighbours, newNeighbor)
			}
		}
		i++
	}
	return len(finishedTrails)
}

func main() {
	data, trailHeads := utils.GetInput()
	trailCount := 0

	for _, trailHead := range trailHeads {
		trailCount += FindAllTrails(trailHead, utils.Copy2dArray(data))
	}

	fmt.Println(trailCount)

}
