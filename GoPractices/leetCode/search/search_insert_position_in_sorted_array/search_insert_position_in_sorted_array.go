package main

import "fmt"

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.

	If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4
*/
func main() {
	arr := []int{1, 3, 5, 6}
	fmt.Println(helper(arr, 0, len(arr)-1, 5))
}

func helper(arr []int, l, h, k int) int {
	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == k {
			return mid
		} else if arr[mid] > k {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
