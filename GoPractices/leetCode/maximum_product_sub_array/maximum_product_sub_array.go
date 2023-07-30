package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-1, 0, -2}
	fmt.Print(maxProduct(nums))
}

func maxProduct(nums []int) int {
	size := len(nums)
	max_product := -999999
	if size <= 0 {
		return 0
	}
	product := 1
	for i := 0; i < size; i++ {
		product = product * nums[i]
		max_product = int(math.Max(float64(product), float64(max_product)))
		if nums[i] == 0 {
			product = 1
		}
	}
	product = 1
	for i := size - 1; i >= 0; i-- {
		product = product * nums[i]
		max_product = int(math.Max(float64(product), float64(max_product)))
		if nums[i] == 0 {
			product = 1
		}
	}
	return max_product
}
