package main

import (
	"math"
)

type Stack struct {
	stack []int
}

func (s *Stack) Push(item int) {
	s.stack = append(s.stack, item)
}
func (s *Stack) Pop() int {
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
func (s *Stack) Top() int {
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
	arr := []int{7, 1, 5, 3, 6, 4}

	print("max Profit: ", maxProfit(arr))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	s := Stack{}
	s.Push(prices[0])
	max := 0

	for i := 1; i < n; i++ {
		if s.Top() > prices[i] {
			for !s.IsEmpty() && s.Top() > prices[i] {
				s.Pop()
			}
			s.Push(prices[i])
		}
		max = int(math.Max(float64(max), float64(prices[i]-s.Top())))
	}
	return max
}
