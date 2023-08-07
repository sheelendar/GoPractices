package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum4(nums, target))
}

func combinationSum4(nums []int, target int) int {
	return noOfWaysSum(nums, len(nums), target)
	//return sumUtil(nums, len(nums), target)
}

func noOfWaysSum(nums []int, n, target int) int {
	if target == 0 {
		return 1
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < n; j++ {
			if nums[j] <= i {
				dp[i] = dp[i] + dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

// recurrsion solution is not good timeout will occurre.
func sumUtil(nums []int, n, target int) int {
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}
	res := 0
	for i := 0; i < n; i++ {
		res = res + sumUtil(nums, n, target-nums[i])
	}
	return res
}
