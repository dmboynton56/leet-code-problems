package main

import "fmt"

// solution returns the maximum area between two vertical lines.
func solution(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		h := min(height[left], height[right])
		width := right - left
		area := h * width
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	fmt.Println(solution([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(solution([]int{1, 1}))                       // 1
}
