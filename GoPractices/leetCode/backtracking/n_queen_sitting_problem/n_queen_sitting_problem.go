package main

import "fmt"

/*
https://www.geeksforgeeks.org/n-queen-problem-backtracking-3/

The N Queen is the problem of placing N chess queens on an N×N chessboard so that no two queens attack each other.
For example, the following is a solution for the 4 Queen problem.

The expected output is in the form of a matrix that has ‘Q‘s for the blocks where queens are placed and the empty spaces are
represented by ‘.’ . For example, the following is the output matrix for the above 4-Queen solution.
*/

func main() {
	n := 4
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	var result [][]int
	solveNQueen(board, 0, n, make([]int, n), &result)
	fmt.Println("Number of solutoin ", result)
}

func solveNQueen(board [][]bool, row, n int, res []int, result *[][]int) {
	if row == n { // found solution here
		arr := make([]int, len(res))
		copy(arr, res)
		*result = append(*result, arr)
		return
	}
	for j := 0; j < n; j++ {
		if isQueenSafe(board, row, j, n) { // if queen is safe then go ahead
			board[row][j] = true // put true in board
			res[row] = j
			solveNQueen(board, row+1, n, res, result)
			board[row][j] = false // backtrack if when back from reverse call
		}
	}

}

func isQueenSafe(board [][]bool, row, col, n int) bool {
	for i := row - 1; i >= 0; i-- { // check column blow
		if board[i][col] {
			return false
		}
	}
	j := col - 1
	for i := row - 1; i >= 0 && j >= 0; i-- { // check diagonal from right to left
		if board[i][j] {
			return false
		}
		j--
	}
	j = col + 1
	for i := row - 1; i >= 0 && j < n; i-- { //  check other diagonal from left to right
		if board[i][j] {
			return false
		}
		j++
	}
	return true
}
