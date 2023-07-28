package main

import "fmt"

/*
Discuss(210+)
Courses
Practice

Given a binary 2D matrix, find the number of islands. A group of connected 1s forms an island. For example, the below matrix contains 5 islands.

Example:

	Input: mat[][] = {{1, 1, 0, 0, 0},
	                           {0, 1, 0, 0, 1},
	                           {1, 0, 0, 1, 1},
	                          {0, 0, 0, 0, 0},
	                         {1, 0, 1, 0, 0}}
	Output: 4
*/

func main() {
	/*mat := [][]int{
	{1, 1, 0, 0, 0},
	{0, 1, 0, 0, 1},
	{1, 0, 0, 1, 1},
	{0, 0, 0, 0, 0},
	{1, 0, 1, 0, 0}}*/

	mat := [][]int{{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1}}
	fmt.Println("number of islands: ", findIslands(mat))
}

func findIslands(mat [][]int) int {
	Row := len(mat)
	Col := len(mat[0])
	if Row == 0 {
		return 0
	}
	count := 0
	for i := 0; i < Row; i++ {
		for j := 0; j < Col; j++ {
			if mat[i][j] == 1 {
				checkUP(mat, i, j, Row, Col)
				count++
			}
		}
	}
	return count
}

func checkUP(mat [][]int, i, j, row, col int) {

	if i < 0 || i >= row {
		return
	}
	if j < 0 || j >= col {
		return
	}
	if mat[i][j] == 0 {
		return
	}
	mat[i][j] = 0
	checkUP(mat, i-1, j-1, row, col)
	checkUP(mat, i-1, j, row, col)
	checkUP(mat, i-1, j+1, row, col)
	checkUP(mat, i, j-1, row, col)
	checkUP(mat, i, j+1, row, col)
	checkUP(mat, i+1, j-1, row, col)
	checkUP(mat, i+1, j, row, col)
	checkUP(mat, i+1, j+1, row, col)
}
