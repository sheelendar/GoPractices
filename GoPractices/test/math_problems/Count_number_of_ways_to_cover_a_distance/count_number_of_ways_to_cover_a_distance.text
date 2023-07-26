package main

import "fmt"

/*
Given a distance ‘dist’, count total number of ways to cover the distance with 1, 2 and 3 steps.
Examples: (https://www.geeksforgeeks.org/count-number-of-ways-to-cover-a-distance/)
Input: n = 3
Output: 4
Explanation:
Below are the four ways

	1 step + 1 step + 1 step
	1 step + 2 step
	2 step + 1 step
	3 step

Input: n = 4
Output: 7
Explanation:
Below are the four ways

	1 step + 1 step + 1 step + 1 step
	1 step + 2 step + 1 step
	2 step + 1 step + 1 step
	1 step + 1 step + 2 step
	2 step + 2 step
	3 step + 1 step
	1 step + 3 step
*/

func main() {
	dist := 4
	fmt.Print("Number of ways to complete distance ", countNumberOfWays(dist))
}

func countNumberOfWays(dist int) int {
	if dist <= 1 {
		return 1
	}
	res := make([]int, dist+1)
	res[0] = 1
	res[1] = 1
	if dist > 2 {
		res[2] = 2
	}
	// here number of stapes is given 3 that's Y we calculate till three steps res[i-1] + res[i-2] + res[i-3]
	// if only two steps are given than we will calculate till two steps res[i-1] + res[i-2]
	// if n steps is given then calculate till n res[i-1] + res[i-2] + res[i-3]...... [i-n]
	for i := 3; i <= dist; i++ {
		res[i] = res[i-1] + res[i-2] + res[i-3]
	}
	return res[dist]
}
