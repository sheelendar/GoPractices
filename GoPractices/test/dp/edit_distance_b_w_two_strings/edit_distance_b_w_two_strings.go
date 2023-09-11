package main

import (
	"fmt"
)

func main() {
	str1 := "sunday"
	str2 := "saturday"
	fmt.Println(editDistance(str1, str2))
}

func editDistance(str1 string, str2 string) int {
	m := len(str1)
	n := len(str2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i-1][j], Min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
