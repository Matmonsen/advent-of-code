package main

import (
	"advent-of-code/2024/8/utils"
	"fmt"
)

func main() {
	inputPuzzle, antennas := utils.GetInput()
	uniqueAntiNodes := make(map[string]utils.Location)

	for key := range antennas {
		for _, location := range antennas[key] {
			antiNodes := utils.FindAntiNodes(location, antennas[key])

			for _, antiNode := range antiNodes {
				if antiNode.NodeA.Y > -1 && antiNode.NodeA.X > -1 && antiNode.NodeA.Y < len(inputPuzzle) && antiNode.NodeA.X < len(inputPuzzle[0]) {
					inputPuzzle[antiNode.NodeA.Y][antiNode.NodeA.X] = antiNode.NodeA.Symbol
					uniqueAntiNodes[antiNode.NodeA.Key] = antiNode.NodeA
				}
				if antiNode.NodeB.Y > -1 && antiNode.NodeB.X > -1 && antiNode.NodeB.Y < len(inputPuzzle) && antiNode.NodeB.X < len(inputPuzzle[0]) {
					inputPuzzle[antiNode.NodeB.Y][antiNode.NodeB.X] = antiNode.NodeB.Symbol
					uniqueAntiNodes[antiNode.NodeB.Key] = antiNode.NodeB
				}
			}
		}

	}

	fmt.Println(len(uniqueAntiNodes))

}
