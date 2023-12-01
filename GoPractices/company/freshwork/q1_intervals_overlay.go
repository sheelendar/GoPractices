package main

import (
	"fmt"
	"sort"
)

/*


 */

func main() {
	//timeList := [][]int{{11, 13}, {10, 12}}
	timeList := [][]int{{15, 18}, {2, 6}, {8, 10}, {1, 3}}

	// output:= {{10,13}}
	fmt.Println(mergeOverlap(timeList))

}
func mergeOverlap(list [][]int) [][]int { // O(1)
	var result [][]int
	n := len(list)
	if n == 0 || n == 1 {
		return list
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i][0] < list[j][0]
	})

	data := list[0]

	for i := 1; i < n; i++ {

		if data[1] > list[i][0] {
			data[1] = list[i][1]
		} else {
			result = append(result, data)
			data = list[i]
		}
	}
	result = append(result, data)
	return result

}
