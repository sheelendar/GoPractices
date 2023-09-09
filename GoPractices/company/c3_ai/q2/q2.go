package main

import "fmt"

func main() {
	prices := []int32{2, 3, 5, 1, 1, 2, 1}
	money := int32(7)
	fmt.Println(getMaxColors(prices, money))
}

func getMaxColors(prices []int32, money int32) int32 {
	start := int32(0)
	end := int32(0)
	max := int32(0)
	sum := int32(0)
	n := int32(len(prices))

	for end < n {
		sum += prices[end]
		for sum > money && start < n {
			sum -= prices[start]
			start++
		}
		max = maxFunc(max, end-start+1)
		end++
	}
	return max
}

func maxFunc(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
