package main

import (
	"fmt"
	"sort"
)

// solution returns all unique triplets that sum to zero.
// Go: sort.Slice with less func; Python: nums.sort() or sorted().
func solution(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate anchors; Python same guard
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	fmt.Println(solution([]int{-1, 0, 1, 2, -1, -4}))
	// [[-1 -1 2] [-1 0 1]]
}
