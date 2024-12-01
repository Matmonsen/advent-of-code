package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
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

	sum := 0
	dict:= make(map[int]int)
	
	for i := 0; i < len(stringArr); i++ {
		dict[right[i]] = dict[right[i]] + 1
	}

	for i := 0; i < len(stringArr); i++ {
		sum += dict[left[i]] * left[i]
	}

	fmt.Print(sum)
}