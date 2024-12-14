package main

import (
	"advent-of-code/2024/11/utils"
	"fmt"
	"strconv"
)

func main() {
	data := utils.GetInput()

	for blinks := 0; blinks < 75; blinks++ {
		next := make([]int64, 0)

		for _, item := range data {
			if item == 0 {
				next = append(next, 1)
			} else {
				s := strconv.FormatInt(item, 10)

				if len(s)%2 == 0 {

					length := len(s) / 2
					first, _ := strconv.ParseInt(s[:length], 10, 64)
					second, _ := strconv.ParseInt(s[length:], 10, 64)

					next = append(next, first, second)
				} else {
					next = append(next, item*2024)
				}
			}
		}
		data = next
	}
	// TODO lagre int as bytes istedet?
	// https://stackoverflow.com/questions/52920032/what-is-the-best-way-to-store-a-byte-array-in-memory
	// out of memory med 75 iterasjoner
	fmt.Println(len(data))
}
