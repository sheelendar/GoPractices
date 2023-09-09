package main

import (
	"fmt"
	"math"
)

/*
We are given hours, a list of the number of hours worked per day for a given employee.
A day is considered to be a tiring day if and only if the number of hours worked is (strictly) greater than 8.
A well-performing interval is an interval of days for which the number of tiring days is strictly larger than the number of non-tiring days.
Return the length of the longest well-performing interval.
Example 1:

Input: hours = [9,9,6,0,6,6,9]
Output: 3
Explanation: The longest well-performing interval is [9,9,6].

Example 2:
Input: hours = [6,6,6]
Output:
*/

func main() {
	//hours := []int{9, 9, 6, 0, 6, 6, 9} // output:3
	///hours := []int{9, 9, 9}                               // output:3
	hours := []int{12, 8, 7, 7, 9, 10, 8, 7, 9, 7, 8, 11} // output:5
	//hours := []int{6, 6, 6} //output:0
	//hours := []int{6, 6, 9} //output:1
	//hours := []int{4, 11, 3, 4} //output :1
	fmt.Println(longestWPI(hours))
}

func longestWPI(hours []int) int {
	n := len(hours)
	if n == 0 {
		return 0
	}
	dp := make(map[int]int)
	sum := 0
	max := 0
	for i := 0; i < n; i++ {
		if hours[i] > 8 {
			sum = sum + 1
		} else {
			sum = sum - 1
		}
		if sum > 0 {
			max = i + 1
		} else {
			if _, ok := dp[sum]; !ok {
				dp[sum] = i
			}
			if index, ok := dp[sum-1]; ok {
				max = int(math.Max(float64(max), float64(i-index)))
			}
		}
	}
	return max
}
