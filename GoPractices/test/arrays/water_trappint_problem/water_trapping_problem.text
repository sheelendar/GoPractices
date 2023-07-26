package main

import (
	"fmt"
	"math"
)

/*
Given an array of N non-negative integers arr[] representing an elevation map where the width of each bar is 1,
compute how much water it is able to trap after raining.(https://www.geeksforgeeks.org/trapping-rain-water/)
Examples:

	Input: arr[] = {2, 0, 2}
	Output: 2
	Explanation: The structure is like below.
	We can trap 2 units of water in the middle gap
*/
func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Print(maxWaterTrapping(arr, len(arr)))
}

func maxWaterTrapping(arr []int, size int) any {
	left := 0
	right := size - 1
	maxL := 0
	maxR := 0
	maxWater := 0
	// itrate untill left is less than or equal to right
	for left <= right {
		// make max as constant and calculate max water from other side where value is less
		// when right side is less value and felt has max value
		if arr[right] <= arr[left] {
			maxWater = maxWater + int(math.Max(float64(0), float64(maxR-arr[right])))
			maxR = int(math.Max(float64(maxR), float64(arr[right])))
			right--
		} else {
			// when left side is less value and right has max value
			maxWater = maxWater + int(math.Max(float64(0), float64(maxL-arr[left])))
			maxL = int(math.Max(float64(maxL), float64(arr[left])))
			left++
		}
	}
	return maxWater
}
