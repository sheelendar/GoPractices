package main

import "fmt"

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product
of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/
func main() {
	num := []int{-1, 1, 0, -3, 3}
	result := productExceptSelf(num)
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		fmt.Print(" ")
	}
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n <= 0 {
		return nil
	}
	out := make([]int, n)
	countzeros := 0
	indexOfzero := -1
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			countzeros++
			indexOfzero = i
		}
	}
	if countzeros > 1 {
		return out
	} else if countzeros == 1 {
		p := 1
		for i := 0; i < n; i++ {
			if i != indexOfzero {
				p = p * nums[i]
			}
		}
		out[indexOfzero] = p
	} else {
		p := 1
		for i := 0; i < n; i++ {
			p = p * nums[i]
		}
		for i := 0; i < n; i++ {
			out[i] = p / nums[i]
		}
	}
	return out
}
