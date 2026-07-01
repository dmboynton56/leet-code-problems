package main

import "fmt"

// solution finds target in a rotated sorted array, or -1.
func solution(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(solution([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(solution([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
	fmt.Println(solution([]int{1}, 0))                   // -1
}
