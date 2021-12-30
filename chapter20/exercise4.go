package main

import "fmt"

func greatestProduct(array []int) int {
	product := array[0] * array[1]
	max := array[0]
	min := array[0]

	for index, num := range array {
		if index == 0 {
			continue
		}

		if num*max > product {
			product = num * max
		}

		if num*min > product {
			product = num * min
		}

		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}
	return product
}

func main() {
	fmt.Println(greatestProduct([]int{5, -10, -6, 9, 4}))
}
