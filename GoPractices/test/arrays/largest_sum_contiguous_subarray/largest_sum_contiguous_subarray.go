package main

import "fmt"

func main() {
	a := []int64{2, -3, 4, -1, -2, 1, 5, -3}
	var largeSum int64
	largeSum = getLargestSumOfContinuesSubArray(a, len(a))
	fmt.Println("sum:", largeSum)
}

func getLargestSumOfContinuesSubArray(arr []int64, size int) int64 {
	largeSum := int64(0)
	tempSum := int64(0)
	for item := range arr {
		tempSum = tempSum + arr[item]
		if tempSum > largeSum {
			largeSum = tempSum
		}
		if tempSum < 0 {
			tempSum = 0
		}
	}
	return largeSum
}

// second solution.
func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	curSum, maxSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		curSum = max(curSum+nums[i], nums[i])
		maxSum = max(curSum, maxSum)
	}
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
