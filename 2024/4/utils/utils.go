package utils

import (
	"os"
	"strings"
)

func GetInput() [][]string {
        fileData, _ := os.ReadFile("input.txt")
        stringArr := strings.Split(string(fileData), "\n")
        data := make([][]string, int(len(stringArr)))

        for i, arr := range stringArr {
                row := make([]string, int(len(arr)))
                for j, c := range strings.Split(arr, "") {
                        row[j] = c
                }

                data[i] = row
        }

        return data
}