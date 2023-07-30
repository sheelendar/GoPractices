package main

import "fmt"

/*
Given an array of integers arr:

	Write a function flip(arr, k) that reverses the order of the first k elements in the array arr.
	Write a function pancakeSort(arr) that sorts and returns the input array. You are allowed to use only the function

flip you wrote in the first step in order to make changes in the array.

Example:

input:  arr = [1, 5, 4, 3, 2]

output: [1, 2, 3, 4, 5] # to clarify, this is pancakeSort's output

Analyze the time and space complexities of your solution.
*/
func main() {
	arr := []int{1, 2}
	flip(arr, len(arr)-1)
	r := PancakeSort(arr)

	for i := 0; i < len(r); i++ {
		fmt.Print(r[i])
		fmt.Print(" ")
	}
}

func PancakeSort(arr []int) []int {
	size := len(arr)
	for j := size - 1; j > 0; j-- {
		max := 0
		n := j
		maxIndex := -1
		for i := 0; i < n; i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		flip(arr, maxIndex)
		flip(arr, n)
	}
	return arr
}
func flip(arr []int, k int) {
	i := 0
	for i < k {
		temp := arr[i]
		arr[i] = arr[k]
		arr[k] = temp
		i++
		k--
	}
}
