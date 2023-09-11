package main

import "fmt"

func main() {
	S1 := "AGGTAB"
	S2 := "GXTXAYB" // output: 4
	fmt.Println(longestSubSequence(S1, S2))
}

func longestSubSequence(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
