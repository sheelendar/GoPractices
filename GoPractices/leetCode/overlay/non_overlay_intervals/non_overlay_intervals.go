package main

import (
	"fmt"
	"sort"
)

/*
Given an array of intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need
to remove to make the rest of the intervals non-overlapping.

Example 1:
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

Example 2:
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

Example 3:
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
*/

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 || n == 1 {
		return 0
	}
	// sort based on second value or end value
	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][1] < intervals[j][1])
	})
	// start count from 1
	count := 1
	data := intervals[0]
	for i := 1; i < n; i++ {
		// check how many pair are not overlap
		if data[1] <= intervals[i][0] {
			count++
			data = intervals[i]
		}
	}
	// minus not overlap pair from total and find overlap pair
	return n - count
}
