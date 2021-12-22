package main

import "fmt"

func duplicates(slice []string) (dup string) {
	stringMap := make(map[string]bool)

	for _, str := range slice {
		if stringMap[str] {
			dup = str
			break
		} else {
			stringMap[str] = true
		}
	}
	return dup
}

func main() {
	fmt.Println(duplicates([]string{"a", "b", "c", "d", "c", "e", "f"}))
}
