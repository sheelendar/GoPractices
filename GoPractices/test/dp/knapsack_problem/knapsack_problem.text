package main

import (
	"fmt"
	"math"
)

/*
0/1 Knapsack Problem

We are given N items where each item has some weight and profit associated with it.
We are also given a bag with capacity W, [i.e., the bag can hold at most W weight in it].
The target is to put the items into the bag such that the sum of profits associated with them is the maximum possible.
Note: The constraint here is we can either put an item completely into the bag or
cannot put it at all [It is not possible to put a part of an item into the bag].

	Input: N = 3, W = 4, profit[] = {1, 2, 3}, weight[] = {4, 5, 1}
	Output: 3
	Explanation: There are two items which have weight less than or equal to 4. If we select the item with weight 4, the possible profit is 1. And if we select the item with weight 1, the possible profit is 3. So the maximum possible profit is 3. Note that we cannot put both the items with weight 4 and 1 together as the capacity of the bag is 4.

	Input: N = 3, W = 3, profit[] = {1, 2, 3}, weight[] = {4, 5, 6}
	Output: 0
*/
func main() {
	profit := []int{60, 100, 120}
	weight := []int{10, 20, 30}
	W := 50
	dp := make(map[string]int)
	fmt.Println("total profit with wait", knapSackProblemsWithRecursion(W, weight, profit, len(profit), dp))
	fmt.Println("total profit with wait", knapSackProblems(W, weight, profit, len(profit)))
}

func knapSackProblemsWithRecursion(W int, weights []int, profits []int, n int, dp map[string]int) int {
	key := fmt.Sprintf("%d_%d", n, W)
	if W == 0 || n == 0 {
		return 0
	}
	if v, ok := dp[key]; ok {
		return v
	}
	profit := 0
	if W < weights[n-1] {
		profit = knapSackProblemsWithRecursion(W, weights, profits, n-1, dp)
		dp[key] = profit
		return profit
	}
	profit = int(math.Max(float64(profits[n-1]+knapSackProblemsWithRecursion(W-weights[n-1], weights, profits, n-1, dp)), float64(knapSackProblemsWithRecursion(W, weights, profits, n-1, dp))))
	dp[key] = profit
	return profit
}

func knapSackProblems(W int, weights []int, profits []int, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= W; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if weights[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(profits[i-1]+dp[i-1][j-weights[i-1]]), float64(dp[i-1][j])))
			}
		}
	}
	//display(dp, n, W)
	return dp[n][W]
}

func display(dp [][]int, n, W int) {
	for i := 0; i <= n; i++ {
		for j := 0; j < W; j++ {
			print(dp[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
