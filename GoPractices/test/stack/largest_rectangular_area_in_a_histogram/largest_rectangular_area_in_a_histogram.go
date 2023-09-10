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
func main() {
	arr := []int{6, 2, 5, 4, 5, 1, 6}
	fmt.Println(largestArea(arr, len(arr)))
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
	var s []int
	i := 0
	for i < n {
		if len(s) == 0 || hist[s[len(s)-1]] <= hist[i] {
			s = append(s, i)
			i++
		} else {
			k := s[len(s)-1]
			s = s[:len(s)-1]
			if len(s) == 0 {
				max_area = hist[k] * i
			} else {
				max_area = hist[k] * (i - s[len(s)-1] - 1)
			}
			max_max_area = int(math.Max(float64(max_area), float64(max_max_area)))
		}
	}
	for len(s) > 0 {
		k := s[len(s)-1]
		s = s[:len(s)-1]
		if len(s) == 0 {
			max_area = hist[k] * i
		} else {
			max_area = hist[k] * (i - s[len(s)-1] - 1)
		}
		max_max_area = int(math.Max(float64(max_area), float64(max_max_area)))
	}
	return max_max_area
}
