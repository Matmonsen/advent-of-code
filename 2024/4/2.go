package main

import (
	"advent-of-code/2024/4/utils"
	"fmt"
)

func DiagonalXmas(topLeft string, topRight string, bottomLeft string, bottomRight string, y int, x int, data [][]string) bool {
        upLeft := y-1 >= 0 && x-1 >= 0 && data[y-1][x-1] == topLeft
        downRight := y+1 < len(data) && x+1 < len(data[0]) && data[y+1][x+1] == bottomRight

        upRight := y-1 >= 0 && x+1 < len(data[0]) && data[y-1][x+1] == topRight
        downLeft := y+1 < len(data) && x-1 >= 0 && data[y+1][x-1] == bottomLeft

        return upLeft && downRight && upRight && downLeft
}

func main() {
        data := utils.GetInput(

        occurences := 0

        for y, row := range data {
                for x, c := range row {

                        if c == "A" {
                                if      DiagonalXmas("M", "S", "M", "S", y, x, data) || 
                                        DiagonalXmas("S", "M", "S", "M", y, x, data) || 
                                        DiagonalXmas("S", "S", "M", "M", y, x, data) || 
                                        DiagonalXmas("M", "M", "S", "S", y, x, data)  {
                                        occurences++

                                }
                        }
                }
        }

        fmt.Println(occurences)

}