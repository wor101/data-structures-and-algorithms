package main

import "fmt"

func fib(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	_, prs := memo[n]
	if !prs {
		memo[n] = fib(n-2, memo) + fib(n-1, memo)
	}

	return memo[n]
}

func main() {
	fmt.Println(fib(6, make(map[int]int)))

}
