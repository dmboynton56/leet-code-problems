package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isAlnum(b byte) bool {
	r := rune(b)
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// solution returns true if s is a palindrome considering only alphanumeric chars.
func solution(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlnum(s[left]) {
			left++
		}
		for left < right && !isAlnum(s[right]) {
			right--
		}
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(solution("A man, a plan, a canal: Panama")) // true
	fmt.Println(solution("race a car"))                     // false
	fmt.Println(solution(" "))                              // true
}
