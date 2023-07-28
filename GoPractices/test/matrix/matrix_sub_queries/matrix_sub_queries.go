package main

import "fmt"

/*
Submatrix Sum Queries

Given a matrix of size M x N, there are large number of queries to find submatrix sums.
Inputs to queries are left top and right bottom indexes of submatrix whose sum is to find out.
How to preprocess the matrix so that submatrix sum queries can be performed in O(1) time.

tli :  Row number of top left of query submatrix
tlj :  Column number of top left of query submatrix
rbi :  Row number of bottom right of query submatrix
rbj :  Column number of bottom right of query submatrix

Input: mat[M][N] = {{1, 2, 3, 4, 6},

	{5, 3, 8, 1, 2},
	{4, 6, 7, 5, 5},
	{2, 4, 8, 9, 4} };

Query1: tli = 0, tlj = 0, rbi = 1, rbj = 1
Query2: tli = 2, tlj = 2, rbi = 3, rbj = 4
Query3: tli = 1, tlj = 2, rbi = 3, rbj = 3;

Output:
Query1: 11  // Sum between (0, 0) and (1, 1)
Query2: 38  // Sum between (2, 2) and (3, 4)
Query3: 38  // Sum between (1, 2) and (3, 3)
*/

func main() {
	mat := [][]int{
		{1, 2, 3, 4, 6},
		{5, 3, 8, 1, 2},
		{4, 6, 7, 5, 5},
		{2, 4, 8, 9, 4}}
	sum := prepareMatrixAndCalculateSum(mat)
	findSum(sum, 0, 0, 1, 1)
	findSum(sum, 2, 2, 3, 4)
}

func findSum(sum [][]int, x1 int, y1 int, x2 int, y2 int) {
	s := sum[x2][y2]
	if x1 > 0 {
		s = s - sum[x1-1][y2]
	}
	if y1 > 0 {
		s = s - sum[x2][y1-1]
	}
	if x1 > 0 && y1 > 0 {
		s = s + sum[x1-1][y1-1]
	}
	fmt.Print(s)
	fmt.Println()
}

func prepareMatrixAndCalculateSum(mat [][]int) [][]int {
	Row := len(mat)
	Col := len(mat[0])
	sum := make([][]int, Row)
	for i := 0; i < Row; i++ {
		sum[i] = make([]int, Col)
	}
	for i := 0; i < Col; i++ {
		sum[0][i] = mat[0][i]
	}
	for i := 0; i < Row; i++ {
		sum[i][0] = mat[i][0]
	}
	for i := 1; i < Row; i++ {
		for j := 0; j < Col; j++ {
			sum[i][j] = sum[i-1][j] + mat[i][j]
		}
	}

	for i := 1; i < Col; i++ {
		for j := 0; j < Row; j++ {
			sum[j][i] = sum[j][i-1] + sum[j][i]
		}

	}

	for i := 0; i < Row; i++ {
		for j := 0; j < Col; j++ {
			fmt.Print(sum[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
	return sum
}
