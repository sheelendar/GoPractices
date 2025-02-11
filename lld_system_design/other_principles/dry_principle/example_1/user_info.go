package main

import "fmt"

func printUserInfo1(name string, age int) {
	fmt.Println("User Name:", name)
	fmt.Println("User Age:", age)
}

func printUserInfo2(name string, age int) {
	fmt.Println("User Name:", name)
	fmt.Println("User Age:", age)
}

func main() {
	printUserInfo1("Alice", 25)
	printUserInfo2("Bob", 30)
}
