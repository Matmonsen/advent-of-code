package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Id     int
	Length int
	Index  int
}

func GetInput() ([]int, []File) {
	fileData, _ := os.ReadFile("input.txt")
	numbers := strings.Split(string(fileData), "")

	blocks := make([]int, 0)
	fileStack := make([]File, 0)

	id := 0

	for i := 0; i < len(numbers); i += 2 {
		fileLength, _ := strconv.Atoi(numbers[i])
		freeSpace := 0
		if i+1 < len(numbers) {
			freeSpace, _ = strconv.Atoi(numbers[i+1])
		}

		fileStack = append(fileStack, File{id, fileLength, len(blocks)})

		for i := 0; i < fileLength; i++ {
			blocks = append(blocks, id)
		}

		for i := 0; i < freeSpace; i++ {
			blocks = append(blocks, Space)
		}

		id++
	}
	return blocks, fileStack
}

func FindFirstFittingSpaceFromLeft(file File, blocks []int) int {
	spaceStartIndex := -1
	for i := 0; i < file.Index; i++ {
		if blocks[i] == Space {
			if spaceStartIndex == -1 {
				spaceStartIndex = i
			}
		} else {
			if spaceStartIndex > -1 {
				spaceSize := i - spaceStartIndex
				if spaceSize >= file.Length {
					fmt.Println(blocks)
					return spaceStartIndex
				} else {
					spaceStartIndex = -1
				}
			}
		}
	}
	return -1
}

const Space = -1

func main() {
	blocks, fileStack := GetInput()

	for f := len(fileStack) - 1; f >= 0; f-- {
		file := fileStack[f]
		spaceStartIndex := FindFirstFittingSpaceFromLeft(file, blocks)
		if spaceStartIndex != -1 {
			for space := range make([]int, file.Length) {
				blocks[spaceStartIndex+space] = file.Id
				blocks[file.Index+space] = Space
			}
		}
	}

	checksum := int64(0)
	for i, val := range blocks {
		if val != -1 {
			checksum += int64(blocks[i] * i)
		}
	}
	fmt.Println(blocks)
	// I have not managed to solve this yet...
	// One edge case I have not yet found...
	// 6415163644414 Is too high...
	fmt.Println(checksum)
}
