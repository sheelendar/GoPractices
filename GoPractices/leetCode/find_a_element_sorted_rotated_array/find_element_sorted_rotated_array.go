package main

import "fmt"

func main() {
	//num := []int{3, 1, 2}
	//num := []int{3, 4, 5, 1, 2}
	//num := []int{2, 1}
	//num := []int{2, 3, 4, 5, 1}
	num := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Print(findMin(num))
}

func findMin(nums []int) int {
	h := len(nums)
	l := 0
	if h == 1 {
		return nums[0]
	}
	return findMinElement(nums, l, h-1)
}

func findMinElement(nums []int, l int, h int) int {
	if nums[l] == nums[h] {
		return nums[l]
	}
	if l+1 == h && nums[l] < nums[h] {
		return nums[l]
	}
	if l+1 == h && nums[h] <= nums[l] {
		return nums[h]
	}
	mid := (l + h) / 2
	if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
		return nums[mid]
	}
	if nums[mid] > nums[mid+1] && nums[mid] < nums[mid-1] {
		return findMinElement(nums, l, mid-1)
	} else {
		return findMinElement(nums, mid+1, h)
	}
}
