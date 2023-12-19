package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)

	for i := 0; i < len(result); i++ {
		fmt.Print("[")
		for j := 0; j < len(result[i]); j++ {
			fmt.Print(result[i][j])
		}
		fmt.Print("]")
	}

}

func subsets(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{{}}
	}
	result := [][]int{}
	r := []int{}
	helper(nums, 0, n, &result, r)

	return result
}

func helper(nums []int, k, n int, result *[][]int, r []int) {
	// need to perform copy otherwise orginal array can modify
	cpy := make([]int, len(r))
	copy(cpy, r)
	*result = append(*result, cpy)
	for i := k; i < n; i++ {
		r = append(r, nums[i])
		helper(nums, i+1, n, result, r)
		r = r[:len(r)-1]
	}
}
