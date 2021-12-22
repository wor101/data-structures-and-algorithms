package main

import (
	"fmt"
)

func twoNumberProducts(numbers []int) []int {
	var products []int

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			products = append(products, (numbers[i] * numbers[j]))
		}
	}
	return products
}

func main() {
	fmt.Println(twoNumberProducts([]int{1, 2, 3, 4, 5}))

}
