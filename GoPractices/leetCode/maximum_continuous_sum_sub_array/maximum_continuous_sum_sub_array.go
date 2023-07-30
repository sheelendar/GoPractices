package main

import (
	"fmt"
)

/*
Given an integer array nums, find the
subarray
with the largest sum, and return its sum.

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/
func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Print(maxSubArray(arr))
}

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
