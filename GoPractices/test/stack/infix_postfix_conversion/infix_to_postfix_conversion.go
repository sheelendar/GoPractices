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

type Stack struct {
	stack []byte
}

func (s *Stack) Push(item byte) {
	s.stack = append(s.stack, item)
}
func (s *Stack) Pop() byte {
	size := len(s.stack)
	if size > 0 {
		temp := s.stack[size-1]
		s.stack = s.stack[0 : size-1]
		if size == 1 {
			s.stack = nil
		}
		return temp
	}
	return '-'
}
func (s *Stack) Top() byte {
	size := len(s.stack)
	if size > 0 {
		return s.stack[size-1]
	}
	return '-'
}
func (s *Stack) IsEmpty() bool {
	size := len(s.stack)
	if size <= 0 {
		return true
	}
	return false
}
func main() {
	prefix := "a+b*(c^d-e)^(f+g*h)-i"
	size := len(prefix)
	fmt.Print(convertIntoPostfix(prefix, size))
}

func convertIntoPostfix(prefix string, size int) string {
	stack := Stack{}
	postfix := ""
	for i := 0; i < size; i++ {
		ch := prefix[i]

		if isAlphaNumeric(ch) {
			postfix = postfix + string(ch)
		} else if ch == '(' {
			stack.Push(ch)
		} else if ch == ')' {
			for !stack.IsEmpty() && stack.Top() != '(' {
				op := stack.Pop()
				postfix = postfix + string(op)
			}
			stack.Pop()
		} else {
			for !stack.IsEmpty() && operatorPriority(ch) <= operatorPriority(stack.Top()) {
				op := stack.Pop()
				postfix = postfix + string(op)
			}
			stack.Push(ch)
		}
	}
	for !stack.IsEmpty() {
		op := stack.Pop()
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
