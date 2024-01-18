package main

import (
	"fmt"
)

/*
A sawtooth sequence is a sequence of numbers that alternate between increasing and decreasing.
In other words, each element is either strictly greater than its neighbouring elements or strictly less than its neighbouring elements.
Given an array of integers arr, your task is to count the number of contiguous subarrays that represent a sawtooth
sequence of at least two elements.
Example

    For arr = [9, 8, 7, 6, 5], the output should be solution(arr) = 4.

    Since all the elements are arranged in decreasing order, it won't be possible to form any sawtooth subarrays of length 3 or more. There are 4 possible subarrays containing two elements, so the answer is 4.

    For arr = [10, 10, 10], the output should be solution(arr) = 0.

    Since all of the elements are equal, none of subarrays can be sawtooth, so the answer is 0.

    For arr = [1, 2, 1, 2, 1], the output should be solution(arr) = 10.

	All contiguous subarrays containing at least two elements satisfy the condition of problem. There are 10 possible contiguous subarrays containing at least two elements, so the answer is 10.
			public int countSeq(int[] arr) {
				int len = arr.length;
				if (len < 2) {
				return 0;
				}
				int s = 0;
				int e = 1;
				int sign = arr[e] - arr[s];
				int count = 0;

				while (e < len) {
				while (e < len && arr[e] - arr[e-1] != 0 && isSameSign(arr[e] - arr[e-1], sign)) {
					sign = -1 * sign;
					e++;
				}
				// the biggest continue subsequence starting from s ends at e-1;
				int size = e - s;
				count = count + (size * (size - 1)/2); // basically doing C(size,2)
				s = e - 1;
				e = s + 1;
				}
				return count;
		}

*/

func main() {
	arr := []int{9, 8, 7, 6, 5} // 4
	//arr := []int{1, 2, 1, 2, 1} //10
	fmt.Println(countSubarray(arr))
}
func countSubarray(arr []int) int {
	count := 0
	n := len(arr)
	if n < 2 {
		return 0
	}
	start := 0
	end := 1
	sign := arr[end] - arr[start]
	for end < n {
		for end < n && arr[end]-arr[end-1] != 0 && !isSameSign(arr[end]-arr[end-1], sign) {
			sign = arr[end] - arr[end-1]
			end++
		}
		m := end - start
		count = count + (m*(m-1))/2

		start = end - 1
		end = end + 1
	}
	return count
}
func isSameSign(a, b int) bool {
	if (a > 0 && b > 0) || (a < 0 && b < 0) {
		return true
	}
	return false
}
