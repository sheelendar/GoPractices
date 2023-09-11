package main

import "fmt"

/*
Given a 2D array(m x n). The task is to check if there is any path from top left to bottom right. In the matrix, -1
is considered as blockage (can’t go through this cell) and 0 is considered path cell (can go through it).

Note: Top left cell always contains 0
Examples:

	Input : arr[][] = {{ 0, 0, 0, -1, 0},
	{-1, 0, 0, -1, -1},
	{ 0, 0, 0, -1, 0},
	{-1, 0, 0, 0, 0},
	{ 0, 0, -1, 0, 0}}

Output : Yes
*/
func main() {
	arr := [][]int{{0, 0, 0, -1, 0},
		{-1, 0, 0, -1, -1},
		{0, 0, 0, -1, 0},
		{-1, 0, 0, 0, 0},
		{0, 0, -1, 0, 0}} // output : true

	arr2 := [][]int{{0, 0, 0, -1, 0},
		{-1, 0, 0, -1, -1},
		{0, 0, 0, -1, 0},
		{-1, 0, -1, 0, -1},
		{0, 0, -1, 0, 0}} // output: false
	fmt.Println(isPossiblePath(arr))
	fmt.Println(isPossiblePath(arr2))
}

func isPossiblePath(arr [][]int) bool {
	m := len(arr)
	if m == 0 {
		return true
	}
	n := len(arr[0])
	if n == 0 {
		return true
	}
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		if arr[i][0] == 0 {
			dp[i][0] = true
		}
	}
	for j := 0; j < n; j++ {
		if arr[0][j] == 0 {
			dp[0][j] = true
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if arr[i][j] == 0 && (dp[i][j-1] || dp[i-1][j]) {
				dp[i][j] = true
			}
		}
	}
	return dp[m-1][n-1]
}
