package main

import "fmt"

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1
*/
func main() {
	height := []int{1, 1}
	fmt.Println(findMaxCon(height))
}
func findMaxCon(arr []int) int {
	right := len(arr) - 1
	left := 0
	ans := 0
	max := 0
	for left < right {
		if arr[left] < arr[right] {
			max = arr[left] * (right - left)
			left++
		} else {
			max = arr[right] * (right - left)
			right--
		}
		if ans < max {
			ans = max
		}
	}
	return ans
}
