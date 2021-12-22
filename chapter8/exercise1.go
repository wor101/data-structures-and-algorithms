package main

import "fmt"

func intersection(slice1 []int, slice2 []int) (intersect []int) {
	var short []int
	var long []int
	longMap := make(map[int]bool)

	if len(slice1) <= len(slice2) {
		short = slice1[:]
		long = slice2[:]

	} else {
		short = slice2[:]
		long = slice1[:]
	}

	for _, element := range long {
		longMap[element] = true
	}

	for _, element := range short {
		if longMap[element] {
			intersect = append(intersect, element)
		}
	}

	// intersect = append(short, long...)
	return intersect
}

func main() {
	fmt.Println(intersection([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(intersection([]int{1, 2, 3, 4, 5}, []int{4, 5, 6}))
}
