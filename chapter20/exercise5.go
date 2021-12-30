package main

import "fmt"

func orderTemps(temperatures []float64) []int {
	tempHash := make(map[int]int)

	for _, tempFloat := range temperatures {
		var temp int = int(tempFloat * 10)

		if _, ok := tempHash[temp]; ok {
			tempHash[temp] += 1
		} else {
			tempHash[temp] = 1
		}
	}

	orderedTemps := make([]int, 0)

	fmt.Println(tempHash)

	for i := 970; i <= 980; i++ {
		fmt.Println("I:", i)
		fmt.Println("tempHash:", tempHash[i])
		if value, ok := tempHash[i]; ok {
			fmt.Println("In hash")
			fmt.Println("Value:", tempHash[i])
			for j := 0; j < value; j++ {
				orderedTemps = append(orderedTemps, i)
			}
		}
		fmt.Println(orderedTemps)
		fmt.Println("")
	}

	return orderedTemps
}

func main() {
	fmt.Println(orderTemps([]float64{97.0, 97.4, 98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1}))
}
