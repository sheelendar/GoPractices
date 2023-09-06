package main

import "fmt"

func main() {
	arr := []int{1, 5, 8, 9, 10, 17, 17, 20}
	fmt.Print(rodCut(arr, len(arr)))
}

func rodCut(arr []int, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = arr[i-1]
		for j := 0; j < i; j++ {
			if dp[i] < dp[i-j-1]+arr[j] {
				dp[i] = dp[i-j-1] + arr[j]
			}
		}
	}
	return dp[n]
}

/*
n=3
dp[1] < dp[0]+arr[0]
dp[2] < dp[1]+arr[0]
dp[2] < dp[0]+arr[1]

*/
