package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetInput() [][]int {
	fileData, _ := os.ReadFile("input.txt")
	stringArr := strings.Split(string(fileData), "\n")

	data := make([][]int, int(len(stringArr)))

	for i := 0; i < len(stringArr); i++ {
		line := strings.Split(stringArr[i], " ")

		row := make([]int, int(len(line)))
		for j := 0; j < len(line); j++ {
			row[j], _ = strconv.Atoi(line[j])
		}
		data[i] = row
	}

	return data
}

const SAFE = 1
const UNSAFE = 0

const ASC = 1
const DESC = -1

func GetDirection(a int, b int) int {
	if a > b {
		return DESC
	}

	return ASC
}

func IsPairSafe(a int, b int, direction int) bool {
	abs := a - b

	if direction == ASC {
		abs = b - a
	}

	return abs > 0 && abs < 4
}

func IsRowSafe(line []int) bool {
	isSafe := true

	if line[0] == line[1] {
		return false
	}

	dir := GetDirection(line[0], line[1])

	for i := 0; i < len(line)-1; i++ {
		if !IsPairSafe(line[i], line[i+1], dir) {
			isSafe = false
			break
		}
	}
	return isSafe
}
