package main

import (
	"fmt"
	"math"
)

func main() {
	num := []int{3, 1, 2}
	//num := []int{3, 4, 5, 1, 2}
	//num := []int{2, 1}
	//num := []int{2, 3, 4, 5, 1}
	//num := []int{4, 5, 6, 7, 0, 1, 2}
	//num := []int{7, 8, 1, 2, 3, 4, 5, 6}
	size := len(num)
	fmt.Print("min: ", findMaximumElement(num, 0, size-1))
	/*fmt.Println()
	fmt.Print("max: ", findMinElement(num, 0, size-1))*/
}

func findMaximumElement(num []int, l int, h int) int {
	//base conditions
	if h < 0 {
		return -1
	}
	if h == 0 {
		return num[h]
	}
	ans := math.MinInt
	// standard condition to check limit for binary search
	for l <= h {
		mid := (l + h) / 2
		// this condition means first part is sorted of array so will take max element and discard others
		if num[mid] >= num[l] {
			ans = int(math.Max(float64(ans), float64(num[mid])))
			l = mid + 1
		} else {
			// num[mid] >= num[l] is false means that second part is sorted of array so take max from sorted part and discard other element.
			ans = int(math.Max(float64(ans), float64(num[h])))
			h = mid - 1
		}
	}
	return ans
}
