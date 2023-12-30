package main

import "fmt"

/*
 Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
Example 1:  https://leetcode.com/problems/majority-element-ii/description/

Input: nums = [3,2,3], Output: [3]

Example 2: Input: nums = [1] , Output: [1]

Example 3: Input: nums = [1,2], Output: [1,2]
Ans:


*/

func main() {

	nums := []int{1, 1, 2, 3, 4, 1, 1, 5, 6, 7, 1, 1, 8, 9, 10, 1, 11, 12, 13, 14}
	fmt.Println(majorityElement(nums))
}
func majorityElement(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	num1 := nums[0]
	num2 := nums[1]
	count1 := 0
	count2 := 0
	for i := 0; i < n; i++ {
		if nums[i] == num1 {
			count1++
		} else if nums[i] == num2 {
			count2++
		} else if count1 == 0 {
			num1 = nums[i]
			count1 = 1
		} else if count2 == 0 {
			num2 = nums[i]
			count2 = 1
		} else {
			count1--
			count2--
		}
	}
	count1 = 0
	count2 = 0
	for i := 0; i < n; i++ {
		if num1 == nums[i] {
			count1++
		}
		if num2 == nums[i] {
			count2++
		}
	}
	var result []int
	if count1 > n/3 {
		result = append(result, num1)
	}
	if count2 > n/3 && num1 != num2 {
		result = append(result, num2)
	}
	return result
}
