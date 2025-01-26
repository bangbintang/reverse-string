package main

import (
	"fmt"
)

func reverseString(s string) string {
	reverses := []rune(s)
	n := len(reverses)
	for i := 0; i < n/2; i++ {
		reverses[i], reverses[n-1-i] = reverses[n-1-i], reverses[i]
	}
	return string(reverses)
}

func main() {
	// Contoh penggunaan fungsi reverseString
	input := "Hello, World!"
	reversed := reverseString(input)
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Reversed: %s\n", reversed)
}
