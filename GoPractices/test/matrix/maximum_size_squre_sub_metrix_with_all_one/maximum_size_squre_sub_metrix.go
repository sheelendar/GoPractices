package main

import (
	"fmt"
)

/*
Given a binary matrix, find out the maximum size square sub-matrix with all 1s.
For example, consider the below binary matrix.

Given a binary matrix, find out the maximum size square sub-matrix with all 1s. For example, consider the below binary matrix.
(https://www.geeksforgeeks.org/maximum-size-sub-matrix-with-all-1s-in-a-binary-matrix/)
*/
func main() {
	M := [][]int64{{0, 1, 1, 0, 1}, {1, 1, 0, 1, 0},
		{0, 1, 1, 1, 0}, {1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}}
	printMexSizeMatrixWithOnes(M, len(M), len(M[0]))
}

func printMexSizeMatrixWithOnes(M [][]int64, R int, C int) {
	var S [][]int64
	S = make([][]int64, R)
	for i := 0; i < R; i++ {
		S[i] = make([]int64, C)
	}
	for i := 0; i < C; i++ {
		S[0][i] = M[0][i]
	}
	for i := 0; i < R; i++ {
		S[i][0] = M[i][0]
	}
	for i := 1; i < R; i++ {
		for j := 1; j < C; j++ {
			if M[i][j] == 1 {
				S[i][j] = Min(Min(S[i-1][j], S[i][j-1]), S[i-1][j-1]) + 1
			} else {
				S[i][j] = 0
			}
		}
	}
	displayMatrix(S, R, C)
	fmt.Println()
	displayMexMatrix(S, R, C)
}
func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func displayMexMatrix(S [][]int64, R int, C int) {
	maxi := 0
	maxj := 0
	max := int64(0)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if max < S[i][j] {
				maxi = i + 1
				maxj = j + 1
				max = S[i][j]
			}
		}
	}
	for i := R - maxj; i < maxi; i++ {
		for j := C - maxj; j < maxj; j++ {
			fmt.Print(S[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func displayMatrix(S [][]int64, R int, C int) {
	for i := 1; i < R; i++ {
		for j := 1; j < C; j++ {
			fmt.Print(S[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
