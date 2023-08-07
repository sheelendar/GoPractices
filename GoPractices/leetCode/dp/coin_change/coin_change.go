package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{2}
	amount := 0
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	dp := make(map[string]int)
	result := coinUtil(coins, len(coins)-1, amount, dp)
	if result == math.MaxInt {
		return -1
	}
	return result
}

func coinUtil(coin []int, n, amount int, dp map[string]int) int {
	if coin == nil {
		return 0
	}
	if amount == 0 {
		return 0
	}
	if v, ok := dp[fmt.Sprintf("%d_%d", n, amount)]; ok {
		return v
	}
	minC := math.MaxInt
	for i := 0; i < len(coin); i++ {
		if amount >= coin[i] {
			res := coinUtil(coin, n, amount-coin[i], dp)
			if res != math.MaxInt && res+1 < minC {
				minC = res + 1
			}
		}
	}
	dp[fmt.Sprintf("%d_%d", n, amount)] = minC
	return minC
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
