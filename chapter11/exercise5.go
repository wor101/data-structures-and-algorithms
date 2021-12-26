package main

import "fmt"

func uniquePaths(rows, columns int) int {
	if rows == 1 || columns == 1 {
		return 1
	}
	return uniquePaths((rows-1), columns) + uniquePaths(rows, (columns-1))
}

func main() {
	fmt.Println("Unique Paths")
	fmt.Println(uniquePaths(2, 2))
	fmt.Println(uniquePaths(2, 3))
	fmt.Println(uniquePaths(2, 7))
	fmt.Println(uniquePaths(3, 7))

}
