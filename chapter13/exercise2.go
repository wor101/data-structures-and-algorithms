package main

import (
	"fmt"
	"sort"
)

func findMissingNumber(array []int) int {
	sort.Ints(array)
	var missingNumber int

	for index, value := range array {
		if index == 0 {
			continue
		}

		if value != array[index-1]+1 {
			missingNumber = array[index-1] + 1
		}
	}

	return missingNumber
}

func main() {
	fmt.Println(findMissingNumber([]int{9, 3, 2, 5, 6, 7, 1, 0, 4}))
}
