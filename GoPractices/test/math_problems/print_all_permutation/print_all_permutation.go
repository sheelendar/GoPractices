package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(permute(arr))
}

func permute(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	if n == 0 {
		return result
	}
	helper(nums, 0, n-1, &result)
	return result
}

func helper(nums []int, l, r int, result *[][]int) {
	if l == r {
		arr := make([]int, len(nums))
		copy(arr, nums)
		*result = append(*result, arr)
	} else {
		for i := l; i <= r; i++ {
			nums[i], nums[l] = nums[l], nums[i]
			helper(nums, l+1, r, result)
			nums[i], nums[l] = nums[l], nums[i]
		}
	}
}
