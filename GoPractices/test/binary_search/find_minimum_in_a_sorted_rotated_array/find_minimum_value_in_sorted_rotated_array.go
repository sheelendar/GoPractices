package main

import (
	"fmt"
	"math"
)

func main() {
	//num := []int{3, 1, 2}
	//num := []int{3, 4, 5, 1, 2}
	//num := []int{2, 1}
	//num := []int{2, 3, 4, 5, 1}
	//	num := []int{4, 5, 6, 7, 0, 1, 2}
	num := []int{7, 8, 1, 2, 3, 4, 5, 6}
	size := len(num)
	fmt.Print("min: ", findMinimumElement(num, 0, size-1))
	fmt.Println()
	fmt.Print("max: ", findMinElement(num, 0, size-1))
}

func findMinimumElement(nums []int, l int, h int) int {
	//	base conditions
	if h < 0 {
		return -1
	}
	if h == 0 {
		return nums[h]
	}
	// ans as default
	ans := math.MaxInt
	// base condition as standard
	for l <= h {
		mid := (l + h) / 2
		// check if first part is sorted then take min element from this part and discard all other elements.
		if nums[l] <= nums[mid] {
			ans = int(math.Min(float64(ans), float64(nums[l])))
			l = mid + 1
		} else {
			// if first part is not sorted then definitely second part will be sorted and take min element and discard other elements.
			ans = int(math.Min(float64(ans), float64(nums[mid])))
			h = mid - 1
		}
	}
	return ans
}

func findMinElement(nums []int, l int, h int) int {
	// base condition
	if h < 0 {
		return -1
	}
	if h == 0 {
		return nums[0]
	}
	// base condition as standard
	for l < h {
		mid := (l + h) / 2
		//	if array is sorted and rotated then nums[mid] > nums[h] is true then minimum exist in second part.
		if nums[mid] > nums[h] {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return nums[l]
}
