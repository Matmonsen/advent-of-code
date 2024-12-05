package main

import (
	"advent-of-code/2024/5/utils"
	"fmt"
)

func main() {
	validationMap, reports := utils.GetInput()

	sum := 0
	for _, report := range reports {
		
		if utils.RowIsCorrectOrder(report, validationMap) {
			sum += report[len(report) / 2]
		}
	}

	fmt.Println(sum)
	
}
