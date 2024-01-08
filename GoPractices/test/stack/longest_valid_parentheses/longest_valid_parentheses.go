/*
Given a string containing just the characters '(' and ')',return the length of the longest valid (well-formed) parentheses Substring.
https://leetcode.com/problems/longest-valid-parentheses/description/

Example 1 :
Input: s = "(()" ,
Output: 2,
Explanation: The longest valid parentheses substring is "()".

Example 2:,
Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
*/
package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

func longestValidParentheses(s string) int {
	n := len(s)
	var stack []int
	left := -1
	res := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			size := len(stack)
			if size > 0 {
				stack = stack[0 : size-1]
			} else {
				left = i
			}
			if size-1 <= 0 {
				res = Max(res, i-left)
			} else {
				res = Max(res, i-stack[size-2])
			}
		}
	}
	return res
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
