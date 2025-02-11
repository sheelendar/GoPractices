package main

import "fmt"

// Simple function to check if a number is even
func isEven(num int) bool {
	return num%2 == 0
}

func main() {
	fmt.Println(isEven(10)) // Output: true
}
