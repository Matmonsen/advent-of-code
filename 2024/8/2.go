package main

import (
	"advent-of-code/2024/8/utils"
	"fmt"
	"math"
)

func main() {
	_, antennas, outerX, outerY := utils.GetInput()
	uniqueAntiNodes := make(map[string]utils.Item)
	antiNodeDistances := make([]int, int(math.Max(float64(outerY), float64(outerX))))

	for key := range antennas {
		for _, location := range antennas[key] {
			antiNodes := utils.FindAntiNodes(location, antennas[key])

			for _, antiNode := range antiNodes {

				uniqueAntiNodes[antiNode.AntennaB.Location.Key] = antiNode.NodeB
				uniqueAntiNodes[antiNode.AntennaA.Location.Key] = antiNode.NodeA

				for distance := range antiNodeDistances {
					distanceY := distance * antiNode.DistanceY
					distanceX := distance * antiNode.DistanceX

					nodeAy := antiNode.NodeA.Location.Y - distanceY
					nodeAx := antiNode.NodeA.Location.X - distanceX
					nodeBy := antiNode.NodeB.Location.Y + distanceY
					nodeBx := antiNode.NodeB.Location.X + distanceX

					if utils.PositionIsNotOutOfBounds(nodeAx, nodeAy, outerX, outerY) {
						uniqueAntiNodes[utils.GetMapKey(nodeAx, nodeAy)] = antiNode.NodeA
					}
					if utils.PositionIsNotOutOfBounds(nodeBx, nodeBy, outerX, outerY) {
						uniqueAntiNodes[utils.GetMapKey(nodeBx, nodeBy)] = antiNode.NodeB
					}
				}

			}
		}

	}

	fmt.Println(len(uniqueAntiNodes))

}
