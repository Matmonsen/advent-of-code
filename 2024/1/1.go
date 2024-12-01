package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
    "slices"
	"math"
)

func main() {
    fileData, _ := os.ReadFile("input.txt")
	stringArr := strings.Split(string(fileData), "\n")

	left := make([]int, int(len(stringArr)))
	right := make([]int, int(len(stringArr)))

	for i := 0; i < len(stringArr); i++ {
		row := strings.Split(stringArr[i], "   ")
		
		left[i], _ = strconv.Atoi(row[0])
		right[i], _ = strconv.Atoi(row[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	
	for i := 0; i < len(stringArr); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}


	fmt.Print(sum)
}