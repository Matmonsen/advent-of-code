package main

import (
	"advent-of-code/2024/5/utils"
	"fmt"
	"slices"
)

func NumberPlacementIsLegal(prevNumbers []int, number int, validationMap map[int][]int) int {
	requiredNumbers, exists := validationMap[number]
	if !exists {
		return -1
	}

	notValidPlacement := -1
	for i, r := range prevNumbers {
		if slices.Contains(requiredNumbers, r) {
			notValidPlacement = i
		}
	}

	return notValidPlacement
}

func main() {
	validationMap, reports := utils.GetInput()

	sum := 0
	for _, report := range reports {
		if !utils.RowIsCorrectOrder(report, validationMap) {
			for i := len(report) - 1; i > -1; i-- {

				num := report[i]
				placement := NumberPlacementIsLegal(report[0:i], num, validationMap)
				if placement > -1 {
					for j := i; j >= placement; j-- {
						if j-1 >= placement {
							report[j] = report[j-1]
						}
					}
					report[placement] = num
					i++
				}
			}

			sum += report[len(report)/2]
		}
	}

	fmt.Println(sum)
}
