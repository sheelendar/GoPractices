package main

import "fmt"

/*
Sudoku Solver

Write the function sudokuSolve that checks whether a given sudoku board (i.e. sudoku puzzle) is solvable. If so, the function will returns true. Otherwise (i.e. there is no valid solution to the given sudoku board), returns false.

In sudoku, the objective is to fill a 9x9 board with digits so that each column, each row, and each of the nine 3x3 sub-boards that compose the board contains all of the digits from 1 to 9. The board setter provides a partially completed board, which for a well-posed board has a unique solution. As explained above, for this problem, it suffices to calculate whether a given sudoku board has a solution. No need to return the actual numbers that make up a solution.

A sudoku board is represented as a two-dimensional 9x9 array of the characters ‘1’,‘2’,…,‘9’ and the '.' character, which represents a blank space. The function should fill the blank spaces with characters such that the following rules apply:

    In every row of the array, all characters ‘1’,‘2’,…,‘9’ appear exactly once.
    In every column of the array, all characters ‘1’,‘2’,…,‘9’ appear exactly once.
    In every 3x3 sub-board that is illustrated below, all characters ‘1’,‘2’,…,‘9’ appear exactly once.

A solved sudoku is a board with no blank spaces, i.e. all blank spaces are filled with characters that abide to the constraints above. If the function succeeds in solving the sudoku board, it’ll return true (false, otherwise).

Example (more examples can be found here):

*/
//https://www.youtube.com/watch?v=uyetDh-DyDg
//https://www.pramp.com/challenge/O5PGrqGEyKtq9wpgw6XP

func main() {
	board := [][]string{
		{".", ".", ".", "7", ".", ".", "3", ".", "1"},
		{"3", ".", ".", "9", ".", ".", ".", ".", "."},
		{".", "4", ".", "3", "1", ".", "2", ".", "."},
		{".", "6", ".", "4", ".", ".", "5", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "1", ".", ".", "8", ".", "4", "."},
		{".", ".", "6", ".", "2", "1", ".", "5", "."},
		{".", ".", ".", ".", ".", "9", ".", ".", "8"},
		{"8", ".", "5", ".", ".", "4", ".", ".", "."}}

	fmt.Println(SudokuSolve(board))
}

func SudokuSolve(board [][]string) bool {
	// your code goes here
	size := len(board)
	if size < 9 {
		return false
	}
	// start from 0,0 and track all element till size-1,size-1
	return SudokuHeler(board, 0, 0, size)

}

func SudokuHeler(board [][]string, row, col, size int) bool {
	// check we reached row as size then sedoku is competed with all values
	if row >= size {
		return true
	}
	// create new row and col for next iteration
	ri := 0
	cj := 0
	//check if column is at end then some to next row and set column as 0
	if col == size-1 {
		ri = row + 1
		cj = 0
	} else {
		// if column is less then the size then just increase column with 1 and row will be same.
		ri = row
		cj = col + 1
	}

	res := false
	// if we don't have dot(.) then have some value then go ahead
	if board[row][col] != "." {
		res = SudokuHeler(board, ri, cj, size)
	} else {
		for i := '1'; i <= '9'; i++ {
			// check into loop which value can set into cell from 1 to 9.
			// isValid function check same value is not there into same row , same col, 3*3 sum matirx.
			if isValid(board, row, col, size, string(i)) {
				// put that value into cell and call again for next iteration
				board[row][col] = string(i)
				if SudokuHeler(board, ri, cj, size) {
					// if we get true then just return true
					res = true
					return res
				}
				// if we don't find any solution then just check another number. and fill same cell with dot(.) again.
				board[row][col] = "."
			}
		}
	}
	return res
}

func isValid(board [][]string, row, col, size int, ch string) bool {
	// check same row value should not be there.
	for i := 0; i < len(board[0]); i++ {
		if board[row][i] == ch {
			return false
		}
	}
	// check same column value should be there.
	for i := 0; i < size; i++ {
		if board[i][col] == ch {
			return false
		}
	}
	// check 3* 3 matrix values should not be there.
	x := (row / 3) * 3
	y := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[x+i][y+j] == ch {
				return false
			}
		}
	}
	return true
}
