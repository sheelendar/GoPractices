package main

import (
	"fmt"
	"math"
)

/*
nums1 = [3,4,5,1,2]

nums = [4,5,6,7,0,1,2]

nums3 = [11,13,15,17]

nums4 = [3,1,2]
*/

func main() {
	//arr := []int{3, 4, 5, 1, 2} //output:1
	//arr := []int{4, 5, 6, 7, 0, 1, 2} //output:0

	//arr := []int{11, 13, 15, 17} //output:11

	//arr := []int{3, 1, 2} // output: 1

	arr := []int{1, 1, 1, 1, 1, 0, 1, 1, 1} //output: 0
	fmt.Println(findMinimumInArray(arr))
}

func findMinimumInArray(arr []int) int {
	high := len(arr) - 1
	low := 0
	res := math.MaxInt

	for low < high {
		mid := (low + high) / 2

		if arr[low] == arr[mid] {
			low++
			res = Min(res, arr[low])
		} else if arr[low] < arr[mid] {
			res = Min(res, arr[low])
			low = mid + 1
		} else { //arr[mid] < arr[high]
			res = Min(res, arr[mid])
			high = mid - 1
		}
	}
	res = Min(res, arr[low])
	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
