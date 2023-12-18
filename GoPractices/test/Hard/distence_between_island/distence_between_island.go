package main

import "fmt"

/*
 */
func main() {

	//grid := [][]int{{0, 1}, {1, 0}}
	grid := [][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}
	fmt.Print(shortestBridge(grid))
}

func shortestBridge(grid [][]int) int {
	N := len(grid)
	dp := make([][]bool, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]bool, N)
	}
	for i := 0; i < N; i++ {
		f := false
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 {
				helper(i, j, N, grid, dp)
				f = true
				break
			}
		}
		if f {
			break
		}
	}
	visited := make([][]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, N)
	}
	min := 999999
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if dp[i][j] {
				res := help(i, j, N, grid, dp, visited)
				if min > res {
					min = res
				}
			}
		}
	}
	return min

}

func help(i, j, N int, grid [][]int, dp, visited [][]bool) int {

	fmt.Print(" i= ", i)
	fmt.Print(" j=", j)
	fmt.Println()
	if !dp[i][j] && grid[i][j] == 1 {
		return 0
	}
	visited[i][j] = true
	min := 999999
	if checker(i, j-1, N) && !dp[i][j-1] && !visited[i][j-1] {
		min = Min(help(i, j-1, N, grid, dp, visited), min)
	}
	if checker(i, j+1, N) && !dp[i][j+1] && !visited[i][j+1] {
		min = Min(help(i, j+1, N, grid, dp, visited), min)
	}
	if checker(i-1, j, N) && !dp[i-1][j] && !visited[i-1][j] {
		min = Min(help(i-1, j, N, grid, dp, visited), min)
	}
	if checker(i+1, j, N) && !dp[i+1][j] && !visited[i+1][j] {
		min = Min(help(i+1, j, N, grid, dp, visited), min)
	}
	visited[i][j] = false
	return min + 1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func helper(i, j, N int, grid [][]int, dp [][]bool) {
	if grid[i][j] == 0 {
		return
	}
	dp[i][j] = true

	if checker(i, j-1, N) && !dp[i][j-1] && grid[i][j-1] == 1 {
		helper(i, j-1, N, grid, dp)
	}
	if checker(i, j+1, N) && !dp[i][j+1] && grid[i][j+1] == 1 {
		helper(i, j+1, N, grid, dp)
	}
	if checker(i-1, j, N) && !dp[i-1][j] && grid[i-1][j] == 1 {
		helper(i-1, j, N, grid, dp)
	}
	if checker(i+1, j, N) && !dp[i+1][j] && grid[i+1][j] == 1 {
		helper(i+1, j, N, grid, dp)
	}
}

func checker(i, j, N int) bool {
	if i >= 0 && i < N && j >= 0 && j < N {
		return true
	}
	return false
}
