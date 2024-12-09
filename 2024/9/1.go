package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput() []int {
	fileData, _ := os.ReadFile("input.txt")
	numbers := strings.Split(string(fileData), "")

	blocks := make([]int, 0)
	id := 0

	for i := 0; i < len(numbers); i += 2 {
		fileLength, _ := strconv.Atoi(numbers[i])
		freeSpace := 0
		if i+1 < len(numbers) {
			freeSpace, _ = strconv.Atoi(numbers[i+1])
		}

		for i := 0; i < fileLength; i++ {
			blocks = append(blocks, id)
		}

		for i := 0; i < freeSpace; i++ {
			blocks = append(blocks, -1)
		}

		id++
	}

	return blocks
}

func main() {
	blocks := GetInput()

	checksum := int64(0)
	lastIndex := len(blocks) - 1
	for i := 0; i < len(blocks) && i < lastIndex; i++ {
		if blocks[i] == -1 {
			last := blocks[lastIndex]
			blocks[i] = last
			blocks[lastIndex] = -1

			lastIndex--
			for lastIndex > i && blocks[lastIndex] == -1 {
				lastIndex--
			}
		}
	}
	for i := 0; i <= lastIndex; i++ {
		checksum += int64(blocks[i] * i)
	}

	fmt.Println(checksum)
}
