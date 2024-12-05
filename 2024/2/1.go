package main

import (
	"advent-of-code/2024/2/utils"
	"fmt"
)

func main() {
	data := utils.GetInput()
	safeReports := 0

	for _, report := range data {
		if report[0] == report[1] {
			continue
		}

		if utils.IsRowSafe(report) {
			safeReports++
		}

	}

	fmt.Print(safeReports)
}
