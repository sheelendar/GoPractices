package main

import "fmt"

/*
Given a rows x cols binary matrix filled with 0's and 1's,
find the largest rectangle containing only 1's and return its area.
Input: matrix =
{{"1","0","1","0","0"},

	{"1","0","1","1","1"},
	{"1","1","1","1","1"},
	{"1","0","0","1","0"}}

Output: 6
Explanation: The maximal rectangle is shown in the above picture.

Example 2:
Input: matrix = [["0"]]
Output: 0

Example 3:
Input: matrix = [["1"]]
Output: 1
*/
func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			grid[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				grid[i][j] = 1 + grid[i-1][j]
			}
		}
	}
	max := 0
	for i := 0; i < m; i++ {
		max = Max(max, maxHistogram(grid[i]))
	}
	return max
}
func maxHistogram(arr []int) int {
	n := len(arr)
	i := 0
	var stack []int
	max := 0
	for i < n {
		if len(stack) == 0 || arr[stack[len(stack)-1]] <= arr[i] {
			stack = append(stack, i)
			i++
		} else {
			size := len(stack)
			index := stack[size-1]
			stack = stack[:size-1]
			size--
			if size == 0 {
				max = Max(max, i*arr[index])
			} else {
				max = Max(max, (i-stack[size-1]-1)*arr[index])
			}
		}
	}
	for len(stack) > 0 {
		size := len(stack)
		index := stack[size-1]
		stack = stack[:size-1]
		size--
		if size == 0 {
			max = Max(max, i*arr[index])
		} else {
			max = Max(max, (i-stack[size-1]-1)*arr[index])
		}
	}
	return max
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
