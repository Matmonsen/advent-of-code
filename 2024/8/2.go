package main

import (
	"advent-of-code/2024/8/utils"
	"fmt"
	"math"
)

func main() {
	inputPuzzle, antennas := utils.GetInput()
	uniqueAntiNodes := make(map[string]utils.Location)
	outerX := len(inputPuzzle[0])
	outerY := len(inputPuzzle)

	length := make([]int, int(math.Max(float64(len(inputPuzzle)), float64(len(inputPuzzle[0])))))

	for key := range antennas {
		for _, location := range antennas[key] {
			antiNodes := utils.FindAntiNodes(location, antennas[key])

			for _, antiNode := range antiNodes {

				uniqueAntiNodes[antiNode.AntennaB.Key] = antiNode.NodeB
				uniqueAntiNodes[antiNode.AntennaA.Key] = antiNode.NodeA
				inputPuzzle[antiNode.AntennaA.Y][antiNode.AntennaA.X] = antiNode.NodeA.Symbol
				inputPuzzle[antiNode.AntennaB.Y][antiNode.AntennaB.X] = antiNode.NodeA.Symbol
				for i := range length {

					nodeAy := antiNode.NodeA.Y - (i * antiNode.Distance.Y)
					nodeAx := antiNode.NodeA.X - (i * antiNode.Distance.X)
					nodeBy := antiNode.NodeB.Y + (i * antiNode.Distance.Y)
					nodeBx := antiNode.NodeB.X + (i * antiNode.Distance.X)

					if nodeAy > -1 && nodeAx > -1 && nodeAy < outerY && nodeAx < outerX {
						inputPuzzle[nodeAy][nodeAx] = antiNode.NodeA.Symbol
						uniqueAntiNodes[utils.GetMapKey(nodeAx, nodeAy)] = antiNode.NodeA
					}
					if nodeBy > -1 && nodeBx > -1 && nodeBy < outerY && nodeBx < outerX {
						inputPuzzle[nodeBy][nodeBx] = antiNode.NodeB.Symbol
						uniqueAntiNodes[utils.GetMapKey(nodeBx, nodeBy)] = antiNode.NodeB
					}
				}

			}
		}

	}
	for _, row := range inputPuzzle {
		fmt.Println(row)
	}

	fmt.Println(len(uniqueAntiNodes))

}
