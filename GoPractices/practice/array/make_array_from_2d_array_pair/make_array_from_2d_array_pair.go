/*
There is an integer array nums that consists of n unique elements, but you have forgotten it. However, you do remember every pair of adjacent elements in nums.

You are given a 2D integer array adjacentPairs of size n - 1 where each adjacentPairs[i] = [ui, vi] indicates that the elements ui and vi are adjacent in nums.

It is guaranteed that every adjacent pair of elements nums[i] and nums[i+1] will exist in adjacentPairs, either as [nums[i], nums[i+1]] or [nums[i+1], nums[i]]. The pairs can appear in any order.

Return the original array nums. If there are multiple solutions, return any of them.



Example 1:

Input: adjacentPairs = [[2,1],[3,4],[3,2]]
Output: [1,2,3,4]
Explanation: This array has all its adjacent pairs in adjacentPairs.
Notice that adjacentPairs[i] may not be in left-to-right order.

Example 2:

Input: adjacentPairs = [[4,-2],[1,4],[-3,1]]
Output: [-2,4,1,-3]
Explanation: There can be negative numbers.
Another solution is [-3,1,4,-2], which would also be accepted.

Example 3:

Input: adjacentPairs = [[100000,-100000]]
Output: [100000,-100000]
*/

package main

import "fmt"

func main() {
	pairs := [][]int{{4, -10}, {-1, 3}, {4, -3}, {-3, 3}}

	fmt.Println(restoreArray(pairs))
}

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs)
	dp := make(map[int][][]int)

	for i := 0; i < n; i++ {
		dp[adjacentPairs[i][0]] = append(dp[adjacentPairs[i][0]], adjacentPairs[i])
		dp[adjacentPairs[i][1]] = append(dp[adjacentPairs[i][1]], adjacentPairs[i])
	}

	var res []int
	num := 0
	end := 0
	var pairStart [][]int
	var pairEnd [][]int
	for key, p := range dp {
		if len(p) == 1 {
			num = end
			end = key
			pairStart = pairEnd
			pairEnd = p
		}
	}
	res = append(res, num)
	num2 := pairStart[0][0]
	if pairStart[0][0] == num {
		num2 = pairStart[0][1]
	}
	delete(dp, num)
	num = num2
	//res=append(res,num)
	for dp[num] != nil {
		pair := dp[num]
		delete(dp, num)
		res = append(res, num)
		temp := pair[0][0]
		if num == pair[0][0] {
			temp = pair[0][1]
			if dp[temp] == nil && len(pair) > 1 {
				temp = pair[1][0]
				if num == pair[1][0] {
					temp = pair[1][1]
				}
			}
		} else {
			if dp[temp] == nil && len(pair) > 1 {
				temp = pair[1][0]
				if num == pair[1][0] {
					temp = pair[1][1]
				}
			}
		}
		num = temp
	}
	//res=append(res,num)
	return res
}
