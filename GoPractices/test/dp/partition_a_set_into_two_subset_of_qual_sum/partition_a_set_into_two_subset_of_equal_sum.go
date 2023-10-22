package main

import (
	"fmt"
	"math"
)

/*
Given a set of integers, the task is to divide it into two sets S1 and S2 such that the absolute difference between their sums is minimum.
If there is a set S with n elements, then if we assume Subset1 has m elements, Subset2 must have n-m elements and the value of abs(sum(Subset1) – sum(Subset2)) should be minimum.
Example: (https://www.geeksforgeeks.org/partition-a-set-into-two-subsets-such-that-the-difference-of-subset-sums-is-minimum/)
Input:  arr[] = {1, 6, 11, 5}
Output: 1
Explanation:
Subset1 = {1, 5, 6}, sum of Subset1 = 12
Subset2 = {11}, sum of Subset2 = 11

*/

func main() {
	arr := []int{1, 6, 11, 5}
	fmt.Println(minSumPartition(arr, len(arr)))
}

func minSumPartition(arr []int, n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
	}
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if arr[i-1] <= j {
				dp[i][j] = dp[i][j] || dp[i-1][j-arr[i-1]]
			}
		}
	}
	min := math.MaxInt
	for j := sum / 2; j >= 0; j-- {
		if dp[n][j] {
			min = Min(min, sum-j*2)
		}
	}
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
