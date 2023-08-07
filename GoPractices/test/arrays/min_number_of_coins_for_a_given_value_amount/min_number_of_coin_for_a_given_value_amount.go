package main

import (
	"fmt"
	"math"
)

/*
Given a value V, if we want to make a change for V cents, and we have an infinite supply of each of C = { C1, C2, .., Cm} valued coins,
what is the minimum number of coins to make the change? If it’s not possible to make a change, print -1.
Examples:(https://www.geeksforgeeks.org/find-minimum-number-of-coins-that-make-a-change/)
Input: coins[] = {25, 10, 5}, V = 30

	Output: Minimum 2 coins required We can use one coin of 25 cents and one of 5 cents
	Input: coins[] = {9, 6, 5, 1}, V = 11
	Output: Minimum 2 coins required We can use one coin of 6 cents and 1 coin of 5 cents
*/

func main() {
	coins := []int{1, 2, 5}
	m := len(coins)
	V := 11
	dp := make(map[int]int)
	fmt.Print(minChange(coins, m, V, dp))
	fmt.Println()
	fmt.Println(minCoinsChanges(coins, len(coins), V))
}

// dp solution
func minCoinsChanges(coins []int, n int, V int) int {
	if V == 0 {
		return 0
	}
	if coins == nil {
		return -1
	}
	dp := make([]int, V+1)
	for i := 1; i <= V; i++ {
		dp[i] = math.MaxInt
	}
	//n := len(coins)
	for i := 1; i <= V; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				res := dp[i-coins[j]]
				if res != math.MaxInt && res+1 < dp[i] {
					dp[i] = res + 1
				}
			}
		}
	}
	if dp[V] == math.MaxInt {
		return -1
	}
	return dp[V]
}

// Recursion
func minChange(coins []int, n int, v int, m map[int]int) int {
	res := minimumNumberOfCoin(coins, n, v, m)
	if res == math.MaxInt {
		return -1
	}
	return res
}
func minimumNumberOfCoin(coins []int, n int, v int, m map[int]int) int {
	if coins == nil {
		return 0
	}
	if v == 0 {
		return 0
	}
	if _, ok := m[v]; ok {
		return m[v]
	}
	res := math.MaxInt
	for i := 0; i < n; i++ {
		if coins[i] <= v {
			max := minimumNumberOfCoin(coins, n, v-coins[i], m)
			if max != math.MaxInt && max+1 < res {
				res = max + 1
			}
		}
	}
	m[v] = res
	return res
}
