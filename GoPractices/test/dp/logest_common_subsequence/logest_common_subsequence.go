package main

import "fmt"

func main() {
	str1 := "abcde"
	str2 := "ace"

	fmt.Println(longestCommonSubsequence(str1, str2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	return commonUtil(text1, text2, len(text1), len(text2))
}

func commonUtil(str1, str2 string, m, n int) int {
	if n == 0 || m == 0 {
		return 0
	}
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {

		for j := 1; j < n+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
