package main

import "sort"

/*
We have n jobs, where every job is scheduled to be done from startTime[i] to endTime[i], obtaining a profit of profit[i].
You're given the startTime, endTime and profit arrays, return the maximum profit you can take such that there are no two jobs in the subset with overlapping time range.  https://leetcode.com/problems/maximum-profit-in-job-scheduling
If you choose a job that ends at time X you will be able to start another job that starts at time X.
Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
Output: 120
Explanation: The subset chosen is the first and fourth job.
Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.

*/

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	maxEnd := 0
	var arr []Job
	for i := 0; i < n; i++ {
		if maxEnd < endTime[i] {
			maxEnd = endTime[i]
		}
		arr = append(arr, Job{start: startTime[i], end: endTime[i], profit: profit[i]})
	}
	sort.Slice(arr, func(i, j int) bool { // sort bases on the end time
		return arr[i].end < arr[j].end
	})
	maxP := make([]int, n)
	maxProfit := 0
	maxP[0] = arr[0].profit
	for i := 1; i < n; i++ {
		maxP[i] = Max(maxP[i-1], arr[i].profit)
		for j := i - 1; j >= 0; j-- { /// using LIS here
			if arr[j].end <= arr[i].start {
				maxP[i] = Max(maxP[i], maxP[j]+arr[i].profit)
				break
			}
		}
		if maxProfit < maxP[i] {
			maxProfit = maxP[i]
		}
	}
	return maxProfit
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Job struct {
	start  int
	end    int
	profit int
}
