package main

import (
	"advent-of-code/2024/3/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data := utils.GetInput()

	regex := regexp.MustCompile(`(?s)don['']?t\(\).*?do\(\)`)

	for _, match := range regex.FindAllString(data, -1) {
		data = strings.Replace(data, match, "", 1)
	}

	fmt.Println(utils.GetSumOfMatches(data))
}
