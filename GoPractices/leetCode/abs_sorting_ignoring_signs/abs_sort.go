package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-1, 0, 2, 1, 1, 4, 3}
	absSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
}
func absSort(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {

		for j := i + 1; j < size; j++ {

			if math.Abs(float64(arr[j])) < math.Abs(float64(arr[i])) {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp

			} else if math.Abs(float64(arr[j])) == math.Abs(float64(arr[i])) {
				if arr[i] > arr[j] {
					temp := arr[j]
					arr[j] = arr[i]
					arr[i] = temp
				}
			}
		}
	}
}
