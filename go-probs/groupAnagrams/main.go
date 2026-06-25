package main

import (
	"fmt"
	"sort"
)

// solution groups strings that are anagrams of each other.
// Go: map[string][]string with sorted runes as key; Python uses sorted tuple or char count key.
func solution(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := anagramKey(s)
		groups[key] = append(groups[key], s) // append creates slice if key new
	}

	result := make([][]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}
	return result
}

func anagramKey(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes) // Python: ''.join(sorted(s))
}

func main() {
	fmt.Println(solution([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
