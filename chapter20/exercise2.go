package main

import "fmt"

func missingNumber(array []int) int {
	var hash map[int]bool = make(map[int]bool)
	max := array[0]

	for _, num := range array {
		hash[num] = true

		if num > max {
			max = num
		}
	}

	for i := 0; i < max; i++ {
		if _, ok := hash[i]; !ok {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(missingNumber([]int{2, 3, 0, 6, 1, 5}))
	fmt.Println(missingNumber([]int{8, 2, 3, 9, 4, 7, 5, 0, 6}))
}
