package main

import "math"

func main() {
	arr := []int64{5, 5, 10, 100, 10, 5}
	//arr := []int64{3, 2, 7, 10}
	sum := findMaxSumNotwoElementsAreAdjacent(arr, len(arr))
	println("max sum is:", sum)

}

func findMaxSumNotwoElementsAreAdjacent(arr []int64, size int) int64 {

	var maxSum, preSum, currSum int64
	preSum = arr[0]
	for i := 1; i < size; i++ {
		maxSum = int64(math.Max(float64(currSum), float64(preSum)))
		preSum = currSum + arr[i]
		currSum = maxSum
	}
	maxSum = int64(math.Max(float64(maxSum), float64(preSum)))
	return maxSum
}
