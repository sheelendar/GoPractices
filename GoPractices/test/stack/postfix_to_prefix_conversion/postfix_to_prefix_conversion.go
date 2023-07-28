package main

import "fmt"

/*
Postfix: An expression is called the postfix expression if the operator appears in the expression after the operands.
Simply of the form (operand1 operand2 operator).
Example : AB+CD-* (Infix : (A+B) * (C-D) )

Prefix : An expression is called the prefix expression if the operator appears in the expression
before the operands. Simply of the form (operator operand1 operand2).
Example : *+AB-CD (Infix : (A+B) * (C-D) )

Conversion of Postfix expression directly to Prefix without going through the process of converting them
first to Infix and then to Prefix is much better in terms of computation and better understanding the expression
*/

type Stack struct {
	stack []string
}

func (s *Stack) Push(item string) {
	s.stack = append(s.stack, item)
}
func (s *Stack) Pop() string {
	size := len(s.stack)
	if size > 0 {
		temp := s.stack[size-1]
		s.stack = s.stack[0 : size-1]
		if size == 1 {
			s.stack = nil
		}
		return temp
	}
	return ""
}
func (s *Stack) Top() string {
	size := len(s.stack)
	if size > 0 {
		return s.stack[size-1]
	}
	return ""
}
func (s *Stack) IsEmpty() bool {
	size := len(s.stack)
	if size <= 0 {
		return true
	}
	return false
}
func main() {
	post_exp := "ABC/-AK/L-*"

	// Function call
	fmt.Println("Prefix :", postToPre(post_exp))
}

func postToPre(postfix string) string {
	if len(postfix) == 0 {
		return ""
	}
	prefix := ""
	s := Stack{}
	s.Push(string(postfix[0]))
	for i := 1; i < len(postfix); i++ {
		if isOperand(postfix[i]) {
			s.Push(string(postfix[i]))
		} else {
			a := s.Pop()
			b := s.Pop()
			prefix = fmt.Sprintf("%s%s%s", string(postfix[i]), b, a)
			s.Push(prefix)
		}
	}
	return prefix
}
func isOperand(ch byte) bool {
	if ch >= 'a' && ch < 'z' || ch >= 'A' && ch < 'Z' {
		return true
	}
	return false
}
