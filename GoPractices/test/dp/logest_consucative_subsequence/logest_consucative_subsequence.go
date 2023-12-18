package main

import "fmt"

func main() {
	arr := []int{1, 0, -1} // 3
	fmt.Print(longestConsecutive1(arr))

	fmt.Println()
	arr1 := []int{100, 4, 200, 1, 3, 2} //4
	fmt.Print(longestConsecutive1(arr1))

	fmt.Println()
	arr2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive1(arr2))
}

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := 0
	dp := make(map[int]int)
	for i := 0; i < n; i++ {
		c := DFS(nums, nums[i], n, 1, dp)
		if max < c {
			max = c
		}
	}
	return max
}

func DFS(nums []int, k, n, count int, dp map[int]int) int {

	if v, ok := dp[k]; ok {
		return v
	}
	for i := 0; i < n; i++ {
		if k+1 == nums[i] {
			res := count + DFS(nums, nums[i], n, count, dp)
			dp[k] = res
			return res
		}
	}
	dp[k] = count
	return count
}

func longestConsecutive1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := nums[0]
	min := nums[0]
	dp := make(map[int]bool)

	for i := 0; i < n; i++ {
		dp[nums[i]] = true
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	count := 0
	maxCount := 0
	for i := min; i <= max; i++ {

		if _, ok := dp[i]; ok {
			count++
		} else {
			count = 0
		}
		if maxCount < count {
			maxCount = count
		}
	}

	return maxCount
}
