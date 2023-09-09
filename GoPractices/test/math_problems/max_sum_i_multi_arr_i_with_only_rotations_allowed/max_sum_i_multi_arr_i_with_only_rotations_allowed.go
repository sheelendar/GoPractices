package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9} //Output: 330
	fmt.Println(maxSumWithRotation(arr))
}

func maxSumWithRotation(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	sum := 0
	currSum := 0
	max := 0
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
		currSum = currSum + i*arr[i]
	}
	max = currSum
	for i := 1; i < n; i++ {
		currSum = sum - n*arr[n-i] + currSum
		if currSum > max {
			max = currSum
		}
	}
	return max
}
