package main

import (
	"fmt"
	"math"
)

/*
Given a rod of length n inches and an array of prices that includes prices of all pieces of size smaller than n.
Determine the maximum value obtainable by cutting up the rod and selling the pieces. For example,
if the length of the rod is 8 and the values of different pieces are given as the following,
then the maximum obtainable value is 22 (by cutting in two pieces of lengths 2 and 6)

length   | 1   2   3   4   5   6   7   8
--------------------------------------------
price    | 1   5   8   9  10  17  17  20
*/

func main() {
	arr := []int{1, 5, 8, 9, 10, 17, 17, 20}
	size := len(arr)
	fmt.Println("Maximum profit is ", cuttingRodInProfit(arr, size))
}

func cuttingRodInProfit(arr []int, size int) int {
	if size == 0 {
		return 0
	}
	if size == 1 {
		return arr[0]
	}
	dp := make([]int, size+1)
	for i := 1; i <= size; i++ {
		maxProfit := 0
		for j := 0; j < i; j++ {
			maxProfit = int(math.Max(float64(maxProfit), float64(arr[j]+dp[i-j-1])))
		}
		dp[i] = maxProfit
	}
	return dp[size]
}
