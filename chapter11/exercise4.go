package main

import "fmt"

func indexOfX(inputString string, index int) int {
	if string(inputString[index]) == "x" {
		return index
	}
	return indexOfX(inputString, index+1)
}

func main() {
	fmt.Println("First Index of X")
	fmt.Println(indexOfX("abcdefghijklmnopqrstuvwxyz", 0))
}
