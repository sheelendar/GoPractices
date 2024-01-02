/*
	There are n points on a road you are driving your taxi on.
	The n points on the road are labeled from 1 to n in the direction you are going,
	 and you want to drive from point 1 to point n to make money by picking up passengers.
	 You cannot change the direction of the taxi.

The passengers are represented by a 0-indexed 2D integer array rides, where rides[i] = [starti, endi, tipi] denotes the ith passenger requesting a ride from point starti to point endi who is willing to give a tipi dollar tip.
For each passenger i you pick up, you earn endi - starti + tipi dollars. You may only drive at most one passenger at a time.
Given n and rides, return the maximum number of dollars you can earn by picking up the passengers optimally.
Note: You may drop off a passenger and pick up a different passenger at the same point.

Example 1: https://leetcode.com/problems/maximum-earnings-from-taxi/description/
Input: n = 5, rides = [[2,5,4],[1,5,1]],     Output: 7
Explanation: We can pick up passenger 0 to earn 5 - 2 + 4 = 7 dollars.

Example 2:
Input: n = 20, rides = [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]    Output: 20
Ans:
*/
package main

import "sort"

func maxTaxiEarnings(N int, rides [][]int) int64 {
	size := len(rides)
	price := make([]int, N+1)

	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})
	j := 0
	for i := 1; i <= N; i++ {
		price[i] = price[i-1]
		for j < size && rides[j][1] == i {
			price[i] = Max(price[i], rides[j][1]-rides[j][0]+rides[j][2]+price[rides[j][0]])
			j++
		}
	}
	return int64(price[N])
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
