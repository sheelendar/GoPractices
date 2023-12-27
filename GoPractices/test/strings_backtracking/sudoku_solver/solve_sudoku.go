package main

import "fmt"

/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

	Each of the digits 1-9 must occur exactly once in each row.
	Each of the digits 1-9 must occur exactly once in each column.
	Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.

The '.' character indicates empty cells.
*/
func main() {
	b := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	solveSudoku(b)

}
func solveSudoku(board [][]byte) {
	board, _ = helper(board, 0, 0, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
	//fmt.Println(board)
}
func helper(board [][]byte, row, col, n int) ([][]byte, bool) {
	if row == n {
		return board, true
	}
	newRow := row
	newCol := col
	if col == n-1 {
		newRow = row + 1
		newCol = 0
	} else {
		newCol = col + 1
	}
	if (board)[row][col] != '.' {
		return helper(board, newRow, newCol, n)
	} else {
		for i := '1'; i <= '9'; i++ {
			if isValidNum(board, byte(i), row, col, n) {
				(board)[row][col] = byte(i)
				if b, ok := helper(board, newRow, newCol, n); ok {
					return b, ok
				}
				(board)[row][col] = '.'
			}
		}
	}
	return board, false
}
func isValidNum(board [][]byte, digit byte, row, col, n int) bool {
	for i := 0; i < n; i++ {
		if board[row][i] == digit {
			return false
		}
	}
	for i := 0; i < n; i++ {
		if board[i][col] == digit {
			return false
		}
	}
	nRow := (row / 3) * 3
	nCol := (col / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[nRow+i][nCol+j] == digit {
				return false
			}
		}
	}
	return true
}
