package main

import "fmt"

// solution returns length of longest consecutive elements sequence.
func solution(nums []int) int {
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		seen[n] = struct{}{}
	}

	maxLen := 0
	for num := range seen {
		if _, ok := seen[num-1]; ok {
			continue
		}
		length := 1
		for curr := num + 1; ; curr++ {
			if _, ok := seen[curr]; !ok {
				break
			}
			length++
		}
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

func main() {
	fmt.Println(solution([]int{100, 4, 200, 1, 3, 2})) // 4
	fmt.Println(solution([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9
}
