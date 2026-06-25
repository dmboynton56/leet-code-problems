package main

import "fmt"

// solution finds two indices whose values sum to target.
// Go: explicit types, map[int]int, and the comma-ok idiom for map lookups.
// Python equivalent: dict with `if complement in seen` (no ok tuple).
func solution(nums []int, target int) []int {
	seen := make(map[int]int) // Go maps must be initialized with make() or a literal; nil maps panic on write.
	for i, num := range nums {
		// range returns (index, value). Python: for i, num in enumerate(nums)
		complement := target - num
		if j, ok := seen[complement]; ok {
			// Go has no ternary; return a slice literal. Python: return [seen[complement], i]
			return []int{j, i}
		}
		seen[num] = i
	}
	return []int{} // Go returns zero-value slice; Python might return [] or raise.
}

func main() {
	fmt.Println(solution([]int{2, 7, 11, 15}, 9))  // [0 1]
	fmt.Println(solution([]int{3, 2, 4}, 6))         // [1 2]
	fmt.Println(solution([]int{3, 3}, 6))            // [0 1]
}
