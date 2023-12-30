package main

import (
	"fmt"
	"math"
)

/*
You want to schedule a list of jobs in d days. Jobs are dependent (i.e To work on the ith job, you have to finish all the
jobs j where 0 <= j < i).
You have to finish at least one task every day. The difficulty of a job schedule is the sum of difficulties of each day of
 the d days. The difficulty of a day is the maximum difficulty of a job done on that day.
You are given an integer array jobDifficulty and an integer d. The difficulty of the ith job is jobDifficulty[i].
Return the minimum difficulty of a job schedule. If you cannot find a schedule for the jobs return -1.

Example 1:
Input: jobDifficulty = [6,5,4,3,2,1], d = 2
Output: 7
Explanation: First day you can finish the first 5 jobs, total difficulty = 6.
Second day you can finish the last job, total difficulty = 1.
The difficulty of the schedule = 6 + 1 = 7

Example 2:
Input: jobDifficulty = [9,9,9], d = 4
Output: -1
Explanation: If you finish a job per day you will still have a free day. you cannot find a schedule for the given jobs.

Example 3:
Input: jobDifficulty = [1,1,1], d = 3
Output: 3
Explanation: The schedule is one job per day. total difficulty will be 3.
*/

func main() {
	jobDifficulty := []int{6, 5, 4, 3, 2, 1}
	d := 2
	fmt.Println(minDifficulty(jobDifficulty, d))
}

func minDifficulty(diff []int, D int) int {
	N := len(diff)
	if N < D {
		return -1
	}
	dp := make([][]int, D)
	for i := 0; i < D; i++ {
		dp[i] = make([]int, N)
	}
	dp[0][0] = diff[0]
	for j := 1; j < N; j++ {
		dp[0][j] = Max(diff[j], dp[0][j-1])
	}
	for d := 1; d < D; d++ {
		for j := d; j < N; j++ {
			localMax := diff[j]
			dp[d][j] = math.MaxInt
			for schedule := j; schedule >= d; schedule-- {
				localMax = Max(localMax, diff[schedule])
				dp[d][j] = Min(dp[d][j], dp[d-1][schedule-1]+localMax)
			}
		}
	}
	return dp[D-1][N-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
