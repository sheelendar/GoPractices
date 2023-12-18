package main

import (
	"fmt"
	"math"
)

/*
Given an integer array nums, find a subarray that has the largest product, and return the product.
The test cases are generated so that the answer will fit in a 32-bit integer.

Example 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/
func main() {
	nums := []int{-1, 0, -2}
	fmt.Print(maxProduct(nums))
}

func maxProduct(nums []int) int {
	size := len(nums)
	max_product := -999999
	if size <= 0 {
		return 0
	}
	product := 1
	for i := 0; i < size; i++ {
		product = product * nums[i]
		max_product = int(math.Max(float64(product), float64(max_product)))
		if nums[i] == 0 {
			product = 1
		}
	}
	product = 1
	for i := size - 1; i >= 0; i-- {
		product = product * nums[i]
		max_product = int(math.Max(float64(product), float64(max_product)))
		if nums[i] == 0 {
			product = 1
		}
	}
	return max_product
}
