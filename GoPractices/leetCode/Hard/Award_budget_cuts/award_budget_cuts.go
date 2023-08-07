package main

import (
	"fmt"
	"sort"
)

func main() {
	/*grantsArray := []float64{2, 100, 50, 120, 1000}
	newBudget := float64(190)*/
	grantsArray := []float64{2, 4, 6}
	newBudget := float64(3)
	fmt.Print(FindGrantsCap(grantsArray, newBudget))
}

func FindGrantsCap(grantsArray []float64, newBudget float64) float64 {
	size := len(grantsArray)
	sort.Float64s(grantsArray)
	previousSum := float64(0)

	for i := 0; i < size; i++ {
		curr := grantsArray[i]
		available := newBudget - previousSum
		remainingIndex := float64(size - i)
		if (remainingIndex * curr) > available {
			return available / remainingIndex
		}
		previousSum = previousSum + curr
	}
	return grantsArray[size-1]
}

func FindGrantsCap1(grantsArray []float64, newBudget float64) float64 {
	currentSum := float64(0)
	l := float64(0)
	h := newBudget
	mid := float64(0)
	for l <= h {

		mid = (l + h) / 2
		currentSum = checkBudget(mid, grantsArray)
		if currentSum == newBudget {
			return mid
		}
		if currentSum > newBudget {
			h = mid
		} else {
			l = mid
		}
	}
	return mid
}

func checkBudget(mid float64, array []float64) float64 {
	currSum := float64(0)
	for i := 0; i < len(array); i++ {
		if array[i] <= mid {
			currSum = currSum + array[i]
		} else {
			currSum = currSum + mid
		}
	}
	return currSum
}
