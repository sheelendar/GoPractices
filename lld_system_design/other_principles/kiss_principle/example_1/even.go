package main

import "fmt"

// Overcomplicated function to check if a number is even
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isEven(10)) // Output: true
}
