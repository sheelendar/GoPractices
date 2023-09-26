/*
Given an unsorted array of size n. Array elements are in the range of 1 to n. One number from set {1, 2, …n} is missing and one number occurs twice in the array. Find these two numbers.
Examples:
Input: arr[] = {3, 1, 3}:  Output: Missing = 2, Repeating = 3
Explanation: In the array, 2 is missing and 3 occurs twice
Input: arr[] = {4, 3, 6, 2, 1, 1}:  Output: Missing = 5, Repeating = 1
// https://www.geeksforgeeks.org/find-a-repeating-and-a-missing-number/

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	//arr := []int{4, 3, 6, 2, 1, 1} //1,5
	arr := []int{3, 1, 3} // 3,2
	//arr := []int{7, 3, 4, 5, 5, 6, 2}  //5,1
	a, b := findMissingAndRepeatingNumbers(arr)
	fmt.Println(a, ", ", b)
}
func findMissingAndRepeatingNumbers(arr []int) (int, int) {
	n := len(arr)
	a := -1
	for i := 0; i < n; i++ {
		index := int(math.Abs(float64(arr[i])))
		if arr[index-1] < 0 {
			a = index
		} else {
			arr[index-1] = -arr[index-1]
		}
	}
	b := -1
	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			b = i + 1
			break
		}
	}
	return a, b
}
