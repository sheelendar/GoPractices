package main

import "fmt"

// https://leetcode.com/problems/pacific-atlantic-water-flow/
/*
There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.

*/
func main() {
	mat := [][]int{{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4}}
	result := pacificAtlantic(mat)
	for i := 0; i < len(result); i++ {
		fmt.Print("{")
		fmt.Print(result[i][0])
		fmt.Print(" ")
		fmt.Print(result[i][1])
		fmt.Print("}")
		fmt.Println()
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	M := len(heights)
	N := len(heights[0])
	// declare two array to check water flow.
	// assign memory for both
	pacific := make([][]bool, M)
	atlantic := make([][]bool, M)
	for i := 0; i < M; i++ {
		pacific[i] = make([]bool, N)
		atlantic[i] = make([]bool, N)
	}

	// calculate DFS for first column and last row. because water can flow from both side.
	for i := 0; i < N; i++ {
		DFS(pacific, heights, 0, i, -1, M, N)
		DFS(atlantic, heights, M-1, i, -1, M, N)
	}

	// calculate DFS for first column and last column. because water can flow from both side.
	for i := 0; i < M; i++ {
		DFS(pacific, heights, i, 0, -1, M, N)
		DFS(atlantic, heights, i, N-1, -1, M, N)
	}

	/*display(pacific, M, N)
	fmt.Println()
	fmt.Println()
	display(atlantic, M, N)*/
	var result [][]int

	//now we got highest height in both of the array for left and right so where both array are true for i,j is
	//highest point and water can flow both the side from here.
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if atlantic[i][j] && pacific[i][j] {
				result = append(result, []int{i, j})
			}
		}
		fmt.Println()
	}

	return result
}

func display(pacific [][]bool, m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(pacific[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func DFS(dp [][]bool, heights [][]int, row int, col int, pre, M, N int) {
	if row >= 0 && col >= 0 && row < M && col < N && !dp[row][col] && heights[row][col] >= pre {
		dp[row][col] = true
		DFS(dp, heights, row-1, col, heights[row][col], M, N)
		DFS(dp, heights, row+1, col, heights[row][col], M, N)
		DFS(dp, heights, row, col-1, heights[row][col], M, N)
		DFS(dp, heights, row, col+1, heights[row][col], M, N)
	}
}
