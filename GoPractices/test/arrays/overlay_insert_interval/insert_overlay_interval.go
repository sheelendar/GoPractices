package main

import (
	"fmt"
	"sort"
)

func main() {
	//intervals := [][]int{{1, 3}, {6, 9}}
	//newInterval := []int{2, 5}

	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}

	//intervals := [][]int{{1, 5}}
	//newInterval := []int{0, 0}
	result := insert(intervals, newInterval)
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i][0])
		fmt.Print(" ")
		fmt.Print(result[i][1])
		fmt.Println()
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {

	n := len(intervals)

	if n == 0 {
		return [][]int{newInterval}
	}

	var result [][]int

	isNotFound := true
	var tempInterval [][]int
	for i := 0; i < n; i++ {

		if intervals[i][0] > newInterval[0] {
			isNotFound = false
			tempInterval = append(tempInterval, newInterval)
		}
		tempInterval = append(tempInterval, intervals[i])
	}
	if isNotFound {
		tempInterval = append(tempInterval, newInterval)
	}
	n = len(tempInterval)
	data := tempInterval[0]

	for i := 1; i < n; i++ {
		if data[1] >= tempInterval[i][0] {

			if data[1] < tempInterval[i][1] {
				data[1] = tempInterval[i][1]
			}

		} else {
			result = append(result, []int{data[0], data[1]})
			data = tempInterval[i]
		}

	}
	result = append(result, []int{data[0], data[1]})
	return result
}

func merge(intervals [][]int) [][]int {

	n := len(intervals)
	data := intervals[0]
	var result [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		fmt.Print(intervals[i][0])
		fmt.Print(" ")
		fmt.Print(intervals[i][1])
	}
	for i := 1; i < n; i++ {

		if data[1] >= intervals[i][0] {
			data[1] = intervals[i][1]
		} else {
			result = append(result, []int{data[0], data[1]})
			data = intervals[i]
		}
	}
	result = append(result, []int{data[0], data[1]})
	return result
}
