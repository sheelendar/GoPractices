package main

import "fmt"

func main() {
	nums := []int{0, 4, 3, 0}
	target := 0
	fmt.Printf("Hello Pramp", twoSum(nums, target))
}

func twoSum(nums []int, t int) []int {
	n := len(nums)
	if n <= 0 {
		return nil
	}
	var r []int
	dp := make(map[int]int)

	for i := 0; i < n; i++ {
		diff := t - nums[i]
		if _, ok := dp[diff]; !ok {
			dp[nums[i]] = i
		} else {
			v, _ := dp[diff]
			r = append(r, v)
			r = append(r, i)

		}
	}
	if t == 0 && r == nil {
		for i := 0; i < n; i++ {
			if nums[i] == 0 {
				r = append(r, i)
			}
		}
		return r
	}
	return r
}
