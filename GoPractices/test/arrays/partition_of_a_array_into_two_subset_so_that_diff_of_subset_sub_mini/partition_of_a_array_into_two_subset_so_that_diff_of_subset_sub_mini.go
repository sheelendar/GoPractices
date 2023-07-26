package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 1, 4, 2, 2, 1}
	n := len(arr)
	dp := make(map[string]int)
	totalSum := 0
	for i := 0; i < n; i++ {
		totalSum = totalSum + arr[i]
	}
	fmt.Print("The minimum difference between two sets is ", findMinDiffOfTwoSubsetSums(arr, n, 0, totalSum, dp))
}

func findMinDiffOfTwoSubsetSums(arr []int, n, sum, totalSum int, dp map[string]int) int {
	// make a key for n and sum because both are changing every time.
	key := fmt.Sprintf("%d_%d", n, sum)
	if n == 0 {
		min := int(math.Abs(float64((totalSum - sum) - sum)))
		dp[key] = min
		return min
	}

	if v, ok := dp[key]; ok {
		return v
	}
	// include current number into sum
	diff1 := findMinDiffOfTwoSubsetSums(arr, n-1, sum+arr[n-1], totalSum, dp)
	// not include current number into sum
	diff2 := findMinDiffOfTwoSubsetSums(arr, n-1, sum, totalSum, dp)
	min := int(math.Min(float64(diff1), float64(diff2)))
	dp[key] = min
	return min
}
