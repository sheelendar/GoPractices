package main

import "fmt"

func main() {
	a := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	fmt.Println(maxSubArray(a))
}

// second solution.
func maxSubArray(nums []int) []int {
	curSum, maxSum := 0, -999999
	start := 0
	end := 0
	start1 := 0

	for i := 0; i < len(nums); i++ {
		curSum = curSum + nums[i]
		if curSum < nums[i] {
			curSum = nums[i]
			start = i
		}
		if maxSum < curSum {
			maxSum = curSum
			start1 = start
			end = i
		}
	}
	return []int{start1, end}
}
