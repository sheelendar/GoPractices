package main

import "fmt"

func main() {
	//grid := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	grid := [][]byte{
		{1, 1, 1, 1, 1, 0, 1, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 0, 0, 1, 0, 1},
		{1, 0, 0, 1, 1, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 1, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 1, 1, 1, 1, 0},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	M := len(grid)
	N := len(grid[0])

	if M == 1 && N == 1 {
		if grid[0][0] == 1 {
			return 1
		}
		return 0
	}

	res := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == byte(1) {
				res++
				DFS(grid, i, j, M, N)
			}
		}
	}
	return res
}

func DFS(grid [][]byte, row, col, M, N int) {
	if grid[row][col] == byte(0) {
		return
	}
	rowArr := []int{-1, 0, 0, 1}
	colArr := []int{0, -1, 1, 0}

	if grid[row][col] == byte(1) {
		grid[row][col] = byte(0)
	}

	for i := 0; i < 4; i++ {
		newRow := row + rowArr[i]
		newCol := col + colArr[i]
		if check(newRow, newCol, M, N) {
			DFS(grid, newRow, newCol, M, N)
		}
	}
}

func check(row, col, M, N int) bool {
	if row >= 0 && row < M && col >= 0 && col < N {
		return true
	}
	return false
}
