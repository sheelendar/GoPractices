package main

import (
	"fmt"
	"sort"
)

/*
Given an array arr[] of size N, the task is to print the lexicographically next greater permutation of the given array. If there does not exist any greater permutation, then print the lexicographically smallest permutation of the given array.

Examples:

	Input: N = 6, arr = {1, 2, 3, 6, 5, 4}
	Output: {1, 2, 4, 3, 5, 6}
	Explanation: The next permutation of the given array is {1, 2, 4, 3, 5, 6}.

	Input: N = 3, arr = {3, 2, 1}
	Output: {1, 2, 3}
	Explanation: As arr[] is the last permutation.
	So, the next permutation is the lowest one.
*/
func main() {
	arr := []int{1, 2, 3, 6, 5, 4} //output : 1 2 4 3 5 6
	fmt.Println(nextPermutation(arr))
}

func nextPermutation(nums []int) []int {
	n := len(nums)
	if n == 0 || n == 1 {
		return nums
	}
	index := -1
	for i := n - 1; i > 0; i-- { // find consucative where arr[i-1] < arr[i] and take i-1 as index
		if nums[i-1] < nums[i] {
			index = i - 1
			break
		}
	}
	if index == -1 {
		return reverseArr(nums, n) // if not found arr[i-1] < arr[i] such pair then reverse array.
	}
	i := n - 1
	for ; index < i; i-- { // check if we have any number that grater then arr[index] from right side till number
		if nums[index] < nums[i] {
			break
		}
	}
	temp := nums[i] // swap index and newly find grater number than arr[index]
	nums[i] = nums[index]
	nums[index] = temp
	arr := nums[index+1 : n]
	sort.Slice(arr, func(i, j int) bool { // sort all element grater than index
		return arr[i] < arr[j]
	})
	for i := 0; i < len(arr); i++ {
		nums[index+1+i] = arr[i]
	}
	return nums
}

func reverseArr(nums []int, n int) []int {
	i := 0
	j := n - 1
	for i < j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}
	return nums
}
