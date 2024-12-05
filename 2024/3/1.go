package main

import (
	"advent-of-code/2024/3/utils"
	"fmt"
)

func main() {
	data := utils.GetInput()

	answer := utils.GetSumOfMatches(data)
	
	fmt.Println(answer)
}
