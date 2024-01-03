package main

import (
	"fmt"
)

func A000124(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = ((i * (i + 1)) / 2) + 1
	}
	return result
}

func main() {
	var inputNumber int
	fmt.Print("Masukkan nilai n: ")
	_, err := fmt.Scan(&inputNumber)

	if err != nil || inputNumber < 1 {
		fmt.Println("Masukkan nilai n harus lebih dari atau sama dengan 1.")
		return
	}

	sequence := A000124(inputNumber)
	output := ""
	for i, num := range sequence {
		output += fmt.Sprint(num)
		if i < len(sequence)-1 {
			output += "-"
		}
	}
	fmt.Printf("Output: %s\n", output)
}
