package main

import (
	"fmt"
	"math"
)

/*
Given an array containing only 0s and 1s, find the largest subarray which contains equal no of 0s and 1s.
The expected time complexity is O(n). https://www.geeksforgeeks.org/largest-subarray-with-equal-number-of-0s-and-1s/

Examples:
Input: arr[] = {1, 0, 1, 1, 1, 0, 0}
Output: 1 to 6
(Starting and Ending indexes of output subarray)
Input: arr[] = {1, 1, 1, 1}
Output: No such subarray
*/

func main() {
	arr := []int{1, 0, 1, 1, 1, 0, 0}
	//Output: 1 to 6

	//arr := []int{1, 1, 1, 1}
	//Output: 0

	//arr := []int{1, 0, 1, 0}
	// output: 4

	fmt.Println(findSubarrayWithEqualNumberOf1And0(arr))
}

func findSubarrayWithEqualNumberOf1And0(arr []int) int {
	n := len(arr)
	if n == 0 || n == 1 {
		return 0
	}
	dp := make(map[int]int)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			arr[i] = -1
		}
	}
	max := 0
	res := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
		if sum == 0 {
			max = int(math.Max(float64(i+1), float64(max)))
		}
		if index, ok := dp[sum]; ok {
			max = int(math.Max(float64(i-index-1), float64(max)))
		} else {
			dp[sum] = i
		}
		if max > res {
			res = max
		}
	}
	return res
}
