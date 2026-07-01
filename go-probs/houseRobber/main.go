package main

import "fmt"

// solution returns max money robbing non-adjacent houses.
func solution(nums []int) int {
	prev, curr := 0, 0
	for _, num := range nums {
		prev, curr = curr, max(curr, prev+num)
	}
	return curr
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 1}))   // 4
	fmt.Println(solution([]int{2, 7, 9, 3, 1})) // 12
}
