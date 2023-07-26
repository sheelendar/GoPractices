package main

type Stack struct {
	stack []int64
}

func (s *Stack) Push(item int64) {
	s.stack = append(s.stack, item)
}
func (s *Stack) Pop() int64 {
	size := len(s.stack)
	if size > 0 {
		temp := s.stack[size-1]
		s.stack = s.stack[0 : size-1]
		if size == 1 {
			s.stack = nil
		}
		return temp
	}
	return -1
}
func (s *Stack) Top() int64 {
	size := len(s.stack)
	if size > 0 {
		return s.stack[size-1]
	}
	return -1
}
func (s *Stack) IsEmpty() bool {
	size := len(s.stack)
	if size <= 0 {
		return true
	}
	return false
}

func main() {
	stack := &Stack{stack: make([]int64, 0)}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	println(stack.stack)
	println(stack.IsEmpty())
	println(stack.Top())
	println(stack.Pop())
	println(stack.Pop())
	println(stack.Pop())
	println(stack.IsEmpty())
	stack.Push(50)
	println(stack.Top())
	println(stack.Pop())
	println(stack.IsEmpty())
	println(stack.stack)
}
