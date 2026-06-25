package main

import (
	"fmt"
	"math"
)

// solution returns max profit from one buy and one sell.
// Go: math.MaxInt for initial min; Python uses float('inf') or nums[0].
func solution(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := math.MaxInt // Go has no float('inf') for ints; use math.MaxInt as sentinel.
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit // if-init in for scope; Python: max(max_profit, price - min_price)
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(solution([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(solution([]int{7, 6, 4, 3, 1}))    // 0
}
