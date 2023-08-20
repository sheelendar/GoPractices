package main

import "fmt"

func main() {
	//Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
	//Output: [[1,0,1],[0,0,0],[1,0,1]]
	input := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(input)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			fmt.Print(input[i][j])
			fmt.Print()
		}
		fmt.Println()
	}
}

// solution with O(1) space complexcity
func setZeroes(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	col := len(matrix[0])
	firstRowZero := false
	firstColumn := false

	// if any i and j has zero then just put zero in i,0 and j,0 to track it later.
	// also check if first column or row has zero at any point if yes then later we have to make whole row or column zero.
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				if i == 0 {
					firstRowZero = true
					firstColumn = true
				}
			}
		}
	}
	// check if first row or column is zero for value of i or j then make that index as zero as well collected earlier.
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// if first row has nay zero then make whole row as zero
	if firstRowZero {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}
	//  check if first column has zero then make whole column as zero
	if firstColumn {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}

func setZeroes2(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}

	col := len(matrix[0])

	r := make([]int, row)
	c := make([]int, col)

	for i := 0; i < row; i++ {

		for j := 0; j < col; j++ {

			if matrix[i][j] == 0 {
				r[i] = 1
				c[j] = 1
			}
		}
	}

	for i := 0; i < row; i++ {
		if r[i] == 0 {
			for j := 0; j < col; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < col; j++ {
		if c[j] == 0 {
			for i := 0; i < row; i++ {
				matrix[i][j] = 0
			}
		}
	}
}
