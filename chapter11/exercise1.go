package main

import "fmt"

func totalCharacters(sliceOfStrings []string) int {
	if len(sliceOfStrings) <= 0 {
		return 0
	}
	currentString := sliceOfStrings[0]

	return len(currentString) + totalCharacters(sliceOfStrings[1:])
}

func main() {
	fmt.Println("A function that accepts an array of strings and returns the total number of characters across all the strings:")
	fmt.Println(totalCharacters([]string{"ab", "c", "def", "ghij"}))
}
