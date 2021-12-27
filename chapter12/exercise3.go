package main

import "fmt"

func unique_paths(rows, columns int, memo map[[2]int]int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	_, prs := memo[[2]int{rows, columns}]

	if !prs {
		memo[[2]int{rows, columns}] = unique_paths(rows-1, columns, memo) + unique_paths(rows, columns-1, memo)
	}

	return memo[[2]int{rows, columns}]
}

func main() {
	fmt.Println(unique_paths(3, 7, make(map[[2]int]int)))

}
