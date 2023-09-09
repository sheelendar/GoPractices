package main

import "fmt"

func main() {
	intervals := [][]int{{1, 5}}
	newInterval := []int{0, 0}
	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {

	n := len(intervals)

	if n == 0 {
		return [][]int{newInterval}
	}
	var result [][]int
	k := -1
	for i := 0; i < n; i++ {
		if newInterval != nil && intervals[i][1] >= newInterval[0] && !(intervals[i][1] > newInterval[1]) {
			var res []int
			if intervals[i][0] > newInterval[0] {
				res = append(res, newInterval[0])
			} else {
				res = append(res, intervals[i][0])
			}
			if intervals[i][1] < newInterval[1] {
				res = append(res, newInterval[1])
			} else {
				res = append(res, intervals[i][1])
			}
			result = append(result, res)
			k++
			newInterval = nil

		} else {

			if len(result) == 0 {
				result = append(result, intervals[i])
				k++
			} else {
				if result[k][1] >= intervals[i][0] {
					if result[k][0] > intervals[i][0] {
						result[k][0] = intervals[i][0]
					}
					if result[k][1] < intervals[i][1] {
						result[k][1] = intervals[i][1]
					}
				} else {
					result = append(result, intervals[i])
					k++
				}
			}
		}
	}
	var result1 [][]int
	if newInterval != nil {
		for i := 0; i <= k; i++ {
			if result[i][0] < newInterval[0] {
				result1 = append(result, result[i])
			} else {
				result = append(result1, newInterval)
			}
		}
	}
	return result1
}
