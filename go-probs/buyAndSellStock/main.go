package main

import (
	"fmt"
	"strconv"
)

func buySell(prices []int) int {
	maxProfit := 0
	l := 0
	for i := range prices {
		profit := prices[i] - prices[l]
		if profit > maxProfit {
			maxProfit = profit
		}
		if profit < 0 {
			l = i
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Test 1: " + strconv.Itoa(buySell(prices)))
}
