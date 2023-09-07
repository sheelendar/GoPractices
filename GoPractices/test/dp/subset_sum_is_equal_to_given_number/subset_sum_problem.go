package main

import "fmt"

func main() {
	set := []int{3, 34, 4, 12, 5, 2}
	sum := 9

	dp := make(map[string]bool)
	fmt.Println(isSubset(set, len(set), sum, dp))
	fmt.Print(isSubsetWithSum(set, len(set), sum))
}

func isSubset(set []int, n int, sum int, dp map[string]bool) bool {
	if n == 0 || sum == 0 {
		if sum == 0 {
			return true
		}
		return false
	}
	if v, ok := dp[fmt.Sprintf("%%", n, sum)]; ok {
		return v
	}
	if set[n-1] > sum {
		res := isSubset(set, n-1, sum, dp)
		dp[fmt.Sprintf("%%", n, sum)] = res
		return res
	}
	res := isSubset(set, n-1, sum, dp) || isSubset(set, n-1, sum-set[n-1], dp)
	dp[fmt.Sprintf("%%", n, sum)] = res
	return res

}
func isSubsetWithSum(set []int, n int, sum int) any {
	if n == 0 || sum == 0 {
		if sum == 0 {
			return true
		}
		return false
	}
	dp := make([][]bool, sum+1)
	for i := 0; i <= sum; i++ {
		dp[i] = make([]bool, n+1)
	}
	for i := 1; i <= sum; i++ {
		dp[i][0] = false
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = true
	}
	for i := 1; i <= sum; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i][j-1]
			if i >= set[j-1] {
				dp[i][j] = dp[i][j] || dp[i-set[j-1]][j-1]
			}
		}
	}
	return dp[sum][n]
}
