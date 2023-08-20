package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Print(largestRectangleArea(heights))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	var s []int
	max := 0
	s = append(s, 0)
	for i := 1; i < n; i++ {
		index := s[len(s)-1]

		if heights[index] <= heights[i] {
			s = append(s, i)
		} else {
			s = s[:len(s)-1]
			for len(s) > 0 && heights[index] > heights[i] {
				max = MAX(max, heights[index]*(i-index))
				index = s[len(s)-1]
				s = s[:len(s)-1]
			}
			s = append(s, i)
		}

	}
	index := s[len(s)-1]
	for len(s) > 0 {
		max = MAX(max, heights[index]*(n-index))
		if len(s) > 1 {
			s = s[:len(s)-1]
			index = s[len(s)-1]
		} else {
			s = nil
		}
	}
	return max
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}
