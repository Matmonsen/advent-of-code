package utils

import (
	"os"
	"strings"
    "strconv"
)

func GetInput() ([]int, []int) {
    fileData, _ := os.ReadFile("input.txt")
	stringArr := strings.Split(string(fileData), "\n")

	left := make([]int, int(len(stringArr)))
	right := make([]int, int(len(stringArr)))

	for i := 0; i < len(stringArr); i++ {
		row := strings.Split(stringArr[i], "   ")
		
		left[i], _ = strconv.Atoi(row[0])
		right[i], _ = strconv.Atoi(row[1])
	}

	return left, right
}
