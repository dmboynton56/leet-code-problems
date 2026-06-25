package main

import "fmt"

// solution counts distinct ways to climb n stairs (1 or 2 steps at a time).
// Go: iterative DP with two ints; Python often uses a loop or @lru_cache recursion.
func solution(n int) int {
	if n <= 2 {
		return n // base cases: 1 way for n=1, 2 ways for n=2
	}

	prev2, prev1 := 1, 2 // fibonacci-style: ways(n-2), ways(n-1)
	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr // rotate window; Python: a, b = b, a + b
	}
	return prev1
}

func main() {
	fmt.Println(solution(2)) // 2
	fmt.Println(solution(3)) // 3
	fmt.Println(solution(5)) // 8
}
