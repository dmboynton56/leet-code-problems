package main

import "fmt"

// solution returns length of longest substring without repeating characters.
// Go: map[rune]int for last index; Python dict works similarly.
func solution(s string) int {
	lastSeen := make(map[rune]int)
	maxLen := 0
	start := 0 // left bound of sliding window (inclusive index in rune slice sense)

	for i, char := range s {
		// range index i is byte offset for ASCII but use map on rune; for ASCII s, i aligns with rune index.
		if prev, ok := lastSeen[char]; ok && prev >= start {
			start = prev + 1 // shrink window from left; Python same logic
		}
		lastSeen[char] = i
		if window := i - start + 1; window > maxLen {
			maxLen = window
		}
	}
	return maxLen
}

func main() {
	fmt.Println(solution("abcabcbb")) // 3
	fmt.Println(solution("bbbbb"))     // 1
	fmt.Println(solution("pwwkew"))   // 3
}
