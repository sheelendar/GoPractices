package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 4}, {0, 4}}
	result := merge(intervals)
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i][0])
		fmt.Print(" ")
		fmt.Println(result[i][1])
	}
}

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	var result [][]int
	// sort by first value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	data := intervals[0]
	for i := 1; i < n; i++ {
		// check if there is nay overlapping if yes then include next overlap into same
		if data[1] >= intervals[i][0] {
			//  before merging check which is grater consider grater one.
			if data[1] < intervals[i][1] {
				data[1] = intervals[i][1]
			}

		} else {
			// if not overlap then collect result and increase data with next index.
			result = append(result, []int{data[0], data[1]})
			data = intervals[i]
		}
	}
	// also collect last element into resutl.
	result = append(result, []int{data[0], data[1]})
	return result
}
