package main

import (
	"advent-of-code/2024/2/utils"
	"fmt"
)

func RemoveAtIndex(i int, line []int) int[] {
	row := make([]int, len(line)-1)
	copy(row, line[:i])
	copy(row[i:], line[i+1:])
	return row
}

func main() {
	data := utils.GetInput()
	safeReports := 0

	for _, line := range data {

		if !utils.IsRowSafe(line) {
			for i, _ := range line {
				row := RemoveAtIndex(i, line)

				if utils.IsRowSafe(row) {
					safeReports++
					break
				}
			}
		} else {
			safeReports++
		}

	}

	fmt.Print(safeReports)
}
