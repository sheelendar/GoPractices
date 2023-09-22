package main

import "math"

/*
Given two strings_and_two_strings_operations str1 and str2 and below operations that can be performed on str1. Find minimum number of edits (operations) required to convert ‘str1’ into ‘str2’.
    Insert
    Remove
    Replace

All of the above operations are of equal cost.
Examples:

    Input:   str1 = “geek”, str2 = “gesek”
    Output:  1
    Explanation: We can convert str1 into str2 by inserting a ‘s’.
*/

func main() {
	//str1 := "geek"
	//str2 := "gesek"
	str1 := "food"
	str2 := "money"
	output := editDistance(str1, str2, len(str1), len(str2))
	println("minimum changes: ", output)
}

func editDistance(str1 string, str2 string, l1 int, l2 int) int64 {
	dp := make([][]int64, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int64, l2+1)
	}
	for i := 0; i < l1; i++ {
		dp[i][0] = int64(i)
	}
	for i := 0; i < l2; i++ {
		dp[0][i] = int64(i)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = int64(math.Min(math.Min(float64(dp[i-1][j-1]), float64(dp[i][j-1])), float64(dp[i-1][j]))) + 1
			}
		}
	}
	displayArray(dp, l1, l2)
	return dp[l1][l2]
}

func displayArray(dp [][]int64, l1, l2 int) {
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			print(dp[i][j])
			print(" ")
		}
		println()
	}
}
