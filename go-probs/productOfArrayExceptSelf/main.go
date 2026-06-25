package main

import "fmt"

// solution returns product of all elements except self without division.
// Go: two passes with output slice; prefix/suffix products.
func solution(nums []int) []int {
	n := len(nums)
	result := make([]int, n) // make allocates zeroed slice; Python [1]*n then overwrite

	// Forward pass: result[i] = product of nums[0..i-1]
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Backward pass: multiply by suffix product on the fly
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}
	return result
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4})) // [24 12 8 6]
	fmt.Println(solution([]int{-1, 1, 0, -3, 3}))
}
