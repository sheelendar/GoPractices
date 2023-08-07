package main

import (
	"fmt"
)

/*
Given an integer array of coins[ ] of size N representing different types of currency and an integer sum,
The task is to find the number of ways to make sum by using different combinations from coins[].

Note: Assume that you have an infinite supply of each type of coin.

Examples:

	Input: sum = 4, coins[] = {1,2,3},
	Output: 4
	Explanation: there are four solutions: {1, 1, 1, 1}, {1, 1, 2}, {2, 2}, {1, 3}.
*/
var (
	table = make(map[string]int64)
)

func main() {
	arr := []int{1, 2, 3}
	sum := 4
	//arr := []int{1, 2, 5}
	//sum := 12
	fmt.Println("total coin solutions : ", coinChanged(arr, len(arr), sum))
	fmt.Println("total coin solutions from recursion: ", coinChangedRecursion(arr, len(arr), sum))
}

func coinChanged(arr []int, n, sum int) any {
	if sum == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	dp := make([]int64, sum+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := arr[i]; j <= sum; j++ {
			dp[j] = dp[j] + dp[j-arr[i]]

		}
	}
	return dp[sum]
}

func coinChangedRecursion(arr []int, n, sum int) int64 {
	if sum == 0 {
		return 1
	}
	if n <= 0 {
		return 0
	}
	if arr[n-1] > sum {
		return coinChangedRecursion(arr, n-1, sum)
	}
	if v, ok := table[fmt.Sprintf("%d_%d", n, sum)]; ok {
		return v
	}
	count := coinChangedRecursion(arr, n-1, sum) + coinChangedRecursion(arr, n, sum-arr[n-1])
	table[fmt.Sprintf("%d_%d", n, sum)] = count
	return count
}
