package utils

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func NumIsInCorrectOrder(num int, previousNumbers []int, validationMap map[int][]int) bool {
	nums := validationMap[num]
	for _, prevNum := range previousNumbers {
		if slices.Contains(nums, prevNum) {
			return false
		}

	}

	return true
}

func RowIsCorrectOrder(report []int, validationMap map[int][]int) bool {
	reportIsCorrect := true
	for i, num := range report {
		reportIsCorrect = reportIsCorrect && NumIsInCorrectOrder(num, report[0:i], validationMap)
	}

	return reportIsCorrect
}

func GetInput() (map[int][]int, [][]int) {
	fileData, _ := os.ReadFile("input.txt")
	sections := strings.Split(string(fileData), "\n\n")

	validation := strings.Split(sections[0], "\n")
	input := strings.Split(sections[1], "\n")

	m := make(map[int][]int)

	for _, line := range validation {
		splittedLine := strings.Split(line, "|")
		first, _ := strconv.Atoi(splittedLine[0])
		sec, _ := strconv.Atoi(splittedLine[1])

		_, exists := m[first]
		if !exists {
			m[first] = make([]int, 0)
		}

		m[first] = append(m[first], sec)
	}

	reports := make([][]int, len(input))
	for lineIndex, line := range input {
		splittedLine := strings.Split(line, ",")
		row := make([]int, len(splittedLine))
		for i, num := range splittedLine {
			row[i], _ = strconv.Atoi(num)
		}
		reports[lineIndex] = row
	}
	return m, reports
}
