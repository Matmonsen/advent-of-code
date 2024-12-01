package main

import (
	"advent-of-code/2024/1/utils"
	"fmt"
)

func main() {
	left, right := utils.GetInput()

	sum := 0
	dict:= make(map[int]int)

	for i := 0; i < len(right); i++ {
		dict[right[i]] = dict[right[i]] + 1
	}

	for i := 0; i < len(right); i++ {
		sum += dict[left[i]] * left[i]
	}

	fmt.Print(sum)
}