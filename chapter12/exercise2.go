package main

import "fmt"

func golomb(n int, memo map[int]int) int {
	fmt.Println("RECURSION")
	if n == 1 {
		return 1
	}

	_, prs := memo[n]

	if !prs {
		memo[n] = 1 + golomb(n-golomb(golomb(n-1, memo), memo), memo)
	}

	return memo[n]
}

func main() {

	fmt.Println(golomb(20, make(map[int]int)))
}
