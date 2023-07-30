package main

import "fmt"

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false
if every element is distinct.
Example 1:
Input: nums = [1,2,3,1]
Output: true

*/

func main() {
	arr := []int{1, 2, 3, 5}
	fmt.Print(containsDuplicate(arr))
}
func containsDuplicate(nums []int) bool {
	n := len(nums)
	if n <= 0 {
		return false
	}
	dp := make(map[int]bool)
	for i := 0; i < n; i++ {
		if !dp[nums[i]] {
			dp[nums[i]] = true
		} else {
			return true
		}
	}
	return false
}
