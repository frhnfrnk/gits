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

	// Memproses setiap karakter dalam string
	for _, char := range input {
		// Jika karakter adalah tanda buka
		if strings.ContainsRune("({[", char) {
			stack = append(stack, char)
		} else if strings.ContainsRune(")}]", char) {
			// Jika karakter adalah tanda tutup
			if len(stack) == 0 || bracketMap[stack[len(stack)-1]] != char {
				// Jika tidak ada tanda buka yang cocok atau tanda buka tidak sesuai
				return "NO"
			}
			// Menghapus tanda buka yang cocok dari stack
			stack = stack[:len(stack)-1]
		}
	}

	// Jika setelah memproses semua karakter, stack masih tidak kosong
	if len(stack) > 0 {
		return "NO"
	}

	// Jika berhasil melewati semua kondisi, berarti tanda kurung seimbang
	return "YES"
}

func main() {
	// Contoh penggunaan
	fmt.Println(isBalanced("{[({}()(){})]}")) // Output: YES
}
