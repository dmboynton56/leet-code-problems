package main

import (
	"fmt"
	"slices"
)

func reverse(nums []int, k int) {
	n := len(nums)
	k = k % n

	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func main() {
	nums := []int{1, 2, 4, 5, 6}
	reverse(nums, 2)

	fmt.Println(nums)
}
