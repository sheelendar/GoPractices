package main

import (
	"fmt"
)

/*
Given an integer array nums, you need to find one continuous subarray that if you only sort this subarray in ascending order,
then the whole array will be sorted in ascending order.

Return the shortest such subarray and output its length.
Input: nums = [2,6,4,8,10,9,15]
Output: 5


2,6,4,1,8,10,9,15


2,6,4,8,10,9,15


start=1,4  -> 6
end =2,5   -> 4


[]


9
8
1

lenght= i-s[size-1]
*/

func main() {
	//arr := []int{2, 6, 4, 8, 10, 9, 15}
	//arr := []int{1, 3, 2, 4, 5}
	arr := []int{1, 20, 3, 4, 0, 7, 18, 9, 6, 10, 11, -1, 12, 13, 25}
	//fmt.Println(findWindow(arr))
	//fmt.Println(findWindow(arr1))
	fmt.Println(findUnsortedSubarray(arr))
}

func findUnsortedSubarray(arr []int) int {
	n := len(arr)
	l, r := -1, -1
	// Find the left boundary of unsorted subarray
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			l = i
			break
		}
	}
	// Find the right boundary of unsorted subarray
	for i := n - 1; i >= 1; i-- {
		if arr[i] < arr[i-1] {
			r = i
			break
		}
	}
	// If l is still -1, it means the array is already sorted
	if l == -1 {
		return 0
	}
	// Find the minimum and maximum values within the unsorted subarray
	mini, maxi := arr[l], arr[l]
	for i := l; i <= r; i++ {
		if mini > arr[i] {
			mini = arr[i]
		}
		if maxi < arr[i] {
			maxi = arr[i]
		}
	}
	// Update the left and right boundaries based on the sorted order
	for i := l; i >= 0; i-- {
		if arr[i] > mini {
			l = i
		}
	}
	for i := r; i < n; i++ {
		if arr[i] < maxi {
			r = i
		}
	}
	return (r - l) + 1
}
