package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetInput() []int64 {
	fileData, _ := os.ReadFile("input.txt")
	items := strings.Split(string(fileData), " ")

	data := make([]int64, len(items))

	for i, item := range items {

		data[i], _ = strconv.ParseInt(item, 10, 64)
	}

	return data
}
