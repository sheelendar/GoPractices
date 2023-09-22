package main

import "fmt"

func main() {
	//price := []int{10, 22, 5, 75, 65, 80}
	//K := 2
	//Output: 87

	//price := []int{12, 14, 17, 10, 14, 13, 12, 15}
	price := []int{100, 30, 15, 10, 8, 25, 80}
	K := 3
	p := maxProfit(price, K)
	fmt.Print(p)
}

func maxProfit(price []int, k int) int {
	n := len(price)
	if n <= 1 {
		return 0
	}
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= k; i++ {
		dp[k][0] = 0
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 0
	}
	for i := 1; i <= k; i++ { // run till k txn first loop
		for j := 1; j < n; j++ { // run loop till the size of day array

			dp[i][j] = dp[i][j-1] // put default value one day back
			for m := 0; m <= j; m++ {
				if dp[i][j] < dp[i-1][m]+price[j]-price[m] { // check for every m days if
					dp[i][j] = dp[i-1][m] + price[j] - price[m]
				}
			}
		}
	}
	return dp[k][n-1]
}
