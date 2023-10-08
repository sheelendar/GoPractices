package main

import "fmt"

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

Input: n = 1
Output: ["()"]
*/

func main() {
	n := 3
	printAllParentheses("", 0, 0, n*2)
}
func printAllParentheses(str string, i, j, n int) {
	if len(str) == n {
		fmt.Println(str)
		return
	}
	if i < n/2 {
		newStr := fmt.Sprintf("%s%s", str, "{")
		printAllParentheses(newStr, i+1, j, n)
	}
	if j < i {
		newStr := fmt.Sprintf("%s%s", str, "}")
		printAllParentheses(newStr, i, j+1, n)
	}
}
