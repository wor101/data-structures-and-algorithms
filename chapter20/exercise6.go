package main

import "fmt"

func longestSequence(array []int) int {
	numberHash := make(map[int]bool)
	min := array[0]
	max := array[0]
	longest := 0

	// fill hash
	for _, num := range array {
		numberHash[num] = true
	}

	// find the min and max
	for _, num := range array {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Println(numberHash)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	for _, num := range array {
		//fmt.Println("Num:", num)
		if num < max {
			for i := num + 1; i <= max; i++ {
				if _, ok := numberHash[i]; ok {
					fmt.Println("Search:", i)
					fmt.Println("Current Lenght:", (i-num)+1)
					if (i-num)+1 > longest {
						longest = (i - num) + 1
					}

				} else {
					break
				}
			}
		}
	}

	return longest
}

func main() {
	fmt.Println(longestSequence([]int{19, 13, 15, 12, 18, 14, 17, 11}))
}
