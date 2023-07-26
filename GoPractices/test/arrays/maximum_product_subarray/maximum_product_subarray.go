package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, -2, -3, 0, 7, -8, -2}
	fmt.Println("maximum product is", maximumProductSubArray(arr, len(arr)))
}

func maximumProductSubArray(arr []int, size int) int {
	if size <= 0 {
		return 0
	}
	maximum_product := 1
	minimum_product := 1
	max_max_product := 0

	for i := 0; i < size; i++ {

		if arr[i] > 0 {
			maximum_product = maximum_product * arr[i]
			minimum_product = int(math.Min(float64(minimum_product*arr[i]), float64(1)))

		} else if arr[i] < 0 {
			temp := maximum_product
			maximum_product = int(math.Max(float64(minimum_product*arr[i]), float64(1)))
			minimum_product = temp * arr[i]

		} else {
			maximum_product = 1
			minimum_product = 1
		}
		if max_max_product < maximum_product {
			max_max_product = maximum_product
		}
	}
	return max_max_product
}
