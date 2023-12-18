package main

import (
	"fmt"
)

/*
You are given an integer array nums. Two players are playing a game with this array: player 1 and player 2.

Player 1 and player 2 take turns, with player 1 starting first. Both players start the game with a score of 0.
 At each turn, the player takes one of the numbers from either end of the array (i.e., nums[0] or nums[nums.length - 1])
 which reduces the size of the array by 1. The player adds the chosen number to their score. The game ends when there
  are no more elements in the array.

Return true if Player 1 can win the game. If the scores of both players are equal, then player 1 is still the winner,
and you should also return true. You may assume that both players are playing optimally.

Example 1:

Input: nums = [1,5,2]
Output: false
Explanation: Initially, player 1 can choose between 1 and 2.
If he chooses 2 (or 1), then player 2 can choose from 1 (or 2) and 5. If player 2 chooses 5, then player 1 will be left with 1 (or 2).
So, final score of player 1 is 1 + 2 = 3, and player 2 is 5.
Hence, player 1 will never be the winner and you need to return false.

*/

func main() {
	a := []int{1, 5, 2}
	fmt.Print(predictTheWinner(a))
}

func predictTheWinner1(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}
	if n == 1 {
		if nums[0] > 0 {
			return true
		}
		return false
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for g := 0; g < n; g++ {
		for i, j := 0, g; i < n-1 && g < n-1; i, j = i+1, j+1 {
			if g == 0 {
				dp[i][j] = nums[i]
			} else if g == 1 {
				dp[i][j] = max(nums[i], nums[j])
			} else {
				v1 := Min(dp[i+2][j], dp[i+1][j-1])
				v2 := Min(dp[i+1][j-1], dp[i][j-2])
				dp[i][j] = max(v1, v2)
			}
		}
	}
	return dp[0][n-1] > 0
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//{1, 5, 2}

func predictTheWinner(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}
	if n == 1 {
		if nums[0] >= 0 {
			return true
		}
		return false
	}
	return helper(0, n-1, 0, 0, nums, true)
}

func helper(l, h, sum1, sum2 int, a []int, t bool) bool {
	if l == h {
		if t {
			if sum1+a[l] >= sum2 {
				return true
			} else {
				return false
			}
		} else {
			if sum2+a[l] <= sum1 {
				return true
			} else {
				return false
			}
		}
	}
	if t {
		return helper(l+1, h, sum1+a[l], sum2, a, !t) || helper(l, h-1, sum1+a[h], sum2, a, !t)
	} else {
		return helper(l+1, h, sum1, sum2+a[l], a, !t) && helper(l, h-1, sum1, sum2+a[h], a, !t)
	}
}
