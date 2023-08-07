package main

// https://leetcode.com/problems/container-with-most-water/submissions/
import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Print(maxArea(arr))
}

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	maxWater := 0
	i := 0
	j := len(height) - 1
	sum := 0
	for i < j {
		if height[i] > height[j] {
			sum = height[j] * (j - i)
			j--
		} else if height[i] <= height[j] {
			sum = height[i] * (j - i)
			i++
		}
		maxWater = int(math.Max(float64(sum), float64(maxWater)))
	}
	return maxWater
}
