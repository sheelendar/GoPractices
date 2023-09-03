package main

import "fmt"

func main() {
	//heights := []int{2, 1, 5, 6, 2, 3}
	heights := []int{1, 1}
	//heights := []int{4, 2}
	fmt.Print(largestRectangleArea2(heights))
}

func largestRectangleArea2(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	mL := make([]int, n)
	mR := make([]int, n)
	mL[0] = -1
	var stack []int
	stack = append(stack, 0)
	for i := 1; i < n; i++ {
		if len(stack) > 0 && heights[stack[len(stack)-1]] < heights[i] {
			mL[i] = stack[len(stack)-1]
		} else {
			for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				mL[i] = stack[len(stack)-1]
			} else {
				mL[i] = -1
			}
		}
		stack = append(stack, i)
	}
	stack = nil
	stack = append(stack, n-1)
	mR[n-1] = n
	for i := n - 2; i >= 0; i-- {
		if len(stack) > 0 && heights[stack[len(stack)-1]] < heights[i] {
			mR[i] = stack[len(stack)-1]
		} else {
			for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				mR[i] = stack[len(stack)-1]
			} else {
				mR[i] = n
			}
		}
		stack = append(stack, i)
	}
	max := 0
	for i := 0; i < n; i++ {
		max = MAX(max, (mR[i]-mL[i]-1)*heights[i])
	}
	return max
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Bar struct {
	Index  int
	Height int
}

func largestRectangleArea(heights []int) int {
	ret := 0
	size := len(heights)
	stack := []Bar{Bar{-1, 0}}

	for i := 0; i < size; i++ {
		last := len(stack) - 1
		height := stack[last].Height

		if heights[i] < height {
			width := i - stack[last-1].Index - 1

			ret = MAX(ret, height*width)
			stack = stack[:last]
			i--
		} else if heights[i] > height {
			stack = append(stack, Bar{i, heights[i]})
		} else {
			stack[last].Index = i
		}
	}

	for i := len(stack) - 1; i > 0; i-- {
		height := stack[i].Height
		width := size - stack[i-1].Index - 1

		ret = MAX(ret, height*width)
	}

	return ret
}
