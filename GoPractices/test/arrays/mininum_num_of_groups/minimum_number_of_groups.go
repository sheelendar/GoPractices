package main

import "fmt"

func main() {

	//nums := []int{10, 10, 10, 3, 1, 1}
	nums := []int{1, 1, 3, 3, 1, 1, 2, 2, 3, 1, 3, 2}
	fmt.Println(minGroupsForValidAssignment(nums))
}

func minGroupsForValidAssignment(nums []int) int {
	size := len(nums)
	dp := make(map[int]int)
	min := 999999
	for i := 0; i < size; i++ {
		dp[nums[i]]++
	}
	for _, fre := range dp {
		if fre < min {
			min = fre
		}
	}

	min = min + 1
	count := 0
	for _, fre := range dp {
		count += fre / min
		if fre%min > 0 {
			count = count + 1
		}
	}
	return count
}
