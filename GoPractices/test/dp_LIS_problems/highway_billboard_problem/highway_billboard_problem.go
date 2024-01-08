package main

import "fmt"

/*
Consider a highway of M miles. The task is to place billboards on the highway such that revenue is maximized.
The possible sites for billboards are given by number x1 < x2 < ….. < xn-1 < xn, specifying positions in

	miles measured from one end of the road. If we place a billboard at position xi, we receive a revenue of ri > 0.
	 There is a restriction that no two billboards can be placed within t miles or less than it.

Note : All possible sites from x1 to xn are in range from 0 to M as need to place billboards on a highway of M miles.
Examples:

Input : M = 20

	x[]       = {6, 7, 12, 13, 14}
	revenue[] = {5, 6, 5,  3,  1}
	t = 5

Output: 10
By placing two billboards at 6 miles and 12
miles will produce the maximum revenue of 10.

Input : M = 15

	x[] = {6, 9, 12, 14}
	revenue[] = {5, 6, 3, 7}
	t = 2

Output : 18
*/
func main() {
	M := 20
	x := []int{6, 7, 12, 13, 14}
	revenue := []int{5, 6, 5, 3, 1}
	D := 5

	/*M := 15
	x := []int{6, 9, 12, 14}
	revenue := []int{5, 6, 3, 7}
	D := 2*/
	fmt.Println(maxProfitByPlacingBillBoard(x, revenue, M, D))
	fmt.Println(maxProfitByBillBoard(x, revenue, M, D))
}

// time complexcity here M(given miles) and space complexcity is O(n) number board.
func maxProfitByBillBoard(loc, revenue []int, M, D int) int {
	n := len(loc)
	dp := make([]int, M)
	dp[0] = revenue[0]

	maxProfit := 0
	for i := 1; i < n; i++ {
		dp[i] = Max(dp[i-1], revenue[i])
		for j := 0; j < i; j++ {
			distence := loc[i] - loc[j]
			if distence > D {
				dp[i] = Max(dp[i], revenue[i]+dp[j])
			}
		}
		if dp[i] > maxProfit {
			maxProfit = dp[i]
		}
	}
	return maxProfit
}

// time complexcity here M(given miles) and space complexcity is O(n) number board.
func maxProfitByPlacingBillBoard(loc, revenue []int, M, D int) int {
	n := len(loc)
	dp := make([]int, M+1)
	dp[0] = 0
	boardLocation := make(map[int]int)
	for i := 0; i < n; i++ {
		boardLocation[loc[i]] = revenue[i]
	}
	maxProfit := 0
	for i := 1; i <= M; i++ {
		if rev, ok := boardLocation[i]; ok {
			p := rev
			if i-D-1 >= 0 {
				p += dp[i-D-1]
			}
			dp[i] = Max(dp[i-1], p)
		} else {
			dp[i] = dp[i-1]
		}
		if maxProfit < dp[i-1] {
			maxProfit = dp[i-1]
		}
	}
	return maxProfit

}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
