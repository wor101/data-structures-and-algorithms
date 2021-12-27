package main

import "fmt"

func add_until_100(array []int) int {
	if len(array) == 0 {
		return 0
	}

	remaining := add_until_100(array[1:])

	if array[0]+remaining > 100 {
		return remaining
	} else {
		return array[0] + remaining
	}
}

func main() {
	fmt.Println(add_until_100([]int{10, 25, 75, 15}))
}
