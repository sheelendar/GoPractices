package main

import "fmt"

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
	exp := "ab*c+"
	fmt.Print(getInfixFromPostfix(exp))
}

func getInfixFromPostfix(exp string) string {
	size := len(exp)
	if size < 3 {
		return exp
	}
	prefix := ""
	stack := Stack{}
	for i := 0; i < len(exp); i++ {
		if isOprend(exp[i]) {
			stack.Push(string(exp[i]))
		} else {
			operator := string(exp[i])
			op2 := stack.Pop()
			op1 := stack.Pop()
			prefix = fmt.Sprintf("(%s%s%s)", op1, operator, op2)
			stack.Push(prefix)
		}
	}

	return prefix
}

func isOprend(u uint8) bool {
	if 'a' <= u && u <= 'z' || 'A' <= u && u <= 'Z' {
		return true
	}
	return false
}
