/*
Write me a code to find out diagnoal difference in array

	func diagonalDifference(arr [][]int32) int32 {
	    // Write your code here

}

	matrix := [][]int32{
	        {11, 2, 4},
	        {4, 5, 6},
	        {10, 8, -12},
	    }
		left dia:= 11, 5, -12 = 4  (i=j)  O(n)
		right 4,5,10 = 19    ({0,2}{1,1} {2,0} ) (++i, --j)  O(n)
		|19-4|= 15
*/
package main

import "fmt"

func main() {
	/*matrix := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}*/

	matrix := [][]int32{
		{16, 2, 3, 13},
		{5, 11, 10, 8},
		{9, 7, 6, 12},
		{4, 14, 15, 1},
	}

	fmt.Println(diagonalDifference(matrix))
}

func diagonalDifference(arr [][]int32) int32 {
	n := len(arr)
	leftDiagonalSum := int32(0)
	rightDiagonalSum := int32(0)
	for i := 0; i < n; i++ {
		leftDiagonalSum += arr[i][i]
		rightDiagonalSum += arr[i][n-i-1]
	}
	diff := leftDiagonalSum - rightDiagonalSum
	if diff < 0 {
		return -diff
	}
	return diff
}
