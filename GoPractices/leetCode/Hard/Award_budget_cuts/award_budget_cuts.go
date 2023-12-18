package main

import (
	"fmt"
	"sort"
)

/*
The awards committee of your alma mater (i.e. your college/university) asked for your assistance with a budget
allocation problem they’re facing. Originally, the committee planned to give N research grants this year.
However, due to spending cutbacks, the budget was reduced to newBudget dollars and now they need to reallocate the grants.
 The committee made a decision that they’d like to impact as few grant recipients as possible by applying
 a maximum cap on all grants. Every grant initially planned to be higher than cap will now be exactly cap dollars.
 Grants less or equal to cap, obviously, won’t be impacted.

Given an array grantsArray of the original grants and the reduced budget newBudget, write a function
findGrantsCap that finds in the most efficient manner a cap such that the least number of recipients is
impacted and that the new budget constraint is met (i.e. sum of the N reallocated grants equals to newBudget).

Analyze the time and space complexities of your solution.
*/

func main() {
	grantsArray := []float64{2, 100, 50, 120, 1000}
	newBudget := float64(190)
	//grantsArray := []float64{2, 4, 6}
	//newBudget := float64(3)
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
