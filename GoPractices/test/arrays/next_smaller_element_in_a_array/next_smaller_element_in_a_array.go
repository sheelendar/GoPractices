package main

import "fmt"

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
	arr := []int{11, 13, 21, 3}
	nextSmallerElement(arr)
}

func nextSmallerElement(arr []int) {
	size := len(arr)
	s := Stack{}
	i := 0
	for i < size {
		if s.IsEmpty() || int(s.Top()) <= arr[i] {
			s.Push(int64(arr[i]))
		} else {
			for !s.IsEmpty() || int(s.Top()) > arr[i] {
				fmt.Println(s.Top(), "-> ", arr[i])
				s.Pop()
			}
			s.Push(int64(arr[i]))
		}
		i++
	}
	for !s.IsEmpty() {
		fmt.Println(s.Top(), "-> ", -1)
		s.Pop()
	}
}
