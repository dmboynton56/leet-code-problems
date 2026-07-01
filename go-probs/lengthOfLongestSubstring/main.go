package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	l := 0
	curr := 0

	lastSeen := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		if idx, ok := lastSeen[s[r]]; ok && idx >= l {
			l = lastSeen[s[r]] + 1
			lastSeen[s[r]] = r
			curr = r - l + 1
		} else {
			lastSeen[s[r]] = r
			curr += 1
		}
		if curr > result {
			result = curr
		}
	}
	return result
}

func lengthOfLongestSubstring1(s string) int {
	maxLength := 0
	tail := 0
	sub := ""
	for i := range s {
		if index := strings.Index(sub, string(s[i])); index != -1 {
			tail += index + 1
		}
		sub = s[tail : i+1]
		maxLength = max(maxLength, len(sub))
	}
	return maxLength
}

func main() {
	fmt.Println("Test: " + strconv.Itoa(lengthOfLongestSubstring("bbbbb")))
}
