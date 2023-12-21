package main

import (
	"fmt"
	"math"
)

/*
Given that there are N books and M students. Also given are the number of pages in each book in ascending order.
 The task is to assign books in such a way that the maximum number of pages assigned to a student is minimum,
 with the condition that every student is assigned to read some consecutive books. Print that minimum number of pages.

Example :

    Input: N = 4, pages[] = {12, 34, 67, 90}, M = 2

    Output: 113

*/

func main() {
	n := 4
	pages := []int{12, 34, 67, 90}
	m := 2

	fmt.Println(findNumberOfPages(pages, n, m))
}

func findNumberOfPages(pages []int, n, m int) int {
	if m > n {
		return -1
	}
	high := 0
	low := -1

	for i := 0; i < n; i++ {
		high += pages[i]
		if low < pages[i] {
			low = pages[i]
		}
	}
	res := math.MaxInt
	for low <= high {
		mid := low + (high-low)/2
		if isPossible(pages, n, m, mid) {
			res = Min(mid, res)
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}
func isPossible(pages []int, n, m, mid int) bool {
	sum := 0
	studentsRequired := 1
	for i := 0; i < n; i++ {
		sum += pages[i]
		if sum > mid {
			studentsRequired++
			sum = pages[i]
		}
	}
	if studentsRequired > m {
		return false
	}
	return true
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
