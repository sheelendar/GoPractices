package main

import (
	"fmt"
	"math"
)

/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray
whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1

Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/
func main() {
	num := []int{1, 2, 3, 4, 5}
	target := 11
	fmt.Println(minSubArrayLen(target, num))
}
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 || target == 0 {
		return 0
	}
	left := 0
	sum := 0
	count := math.MaxInt
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
		if sum >= target && count > i-left+1 {
			count = i - left + 1
		}
		for sum > target {
			sum = sum - nums[left]
			left++
			if sum >= target && count > i-left+1 {
				count = i - left + 1
			}
		}
	}
	if count == math.MaxInt {
		return 0
	}
	return count
}
