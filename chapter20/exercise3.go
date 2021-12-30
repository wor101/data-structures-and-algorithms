package main

import "fmt"

func maxBuySell(prices []int) (profit int) {
	profit = prices[1] - prices[0]
	low := prices[0]

	for i, price := range prices {
		if price < low && profit <= 0 && i < len(prices)-1 {
			low = price
		}

		if price-low > profit {
			profit = price - low
		}
	}

	return profit
}

func main() {
	fmt.Println(maxBuySell([]int{10, 7, 5, 8, 11, 2, 6}))
}
