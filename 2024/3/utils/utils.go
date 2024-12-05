package utils

import (
	"os"
	"strconv"
	"regexp"
)

func GetInput() string {
	fileData, _ := os.ReadFile("input.txt")

	return string(fileData)
}

func GetSumOfMatches(data string) int {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := regex.FindAllStringSubmatch(data, -1)

	sum := 0
	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		sum += left * right
	}

	return sum
}