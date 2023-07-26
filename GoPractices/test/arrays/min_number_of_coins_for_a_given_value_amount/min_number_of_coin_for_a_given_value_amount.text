package main

import (
	"fmt"
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
	coins := []int{9, 6, 5, 1}
	m := len(coins)
	V := 11
	dp := make(map[int]int)
	fmt.Print(minimumNumberOfCoin(coins, m, V, dp))
}

func minimumNumberOfCoin(coins []int, n int, v int, m map[int]int) int {
	if v == 0 {
		return 0
	}
	if _, ok := m[v]; ok {
		return m[v]
	}
	res := 99999999
	for i := 0; i < n; i++ {
		if coins[i] <= v {
			max := minimumNumberOfCoin(coins, n, v-coins[i], m)
			if max+1 < res {
				res = max + 1
			}
		}
	}
	m[v] = res
	return res
}
