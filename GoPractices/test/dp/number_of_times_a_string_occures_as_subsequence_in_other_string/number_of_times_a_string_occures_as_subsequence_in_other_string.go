package main

import "fmt"

func main() {

	//a := "GeeksforGeeks"  // output: 4
	//b := "Gks"

	a := "geeksforgeeks" // output: 9
	b := "geeks"
	fmt.Println(countSubSequences(a, b))
}

func countSubSequences(a string, b string) int {
	m := len(a)
	n := len(b)
	if m == 0 || n == 0 {
		if n == 0 {
			return 1
		}
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
