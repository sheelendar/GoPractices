package main

import "fmt"

func main() {
	num := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(num))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	lis := make([]int, n)
	lis[0] = 1
	for i := 1; i < n; i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && lis[j]+1 > lis[i] {
				lis[i] = lis[j] + 1
			}
		}
	}
	max := 1
	for i := 0; i < n; i++ {
		if max < lis[i] {
			max = lis[i]
		}
	}
	return max
}
