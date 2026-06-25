package main

import "fmt"

// solution returns the largest sum of any contiguous subarray (Kadane's algorithm).
// Go: track running sum and global max; Python's max() is built-in.
func solution(nums []int) int {
	maxSum := nums[0]   // assume non-empty per problem
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// extend current subarray or start fresh at nums[i]
		if currentSum+nums[i] > nums[i] {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(solution([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(solution([]int{1}))                                 // 1
	fmt.Println(solution([]int{5, 4, -1, 7, 8}))                  // 23
}
