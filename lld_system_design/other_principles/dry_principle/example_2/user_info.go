package main

import "fmt"

// Single reusable function
func printUserInfo(name string, age int) {
	fmt.Println("User Name:", name)
	fmt.Println("User Age:", age)
}

func main() {
	printUserInfo("Alice", 25)
	printUserInfo("Bob", 30)
}
