package main

import "fmt"

/*
Given a set of non-negative integers and a value sum, the task is to check if there is a subset of the given set whose sum is equal to the given sum.
Examples:

	Input: set[] = {3, 34, 4, 12, 5, 2}, sum = 9
	Output: True
	Explanation: There is a subset (4, 5) with sum 9.
*/
func main() {
	set := []int{3, 34, 4, 12, 5, 2}
	//sum := 9
	sum := 30
	size := len(set)
	dp := make(map[string]bool, size+1)
	fmt.Println("is there any subset with given sum:", isSumSetThere(set, sum, size, dp))
}

func isSumSetThere(set []int, sum int, n int, dp map[string]bool) bool {
	if sum == 0 {
		return true
	}
	if n <= 0 {
		return false
	}
	key := fmt.Sprintf("%d_%d", sum, n-1)
	if v, ok := dp[key]; ok {
		return v
	}
	if sum < set[n-1] {
		dp[key] = isSumSetThere(set, sum, n-1, dp)
		return dp[key]
	}
	isSum := isSumSetThere(set, sum-set[n-1], n-1, dp) || isSumSetThere(set, sum, n-1, dp)
	dp[key] = isSum
	return isSum
}
