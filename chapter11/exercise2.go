package main

import "fmt"

func justEvenNumbers(sliceOfNumbers []int, sliceOfEvens []int) []int {

	if len(sliceOfNumbers) <= 0 {
		return sliceOfEvens
	}

	if sliceOfNumbers[0]%2 == 0 {
		sliceOfEvens = append(sliceOfEvens, sliceOfNumbers[0])
	}
	return justEvenNumbers(sliceOfNumbers[1:], sliceOfEvens)
}

func main() {
	fmt.Println("Accept an array of #s and return new array containing just event #s")

	fmt.Println(justEvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}))
}
