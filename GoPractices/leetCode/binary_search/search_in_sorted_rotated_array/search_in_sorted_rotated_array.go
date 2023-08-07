package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(arr, 0))
}

func search(nums []int, target int) int {
	return searchUtil(nums, 0, len(nums)-1, target)
}

func searchUtil(nums []int, l, h, target int) int {
	for l <= h {
		mid := (l + h) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if target >= nums[l] && target < nums[mid] {
				h = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[h] {
				l = mid + 1
			} else {
				h = mid - 1
			}
		}
	}
	return -1
}
