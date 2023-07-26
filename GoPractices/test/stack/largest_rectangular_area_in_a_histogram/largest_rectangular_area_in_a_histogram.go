package main

import (
	"fmt"
	"math"
)

/*
Find the largest rectangular area possible in a given histogram where the largest rectangle can be made of a number of contiguous
bars whose heights are given in an array. For simplicity, assume that all bars have the same width and the width is 1 unit.
Example:

	Input: histogram = {6, 2, 5, 4, 5, 1, 6} Output: 12
*/

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
	hist := []int{6, 2, 5, 4, 5, 1, 6}
	fmt.Print("Largest Area: ", largestArea(hist, len(hist)))
}

func largestArea(hist []int, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return hist[0]
	}
	max_max_area := 0
	max_area := 0
	s := Stack{}
	i := 0
	for i < n {
		if s.IsEmpty() || hist[s.Top()] <= hist[i] {
			s.Push(int64(i))
			i++
		} else {
			k := s.Pop()
			if s.IsEmpty() {
				max_area = hist[k] * i
			} else {
				max_area = hist[k] * (i - int(s.Top()) - 1)
			}
			max_max_area = int(math.Max(float64(max_area), float64(max_max_area)))
		}
	}
	for !s.IsEmpty() {
		k := s.Pop()
		if s.IsEmpty() {
			max_area = hist[k] * i
		} else {
			max_area = hist[k] * (i - int(s.Top()) - 1)
		}
		max_max_area = int(math.Max(float64(max_area), float64(max_max_area)))
	}
	return max_max_area
}
