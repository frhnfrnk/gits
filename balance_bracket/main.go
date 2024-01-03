package main

import (
	"fmt"
	"strings"
)

func isBalanced(input string) string {
	stack := make([]rune, 0)

	bracketMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range input {
		if strings.ContainsRune("({[", char) {
			stack = append(stack, char)
		} else if strings.ContainsRune(")}]", char) {
			if len(stack) == 0 || bracketMap[stack[len(stack)-1]] != char {
				// Jika tidak ada tanda buka yang cocok atau tanda buka tidak sesuai
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return "NO"
	}
	return "YES"
}

func main() {
	fmt.Println(isBalanced(" { [ ( ) ] } "))
	fmt.Println(isBalanced(" { [ ( ] ) } "))
	fmt.Println(isBalanced(" { ( ( [ ] ) [ ] ) [ ] }  "))
}
