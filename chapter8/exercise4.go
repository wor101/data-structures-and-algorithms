package main

import "fmt"

func nonDuplicate(inputString string) string {
	var nonDup string
	var inputMap = make(map[string]int)

	for _, char := range inputString {
		if inputMap[string(char)] >= 1 {
			inputMap[string(char)] += 1
		} else {
			inputMap[string(char)] = 1
		}
	}

	for key, value := range inputMap {
		fmt.Printf("Key: %s, Value: %v", key, value)
		fmt.Println("")

		if value == 1 {
			nonDup = key
			break
		}
	}
	return nonDup
}

func main() {
	fmt.Println(nonDuplicate("minimum"))
}
