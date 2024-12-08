package main

import (
	"advent-of-code/2024/8/utils"
	"fmt"
)

func main() {
	_, antennas, outerX, outerY := utils.GetInput()
	uniqueAntiNodes := make(map[string]utils.Item)

	for key := range antennas {
		for _, location := range antennas[key] {
			antiNodes := utils.FindAntiNodes(location, antennas[key])

			for _, antiNode := range antiNodes {

				if utils.PositionIsNotOutOfBounds(antiNode.NodeA.Location.X, antiNode.NodeA.Location.Y, outerX, outerY) {
					uniqueAntiNodes[utils.GetMapKey(antiNode.NodeA.Location.X, antiNode.NodeA.Location.Y)] = antiNode.NodeA
				}
				if utils.PositionIsNotOutOfBounds(antiNode.NodeB.Location.X, antiNode.NodeB.Location.Y, outerX, outerY) {
					uniqueAntiNodes[utils.GetMapKey(antiNode.NodeB.Location.X, antiNode.NodeB.Location.Y)] = antiNode.NodeB
				}
			}
		}

	}

	fmt.Println(len(uniqueAntiNodes))

}
