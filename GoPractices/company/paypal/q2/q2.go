/*
Consider a rectangular h × w board with all cells initially white. You are to process several queries of the following types:

    "x a b" - color the white cell (a, b) (0-based coordinates, the first one is a row index, and the second one is a column index) black;
    "> a b" - find the leftmost white cell to the right of the white cell (a, b);
    "< a b" - find the rightmost white cell to the left of the white cell (a, b);
    "v a b" - the same, but the search should be done downwards;
    "^ a b" - the same, but the search should be done upwards;

For each query, except the ones of the first type, find the answer.

Example

For h = 3, w = 5, and
queries = ["v 1 2", "x 2 2", "v 1 2", "> 2 1", "x 2 3", "> 2 1", "< 2 0"],
the output should be
solution(h, w, queries) = [[2, 2], [-1, -1], [2, 3], [2, 4], [-1, -1]].
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	h := 3
	w := 5
	queries := []string{"v 1 2", "x 2 2", "v 1 2", "> 2 1", "x 2 3", "> 2 1", "< 2 0"}

	//solution(h, w, queries) = [[2, 2], [-1, -1], [2, 3], [2, 4], [-1, -1]].
	fmt.Println(solution(h, w, queries))
}
func solution(h int, w int, queries []string) [][]int {
	dp := make(map[byte]int)
	dp['x'] = 0
	dp['>'] = 1
	dp['<'] = 2
	dp['v'] = 3
	dp['^'] = 4
	board := make([][]int, h)
	for i := 0; i < h; i++ {
		board[i] = make([]int, w)
	}

	var result [][]int
	for k := 0; k < len(queries); k++ {
		arr := strings.Split(queries[k], " ")
		i, _ := strconv.Atoi(arr[1])
		j, _ := strconv.Atoi(arr[2])

		if arr[0] == "x" {
			board[i][j] = 1
		} else {
			res := DFS(board, i, j, h, w, arr[0])
			result = append(result, res)
		}

	}
	return result
}
func DFS(board [][]int, row, col, m, n int, symbol string) []int {

	res := []int{-1, -1}
	if symbol == ">" {
		for j := col + 1; j < n; j++ {
			if board[row][j] == 0 {
				return []int{row, j}
			}
		}

	} else if symbol == "<" {
		for j := col - 1; j >= 0; j-- {
			if board[row][j] == 0 {
				return []int{row, j}
			}
		}
	} else if symbol == "v" {
		for i := row + 1; i < m; i++ {
			if board[i][col] == 0 {
				return []int{i, col}
			}
		}
	} else if symbol == "^" {
		for i := row - 1; i >= 0; i-- {
			if board[i][col] == 0 {
				return []int{i, col}
			}
		}
	}
	return res
}
