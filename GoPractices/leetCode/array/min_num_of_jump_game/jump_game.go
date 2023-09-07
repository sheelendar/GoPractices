package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9}
	n := len(arr)
	//Output: 3
	fmt.Print(minNumOfJump(arr, n))
}

func minNumOfJump(arr []int, n int) int {
	if n <= 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
		for j := 0; j < i; j++ {
			if i <= arr[j]+j && dp[i] > dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[n-1]
}
