package main

import "fmt"

func convRoman(numerals string) int {
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	integer := 0

	for _, ch := range numerals {
		integer += values[ch]
	}

	return integer
}

func main() {
	fmt.Println(convRoman("XXVII"))
}
