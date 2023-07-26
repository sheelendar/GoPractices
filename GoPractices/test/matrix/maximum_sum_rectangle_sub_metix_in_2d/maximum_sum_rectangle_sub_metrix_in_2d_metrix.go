package main

import "fmt"

func main() {
	metrix := [][]int{{1, 2, -1, -4, -20},
		{-8, -3, 4, 2, 1},
		{3, 8, 10, 1, 3},
		{-4, -1, 1, 7, -6}}
	maximum_sub_metrix(metrix, len(metrix), len(metrix[0]))
}

func maximum_sub_metrix(metrix [][]int, n, m int) {
	row1 := 0
	row2 := 0
	col1 := 0
	col2 := 0
	colS := 0
	colE := 0
	maxSum := 0
	res := make([]int, m)

	for k := 0; k < n; k++ {
		res = make([]int, m)
		for i := k; i < n; i++ {

			for j := 0; j < m; j++ {
				res[j] = res[j] + metrix[i][j]
			}
			sum := getLargestSumOfContinuesSubArray(res, m, &colS, &colE)
			if sum > maxSum {
				col1 = colS
				col2 = colE
				row1 = k
				row2 = i
				maxSum = sum
			}
			fmt.Println(res)
		}
	}
	fmt.Println("Start row, col", row1, col1)
	fmt.Println("End row, col", row2, col2)
	fmt.Println("max sum", maxSum)

}

func getLargestSumOfContinuesSubArray(arr []int, size int, colS, colE *int) int {
	largeSum := int(0)
	tempSum := int(0)
	*colS = 0
	*colE = -1
	for item := range arr {
		tempSum = tempSum + arr[item]
		if tempSum > largeSum {
			largeSum = tempSum
			*colE = item
		}
		if tempSum < 0 {
			tempSum = 0
			*colS = item + 1
		}
	}

	if *colE != -1 {
		return largeSum
	}

	// Special Case: When all numbers in arr[]
	// are negative
	largeSum = arr[0]
	*colS = 0
	*colE = 0

	// Find the maximum element in array
	for i := 1; i < size; i++ {
		if arr[i] > largeSum {
			largeSum = arr[i]
			*colS = i
			*colE = i
		}
	}
	return largeSum
}
