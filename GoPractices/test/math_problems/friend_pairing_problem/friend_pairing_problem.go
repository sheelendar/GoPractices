package main

import "fmt"

func main() {
	//n := 3 //Output : 4
	n := 4 // output: 10

	fmt.Println(friendPairing(n))
}

func friendPairing(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + i*dp[i-2]
	}
	return dp[n-1]
}
