package main

import "fmt"

// solution checks whether brackets are properly opened and closed.
// Go: iterate runes (Unicode code points), not bytes — important for non-ASCII.
// Python: iterating a str also yields single-character strings; here ASCII-only is fine.
func solution(s string) bool {
	stack := []rune{} // Go slice used as stack; Python uses list with append/pop.
	for _, char := range s {
		// range over string yields runes in Go; `char` is rune (int32), not byte.
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char) // append returns new slice; reassign is idiomatic.
		} else {
			if len(stack) == 0 {
				return false // Go bool is explicit; no truthiness on empty collections like Python.
			}
			top := stack[len(stack)-1] // Go: no negative indexing; stack[len(stack)-1] is last element.
			if (char == ')' && top != '(') ||
				(char == ']' && top != '[') ||
				(char == '}' && top != '{') {
				return false
			}
			stack = stack[:len(stack)-1] // slice reslice pops; Python: stack.pop()
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(solution("()"))       // true
	fmt.Println(solution("()[]{}"))   // true
	fmt.Println(solution("(]"))       // false
	fmt.Println(solution("([)]"))    // false
	fmt.Println(solution("{[]}"))     // true
}
