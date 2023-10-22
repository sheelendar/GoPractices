package main

import (
	"fmt"
)

// maximum_sum_sub_array_with_a_pair_where_a_<_b_>_c
func main() {
	//arr := []int32{3, 4, 5, 1, 2, 3, 1}
	//arr := []int32{3, 6, 1, 1, 1}
	arr := []int32{7, 3, 4, 3, 6, 1, 1, 1} // 8
	println("", maxSumSubsequence(arr, len(arr)))
}

func maxSumSubsequence(arr []int32, size int) int64 {
	size = len(arr)
	if size == 0 {
		return -1
	}

	left := make([]int32, size+1)
	right := make([]int32, size+1)
	left[0] = arr[0]
	right[size-1] = arr[size-1]

	for i := 1; i < size; i++ {
		left[i] = Min(arr[i], left[i-1])
	}
	left[0] = 0
	for i := size - 2; i >= 0; i-- {
		right[i] = Min(arr[i], right[i+1])
	}
	right[size-1] = 0
	for i := 0; i < size; i++ {
		fmt.Print(left[i])
		fmt.Print(" ")
	}
	fmt.Println()
	for i := 0; i < size; i++ {
		fmt.Print(right[i])
		fmt.Print(" ")
	}
	fmt.Println()
	max := int64(999999)

	for i := 1; i < size; i++ {
		if max > int64(arr[i]+left[i]+right[i]) && left[i] != arr[i] && right[i] != arr[i] {
			max = int64(arr[i] + left[i] + right[i])
		}
	}
	if max == int64(999999) {
		return -1
	}
	return max
}
func Min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
