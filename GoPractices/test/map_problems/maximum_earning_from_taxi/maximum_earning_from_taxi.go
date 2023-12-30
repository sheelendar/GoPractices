package main

import "fmt"

/*

There are n points on a road you are driving your taxi on. The n points on the road are labeled from 1 to n in the
 direction you are going, and you want to drive from point 1 to point n to make money by picking up passengers.
 You cannot change the direction of the taxi.

The passengers are represented by a 0-indexed 2D integer array rides, where rides[i] = [starti, endi, tipi]
 denotes the ith passenger requesting a ride from point starti to point endi who is willing to give a tipi dollar tip.

For each passenger i you pick up, you earn endi - starti + tipi dollars. You may only drive at most one passenger at a time.

Given n and rides, return the maximum number of dollars you can earn by picking up the passengers optimally.

Note: You may drop off a passenger and pick up a different passenger at the same point.
https://leetcode.com/problems/maximum-earnings-from-taxi/description/


Example 1:
Input: n = 5, rides = [[2,5,4],[1,5,1]]
Output: 7
Explanation: We can pick up passenger 0 to earn 5 - 2 + 4 = 7 dollars.

Example 2:
Input: n = 20, rides = [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
Output: 20
Explanation: We will pick up the following passengers:
- Drive passenger 1 from point 3 to point 10 for a profit of 10 - 3 + 2 = 9 dollars.
- Drive passenger 2 from point 10 to point 12 for a profit of 12 - 10 + 3 = 5 dollars.
- Drive passenger 5 from point 13 to point 18 for a profit of 18 - 13 + 1 = 6 dollars.
We earn 9 + 5 + 6 = 20 dollars in total.
*/

func main() {
	rides := [][]int{{2, 3, 6}, {8, 9, 8}, {5, 9, 7}, {8, 9, 1}, {2, 9, 2}, {9, 10, 6}, {7, 10, 10}, {6, 7, 9}, {4, 9, 7}, {2, 3, 1}}
	n := 10
	fmt.Println(maxTaxiEarnings(n, rides))
}

func maxTaxiEarnings(N int, rides [][]int) int64 {
	size := len(rides)
	price := make([]int, N+1)
	dp := make(map[int][][]int)
	for d := 0; d < size; d++ { // collect all ending points at a distance d
		dp[rides[d][1]] = append(dp[rides[d][1]], rides[d])
	}
	for dis := 1; dis <= N; dis++ {
		max := price[dis-1]
		if allEndingPoints, ok := dp[dis]; ok { // check if any points end at distance dis
			for _, point := range allEndingPoints { // if found then check which one is max
				max = Max(max, point[1]-point[0]+point[2]+price[point[0]])
			}
		}
		price[dis] = max
	}
	return int64(price[N])
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
