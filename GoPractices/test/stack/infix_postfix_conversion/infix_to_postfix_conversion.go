package main

import (
	"fmt"
	"regexp"
)

/*
Write a program to convert an Infix expression to Postfix form.
Examples:

	Input: A + B * C + D
	Output: ABC*+D+
*/

func main() {
	prefix := "a+b*(c^d-e)^(f+g*h)-i" // output : abcd^e-fgh*+^*+i-
	size := len(prefix)
	fmt.Print(convertIntoPostfix(prefix, size))
}

func convertIntoPostfix(prefix string, size int) string {
	var stack []byte
	postfix := ""
	for i := 0; i < size; i++ {
		ch := prefix[i]

		if isAlphaNumeric(ch) {
			postfix = postfix + string(ch)
		} else if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' {
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				op := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				postfix = postfix + string(op)
			}
			stack = stack[:len(stack)-1]
		} else {
			for len(stack) > 0 && operatorPriority(ch) <= operatorPriority(stack[len(stack)-1]) {
				op := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				postfix = postfix + string(op)
			}
			stack = append(stack, ch)
		}
	}
	for len(stack) > 0 {
		op := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		postfix = postfix + string(op)
	}
	return postfix
}
func isAlphaNumeric(byte byte) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(string(byte))
}
func operatorPriority(ch byte) int {
	switch ch {
	case '+':
		return 1
	case '-':
		return 1
	case '*':
		return 2
	case '/':
		return 2
	case '^':
		return 3
	case '%':
		return 3
	case '(':
		return 0
	default:
		return -1
	}
}
