package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	output := threeSum(arr)
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[0]); j++ {
			fmt.Print(output[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	size := len(nums)
	for i := 0; i < size; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := size - 1
		for j := i + 1; j < k; {

			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for ; j < k && nums[j] == nums[j-1]; j++ {
				}
				k--
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++

			}
		}
	}
	return result
}
