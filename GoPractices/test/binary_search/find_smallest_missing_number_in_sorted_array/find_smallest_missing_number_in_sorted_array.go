package main

import "fmt"

/*
Given a sorted array of n distinct integers where each integer is in the range from 0 to m-1 and m > n.
Find the smallest number that is missing from the array.
Input: {0, 1, 2, 6, 9}, n = 5, m = 10
Output: 3

Input: {4, 5, 10, 11}, n = 4, m = 12
Output: 0

Input: {0, 1, 2, 3}, n = 4, m = 5
Output: 4
*/
func main() {
	//input := []int{0, 1, 2, 6, 9} // output: 3
	// input := []int{4, 5, 10, 11} // output:0
	//input := []int{0, 1, 2, 3} // output:4
	input := []int{} // output:4
	fmt.Println(finsSmallestMissingNumber(input))
}
func finsSmallestMissingNumber(arr []int) int {
	l := 0
	h := len(arr) - 1
	for l <= h {
		if arr[l] != l {
			return l
		}
		mid := (l + h) / 2
		if arr[mid] != mid {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return h + 1
}
