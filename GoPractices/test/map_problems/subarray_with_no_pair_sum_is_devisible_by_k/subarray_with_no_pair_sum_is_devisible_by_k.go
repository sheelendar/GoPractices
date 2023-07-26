package main

import "fmt"

/*
Given an array of N non-negative integers, task is to find the maximum size of a subarray such that the pairwise
sum of the elements of this subarray is not divisible by a given integer, K. Also, print this subarray as well.
If there are two or more subarrays that follow the above stated condition, then print the first one from the left.
(https://www.geeksforgeeks.org/subarray-no-pair-sum-divisible-k/)

Input : arr[] = [3, 7, 1, 9, 2]

	K = 3

Output : 3

	[3, 7, 1]

3 + 7 = 10, 3 + 1 = 4, 7 + 1 = 8, all are
not divisible by 3.
It is not possible to get a subarray of size bigger
than 3 with the above-mentioned property.
[7, 1, 9] is also of the same size but [3, 7, 1] comes first.
*/

func main() {
	//arr := []int{3, 7, 1, 9, 2}
	arr := []int{5, 10, 15, 20, 25}
	k := 3
	printSubArrayWithNoPairDevisibleByK(arr, len(arr), k)
}

func printSubArrayWithNoPairDevisibleByK(arr []int, size int, k int) {
	dp := make(map[int]int)
	start := 0
	end := 0
	startI := 0
	endI := 0
	for i := 0; i < size; i++ {

		mod := arr[i] % k

		for dp[k-mod] != 0 || (dp[mod] != 0 && mod == 0) {
			dp[arr[start]%k]--
			start++
		}

		dp[mod]++
		endI++

		if endI-startI > end-start {
			end = endI
			startI = start
		}
	}
	fmt.Print(end - start + 1)
}
