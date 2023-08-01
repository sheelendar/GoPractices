package main

/*
For

matrix = [[9, 7, 4, 5],
          [1, 6, 2, -6],
          [12, 20, 2, 0],
          [-5, -6, 7, -2]]

the output should be

solution(matrix) = [[-6, -6, -5, -2],
                    [12, 2, 2, 0],
                    [9, 20, 6, 1],
                    [7, 7, 5, 4]]

    The 0-border consists of the elements [9, 7, 4, 5, -6, 0, -2, 7, -6, -5, 12, 1]. After sorting, the elements order is [-6, -6, -5, -2, 0, 1, 4, 5, 7, 7, 9, 12]. So we'll place these back on the matrix clockwise, starting from the top-left corner.
    The 1-border consists of the elements [6, 2, 2, 20]. After sorting, the elements order is [2, 2, 6, 20].


*/
import (
	"fmt"
	"sort"
)

func main() {
	matrix := [][]int{
		{9, 7, 4, 5},
		{1, 6, 2, -6},
		{12, 20, 2, 0},
		{-5, -6, 7, -2}}

	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(matrix[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
	solution(matrix)
}
func solution(matrix [][]int) [][]int {

	size := len(matrix)
	if size == 1 || size == 0 {
		return matrix
	}

	sRow := 0
	eRow := size - 1
	sCol := 0
	eCol := size - 1

	osRow := 0
	oeRow := size - 1
	osCol := 0
	oeCol := size - 1

	for i := 0; i < size/2; i++ {
		var temp []int

		for j := sCol; j <= eCol; j++ {
			temp = append(temp, matrix[sRow][j])
		}
		sRow++

		for j := sRow; j <= eRow; j++ {
			temp = append(temp, matrix[j][eCol])
		}
		eCol--

		for j := eCol; j >= sCol; j-- {
			temp = append(temp, matrix[eRow][j])
		}
		eRow--

		for j := eRow; j >= sRow; j-- {
			temp = append(temp, matrix[j][sCol])
		}
		sCol++

		sort.Ints(temp)

		k := 0
		for j := osCol; j <= oeCol; j++ {
			matrix[osRow][j] = temp[k]
			k++
		}
		osRow++

		for j := osRow; j <= oeRow; j++ {
			matrix[j][oeCol] = temp[k]
			k++
		}
		oeCol--

		for j := oeCol; j >= osCol; j-- {
			matrix[oeRow][j] = temp[k]
			k++
		}
		oeRow--

		for j := oeRow; j >= osRow; j-- {
			matrix[j][osCol] = temp[k]
			k++
		}
		osCol++

	}
	return matrix
}
