package main

import "fmt"

func reverseArray(array []int) []int {

	for i := len(array) - 1; i >= 0; i-- {
		toMove := array[i]
		array = append(array[:i], array[i+1:]...)
		array = append(array, toMove)
		fmt.Println(toMove)
		fmt.Println(array)
	}

	return array
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(reverseArray(array))
}
