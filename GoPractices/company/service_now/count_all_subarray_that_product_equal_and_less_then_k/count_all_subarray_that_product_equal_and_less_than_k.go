package main

import (
	"fmt"
)

/*
Given an array of positive numbers, calculate the number of possible contiguous subarrays having product lesser than a given number K.

Examples :

    Input : arr[] = [1, 2, 3, 4]
                K = 10
    Output : 7
    The subarrays are {1}, {2}, {3}, {4}, {1, 2}, {1, 2, 3} and {2, 3}

    Input  : arr[] = [1, 9, 2, 8, 6, 4, 3]
                 K = 100
    Output : 16

*/

func main() {
	//arr := []int{1, 2, 3, 4}
	//k := 10

	arr := []int{1, 9, 2, 8, 6, 4, 3}
	k := 100
	fmt.Println(getAllSubArray(arr, k))
}
func getAllSubArray(arr []int, k int) int {
	count := 0
	n := len(arr)
	if n == 0 {
		return 0
	}
	pro := 1
	j := 0
	for i := 0; i < n; i++ {
		pro = pro * arr[i]
		for pro > k && j < n {
			pro = pro / arr[j]
			j++
		}
		if pro <= k {
			count = count + (i - j) + 1
		}
	}
	return count
}
