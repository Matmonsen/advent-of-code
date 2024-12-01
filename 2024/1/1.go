package main

import (
	"advent-of-code/2024/1/utils"
	"fmt"
    "slices"
	"math"
)

func main() {
	left, right := utils.GetInput()

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	
	for i := 0; i < len(right); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}


	fmt.Print(sum)
}