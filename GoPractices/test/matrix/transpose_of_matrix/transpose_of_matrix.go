package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix = transpose(matrix)
	fmt.Println(matrix)
}

func transpose(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
			fmt.Print(temp, " - ", matrix[i][j], ",  ")
		}
	}
	return matrix
}
