package main

import (
	"fmt"
	"sort"
)

func greatestON2(array []int) int {
	var greatest int = array[0]
	// fmt.Println("RECURSION")

	if len(array) == 1 {
		return greatest
	}

	if greatest > greatestON2(array[1:]) {
		return greatest
	}
	return greatestON2(array[1:])
}

func greatestON(array []int) int {
	var greatest int = array[0]

	for _, value := range array {
		if value > greatest {
			greatest = value
		}
	}

	return greatest
}

func greatestONlogN(array []int) int {
	sort.Ints(array)
	return array[len(array)-1]
}

func main() {
	fmt.Println("ON2:", greatestON2([]int{1, 2, 30, 14, 17, 4}))
	fmt.Println("ON:", greatestON([]int{1, 2, 30, 14, 17, 4}))
	fmt.Println("ONlogN:", greatestONlogN([]int{1, 2, 30, 14, 17, 4}))
}
