package main

import "fmt"

func main() {
	a := []int64{2, -3, 4, -1, -2, 1, 5, -3}
	var largeSum int64
	largeSum = getLargestSumOfContinuesSubArray(a, len(a))
	fmt.Println("sum:", largeSum)
}

func getLargestSumOfContinuesSubArray(arr []int64, size int) int64 {
	largeSum := int64(0)
	tempSum := int64(0)
	for item := range arr {
		tempSum = tempSum + arr[item]
		if tempSum > largeSum {
			largeSum = tempSum
		}
		if tempSum < 0 {
			tempSum = 0
		}
	}
	return largeSum
}
