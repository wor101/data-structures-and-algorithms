package main

import "fmt"

func triangularNumber(n int) int {
	if n == 1 {
		return 1
	}
	return n + triangularNumber(n-1)
}

func main() {
	fmt.Println("Triangular Numbers")
	fmt.Println(triangularNumber(7))
}
