package main

import (
	"advent-of-code/2024/4/utils"
	"fmt"
)

func OccuresForward(x int, row []string) bool {
        return x + 3 < len(row) && row[x + 1] == "M" && row[x + 2] == "A" && row[x + 3] == "S"
}

func OccuresBackwards(x int, row []string) bool {
        return x-3 >= 0 && row[x-1] == "M" && row[x-2] == "A" && row[x-3] == "S"
}

func OccursUpwards(y int, x int, data [][]string) bool {
        return y-3 >= 0 && data[y-1][x] == "M" && data[y-2][x] == "A" && data[y-3][x] == "S"
}
func OccursDownwards(y int, x int, data [][]string) bool {
        return y+3 < len(data) && data[y+1][x] == "M" && data[y+2][x] == "A" && data[y+3][x] == "S"
}

func OccursDiagonallyUpRight(y int, x int, data [][]string) bool {
        return y-3 >= 0 && x+3 < len(data[0]) && data[y-1][x+1] == "M" && data[y-2][x+2] == "A" && data[y-3][x+3] == "S"
}

func OccursDiagonallyUpLeft(y int, x int, data [][]string) bool {
        return y-3 >= 0 && x-3 >= 0 && data[y-1][x-1] == "M" && data[y-2][x-2] == "A" && data[y-3][x-3] == "S"
}

func OccursDiagonallyDownLeft(y int, x int, data [][]string) bool {
        return y+3 < len(data) && x-3 >= 0 && data[y+1][x-1] == "M" && data[y+2][x-2] == "A" && data[y+3][x-3] == "S"
}

func OccursDiagonallyDownRight(y int, x int, data [][]string) bool {
        return y+3 < len(data) && x+3 < len(data[0]) && data[y+1][x+1] == "M" && data[y+2][x+2] == "A" && data[y+3][x+3] == "S"
}

func main() {
        data := utils.GetInput()

        occurences := 0

        for y, row := range data {
                for x, c := range row {

                        if c == "X" {
                                if OccuresForward(x, row) {
                                        occurences++
                                } 
                                if OccuresBackwards(x, row) {
                                        occurences++
                                } 
                                if OccursUpwards(y, x, data) {
                                        occurences++
                                } 
                                if OccursDownwards(y, x, data) {
                                        occurences++
                                } 
                                if OccursDiagonallyDownRight(y, x, data) {
                                        occurences++
                                } 
                                if OccursDiagonallyDownLeft(y, x, data) {
                                        occurences++
                                } 
                                if OccursDiagonallyUpRight(y, x, data) {
                                        occurences++
                                } 
                                if OccursDiagonallyUpLeft(y, x, data) {
                                        occurences++
                                }
                        }
                }
        }

        fmt.Println(occurences)
}