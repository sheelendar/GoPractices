package main

/*
Given a value V, if we want to make a change for V cents, and we have an infinite supply of
each of C = { C1, C2, .., Cm} valued coins, what is the minimum number of coins to make the change?
If it’s not possible to make a change, print -1.

Examples:(https://www.geeksforgeeks.org/find-minimum-number-of-coins-that-make-a-change/)

Input: coins[] = {25, 10, 5}, V = 30
	Output: Minimum 2 coins required We can use one coin of 25 cents and one of 5 cents
	Input: coins[] = {9, 6, 5, 1}, V = 11
	Output: Minimum 2 coins required We can use one coin of 6 cents and 1 coin of 5 cents


*/
import (
	"fmt"
	"math"
)

func main() {
	//coins := []int{25, 10, 5}
	//V := 30
	coins := []int{8, 4, 5, 1}
	V := 11
	fmt.Print(minNumberOfChange(coins, 0, len(coins), V))
	fmt.Print(minNumberOfChangeDP(coins, len(coins), V))
}

func minNumberOfChangeDP(coins []int, n, v int) int {
	if v == 0 || n == 0 {
		return 0
	}
	dp := make([]int, v+1)
	dp[0] = 0

	for i := 1; i <= v; i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i <= v; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				res := dp[i-coins[j]]
				if res != math.MaxInt && res+1 < dp[i] {
					dp[i] = res + 1
				}

			}
		}
	}
	return dp[v]
}
func minNumberOfChange(coins []int, l, h int, v int) int {
	if v == 0 {
		return 0
	}
	res := 999999
	for i := l; i < h; i++ {
		if coins[i] <= v {
			min := minNumberOfChange(coins, i, h, v-coins[i])
			if min != 999999 && min+1 < res {
				res = min + 1
			}
		}
	}
	return res
}
