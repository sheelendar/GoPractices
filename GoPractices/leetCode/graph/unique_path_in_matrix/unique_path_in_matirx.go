package main

import "fmt"

func uniquePaths(m int, n int, dp map[string]int) int {
	if m == 1 || n == 1 {
		return 1
	}
	if v, ok := dp[fmt.Sprintf("%d_%d", m, n)]; ok {
		return v
	}
	res := uniquePaths(m-1, n, dp) + uniquePaths(m, n-1, dp)
	dp[fmt.Sprintf("%d_%d", m, n)] = res
	return res

}

func uniquePathsIterative(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7, map[string]int{}))
	fmt.Println(uniquePathsIterative(3, 7))
}
