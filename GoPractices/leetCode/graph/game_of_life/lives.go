package main

import "fmt"

// https://leetcode.com/problems/game-of-life/
/*


According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

    Any live cell with fewer than two live neighbors dies as if caused by under-population.
    Any live cell with two or three live neighbors lives on to the next generation.
    Any live cell with more than three live neighbors dies, as if by over-population.
    Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously. Given the current state of the m x n grid board, return the next state.

*/
func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func gameOfLife(board [][]int) {
	M := len(board)
	N := len(board[0])

	// base condition
	if M == 1 && N == 1 {
		if board[0][0] == 1 {
			board[0][0] = 0
		}
		return
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			// get call once from their neiboures and check givin conditions
			onces := getOnesCount(board, i, j, M, N)

			// assign 2 so we can identify that previous values was 1 and we need to put 0 after completed this.
			if ((onces > 3) || (onces < 2)) && board[i][j] == 1 {
				board[i][j] = 2
				// assign 3 so we can identify that previous values was 0 and we need to assign 1 after completed all.
			} else if (onces == 3) && board[i][j] == 0 {
				board[i][j] = 3
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			}
			if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}

// function check all 1 to all of 8 positons.
func getOnesCount(board [][]int, row, col, M, N int) int {
	rowArr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	colArr := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	count := 0

	for i := 0; i < 8; i++ {
		newR := row + rowArr[i]
		newC := col + colArr[i]
		if newR >= 0 && newR < M && newC >= 0 && newC < N && (board[newR][newC] == 2 || board[newR][newC] == 1) {
			count++
		}

	}

	return count
}

//(1) 2,3 game_of_life    = 2
//(1) < 2 dead     =3
//(1) > 3 dead     =4
//(0) == 3        =5
