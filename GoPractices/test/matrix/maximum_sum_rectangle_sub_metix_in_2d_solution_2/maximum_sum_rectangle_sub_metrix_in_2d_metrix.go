package main

import (
	"fmt"
)

/*
Given a 2D array, find the maximum sum submatrix in it. For example, in the following 2D array,
the maximum sum submatrix is highlighted with blue rectangle and sum of all elements in this submatrix is 29.
https://www.geeksforgeeks.org/maximum-sum-rectangle-in-a-2d-matrix-dp-27/

	{ { 1, 2, -1, -4, -20 },
	  { -8, -3, 4, 2, 1 },
	  { 3, 8, 10, 1, 3 },
	  { -4, -1, 1, 7, -6 } };

Solution: (Top, Left) (1, 1)
(Bottom, Right) (3, 3)
Max sum is: 29
*/
func main() {
	metrix := [][]int{{1, 2, -1, -4, -20},
		{-8, -3, 4, 2, 1},
		{3, 8, 10, 1, 3},
		{-4, -1, 1, 7, -6}}
	maximum_sub_metrix(metrix, len(metrix), len(metrix[0]))
}

func maximum_sub_metrix(metrix [][]int, m, n int) {
	maxSum := -1
	rowS := 0
	rowE := 0
	colS := 0
	colE := 0
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			metrix[i][j] += metrix[i-1][j]
		}
	}
	for k := 0; k < m; k++ {
		for i := k; i < m; i++ {
			curr := 0
			colStart := 0
			for j := 0; j < n; j++ {
				item := metrix[i][j]
				if k > 0 {
					item -= metrix[k-1][j]
				}
				curr = Max(item, curr+item)
				if curr > maxSum {
					maxSum = curr
					rowE = i
					colE = j
					rowS = k
					colS = colStart
				} else if curr < 0 {
					colStart = j + 1
					curr = 0
				}
			}
		}
	}
	fmt.Println("Start row, col", rowS, colS)
	fmt.Println("End row, col", rowE, colE)
	fmt.Println("max sum", maxSum)

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
