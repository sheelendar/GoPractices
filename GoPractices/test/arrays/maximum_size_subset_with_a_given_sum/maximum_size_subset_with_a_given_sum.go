package main

import (
	"fmt"
	"math"
)

func main() {
	set := []int{2, 3, 5, 7, 10, 15}
	sum := 10
	fmt.Print("max size subset ", maximumSumSize(set, sum))
}

func maximumSumSize(set []int, sum int) int {
	size := len(set)
	dp := make([][]bool, sum+1)
	dpCount := make([][]int, sum+1)
	for i := 0; i < sum+1; i++ {
		dp[i] = make([]bool, size+1)
		dpCount[i] = make([]int, size+1)
	}
	for i := 0; i <= size; i++ {
		dp[0][i] = true
		dpCount[0][i] = 0
	}

	for i := 1; i <= sum; i++ {
		dp[i][0] = false
		dpCount[i][0] = -1
	}

	for i := 1; i <= sum; i++ {
		for j := 1; j <= size; j++ {
			dp[i][j] = dp[i][j-1]
			dpCount[i][j] = dpCount[i][j-1]

			if i >= set[j-1] {
				dp[i][j] = dp[i][j] || dp[i-set[j-1]][j-1]

				if dp[i][j] {
					dpCount[i][j] = int(math.Max(float64(dpCount[i][j-1]), float64(dpCount[i-set[j-1]][j-1]+1)))
				}
			}
		}
	}
	return dpCount[sum][size]
}
