package main

import "fmt"

/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [

	["1","1","1","1","0"],
	["1","1","0","1","0"],
	["1","1","0","0","0"],
	["0","0","0","0","0"]

]
Output: 1

Example 2:
Input: grid = [

	["1","1","0","0","0"],
	["1","1","0","0","0"],
	["0","0","1","0","0"],
	["0","0","0","1","1"]

]
Output: 3
*/
func main() {
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1}}

	m := len(grid)
	n := len(grid[0])
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				countIslands(grid, i, j, m, n)
				count++
			}
		}
	}
	fmt.Println(count)
}

func countIslands(grid [][]int, i, j, m, n int) {
	grid[i][j] = 0

	if checkIndex(i-1, j, m, n, grid) {
		countIslands(grid, i-1, j, m, n)
	}
	if checkIndex(i+1, j, m, n, grid) {
		countIslands(grid, i+1, j, m, n)
	}
	if checkIndex(i, j-1, m, n, grid) {
		countIslands(grid, i, j-1, m, n)
	}
	if checkIndex(i, j+1, m, n, grid) {
		countIslands(grid, i, j+1, m, n)
	}
}

func checkIndex(row, col, m, n int, grid [][]int) bool {
	if row < m && row >= 0 && col < n && col >= 0 && grid[row][col] == 1 {
		return true
	}
	return false
}
