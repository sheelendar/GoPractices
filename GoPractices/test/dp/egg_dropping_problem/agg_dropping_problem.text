package main

import (
	"fmt"
	"math"
)

func main() {

	agg := 2
	fl := 36
	dp := make([][]int, agg+1)
	for i := 0; i <= agg; i++ {
		dp[i] = make([]int, fl+1)
	}
	fmt.Println("Min attempt done", minAggDropAttempt(agg, fl, dp))
}

func minAggDropAttempt(agg int, fl int, dp [][]int) int {
	if agg == 1 {
		dp[agg][fl] = fl
		return fl
	}
	if fl == 0 || fl == 1 {
		dp[agg][fl] = fl
		return fl
	}
	if dp[agg][fl] != 0 {
		return dp[agg][fl]
	}
	minCount := 9999999999
	for i := 1; i <= fl; i++ {
		mexCount := int(math.Max(float64(minAggDropAttempt(agg-1, i-1, dp)), float64(minAggDropAttempt(agg, fl-i, dp))))
		if minCount > mexCount {
			minCount = mexCount
		}
	}

	dp[agg][fl] = minCount + 1
	return minCount + 1
}
