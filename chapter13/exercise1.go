package main

import (
	"fmt"
	"sort"
)

func greatestProduct(array []int) int {
	sort.Ints(array)

	return array[len(array)-1] * array[len(array)-2] * array[len(array)-3]
}

func main() {
	fmt.Println(greatestProduct([]int{1, 6, 3, 4, 5, 1}))
}
