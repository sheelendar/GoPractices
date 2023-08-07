package main

import "math"

func main() {
	arr := []int64{5, 5, 10, 100, 10, 5}
	//arr := []int64{3, 2, 7, 10}
	sum := findMaxSumNotwoElementsAreAdjacent(arr, len(arr))
	println("max sum is:", sum)

}

func findMaxSumNotwoElementsAreAdjacent(arr []int64, size int) int64 {

	var maxSum, currSum, preSum int64
	preSum = 0
	currSum = arr[0]
	for i := 1; i < size; i++ {
		maxSum = int64(math.Max(float64(preSum), float64(currSum)))
		currSum = preSum + arr[i]
		preSum = maxSum
	}
	maxSum = int64(math.Max(float64(maxSum), float64(currSum)))
	return maxSum
}
